package src

type Email struct {
	Field         *Field
	ValidatorSpec *ValidatorSpec
}

func NewEmail(f *Field, v *ValidatorSpec) *Email {

	return &Email{
		Field:         f,
		ValidatorSpec: v,
	}
}
