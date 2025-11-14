package entity

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type User struct {
	ID       int
	Username string
	Password string
	Role     Role
}
