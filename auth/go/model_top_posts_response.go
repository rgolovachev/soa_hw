package auth

type TopPostsResponse struct {
	PostIds []string `json:"post_ids"`
	Values  []uint64 `json:"values"`
}
