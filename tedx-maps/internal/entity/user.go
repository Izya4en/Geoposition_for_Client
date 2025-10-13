package entity

// User хранит базовую информацию о пользователе
type User struct {
	ID        int64  `db:"id" json:"id"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password,omitempty"`
	Role      string `db:"role" json:"role"` // "admin", "viewer", "operator"
	CreatedAt string `db:"created_at" json:"created_at"`
}
