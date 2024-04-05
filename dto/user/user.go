package dto

import "github.com/greetinc/greet-auth-srv/entity"

type UserRequest struct {
	ID             string `param:"id" validate:"required"`
	ProfileID      string `query:"profile_id" validate:"required"`
	SwitcherAccess string `param:"switcher_access" validate:"required"`
}

type UserResponse struct {
	ID             string                `json:"id"`
	ProfileID      string                `json:"profile_id"`
	FullName       string                `json:"full_name"`
	TotalFiles     int                   `json:"total_files"`
	Age            int                   `json:"age"`
	ProfilePicture entity.ProfilePicture `json:"profile_picture,omitempty"`
	Range          entity.RadiusRange    `json:"range,omitempty"`
	Distance       string                `json:"distance"`
}
