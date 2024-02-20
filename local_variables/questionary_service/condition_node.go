package local_variables

type ConditionNode struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	X           int32  `json:"x"`
	Y           int32  `json:"y"`
	CreatedAt   string `json:"created_at"`
	UpdateAt    string `json:"updated_at"`
}

// create
type CreateConditionNodeRequest struct {
	Content []ConditonNodeContent `json:"content"`
	X       int32                 `json:"x"`
	Y       int32                 `json:"y"`
}

type ConditonNodeContent struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Language    string `json:"language"`
}

// update
type UpdateConditionNodeRequest struct {
	ID            string                     `json:"id"`
	ConditionNode CreateConditionNodeRequest `json:"condition_node"`
}

type UpdateConditionNodeRequestForSwagger struct {
	ConditionNode CreateConditionNodeRequest `json:"condition_node"`
}

// delete
type DeleteConditionNodeRequest struct {
	ID string `json:"id"`
}

// has condition
type ConditionNodeHasConditionRequest struct {
	ConditionNodeID string `json:"condition_node_id"`
}
