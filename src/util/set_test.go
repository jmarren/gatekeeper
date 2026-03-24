package util

import "testing"

func TestSet(t *testing.T) {

	set := NewStringSet()

	set.Add("hi")

	if !set.Has("hi") {
		t.Errorf("set.Has(\"hi\") returned false after adding \"hi\"")
	}
	if set.Has("bye") {
		t.Errorf("set.Has(\"bye\") returned true while not in set")
	}

	set.Delete("hi")

	if set.Has("hi") {
		t.Errorf("set.Has(\"hi\") returned true after deleting \"hi\"")
	}

}
