package auth

type PostStatsResponse struct {
	PostId string `json:"post_id,omitempty"`

	Likes uint64 `json:"likes,omitempty"`

	Views uint64 `json:"views,omitempty"`
}
