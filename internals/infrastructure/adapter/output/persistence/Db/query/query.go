package query

import (
	"errors"
	"fmt"
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto"
	"github.com/jinzhu/gorm"
)

func (t *TourDB) InsertTourist(tourist dto.TouristObject) (*dto.TouristObject, int64, error) {
	if t.DB == nil {
		return nil, -1, fmt.Errorf("database connection is not initialized")
	}
	if err := t.DB.Create(&tourist).Error; err != nil {
		return nil, -1, fmt.Errorf("exception inserting new tourist: %w", err)
	}
	return &tourist, 1, nil
}

// SearchTouristByEmail finds a tourist by email.
func (t *TourDB) SearchTouristByEmail(email string) (dto.TouristObject, error) {
	if t.DB == nil {
		return dto.TouristObject{}, fmt.Errorf("database connection is not initialized")
	}
	var user dto.TouristObject
	//err := t.DB.Where("email = ?", email).First(&user).Error
	//err := t.DB.Select("id", "first_name", "last_name", "email", "password", "profile_pic", "username").Where("email = ?", email).First(&user).Error
	err := t.DB.Select("id, first_name, last_name, email, password, profile_pic, username").
		Where("email = ?", email).
		First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return dto.TouristObject{}, err
	} else if err != nil {
		return dto.TouristObject{}, fmt.Errorf("error querying database: %w", err)
	}
	return user, nil
}

// GetTouristByID retrieves a tourist by ID.
func (t *TourDB) GetTouristByID(userID string) (dto.TouristObject, error) {
	if t.DB == nil {
		return dto.TouristObject{}, fmt.Errorf("database connection is not initialized")
	}
	var user dto.TouristObject
	err := t.DB.Preload("Wallet").First(&user, userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.TouristObject{}, fmt.Errorf("tourist not found")
		}
		return dto.TouristObject{}, fmt.Errorf("failed to find user: %w", err)
	}
	return user, nil
}

// GetAllTourists retrieves all tourists from the database.
func (t *TourDB) GetAllTourists() ([]dto.TouristObject, error) {
	if t.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}
	var users []dto.TouristObject
	err := t.DB.Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve tourists: %w", err)
	}
	return users, nil
}

// DeleteTouristByID deletes a tourist by their ID.
func (t *TourDB) DeleteTouristByID(userID string) error {
	if t.DB == nil {
		return fmt.Errorf("database connection is not initialized")
	}
	err := t.DB.Delete(&dto.TouristObject{}, userID).Error
	if err != nil {
		return fmt.Errorf("failed to delete tourist by ID: %w", err)
	}
	return nil
}

// DeleteTouristByEmail deletes a tourist by their email.
func (t *TourDB) DeleteTouristByEmail(email string) error {
	if t.DB == nil {
		return fmt.Errorf("database connection is not initialized")
	}
	err := t.DB.Where("email = ?", email).Delete(&dto.TouristObject{}).Error
	if err != nil {
		return fmt.Errorf("failed to delete tourist by email: %w", err)
	}
	return nil
}
