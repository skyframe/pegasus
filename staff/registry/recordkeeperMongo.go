package registry

import (
	"github.com/skyframe/pegasus/address"
	"github.com/skyframe/pegasus/construction"
	"github.com/skyframe/pegasus/employee"
	"github.com/skyframe/pegasus/exceptions"
	"time"
)

type MongoRecordkeeper struct {
}

func (r *MongoRecordkeeper) Create(request CreateRequest) (*CreateResponse, error) {
	return &CreateResponse{
		Registry: Registry{
			Date: time.Now(),
			Site: construction.Site{
				Name: "Tyrwit",
				Address: address.Address{
					City: "Rosebank",
				},
			},
			Entries: []Entry{
				{
					Status: Awol,
					Employee: employee.Employee{
						Name: "John Smith",
					},
				},
			},
		},
	}, nil
}

func (r *MongoRecordkeeper) Retrieve(request RetrieveRequest) (*RetrieveResponse, error) {
	return nil, exceptions.NotImplemented{}
}

func (r *MongoRecordkeeper) Find(request FindRequest) (*FindResponse, error) {
	return nil, exceptions.NotImplemented{}
}

func (r *MongoRecordkeeper) Update(request UpdateRequest) (*UpdateResponse, error) {
	return nil, exceptions.NotImplemented{}
}

func (r *MongoRecordkeeper) Delete(request DeleteRequest) (*DeleteResponse, error) {
	return nil, exceptions.NotImplemented{}
}
