package dto

type UserRequest struct {
	Email             string `json:"email"`
	Mobilephone       string `json:"mobilephone"`
	Active            bool   `json:"active"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	Gender            string `json:"gender"`
	BirthPlace        string `json:"birth_place"`
	DateOfBirth       string `json:"date_of_birth"`
	RegisterProvider  string `json:"register_provider"`
	Photo             string `json:"photo"`
	Profession        string `json:"profession"`
	MobilephonePrefix string `json:"mobilephone_prefix"`
}
