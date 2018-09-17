package security

type AuthenticateRequest struct {
	Credentials Credentials
}

type AuthenticateResponse struct {
	Claims Claims
}

type Authenticator interface {
	Authenticate(AuthenticateRequest) (*AuthenticateResponse, error)
}

type Claims interface{}

type LoginClaims struct {
}
