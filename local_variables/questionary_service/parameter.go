package local_variables

type Parameter struct {
	ID        string `json:"id"`
	Key       string `json:"key"`
	Row       string `json:"row"`
	Operator  string `json:"operator"`
	Answer    string `json:"answer"`
	Answers   string `json:"answers"`
	Column    string `json:"column"`
	Columns   string `json:"columns"`
	Text      string `json:"text"`
	ValueType string `json:"value_type"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateParameterRequest struct {
	ConditionID string              `json:"condition_id"`
	Parameter   CreateParameterBody `json:"parameter"`
}

type CreateParameterRequestForSwagger struct {
	Parameter CreateParameterBody `json:"parameter"`
}

type CreateParameterBody struct {
	Key       string `json:"key"`
	Row       string `json:"row"`
	Operator  string `json:"operator"`
	Answer    string `json:"answer"`
	Answers   string `json:"answers"`
	Column    string `json:"column"`
	Columns   string `json:"columns"`
	Text      string `json:"text"`
	ValueType string `json:"value_type"`
}

// update
type UpdateParameterRequest struct {
	ConditionID string              `json:"condition_id"`
	ParameterID string              `json:"parameter_id"`
	Parameter   CreateParameterBody `json:"parameter"`
}

// delete
type DeleteParameterRequest struct {
	ConditionID string `json:"condition_id"`
	ParameterID string `json:"parameter_id"`
}
