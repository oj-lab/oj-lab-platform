package mapper

import (
	"context"

	"github.com/OJ-lab/oj-lab-services/packages/application"
	"github.com/OJ-lab/oj-lab-services/packages/model"
	"github.com/OJ-lab/oj-lab-services/packages/utils"
	"github.com/alexedwards/argon2id"
)

func CreateUser(ctx context.Context, user model.User) error {
	db := application.GetDefaultDB()
	hashedPassword, err := utils.GetHashedPassword(*user.Password, argon2id.DefaultParams)
	if err != nil {
		return err
	}

	DbUser := model.DbUser{
		Account:        user.Account,
		HashedPassword: hashedPassword,
		Roles:          user.Roles.ToPQArray(),
	}

	return db.Create(&DbUser).Error
}

func DeleteUser(ctx context.Context, user model.User) error {
	db := application.GetDefaultDB()
	return db.Delete(&model.DbUser{Account: user.Account}).Error
}

func UpdateUser(ctx context.Context, update model.User) error {
	db := application.GetDefaultDB()

	old := model.DbUser{}
	err := db.Where("account = ?", update.Account).First(&old).Error
	if err != nil {
		return err
	}

	hashedPassword := ""
	if update.Password != nil {
		hashedPassword, err = utils.GetHashedPassword(*update.Password, argon2id.DefaultParams)
	}
	if err != nil {
		return err
	}

	new := old
	if update.Password != nil {
		new.HashedPassword = hashedPassword
	}
	if update.Roles != nil {
		new.Roles = update.Roles.ToPQArray()
	}

	return db.Model(&model.DbUser{Account: new.Account}).Updates(new).Error
}

type GetUserOptions struct {
	Account string
	Email   string
	Mobile  string
	Offset  *int
	Limit   *int
}

// Count the total number of users that match the options,
// ignoring the offset and limit.
func CountUserByOptions(ctx context.Context, options GetUserOptions) (int64, error) {
	db := application.GetDefaultDB()
	var count int64

	tx := db.
		Model(&model.DbUser{}).
		Where("account = ?", options.Account).
		Or("email = ?", options.Email).
		Or("mobile = ?", options.Mobile)

	err := tx.Count(&count).Error

	return count, err
}

func GetUserByOptions(ctx context.Context, options GetUserOptions) ([]model.User, int64, error) {
	total, err := CountUserByOptions(ctx, options)
	if err != nil {
		return nil, 0, err
	}

	db := application.GetDefaultDB()
	db_users := []model.DbUser{}

	tx := db.
		Where("account = ?", options.Account).
		Or("email = ?", options.Email).
		Or("mobile = ?", options.Mobile)
	if options.Offset != nil {
		tx = tx.Offset(*options.Offset)
	}
	if options.Limit != nil {
		tx = tx.Limit(*options.Limit)
	}

	err = tx.Find(&db_users).Error
	if err != nil {
		return nil, 0, err
	}

	users := []model.User{}
	for _, db_user := range db_users {
		users = append(users, db_user.ToUser())
	}

	return users, total, nil
}

func CheckUserPassword(ctx context.Context, account string, password string) (bool, error) {
	db := application.GetDefaultDB()
	db_user := model.DbUser{}
	err := db.Where("account = ?", account).First(&db_user).Error
	if err != nil {
		return false, err
	}
	return utils.CompareWithHashedPassword(password, db_user.HashedPassword)
}