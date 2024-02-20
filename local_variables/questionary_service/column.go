package local_variables

type Column struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Label     string `json:"label"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateColumnRequest struct {
	QuestionID string     `json:"question_id"`
	Column     ColumnBody `json:"column"`
}

type CreateColumnRequestForSwagger struct {
	Column ColumnBody `json:"column"`
}

type ColumnBody struct {
	Content []ColumnContent `json:"content"`
}

type ColumnContent struct {
	Title    string `json:"title"`
	Value    string `json:"value"`
	Langauge string `json:"language"`
}

type SortColumnsRequest struct {
	ColumnOrders map[string]int32 `json:"column_orders"`
}

type DeleteColumnRequest struct {
	QuestionID string `json:"question_id"`
	ColumnID   string `json:"column_id"`
}
