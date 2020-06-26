package forms

type errors map[string][]string

// method to add an error field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get first error for a field
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}

	return es[0]
}
