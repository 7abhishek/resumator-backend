package services

import (
	"bytes"
	"encoding/json"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/url"
	"resumator-backend/config"
	"resumator-backend/logger"
	"resumator-backend/models"
)

// LinkedInAuthService ...
type LinkedInAuthService struct {
	Client *http.Client
}

var log *zap.SugaredLogger

// GetAccessToken ...
func (linkedInAuthService *LinkedInAuthService) GetAccessToken(code string) (*models.AuthResponse, error) {
	log = logger.GetLogger()
	log.Infof("GetAccessToken called with code %s", code)
	authRequest := &models.AuthRequest{
		GrantType:    "authorization_code",
		Code:         code,
		RedirectURI:  config.GetRedirectURI(),
		ClientID:     config.GetClientID(),
		ClientSecret: config.GetClientSecret(),
	}
	data := url.Values{}
	data.Set("client_id", authRequest.ClientID)
	data.Add("client_secret", authRequest.ClientSecret)
	data.Add("grant_type", authRequest.GrantType)
	data.Add("code", authRequest.Code)
	data.Add("redirect_uri", authRequest.RedirectURI)
	request, err := http.NewRequest("POST", config.GetLinkedInAccessTokenURL(), bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Errorf("unable to create request error ", err)
		return nil, err
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response, err := linkedInAuthService.Client.Do(request)
	if err != nil {
		log.Errorf("error ", err)
		return nil, err
	}
	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Errorf("error during reading response body, %s", err.Error())
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		log.Errorf("error %s", string(bytes))
	}

	var authResponse models.AuthResponse

	err = json.Unmarshal(bytes, &authResponse)
	if err != nil {
		log.Errorf("unmarshal error, %s", err.Error())
		return nil, err
	}
	log.Infof("authResponse %v", authResponse)
	return &authResponse, nil
}
