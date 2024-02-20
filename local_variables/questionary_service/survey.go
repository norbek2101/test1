package local_variables

type Survey struct {
	ID        string `json:"id"`
	Title     string `json:"title" binding:"required"`
	UserID    string `json:"user_id" binding:"required"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateSurveyRequest struct {
	Title    string `json:"title" binding:"required"`
	IsActive bool   `json:"is_active"`
}

// starts with
type SurveyStartsWith struct {
	SurveyID   string `json:"survey_id"`
	QuestionID string `json:"question_id"`
}

// delete
type DeleteSurveyRequest struct {
	SurveyID string `json:"survey_id"`
}
