package local_variables

type Question struct {
	ID        string            `json:"id"`
	Number    string            `json:"number"`
	Content   []QuestionContent `json:"content"`
	X         int32             `json:"x"`
	Y         int32             `json:"y"`
	CreatedAt string            `json:"created_at"`
	UpdateAt  string            `json:"updated_at"`
}

type QuestionContent struct {
	Label       string `json:"label"`
	Description string `json:"description"`
	Language    string `json:"language"`
}

// create
type CreateQuestionRequest struct {
	Type    string            `json:"type"`
	X       int32             `json:"x"`
	Y       int32             `json:"y"`
	Number  string            `json:"number"`
	Content []QuestionContent `json:"content"`
}

//update
type UpdateQuestionRequest struct {
	ID       string             `json:"id"`
	Question UpdateQuestionBody `json:"question"`
}

type UpdateQuestionRequestForSwagger struct {
	Question UpdateQuestionBody `json:"question"`
}

type UpdateQuestionBody struct {
	Type    string            `json:"type"`
	X       int32             `json:"x"`
	Y       int32             `json:"y"`
	Number  string            `json:"number"`
	Content []QuestionContent `json:"content"`
}

type GetAllQuestions struct {
	Questions []Question `json:"questions"`
}

type QuestionSendsQuestionRequest struct {
	FromID string `json:"from_id"`
	ToID   string `json:"to_id"`
}

// delete request
type DeleteQuestionRequest struct {
	QuestionID string `json:"question_id"`
}
