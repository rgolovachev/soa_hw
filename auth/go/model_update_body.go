package auth

type UpdateBody struct {
	FirstName string `json:"firstName,omitempty"`

	LastName string `json:"lastName,omitempty"`

	BirthDate string `json:"birthDate,omitempty"`

	Email string `json:"email,omitempty"`

	Telephone string `json:"telephone,omitempty"`
}
