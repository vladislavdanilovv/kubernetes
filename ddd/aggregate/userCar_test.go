package aggregate

import (
	"errors"
	"fmt"
	"testing"
)

type testCase struct {
	test        string
	userID      int
	firstName   string
	expectedErr error
}

func TestUserCar(t *testing.T) {
	testCases := []testCase{
		{
			test:        "Zero userID",
			userID:      0,
			firstName:   "Tester",
			expectedErr: customZeroError,
		},
		{
			test:        "valid",
			userID:      1,
			firstName:   "ValidName",
			expectedErr: nil,
		},
	}

	for _, value := range testCases {
		t.Run(value.test, func(t *testing.T) {
			testUser, err := NewUserCar(value.userID, value.firstName)

			fmt.Sprintf("New object of user with car %v", testUser)

			if !errors.Is(err, value.expectedErr) {
				t.Errorf("err = %v, expected = %v", err, value.expectedErr)
			}
		})
	}

}
