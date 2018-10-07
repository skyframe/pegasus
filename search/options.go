package search

import "strings"

type Options struct {
	SortBy []string `json:"sortby"`
	Order  []string `json:"order"`
	Limit  int      `json:"limit"`
	Offset int      `json:"offset"`
}

func (q Options) SortByMongoFormat() []string {
	var sortOrder []string
	for i, field := range q.SortBy {
		if len(q.Order) > i && strings.ToLower(q.Order[i]) == "desc" {
			sortOrder = append(sortOrder, "-"+field)
		} else {
			sortOrder = append(sortOrder, field)
		}
	}
	return sortOrder
}