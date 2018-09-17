package exceptions

type UnsupportedCredentials struct{}
func (e UnsupportedCredentials) Error() string {
	return "credentials not supported"
}

type Unauthorized struct{}
func (e Unauthorized) Error() string {
	return "Unauthorized"
}