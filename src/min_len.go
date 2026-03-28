package src

import (
	"github.com/jmarren/gatekeeper/src/util"
)

type MinLen struct {
	field         *Field
	ValidatorSpec *ValidatorSpec
	value         int
}

func (m *MinLen) imports() util.StringSet {
	return util.NewStringSet()
}

func NewMinLen(f *Field, v *ValidatorSpec) *MinLen {

	val, ok := v.Value.(int)

	if !ok {
		panic("minLen value must be an int")
	}

	return &MinLen{
		field:         f,
		ValidatorSpec: v,
		value:         val,
	}
}
