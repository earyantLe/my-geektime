package collect

type DeleteRequest struct {
	Ids []int `json:"ids,omitempty" form:"ids"`
}
