package netatgo

import (
	"context"
	"net/http"
	"net/url"

	"golang.org/x/oauth2/clientcredentials"
)

type (
	Client struct {
		HTTPClient *http.Client
	}

	ClientCredentialsConfig struct {
		// ClientID is the application's ID.
		ClientID string

		// ClientSecret is the application's secret.
		ClientSecret string

		Username string

		Password string

		// Scope specifies optional requested permissions.
		Scopes []Scope
	}
)

const (
	tokenURL = "https://api.netatmo.com/oauth2/token"
)

func NewClientCredentialsClient(cfg ClientCredentialsConfig) Client {
	cc := clientcredentials.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		EndpointParams: url.Values{
			"username":   []string{cfg.Username},
			"password":   []string{cfg.Password},
			"grant_type": []string{"password"},
		},
		Scopes:   cfg.Scopes,
		TokenURL: tokenURL,
	}

	return Client{
		HTTPClient: cc.Client(context.Background()),
	}
}
