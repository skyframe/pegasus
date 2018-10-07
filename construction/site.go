package construction

import "github.com/skyframe/pegasus/address"

type Site struct {
	Name    string          `json:"name"`
	Address address.Address `json:"address"`
}
