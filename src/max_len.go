package src

import "github.com/jmarren/gatekeeper/src/util"

type MaxLen struct {
	FieldName string
	FormName  string
	Value     int
	FmtError  string
}

func (m *MaxLen) imports() util.StringSet {
	return util.NewStringSet()
}

func NewMaxLen(f *FieldSpec, v *ValidatorSpec) *MaxLen {

	val, ok := v.Value.(int)

	if !ok {
		panic("maxLen value must be an int")
	}

	return &MaxLen{
		FieldName: f.Name,
		FormName:  f.FormName,
		Value:     val,
		FmtError:  v.FmtErr,
	}
}
