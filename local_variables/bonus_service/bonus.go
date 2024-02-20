package local_variables

type Bonus struct {
	ID              string `json:"id"`
	CustomerID      string `json:"customer_id" binding:"required"`
	Gender          string `json:"gender"`
	Birthday        string `json:"birthday"`
	FamilyCondition string `json:"family_condition"`
	Children        int32  `json:"children"`
	Workplace       string `json:"workplace"`
}

type CreateBonus struct {
	CustomerID      string `json:"customer_id" binding:"required"`
	Gender          string `json:"gender"`
	Birthday        string `json:"birthday"`
	FamilyCondition string `json:"family_condition"`
	Children        int32  `json:"children"`
	Workplace       string `json:"workplace"`
}
