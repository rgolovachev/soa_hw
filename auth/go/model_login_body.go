package auth

type LoginBody struct {
	Username string `json:"username,omitempty"`

	Password string `json:"password,omitempty"`
}
