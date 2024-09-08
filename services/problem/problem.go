package problem_service

import (
	"context"

	problem_model "github.com/oj-lab/platform/models/problem"
	gorm_agent "github.com/oj-lab/platform/modules/agent/gorm"

	"gorm.io/gorm"
)

func GetProblem(ctx context.Context, slug string) (*problem_model.Problem, error) {
	db := gorm_agent.GetDefaultDB()
	problem, err := problem_model.GetProblem(db, slug)
	if err != nil {
		return nil, err
	}
	return problem, nil
}

func PutProblem(ctx context.Context, problem problem_model.Problem) error {
	db := gorm_agent.GetDefaultDB()
	err := problem_model.CreateProblem(db, problem)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProblem(ctx context.Context, slug string) error {
	db := gorm_agent.GetDefaultDB()
	err := problem_model.DeleteProblem(db, slug)
	if err != nil {
		return err
	}
	return nil
}

func CheckProblemSlug(ctx context.Context, slug string) (bool, error) {
	db := gorm_agent.GetDefaultDB()
	problem, err := problem_model.GetProblem(db, slug)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return true, nil
		}
		return false, err
	}
	return problem == nil, nil
}

func PutProblemPackage(ctx context.Context, slug, zipFile string) error {
	localDir := "/tmp/" + slug
	err := unzipProblemPackage(ctx, zipFile, localDir)
	if err != nil {
		return err
	}

	err = putProblemPackage(ctx, slug, localDir)
	if err != nil {
		return err
	}

	return nil
}

// func Judge(ctx context.Context, slug string, code string, language string) (
// 	[]map[string]interface{}, error,
// ) {
// 	request := judger.JudgeRequest{
// 		Code:     code,
// 		Language: language,
// 	}
// 	responseBody, err := judger.PostJudgeSync(slug, request)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return responseBody, nil
// }
