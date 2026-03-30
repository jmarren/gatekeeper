package other

import (
	"time"

	"github.com/google/uuid"
	"github.com/ttacon/libphonenumber"
)

type IUser interface {
	GetFirstName() string
	GetLastName() string
	GetEmail() string
	GetId() uuid.UUID
	GetCreatedAt() time.Time
	GetPhoneNumber() *libphonenumber.PhoneNumber
	GetAge() int
	GetBirthMonth() string
}
