package data_structure

type StringSet map[string]struct{}

func (s StringSet) Add(stringValue string) {
	s[stringValue] = struct{}{}
}

func (s StringSet) Remove(stringValue string) {
	delete(s, stringValue)
}

func (s StringSet) Has(stringValue string) bool {
	_, ok := s[stringValue]
	return ok
}

func (s StringSet) Length() int {
	return len(s)
}
