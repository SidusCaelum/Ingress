package rest

// User - struct for user data exchange between client and server
type User struct {
	Email    string `json:"Email" binding:"required"`
	Username string `json:"Username" binding:"required"`
}
