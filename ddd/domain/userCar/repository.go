package userCar

import "gis/ddd/aggregate"

type userCarRepository interface {
	Get(userID int) (aggregate.UserCar, error)
}
