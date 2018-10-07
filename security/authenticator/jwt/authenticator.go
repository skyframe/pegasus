package jwt

import "github.com/skyframe/pegasus/security"

type Authenticator struct {
}

func (a *Authenticator) Authenticate(request security.AuthenticateRequest) (*security.AuthenticateResponse, error) {
	return nil, nil
}
