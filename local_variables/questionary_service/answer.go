package local_variables

type Answer struct {
	ID        string `json:"id"`
	Label     string `json:"label"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateAnswerRequest struct {
	QuestionID string     `json:"question_id"`
	Answer     AnswerBody `json:"answer"`
}

type CreateAnswerRequestForSwagger struct {
	Answer AnswerBody `json:"answer"`
}

type AnswerBody struct {
	Content []AnswerContent `json:"content"`
}

type AnswerContent struct {
	Label    string `json:"label"`
	Value    string `json:"value"`
	Langauge string `json:"language"`
}

type SortAnswersRequest struct {
	AnswerOrders map[string]int32 `json:"answer_orders"`
}

type DeleteAnswerRequest struct {
	QuestionID string `json:"question_id"`
	AnswerID   string `json:"answer_id"`
}
