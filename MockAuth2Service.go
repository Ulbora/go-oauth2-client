package services

import (
	"net/http"
)

//MockAuthCodeAuthorize auth code
type MockAuthCodeAuthorize struct {
	OauthHost   string
	RedirectURI string
	ClientID    string
	Scope       string
	State       string
	OverrideURI string
	Req         *http.Request
	Res         http.ResponseWriter
	MockRtn     bool
}

//GetNew GetNew
func (a *MockAuthCodeAuthorize) GetNew() AuthCodeUser {
	var ac AuthCodeUser
	ac = a
	return ac
}

//AuthCodeAuthorizeUser authorize a user with grant type code
func (a *MockAuthCodeAuthorize) AuthCodeAuthorizeUser() bool {
	return a.MockRtn
}

//MockAuthCodeToken auth code token
type MockAuthCodeToken struct {
	OauthHost    string
	RedirectURI  string
	ClientID     string
	Secret       string
	Code         string
	RefreshToken string
	OverrideURI  string
	MockToken    *Token
}

//GetNew GetNew
func (t *MockAuthCodeToken) GetNew() AuthToken {
	var at AuthToken
	at = t
	return at
}

//AuthCodeToken auth code token
func (t *MockAuthCodeToken) AuthCodeToken() *Token {
	return t.MockToken
}

// AuthCodeRefreshToken get refresh token
func (t *MockAuthCodeToken) AuthCodeRefreshToken() *Token {
	return t.MockToken
}

//MockImplicitAuthorize implicit authorize
type MockImplicitAuthorize struct {
	OauthHost   string
	RedirectURI string
	ClientID    string
	Scope       string
	State       string
	OverrideURI string
	Req         *http.Request
	Res         http.ResponseWriter
	MockRtn     bool
}

//GetNew GetNew
func (i *MockImplicitAuthorize) GetNew() Implicit {
	var it Implicit
	it = i
	return it
}

//ImplicitAuthorize implicit authorize
func (i *MockImplicitAuthorize) ImplicitAuthorize() bool {
	return i.MockRtn
}

// MockClientCredentialsToken client credentials token
type MockClientCredentialsToken struct {
	OauthHost   string
	ClientID    string
	Secret      string
	OverrideURI string
	MockToken   *Token
}

//GetNew GetNew
func (c *MockClientCredentialsToken) GetNew() Credentials {
	var cc Credentials
	cc = c
	return cc
}

//ClientCredentialsToken get client credentials token
func (c *MockClientCredentialsToken) ClientCredentialsToken() *Token {
	return c.MockToken
}
