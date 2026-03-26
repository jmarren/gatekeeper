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
	form.Add("FirstName", "xy")
	form.Add("Age", "22")

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
	t.Logf("errs = %v\n", errs.String())

}
