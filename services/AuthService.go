package services

import "resumator-backend/models"

// AuthService ...
type AuthService interface {
	GetAccessToken(code string) (*models.AuthResponse, error)
}
