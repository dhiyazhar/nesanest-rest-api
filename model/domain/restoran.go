package domain

type Restoran struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	ImageUrl    string `json:"image_url"`
}
