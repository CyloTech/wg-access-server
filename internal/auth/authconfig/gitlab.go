package authconfig

import "github.com/place1/wg-access-server/internal/auth/authruntime"

type GitlabConfig struct {
	Name         string `yaml:"name"`
	BaseURL      string `yaml:"baseURL"`
	ClientID     string `yaml:"clientID"`
	ClientSecret string `yaml:"clientSecret"`
	RedirectURL  string `yaml:"redirectURL"`
}

func (c *GitlabConfig) Provider() *authruntime.Provider {
	o := OIDCConfig{
		Name:         c.Name,
		Issuer:       c.BaseURL,
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		RedirectURL:  c.RedirectURL,
		Scopes:       []string{"openid"},
	}
	p := o.Provider()
	p.Type = "Gitlab"
	return p
}
