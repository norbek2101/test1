package local_variables

type Handbook struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}
type HandbookWithAnswers struct {
	ID          string           `json:"id"`
	UserID      string           `json:"user_id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	CreatedAt   string           `json:"created_at"`
	UpdatedAt   string           `json:"updated_at"`
	DeletedAt   string           `json:"deleted_at"`
	Answers     []HandbookAnswer `json:"answers"`
}

// create
type CreateHandbookRequest struct {
	UserId      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateHandbookRequest struct {
	ID       string             `json:"id"`
	Handbook UpdateHandbookBody `json:"handbook"`
}

type UpdateHandbookRequestForSwagger struct {
	Handbook UpdateHandbookBody `json:"handbook"`
}

type UpdateHandbookBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// delete request
type DeleteHandbookRequest struct {
	HandbookID string `json:"handbook_id"`
}

type GetAllHandbooks struct {
	Handbooks []Handbook `json:"handbooks"`
}

type GetHandbook struct {
	Handbook HandbookWithAnswers `json:"handbook"`
}
