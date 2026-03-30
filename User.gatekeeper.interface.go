
package src
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
