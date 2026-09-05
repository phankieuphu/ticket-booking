package entity

type User struct {
	ID       string
	Username string
	// Email    string
	Password string
}

type UserProfile struct {
	UserID      int
	FirstName   string
	LastName    string
	Address     string
	PhoneNumber string
	Email       string
}

type UserRole struct {
	UserID int
	RoleID int
}
