package user

import (
	"github.com/go-errors/errors"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/skyframe/pegasus/exceptions"
	"github.com/skyframe/pegasus/log"
)

type MongoRecordkeeper struct {
	UserCollection *mongo.Collection
}

type Fields struct {
	Id        string
	FirstName string
	LastName  string
	Password  string
	Email     string
	Username  string
}

func createFields(u User) Fields {
	return Fields{
		Id:        u.Id,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Password:  u.Password,
		Email:     u.Email,
		Username:  u.Username,
	}
}

func createFilter(fields Fields) *bson.Document {
	var elems []*bson.Element

	elems = append(elems, bson.EC.String("id", fields.Id))
	elems = append(elems, bson.EC.String("firstname", fields.FirstName))
	elems = append(elems, bson.EC.String("lastname", fields.LastName))
	elems = append(elems, bson.EC.String("password", fields.Password))
	elems = append(elems, bson.EC.String("email", fields.Email))
	elems = append(elems, bson.EC.String("username", fields.Username))

	return bson.NewDocument()
}

func (r *MongoRecordkeeper) Create(request CreateRequest) (*CreateResponse, error) {
	_, err := r.UserCollection.InsertOne(nil, request.User, nil)
	if err != nil {
		log.Error("Failed to create user: ", err)
		return nil, errors.New("could not create the user")
	}
	return &CreateResponse{request.User}, nil
}

func (r *MongoRecordkeeper) Retrieve(request RetrieveRequest) (*RetrieveResponse, error) {
	filter := bson.NewDocument(bson.EC.String("id", request.Id))

	u := User{}
	err := r.UserCollection.FindOne(nil, filter, nil).Decode(&u)
	if err != nil {
		log.Error("Failed to retrieve user: ", err)
		return nil, errors.New("could not retrieve the user")
	}

	return &RetrieveResponse{u}, nil
}

func (r *MongoRecordkeeper) Find(request FindRequest) (*FindResponse, error) {
	fields := createFields(request.Query)
	filter := createFilter(fields)

	curr, err := r.UserCollection.Find(nil, filter, nil)
	if err != nil {
		log.Error("Failed to retrieve user: ", err)
		return nil, errors.New("could not retrieve the user")
	}

	return &FindResponse{u}, nil
}

func (r *MongoRecordkeeper) Update(request UpdateRequest) (*UpdateResponse, error) {
	return nil, exceptions.NotImplemented{}
}

func (r *MongoRecordkeeper) Delete(request DeleteRequest) (*DeleteResponse, error) {
	return nil, exceptions.NotImplemented{}
}
