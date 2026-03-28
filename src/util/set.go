package util

type StringSet map[string]struct{}

func NewStringSet(vals ...string) StringSet {
	s := StringSet(map[string]struct{}{})
	for _, val := range vals {
		s.Add(val)
	}
	return s
}

func (s StringSet) Add(str string) StringSet {
	s[str] = struct{}{}
	return s
}

func (s StringSet) Merge(other StringSet) {
	for key, val := range other {
		s[key] = val
	}
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
