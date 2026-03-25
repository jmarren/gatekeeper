package example

import (
	"testing"
)

func TestSet(t *testing.T) {

	vErr := newValidationErr("age", 22, 10, "%(field) must be >= %(expected)")

	msg := vErr.Error()

	t.Logf("msg = %s\n", msg)

	// set := ()
	//
	// set.Add("hi")
	//
	// if !set.Has("hi") {
	// 	t.Errorf("set.Has(\"hi\") returned false after adding \"hi\"")
	// }
	// if set.Has("bye") {
	// 	t.Errorf("set.Has(\"bye\") returned true while not in set")
	// }
	//
	// set.Delete("hi")
	//
	// if set.Has("hi") {
	// 	t.Errorf("set.Has(\"hi\") returned true after deleting \"hi\"")
	// }
	//
}
