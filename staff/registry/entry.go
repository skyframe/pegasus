package registry

import "github.com/skyframe/pegasus/employee"

type Entry struct {
	Employee employee.Employee `json:"employee"`
	Status   Status            `json:"status"`
	Note     string            `json:"note"`
}
