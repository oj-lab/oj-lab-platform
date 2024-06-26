package judge_model

import (
	"github.com/google/uuid"
	"github.com/oj-lab/oj-lab-platform/models"
	"gorm.io/gorm"
)

func CreateJudgeResult(tx *gorm.DB, result JudgeResult) (*JudgeResult, error) {
	result.UID = uuid.New()
	result.MetaFields = models.NewMetaFields()
	if !result.Verdict.IsValid() {
		return nil, ErrInvalidJudgeVerdict
	}

	return &result, tx.Create(&result).Error
}
