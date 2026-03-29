package src

type MinLen struct {
	Field         *Field
	ValidatorSpec *ValidatorSpec
	Value         int
}

func NewMinLen(f *Field, v *ValidatorSpec) *MinLen {

	val, ok := v.Value.(int)

	if !ok {
		panic("minLen value must be an int")
	}

	return &MinLen{
		Field:         f,
		ValidatorSpec: v,
		Value:         val,
	}
}
