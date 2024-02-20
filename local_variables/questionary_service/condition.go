package local_variables

type Condition struct {
	ID        string `json:"id"`
	Order     int32  `json:"order"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"updated_at"`
}

type CreateConditionRequest struct {
	ConditionNodeID string `json:"condition_node_id"`
}

type DeleteConditionRequest struct {
	ConditionNodeID string `json:"condition_node_id"`
	ConditionID     string `json:"condition_id"`
}
