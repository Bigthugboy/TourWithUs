package OperatorDb

import (
	"errors"
	"fmt"
	dto "github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto/operatorDto"
)

func (o *OperatorDb) SaveTourOperator(object dto.OperatorDto) (dto.SavedOperatorRes, error) {
	if o.DB == nil {
		return dto.SavedOperatorRes{}, errors.New("database not initialized")
	}
	if err := o.DB.Create(&object).Error; err != nil {
		return dto.SavedOperatorRes{}, fmt.Errorf("error inserting new tourModel: %w", err)
	}
	return dto.SavedOperatorRes{
		Id:      object.ID,
		Message: "tour operator created",
		Email:   object.Email,
	}, nil
}

func (o *OperatorDb) GetTourOperatorById(id string) (dto.OperatorDto, error) {
	if o.DB == nil {
		return dto.OperatorDto{}, errors.New("operatorModel DB Not Initialized")
	}
	var operator dto.OperatorDto
	if err := o.DB.Where("id = ?", id).First(&operator).Error; err != nil {
		return dto.OperatorDto{}, fmt.Errorf("exception getting Operator Model: %w", err)
	}
	return operator, nil
}

func (o *OperatorDb) GetAllTourOperator() ([]dto.OperatorDto, error) {
	if o.DB == nil {
		return nil, errors.New("operatorModel DB not initialized")
	}
	var operators []dto.OperatorDto
	if err := o.DB.Find(&operators).Error; err != nil {
		return nil, fmt.Errorf("exception getting tour operators: %w", err)
	}
	return operators, nil
}

func (o *OperatorDb) GetTourOperatorByRating(limit int) ([]dto.OperatorDto, error) {
	var operators []dto.OperatorDto
	if err := o.DB.Order("rating DESC").Limit(limit).Find(&operators).Error; err != nil {
		return nil, fmt.Errorf("error retrieving operators by rating: %w", err)
	}
	return operators, nil
}

func (o *OperatorDb) DeleteTourOperator(id string) (string, error) {
	if o.DB == nil {
		return "", errors.New("database not initialized")
	}
	if err := o.DB.Where("id = ?", id).Delete(&dto.OperatorDto{}).Error; err != nil {
		return "", fmt.Errorf("exception deleting tourModel: %w", err)
	}
	return "tour operator deleted", nil
}

func (o *OperatorDb) GetTourOperatorByEmail(email string) (dto.OperatorDto, error) {
	if o.DB == nil {
		return dto.OperatorDto{}, errors.New("database not initialized")
	}
	var operator dto.OperatorDto
	if err := o.DB.Where("email = ?", email).First(&operator).Error; err != nil {
		return dto.OperatorDto{}, fmt.Errorf("exception getting tour model: %w", err)
	}
	return operator, nil
}
