package security

type Permission struct {
	Name string
}

type Role struct {
	Name string
	Permissions []Permission
}
