package auth

type RegisterBody struct {
	Username string `json:"username,omitempty"`

	Password string `json:"password,omitempty"`
}
