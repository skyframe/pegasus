package employee

import "github.com/skyframe/pegasus/address"

type Employee struct {
	Name    string          `json:"name"`
	Address address.Address `json:"address"`
}
