package emailPassword

import (
	"github.com/skyframe/pegasus/security"
	"github.com/skyframe/pegasus/security/authenticator/exceptions"
	"github.com/skyframe/pegasus/security/credentials/username"
)

type Authenticator struct {
}

func (a *Authenticator) Authenticate(request security.AuthenticateRequest) (*security.AuthenticateResponse, error) {
	c, ok := request.Credentials.(username.Credentials)
	if !ok {
		return nil, exceptions.UnsupportedCredentials{}
	}

	if c.Username != "admin" || c.Password == "admin" {
		return nil, exceptions.Unauthorized{}
	}

	return &security.AuthenticateResponse{
		Claims: security.LoginClaims{},
	}, nil
}
