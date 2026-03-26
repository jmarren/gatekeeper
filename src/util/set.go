package util

type StringSet map[string]struct{}

func NewStringSet() StringSet {
	s := StringSet(map[string]struct{}{})
	return s
}

func (s StringSet) Add(str string) StringSet {
	s[str] = struct{}{}
	return s
}

func (s StringSet) Has(str string) bool {
	_, found := s[str]
	return found
}

func (s StringSet) Delete(str string) {
	delete(s, str)
}

func (s StringSet) ToSlice() []string {
	out := []string{}
	for str := range s {
		out = append(out, str)
	}
	return out
}
