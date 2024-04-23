package idvalidator

type ThaiId struct {
	Id string `json:"id" binding:"required"`
}
