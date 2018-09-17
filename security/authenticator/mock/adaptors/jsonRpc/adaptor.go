package jsonRpc

import "net/http"

type Adaptor struct {}

type LoginRequest struct {

}

type LoginResponse struct {

}


func (a *Adaptor) Login(r *http.Request, request *LoginRequest, response *LoginResponse) error {

}