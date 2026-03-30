package example

import (
	"bytes"
	"net/http"
	"net/url"
	"testing"
)

func TestUser(t *testing.T) {

	form := url.Values{}
	form.Add("Email", "test@example.com")
	form.Add("FirstName", "peach")
	form.Add("LastName", "wxyz")
	form.Add("Age", "22")
	form.Add("BirthMonth", "January")
	form.Add("Id", "03eefd05-c2e3-4917-a1bf-6f112e78295a")
	form.Add("CreatedAt", "2006-01-02 15:04:05")
	form.Add("PhoneNumber", "18476512236")

	// phoneNumber, err := libphonenumber.Parse("18476512236", "US")
	//
	// isValid := libphonenumber.IsValidNumber(phoneNumber)
	//
	// if !isValid {
	// 	t.Log("invalid phone number")
	// }
	//
	// if err != nil {
	// 	t.Logf("err parsing phone number = %s\n", err)
	// } else {
	// 	t.Logf("phoneNumber = %v\n", libphonenumber.Format(phoneNumber, libphonenumber.NATIONAL))
	// }

	// 2. Encode the form data into a bytes.Buffer
	// The body of an http.Request needs to be an io.Reader
	payload := bytes.NewBufferString(form.Encode())

	// 3. Create a mock HTTP request
	req, err := http.NewRequest("POST", "/submit", payload)
	if err != nil {
		t.Fatal(err)
	}

	// 4. Set the Content-Type header to indicate form data
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	user, errs := NewUser(req)

	t.Logf("user = %v\n", user)

	if errs.Any() {
		t.Logf("errs = %v\n", errs.String())
	} else {
		t.Log("No Errors")
	}

}
