package address

type Address struct {
	Name        string `json:"name"`         // John Smith
	CompanyName string `json:"companyName"`  // AA Solutions
	Street      string `json:"street"`       // 9 Black Ave
	Suburb      string `json:"suburb"`       // Rosebank
	City        string `json:"city"`         // Johannesburg
	PostalCode  string `json:"postalCode"`   // 2196
	Country     string `json:"country"`      // South Africa
}
