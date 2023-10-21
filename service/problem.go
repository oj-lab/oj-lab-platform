package service

import (
	"context"

	gormAgent "github.com/OJ-lab/oj-lab-services/core/agent/gorm"
	"github.com/OJ-lab/oj-lab-services/service/business"
	"github.com/OJ-lab/oj-lab-services/service/mapper"
	"github.com/OJ-lab/oj-lab-services/service/model"
)

func GetProblem(ctx context.Context, slug string) (*model.Problem, error) {
	db := gormAgent.GetDefaultDB()
	problem, err := mapper.GetProblem(db, slug)
	if err != nil {
		return nil, err
	}
	return problem, nil
}

func GetProblemInfoList(ctx context.Context) ([]model.ProblemInfo, int64, error) {
	return business.GetProblemInfoList(ctx)
}

func PutProblemPackage(ctx context.Context, slug, zipFile string) error {
	localDir := "/tmp/" + slug
	err := business.UnzipProblemPackage(ctx, zipFile, localDir)
	if err != nil {
		return err
	}

	err = business.PutProblemPackage(ctx, slug, localDir)
	if err != nil {
		return err
	}

	return nil
}

func PostSubmission(ctx context.Context, problemSlug, code, language string) (*model.JudgeTaskSubmission, error) {
	submission := model.NewSubmission("", problemSlug, code, language)
	db := gormAgent.GetDefaultDB()
	result, err := mapper.CreateSubmission(db, submission)
	if err != nil {
		return nil, err
	}

	return result, nil
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
