package security

type AuthenticateRequest struct {
	Credentials Credentials
}

type AuthenticateResponse struct {
	LoginClaims LoginClaims
}

type Authenticator interface {
	Authenticate(AuthenticateRequest) (*AuthenticateResponse, error)
}

type LoginClaims struct {
	Username string
	Roles []Role
}
