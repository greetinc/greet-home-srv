package dto

type LikeRequest struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	LikeID      string `json:"like_id"`
	IsFollowing bool
}

type LikeResponse struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	LikeID      string `json:"like_id"`
	FullName    string `json:"full_name"`
	IsFollowing bool
}
