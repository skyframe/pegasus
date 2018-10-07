package exceptions

type NotImplemented struct {
	Method string
}

func (e NotImplemented) Error() string {
	if e.Method == "" {
		return "not implemented"
	}
	return e.Method + " not implemented"
}
