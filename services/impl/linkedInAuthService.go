package services

import (
	"net/http"
)

// LinkedInAuthService ...
type LinkedInAuthService struct {
	Client *http.Client
}

// GetAccessToken ...
func (linkedInAuthService *LinkedInAuthService) GetAccessToken(code string) {
	response, err := linkedInAuthService.Client.Get()
}
