package local_variables

type Response struct {
	StatusCode   int          `json:"response_code"`
	ResponseData ResponseData `json:"response_data"`
	Error        Error        `json:"error"`
}

type ResponseData struct {
	ID string `json:"id"`
}

type ExistsResponse struct {
	Exists bool `json:"exists"`
}
