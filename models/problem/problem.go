package problem_model

import "github.com/oj-lab/oj-lab-platform/models"

type Problem struct {
	models.MetaFields
	Slug        string        `json:"slug" gorm:"primaryKey"`
	Title       string        `json:"title" gorm:"not null"`
	Description *string       `json:"description,omitempty"`
	Tags        []*ProblemTag `json:"tags" gorm:"many2many:problem_problem_tags;"`
	Solved      bool          `json:"solved,omitempty" gorm:"-"`
}

type ProblemTag struct {
	models.MetaFields
	Name     string     `json:"name" gorm:"primaryKey"`
	Problems []*Problem `json:"problems,omitempty" gorm:"many2many:problem_problem_tags;"`
}

var ProblemInfoSelection = append([]string{"slug", "title"}, models.MetaFieldsSelection...)
