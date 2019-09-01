package services

// AuthService ...
type AuthService interface {
	GetAccessToken(code string)
}
