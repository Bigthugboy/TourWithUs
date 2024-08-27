package query

import (
	"errors"
	"fmt"
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/mapper"

	"github.com/Bigthugboy/TourWithUs/internals/domain/model"
	"gorm.io/gorm"
)

type TouristObject struct {
	gorm.Model
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	ProfilePic string `json:"profilePic"`
	Username   string `json:"username"`
}

// CreateTourist maps model.TouristDetails to TouristObject and inserts into the database.
func (t *TourDB) CreateTourist(tourist *model.TouristDetails) (*model.TouristDetails, error) {
	touristObj := mapper.MapModelToObject(tourist)
	if err := t.DB.Create(&touristObj).Error; err != nil {
		return nil, err
	}
	return mapper.MapObjectToModel(&touristObj), nil
}

// InsertTourist inserts a tourist into the database.
func (t *TourDB) InsertTourist(tourist model.TouristDetails) (int64, error) {
	if t.DB == nil {
		return -1, fmt.Errorf("database connection is not initialized")
	}
	var existingUser TouristObject
	if err := t.DB.Where("email = ?", tourist.Email).First(&existingUser).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return -1, err
	}
	if existingUser.ID != 0 {
		return -1, fmt.Errorf("user with email '%s' already exists", tourist.Email)
	}
	touristObj := mapper.MapModelToObject(&tourist)
	result := t.DB.Create(&touristObj)
	if err := result.Error; err != nil {
		return -1, err
	}
	return result.RowsAffected, nil
}

func (t *TourDB) SearchTouristByEmail(email string) (model.TouristDetails, error) {
	if t.DB == nil {
		return model.TouristDetails{}, fmt.Errorf("database connection is not initialized")
	}
	var user TouristObject
	if err := t.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.TouristDetails{}, nil
		}
		return model.TouristDetails{}, err
	}

	return *mapper.MapObjectToModel(&user), nil
}

func (t *TourDB) GetTouristByID(userID string) (model.TouristDetails, error) {
	var user TouristObject
	if err := t.DB.Preload("Wallet").First(&user, userID).Error; err != nil {
		return model.TouristDetails{}, fmt.Errorf("failed to find user: %v", err)
	}
	return *mapper.MapObjectToModel(&user), nil
}

func (t *TourDB) GetAllTourists() ([]model.TouristDetails, error) {
	var users []TouristObject
	if err := t.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	var tourists []model.TouristDetails
	for _, user := range users {
		tourists = append(tourists, *mapper.MapObjectToModel(&user))
	}
	return tourists, nil
}

func (t *TourDB) DeleteTouristByID(userID string) error {
	if err := t.DB.Delete(&TouristObject{}, userID).Error; err != nil {
		return err
	}
	return nil
}

func (t *TourDB) DeleteTouristByEmail(email string) error {
	if err := t.DB.Where("email = ?", email).Delete(&TouristObject{}).Error; err != nil {
		return err
	}
	return nil
}
