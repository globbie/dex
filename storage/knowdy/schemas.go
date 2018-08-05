package knowdy

import "github.com/globbie/dex/storage"

type claims struct {
	UserID        string `json:"user-id"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	EmailVerified bool `json:"email-verified"`

	Groups []string `json:"groups"`
}

type authCode struct {
	ID string `json:"id"`

	ClientID string `json:"client-id"`

	RedirectURI string `json:"redirect-uri"`

	Nonce string `json:"nonce"`

	Scopes []string `json:"scopes"`

	ConnectorID   string `json:"connector-id"`
	ConnectorData []byte `json:"connector-data"`
	Claims        claims `json:"claims"`
}

func newAuthCode(a storage.AuthCode) (*authCode, error) {
	ac := &authCode{
		ID: a.ID,
		ClientID: a.ClientID,
		RedirectURI: a.RedirectURI,
		Nonce: a.Nonce,
		Scopes: a.Scopes,
		ConnectorID: a.ConnectorID,
		ConnectorData: a.ConnectorData,
		Claims: claims {
			UserID: a.Claims.UserID,
			Username: a.Claims.Username,
			Email: a.Claims.Email,
			EmailVerified: a.Claims.EmailVerified,
			Groups: a.Claims.Groups,
		},
	}
	return ac, nil
}
