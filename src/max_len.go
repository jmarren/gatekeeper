package src

type MaxLen struct {
	Field         *Field
	ValidatorSpec *ValidatorSpec
	Value         int
}

func NewMaxLen(f *Field, v *ValidatorSpec) *MaxLen {

	val, ok := v.Value.(int)

	if !ok {
		panic("maxLen value must be an int")
	}

	return &MaxLen{
		Field:         f,
		ValidatorSpec: v,
		Value:         val,
	}
}
