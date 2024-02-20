package local_variables

type Country struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Code      string `json:"code"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetAllCountryResponse struct {
	Countries []Country `json: "countries`
}
