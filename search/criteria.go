package search

type Criteria struct {
	Filters []Criterion
}

type Criterion struct {
	Field string

}

//  { item: "journal", status: "A", size: { h: 14, w: 21, uom: "cm" }, instock: [ { warehouse: "A", qty: 5 } ] },