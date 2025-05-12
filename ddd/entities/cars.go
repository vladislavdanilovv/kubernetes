package entities

import (
	"time"
)

type Car struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int
	Brand     int
	Model     int
	Color     int
	IssueDate int
	CarNumber string
	IsDeleted bool
	DeletedAt time.Time
}

//
//const (
//	customError string = "ErrorValidationCarBrand"
//)

//func (c *Car) NewCar() (*Car, error) {
//	// Example
//	if c.Brand == 0 {
//		return nil, errors.New(customError)
//	}
//
//	return &Car{
//		ID:        c.ID,
//		CreatedAt: c.CreatedAt,
//		UpdatedAt: c.UpdatedAt,
//		UserID:    c.UserID,
//		Brand:     c.Brand,
//		Model:     c.Model,
//		Color:     c.Color,
//		IssueDate: c.IssueDate,
//		CarNumber: c.CarNumber,
//		IsDeleted: c.IsDeleted,
//		DeletedAt: c.DeletedAt,
//	}, nil
//}
