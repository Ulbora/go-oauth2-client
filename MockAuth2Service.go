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

//SetOauthHost SetOauthHost
func (a *MockAuthCodeAuthorize) SetOauthHost(host string) {
	a.OauthHost = host
}

//SetRedirectURI SetRedirectURI
func (a *MockAuthCodeAuthorize) SetRedirectURI(uri string) {
	a.RedirectURI = uri
}

//SetClientID SetClientID
func (a *MockAuthCodeAuthorize) SetClientID(id string) {
	a.ClientID = id
}

//SetScope SetScope
func (a *MockAuthCodeAuthorize) SetScope(scope string) {
	a.Scope = scope
}

//SetState SetState
func (a *MockAuthCodeAuthorize) SetState(state string) {
	a.State = state
}

//SetOverrideURI SetOverrideURI
func (a *MockAuthCodeAuthorize) SetOverrideURI(uri string) {
	a.OverrideURI = uri
}

//SetReq SetReq
func (a *MockAuthCodeAuthorize) SetReq(r *http.Request) {
	a.Req = r
}

//SetRes SetRes
func (a *MockAuthCodeAuthorize) SetRes(w http.ResponseWriter) {
	a.Res = w
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

//SetOauthHost SetOauthHost
func (t *MockAuthCodeToken) SetOauthHost(host string) {
	t.OauthHost = host
}

//SetRedirectURI SetRedirectURI
func (t *MockAuthCodeToken) SetRedirectURI(uri string) {
	t.RedirectURI = uri
}

//SetClientID SetClientID
func (t *MockAuthCodeToken) SetClientID(id string) {
	t.ClientID = id
}

//SetSecret SetSecret
func (t *MockAuthCodeToken) SetSecret(sec string) {
	t.Secret = sec
}

//SetCode SetCode
func (t *MockAuthCodeToken) SetCode(code string) {
	t.Code = code
}

//SetRefreshToken SetRefreshToken
func (t *MockAuthCodeToken) SetRefreshToken(tkn string) {
	t.RefreshToken = tkn
}

//SetOverrideURI SetOverrideURI
func (t *MockAuthCodeToken) SetOverrideURI(uri string) {
	t.OverrideURI = uri
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

//SetOauthHost SetOauthHost
func (i *MockImplicitAuthorize) SetOauthHost(host string) {
	i.OauthHost = host
}

//SetRedirectURI SetRedirectURI
func (i *MockImplicitAuthorize) SetRedirectURI(uri string) {
	i.RedirectURI = uri
}

//SetClientID SetClientID
func (i *MockImplicitAuthorize) SetClientID(id string) {
	i.ClientID = id
}

//SetScope SetScope
func (i *MockImplicitAuthorize) SetScope(scope string) {
	i.Scope = scope
}

//SetState SetState
func (i *MockImplicitAuthorize) SetState(state string) {
	i.State = state
}

//SetOverrideURI SetOverrideURI
func (i *MockImplicitAuthorize) SetOverrideURI(uri string) {
	i.OverrideURI = uri
}

//SetReq SetReq
func (i *MockImplicitAuthorize) SetReq(r *http.Request) {
	i.Req = r
}

//SetRes SetRes
func (i *MockImplicitAuthorize) SetRes(w http.ResponseWriter) {
	i.Res = w
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

//SetOauthHost SetOauthHost
func (c *MockClientCredentialsToken) SetOauthHost(host string) {
	c.OauthHost = host
}

//SetClientID SetClientID
func (c *MockClientCredentialsToken) SetClientID(id string) {
	c.ClientID = id
}

//SetSecret SetSecret
func (c *MockClientCredentialsToken) SetSecret(sec string) {
	c.Secret = sec
}

//SetOverrideURI SetOverrideURI
func (c *MockClientCredentialsToken) SetOverrideURI(uri string) {
	c.OverrideURI = uri
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
