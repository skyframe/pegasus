package registry

import (
	"github.com/skyframe/pegasus/construction"
	"time"
)

type Registry struct {
	Date    time.Time `json:"date"`
	Site    construction.Site `json:"site"`
	Entries []Entry `json:"entries"`
}
