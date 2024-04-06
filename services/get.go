package services

import (
	"fmt"

	entity "github.com/greetinc/greet-auth-srv/entity"

	dto "greet-home-srv/dto/user"

	util "github.com/greetinc/greet-util/s"
)

func (s *userService) GetAll(req dto.UserRequest) ([]dto.UserResponse, error) {
	currentUserCoordinates, err := s.UserR.GetUserCoordinates(req.ID)
	if err != nil {
		return nil, err
	}

	data, err := s.UserR.GetAll(req)
	if err != nil {
		return nil, err
	}

	var responses []dto.UserResponse
	var totalFiles int

	for _, msg := range data {

		totalFiles++

		// Process ProfilePicture if available
		var profilePicture entity.ProfilePicture
		if msg.ProfilePicture.ID != "" {
			profilePicture = entity.ProfilePicture{
				ID:       msg.ProfilePicture.ID,
				UserID:   msg.ID,
				FileName: msg.ProfilePicture.FileName,
				FilePath: msg.ProfilePicture.FilePath,
			}
		}
		// Process Range if available
		rangeData := entity.RadiusRange{
			Longitude: msg.Range.Longitude,
			Latitude:  msg.Range.Latitude,
		}

		otherUserCoordinates := entity.RadiusRange{
			Longitude: msg.Range.Longitude,
			Latitude:  msg.Range.Latitude,
		}

		distance := int(util.Haversine(currentUserCoordinates, otherUserCoordinates)/1000 + 1)

		response := dto.UserResponse{
			ID:             msg.ID,
			ProfileID:      msg.ProfileID,
			FullName:       msg.FullName,
			Age:            msg.Age,
			ProfilePicture: profilePicture,
			TotalFiles:     totalFiles, // Add the totalFiles count to the UserResponse
			Range:          rangeData,
			Distance:       fmt.Sprintf("%d Km", distance), // Add the distance to the UserResponse

		}
		responses = append(responses, response)

	}

	return responses, nil
}

func (s *userService) IsValidProfileID(profileID string) bool {
	// Implementasikan logika validasi sesuai dengan akses ke database
	// Contoh sederhana: cek apakah profileID ada di database
	_, err := s.UserR.GetUserByProfileID(profileID)
	return err == nil
}
