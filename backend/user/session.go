package user

type AuthAttempt struct {
	Email    string
	Password string
}

type NewUserAttempt struct {
	Email   string
	Password string
	Name     string
}