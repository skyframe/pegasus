package pegasus

type Recordkeeper interface {
	Create(request CreateRequest) (*CreateResponse, error)
	Retrieve(request RetrieveRequest) (*RetrieveResponse, error)
	Find(request FindRequest) (*FindResponse, error)
	Update(request UpdateRequest) (*UpdateResponse, error)
	Delete(request DeleteRequest) (*DeleteResponse, error)
}

type CreateResponse struct {
}
type CreateRequest struct {
}

type RetrieveResponse struct {
}
type RetrieveRequest struct {
}

type FindRequest struct {
}
type FindResponse struct {
}

type UpdateResponse struct {
}
type UpdateRequest struct {
}

type DeleteResponse struct {
}
type DeleteRequest struct {
}
