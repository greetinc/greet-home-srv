package repositories

import (
	"fmt"

	dto "github.com/greetinc/greet-auth-srv/dto/auth"
	entity "github.com/greetinc/greet-auth-srv/entity"
	util "github.com/greetinc/greet-util/s"
)

// func (r *userRepository) GetAll(req dto.UserRequest) ([]dto.UserResponse, error) {
// 	var data []dto.UserResponse

// 	err := r.DB.
// 		Model(&entity.User{}).
// 		Joins("LEFT JOIN like_relations ON users.id = like_relations.like_id AND like_relations.user_id = ?", req.ID).
// 		Joins("LEFT JOIN friends ON friends.like_id = users.id AND friends.user_id = ?", req.ID).
// 		Where("like_relations.is_following IS NULL OR like_relations.is_following = ?", false).
// 		Where("friends.id IS NULL").
// 		Where("users.id != ?", req.ID).
// 		Find(&data).
// 		Error
// 	if err != nil {
// 		return []dto.UserResponse{}, err
// 	}

// 	return data, nil
// }

func (r *userRepository) GetAll(req dto.UserRequest) ([]dto.UserResponse, error) {
	var data []struct {
		ID                 string  `db:"id"`
		ProfileID          string  `db:"profile_id"`
		FullName           string  `db:"full_name"`
		Age                int     `db:"age"`
		ProfilePictureID   string  `db:"profile_picture_id"`
		ProfilePictureName string  `db:"profile_picture_name"`
		ProfilePicturePath string  `db:"profile_picture_path"`
		MinAge             int     `db:"min_age"`
		MaxAge             int     `db:"max_age"`
		Longitude          float64 `db:"longitude"`
		Latitude           float64 `db:"latitude"`
	}

	// Query
	err := r.DB.
		Model(&entity.User{}).
		Select("users.id, user_details.profile_id, user_details.full_name, user_details.age, profile_pictures.id as profile_picture_id, profile_pictures.file_name as profile_picture_name, profile_pictures.file_path as profile_picture_path, COALESCE(MIN(interest_ages.min_age), 0) as min_age, COALESCE(MAX(interest_ages.max_age), 0) as max_age, COALESCE(radius_ranges.longitude, 0) as longitude, COALESCE(radius_ranges.latitude, 0) as latitude").
		Joins("LEFT JOIN user_details ON user_details.user_id = users.id"). // JOIN dengan UserDetail
		Joins("LEFT JOIN like_relations ON users.id = like_relations.like_id AND like_relations.user_id = ?", req.ID).
		Joins("LEFT JOIN friends ON friends.like_id = users.id AND friends.user_id = ?", req.ID).
		Joins("LEFT JOIN profile_pictures ON profile_pictures.user_id = users.id").
		Joins("LEFT JOIN interest_ages ON user_details.age BETWEEN interest_ages.min_age AND interest_ages.max_age AND interest_ages.user_id = ?", req.ID).
		Joins("LEFT JOIN boost_profiles ON boost_profiles.user_id = users.id").
		Joins("LEFT JOIN radius_ranges ON radius_ranges.user_id = users.id").
		Where("like_relations.is_following IS NULL OR like_relations.is_following = ?", false).
		Where("friends.id IS NULL").
		Where("users.id != ?", req.ID).
		Group("users.id, user_details.profile_id, user_details.full_name, user_details.age, profile_pictures.id, profile_pictures.file_name, profile_pictures.file_path, boost_profiles.id, radius_ranges.longitude, radius_ranges.latitude").
		Having("MAX(interest_ages.max_age) >= user_details.age"). // Filter berdasarkan rentang usia dari tabel interest_ages.
		Order("boost_profiles.id IS NOT NULL DESC, RANDOM()").    // Prioritize users who boosted their profiles
		Find(&data).
		Error

	if err != nil {
		return nil, err
	}

	// Map untuk menyimpan data file dan profil
	fileMap := make(map[string]map[string]entity.File)
	profilePictureMap := make(map[string]entity.ProfilePicture)

	// Mengelompokkan data berdasarkan user ID
	for _, entry := range data {
		if _, found := fileMap[entry.ID]; !found {
			fileMap[entry.ID] = make(map[string]entity.File)
		}

		// Menambahkan data profil ke map
		profilePicture := entity.ProfilePicture{
			ID:       entry.ProfilePictureID,
			FileName: entry.ProfilePictureName,
			FilePath: entry.ProfilePicturePath,
		}

		profilePictureMap[entry.ID] = profilePicture
	}
	currentUserCoordinates := entity.RadiusRange{
		Longitude: 107.0477309,
		Latitude:  -6.7341488,
	}

	// Mengonversi map menjadi slice
	var result []dto.UserResponse
	for _, entry := range data {
		otherUserCoordinates := entity.RadiusRange{
			Longitude: entry.Longitude,
			Latitude:  entry.Latitude,
		}
		distance := int(util.Haversine(currentUserCoordinates, otherUserCoordinates) / 1000)

		userResponse := dto.UserResponse{
			ID:             entry.ID,
			ProfileID:      entry.ProfileID,
			FullName:       entry.FullName,
			Age:            entry.Age,
			ProfilePicture: profilePictureMap[entry.ID],
			Range: entity.RadiusRange{
				Longitude: entry.Longitude,
				Latitude:  entry.Latitude,
			},
			Distance: fmt.Sprintf("%d Km", distance), // Menambahkan satuan "KM",
		}

		result = append(result, userResponse)
	}

	return result, nil

}

func (r *userRepository) GetUserByProfileID(profileID string) (*entity.User, error) {
	var user entity.User
	err := r.DB.Where("profile_id = ?", profileID).First(&user).Error
	return &user, err
}

func (r *userRepository) GetUserCoordinates(userID string) (entity.RadiusRange, error) {
	var coordinates entity.RadiusRange

	if err := r.DB.
		Model(&entity.RadiusRange{}).
		Where("user_id = ?", userID).
		First(&coordinates).
		Error; err != nil {
		return entity.RadiusRange{}, err
	}

	return coordinates, nil
}
