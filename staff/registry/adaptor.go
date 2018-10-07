package registry

import (
	"net/http"
)

type Adaptor struct {
	Recordkeeper Recordkeeper
}

func (a *Adaptor) Create(r *http.Request, request *CreateRequest, response *CreateResponse) error {
	createResponse, err := a.Recordkeeper.Create(*request)
	response.Registry = createResponse.Registry
	return err
}

func (a *Adaptor) Retrieve(r *http.Request, request *RetrieveRequest, response *RetrieveResponse) error {
	_, err := a.Recordkeeper.Retrieve(*request)
	return err
}

func (a *Adaptor) Find(r *http.Request, request *FindRequest, response *FindResponse) error {
	_, err := a.Recordkeeper.Find(*request)
	return err
}

func (a *Adaptor) Update(r *http.Request, request *UpdateRequest, response *UpdateResponse) error {
	_, err := a.Recordkeeper.Update(*request)
	return err
}

func (a *Adaptor) Delete(r *http.Request, request *DeleteRequest, response *DeleteRequest) error {
	_, err := a.Recordkeeper.Delete(*request)
	return err
}
