package src

type SimpleValidator struct {
	Field         *Field
	ValidatorSpec *ValidatorSpec
}

func NewSimpleValidator(f *Field, v *ValidatorSpec) *SimpleValidator {

	return &SimpleValidator{
		Field:         f,
		ValidatorSpec: v,
	}
}
