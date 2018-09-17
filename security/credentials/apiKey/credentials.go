package apiKey


//Store the sha256 of the apiKey as the "username"
type Credentials struct {
	ApiKey string `json:"apiKey"`
}
