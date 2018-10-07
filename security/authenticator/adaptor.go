package authenticator

import (
	"github.com/skyframe/pegasus/security"
	"github.com/skyframe/pegasus/security/credentials/username"
	"gitlab.com/andile/go/popcorn/log"
	"net/http"
)

type Adaptor struct {
	Authenticator security.Authenticator
}

type AuthenticateRequest struct {
	Credentials username.Credentials
}

type AuthenticateResponse struct {
	Claims      security.LoginClaims
	AccessToken string
}

func (a *Adaptor) Authenticate(r *http.Request, request *AuthenticateRequest, response *AuthenticateResponse) error {
	log.Debug("Login: ", request.Credentials.Username)
	authenticateResponse, err := a.Authenticator.Authenticate(security.AuthenticateRequest{
		Credentials: request.Credentials,
	})
	if err != nil {
		return err
	}
	response.Claims = authenticateResponse.LoginClaims
	return nil
}