package web

type RestoranCreateRequest struct {
	Name        string `validate:"required,min=1,max=100" json:"name"`
	Description string `validate:"required,min=1,max=100" json:"description"`
	Address     string `validate:"required,min=1,max=200" json:"address"`
	ImageUrl    string `validate:"required,min=1,max=200" json:"image_url"`
}
