package local_variables

type HandbookAnswer struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Value     string `json:"value"`
	Order     int32  `json:"order"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

// create
type CreateHandbookAnswer struct {
	ID             string             `json:"id"`
	HandbookAnswer HandbookAnswerBody `json:"handbook_answer"`
}

type CreateHandbookAnswerForSwagger struct {
	HandbookAnswer HandbookAnswerBody `json:"handbook_answer"`
}

type HandbookAnswerBody struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

type UpdateHandbookAnswerRequest struct {
	ID             string                   `json:"id"`
	HandbookAnswer UpdateHandbookAnswerBody `json:"handbook_answer"`
}

type UpdateHandbookAnswerRequestForSwagger struct {
	HandbookAnswer UpdateHandbookAnswerBody `json:"handbook_answer"`
}

type UpdateHandbookAnswerBody struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

// delete request
type DeleteHandbookAnswerRequest struct {
	HandbookAnswerID string `json:"handbook_answer_id"`
}

type GetAllHandbookAnswers struct {
	HandbookAnswers []HandbookAnswer `json:"handbookanswers"`
}
