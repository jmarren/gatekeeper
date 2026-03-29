package src

import (
	"github.com/jmarren/gatekeeper/src/util"
)

type Option struct {
	ValidatorSpec *ValidatorSpec
	Field         *Field
	Value         []string
}

func NewOption(f *Field, v *ValidatorSpec) *Option {

	f.Obj.imports.Add(SLICES)

	val := util.AnyToStrSlice(v.Value)

	return &Option{
		ValidatorSpec: v,
		Field:         f,
		Value:         val,
	}
}
