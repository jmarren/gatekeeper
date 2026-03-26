package gkerror

import (
	"testing"
)

func TestErrs(t *testing.T) {

	errGroup := NewValidationErrGroup()

	// vErr := newAgeErr(22)
	vErr := NewValidationErr("age", 22, 10, "%(field) must be >= %(expected). Got %(received)")

	errGroup.Add(vErr)

	msgs := errGroup.Errors()

	t.Logf("msgs = %s\n", msgs)
}
