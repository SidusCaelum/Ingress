package rest

// User - struct for user data exchange between client and server
type User struct {
	Email    string `json:"Email" binding:"required"`
	Username string `json:"Username" binding:"required"`
}

// UserCheck - struct for checking user submission content
type UserCheck struct {
	Empty    bool `json:"Empty"`
	Username bool `json:"Username"`
	Email    bool `json:"Email"`
}
