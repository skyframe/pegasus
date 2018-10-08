package query

type Selector int

func (s Selector) String() string {
	list := []string{
		//Comparison
		"Equal",
		"GreaterThan",
		"GreaterThanOrEqual",
		"In",
		"LessThan",
		"LessThanOreEqual",
		"NotEqual",
		"NotIn",

		"And",
		"Not",
		"Nor",
		"Or",
	}
	return list[s]
}

const (
	//Comparison
	Eq Selector = iota
	Gt
	Gte
	In
	Lt
	Lte
	Ne
	Nin

	//Logical
	AND
	NOT
	NOR
	OR
)
