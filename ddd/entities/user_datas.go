package entities

import "time"

type UserData struct {
	UserID               int
	CreatedAt            time.Time
	UpdatedAt            time.Time
	FirstName            string
	LastName             string
	PhotoID              string
	BirthDate            time.Time
	Sex                  int
	DriverLicensePhotoID string
}
