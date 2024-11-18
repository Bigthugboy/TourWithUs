package query

import (
	"errors"
	"fmt"
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto/touristDto"
	"github.com/jinzhu/gorm"
)

func (t *TourDB) InsertTourist(tourist touristDto.TouristObject) (*touristDto.TouristObject, int64, error) {
	if t.DB == nil {
		return nil, -1, fmt.Errorf("database connection is not initialized")
	}
	if err := t.DB.Create(&tourist).Error; err != nil {
		return nil, -1, fmt.Errorf("exception inserting new touristModel: %w", err)
	}
	return &tourist, 1, nil
}

// SearchTouristByEmail finds a touristModel by email.
func (t *TourDB) SearchTouristByEmail(email string) (touristDto.TouristObject, error) {
	if t.DB == nil {
		return touristDto.TouristObject{}, fmt.Errorf("database connection is not initialized")
	}
	var user touristDto.TouristObject
	//err := t.DB.Where("email = ?", email).First(&user).Error
	//err := t.DB.Select("id", "first_name", "last_name", "email", "password", "profile_pic", "username").Where("email = ?", email).First(&user).Error
	err := t.DB.Select("id, first_name, last_name, email, password, profile_pic, username").
		Where("email = ?", email).
		First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return touristDto.TouristObject{}, err
	} else if err != nil {
		return touristDto.TouristObject{}, fmt.Errorf("error querying database: %w", err)
	}
	return user, nil
}

// GetTouristByID retrieves a touristModel by ID.
func (t *TourDB) GetTouristByID(userID string) (touristDto.TouristObject, error) {
	if t.DB == nil {
		return touristDto.TouristObject{}, fmt.Errorf("database connection is not initialized")
	}
	var user touristDto.TouristObject
	err := t.DB.Preload("Wallet").First(&user, userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return touristDto.TouristObject{}, fmt.Errorf("touristModel not found")
		}
		return touristDto.TouristObject{}, fmt.Errorf("failed to find user: %w", err)
	}
	return user, nil
}

// GetAllTourists retrieves all tourists from the database.
func (t *TourDB) GetAllTourists() ([]touristDto.TouristObject, error) {
	if t.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}
	var users []touristDto.TouristObject
	err := t.DB.Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve tourists: %w", err)
	}
	return users, nil
}

// DeleteTouristByID deletes a touristModel by their ID.
func (t *TourDB) DeleteTouristByID(userID string) error {
	if t.DB == nil {
		return fmt.Errorf("database connection is not initialized")
	}
	err := t.DB.Delete(&touristDto.TouristObject{}, userID).Error
	if err != nil {
		return fmt.Errorf("failed to delete touristModel by ID: %w", err)
	}
	return nil
}

// DeleteTouristByEmail deletes a touristModel by their email.
func (t *TourDB) DeleteTouristByEmail(email string) error {
	if t.DB == nil {
		return fmt.Errorf("database connection is not initialized")
	}
	err := t.DB.Where("email = ?", email).Delete(&touristDto.TouristObject{}).Error
	if err != nil {
		return fmt.Errorf("failed to delete touristModel by email: %w", err)
	}
	return nil
}
