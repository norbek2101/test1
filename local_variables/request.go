package local_variables

type GetRequest struct {
	ID string `json:"id"`
}

type DeleteRequest struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
}
