package auth

type TopUsersResponse struct {
	Usernames []string `json:"usernames"`
	Likes     []uint64 `json:"likes"`
}
