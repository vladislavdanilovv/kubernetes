package aggregate

import (
	"errors"
	"gis/ddd/entities"
)

type UserCar struct {
	userData *entities.UserData
	car      *entities.Car
}

var (
	customZeroError = errors.New("Zero")
)

func NewUserCar(userID int, firstName string) (UserCar, error) {
	// Example
	if userID == 0 {
		return UserCar{}, customZeroError
	}

	userData := &entities.UserData{
		UserID:    userID,
		FirstName: firstName,
	}

	return UserCar{
			userData: userData,
			car:      nil},
		nil
}
