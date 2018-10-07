package user

import "github.com/skyframe/pegasus/search"

type Recordkeeper interface {
	Create(request CreateRequest) (*CreateResponse, error)
	Retrieve(request RetrieveRequest) (*RetrieveResponse, error)
	Find(request FindRequest) (*FindResponse, error)
	Update(request UpdateRequest) (*UpdateResponse, error)
	Delete(request DeleteRequest) (*DeleteResponse, error)
}

type CreateResponse struct {
	User User `json:"user"`
}
type CreateRequest struct {
	User User `json:"user"`
}

type RetrieveResponse struct {
	User User `json:"user"`
}
type RetrieveRequest struct {
	Id string `json:"id"`
}

type FindRequest struct {
	Query search.Query
	Criteria search.Criteria
}
type FindResponse struct {
	Users []User `json:"users"`
	Total int    `json:"total"`
}

type UpdateResponse struct {
	User User `json:"user"`
}
type UpdateRequest struct {
	User User `json:"user"`
	Id   User `json:"id"`
}

type DeleteResponse struct {
}
type DeleteRequest struct {
	Id User `json:"id"`
}
