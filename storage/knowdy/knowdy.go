package knowdy

import (
	"errors"
	"github.com/globbie/dex/storage"
	"github.com/sirupsen/logrus"
	"time"
	"encoding/json"
)

var notImplemented = errors.New("not implemented")

type Knowdy struct {

	authCodes map[string][]byte

	logger logrus.FieldLogger
}

type Config struct {
	Address string `json:"address" yaml:"address"`
}

// Open always returns a new in memory storage.
func (c *Config) Open(logger logrus.FieldLogger) (storage.Storage, error) {
	return New(logger), nil
}

func New(logger logrus.FieldLogger) storage.Storage {
	return &Knowdy{
		authCodes: make(map[string][]byte),
		logger: logger,
	}
}

func (s *Knowdy) Close() error { return notImplemented }

func (s *Knowdy) CreateAuthRequest(a storage.AuthRequest) error         { return notImplemented }
func (s *Knowdy) CreateClient(c storage.Client) error                   { return notImplemented }

func (s *Knowdy) CreateAuthCode(c storage.AuthCode) error {
	if _, ok := s.authCodes[c.ID]; ok {
		return storage.ErrAlreadyExists
	}
	ac, err := newAuthCode(c)
	if err != nil {
		return err
	}
	data, err := json.Marshal(ac)
	if err != nil {
		return err
	}
	s.authCodes[c.ID] = data

	return nil
}

func (s *Knowdy) CreateRefresh(r storage.RefreshToken) error            { return notImplemented }
func (s *Knowdy) CreatePassword(p storage.Password) error               { return notImplemented }
func (s *Knowdy) CreateOfflineSessions(o storage.OfflineSessions) error { return notImplemented }
func (s *Knowdy) CreateConnector(c storage.Connector) error             { return notImplemented }

func (s *Knowdy) GetAuthRequest(id string) (a storage.AuthRequest, err error) {
	return a, notImplemented
}
func (s *Knowdy) GetAuthCode(id string) (a storage.AuthCode, err error)    { return a, notImplemented }
func (s *Knowdy) GetClient(id string) (c storage.Client, err error)        { return c, notImplemented }
func (s *Knowdy) GetKeys() (k storage.Keys, err error)                     { return k, notImplemented }
func (s *Knowdy) GetRefresh(id string) (r storage.RefreshToken, err error) { return r, notImplemented }
func (s *Knowdy) GetPassword(email string) (p storage.Password, err error) { return p, notImplemented }
func (s *Knowdy) GetOfflineSessions(userID string, connID string) (o storage.OfflineSessions, err error) {
	return o, notImplemented
}
func (s *Knowdy) GetConnector(id string) (c storage.Connector, err error) { return c, notImplemented }

func (s *Knowdy) ListClients() (c []storage.Client, err error)             { return c, notImplemented }
func (s *Knowdy) ListRefreshTokens() (r []storage.RefreshToken, err error) { return r, notImplemented }
func (s *Knowdy) ListPasswords() (p []storage.Password, err error)         { return p, notImplemented }
func (s *Knowdy) ListConnectors() (c []storage.Connector, err error)       { return c, notImplemented }

func (s *Knowdy) DeleteAuthRequest(id string) error                        { return notImplemented }
func (s *Knowdy) DeleteAuthCode(code string) error                         { return notImplemented }
func (s *Knowdy) DeleteClient(id string) error                             { return notImplemented }
func (s *Knowdy) DeleteRefresh(id string) error                            { return notImplemented }
func (s *Knowdy) DeletePassword(email string) error                        { return notImplemented }
func (s *Knowdy) DeleteOfflineSessions(userID string, connID string) error { return notImplemented }
func (s *Knowdy) DeleteConnector(id string) error                          { return notImplemented }

func (s *Knowdy) UpdateClient(id string, updater func(old storage.Client) (storage.Client, error)) error {
	return notImplemented
}
func (s *Knowdy) UpdateKeys(updater func(old storage.Keys) (storage.Keys, error)) error {
	return notImplemented
}
func (s *Knowdy) UpdateAuthRequest(id string, updater func(a storage.AuthRequest) (storage.AuthRequest, error)) error {
	return notImplemented
}
func (s *Knowdy) UpdateRefreshToken(id string, updater func(r storage.RefreshToken) (storage.RefreshToken, error)) error {
	return notImplemented
}
func (s *Knowdy) UpdatePassword(email string, updater func(p storage.Password) (storage.Password, error)) error {
	return notImplemented
}
func (s *Knowdy) UpdateOfflineSessions(userID string, connID string, updater func(s storage.OfflineSessions) (storage.OfflineSessions, error)) error {
	return notImplemented
}
func (s *Knowdy) UpdateConnector(id string, updater func(c storage.Connector) (storage.Connector, error)) error {
	return notImplemented
}

func (s *Knowdy) GarbageCollect(now time.Time) (g storage.GCResult, e error) { return g, notImplemented }
