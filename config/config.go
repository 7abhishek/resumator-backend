package config

import (
	v "github.com/spf13/viper"
	"log"
)

type Config struct {
	LinkedInAccessTokenURL string
	RedirectURI            string
	ClientID               string
	ClientSecret           string
}

var config *Config

func Initialize() {
	viper := v.New()
	viper.AutomaticEnv()
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath("/Users/abhishek/freetime/resumator-backend")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error occurred while reading condig %v", err)
	}
	config = &Config{
		LinkedInAccessTokenURL: viper.GetString("ACCESS_TOKEN_URL"),
		RedirectURI:            viper.GetString("REDIRECT_URI"),
		ClientID:               viper.GetString("CLIENT_ID"),
		ClientSecret:           viper.GetString("CLIENT_SECRET"),
	}
}

// GetLinkedInAccessTokenURL ...
func GetLinkedInAccessTokenURL() string {
	return config.LinkedInAccessTokenURL
}

// GetRedirectURI ...
func GetRedirectURI() string {
	return config.RedirectURI
}

// GetClientID ...
func GetClientID() string {
	return config.ClientID
}

// GetClientSecret ...
func GetClientSecret() string {
	return config.ClientSecret
}
