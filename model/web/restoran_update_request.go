package web

type RestoranUpdateRequest struct {
	Id          int    `validate:"required"`
	Name        string `validate:"required,min=1,max=200" json:"name"`
	Description string `validate:"required,min=1,max=200" json:"description"`
}
