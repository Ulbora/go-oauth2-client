package services

import (
	"encoding/json"
	"fmt"

	"net/http"
)

//AuthCodeUser AuthCodeUser
type AuthCodeUser interface {
	AuthCodeAuthorizeUser() bool
	SetOauthHost(host string)
	SetRedirectURI(uri string)
	SetClientID(id string)
	SetScope(scope string)
	SetState(state string)
	SetOverrideURI(uri string)
	SetReq(r *http.Request)
	SetRes(w http.ResponseWriter)
}

//AuthCodeAuthorize auth code
type AuthCodeAuthorize struct {
	OauthHost   string
	RedirectURI string
	ClientID    string
	Scope       string
	State       string
	OverrideURI string
	Req         *http.Request
	Res         http.ResponseWriter
}

//SetOauthHost SetOauthHost
func (a *AuthCodeAuthorize) SetOauthHost(host string) {
	a.OauthHost = host
}

//SetRedirectURI SetRedirectURI
func (a *AuthCodeAuthorize) SetRedirectURI(uri string) {
	a.RedirectURI = uri
}

//SetClientID SetClientID
func (a *AuthCodeAuthorize) SetClientID(id string) {
	a.ClientID = id
}

//SetScope SetScope
func (a *AuthCodeAuthorize) SetScope(scope string) {
	a.Scope = scope
}

//SetState SetState
func (a *AuthCodeAuthorize) SetState(state string) {
	a.State = state
}

//SetOverrideURI SetOverrideURI
func (a *AuthCodeAuthorize) SetOverrideURI(uri string) {
	a.OverrideURI = uri
}

//SetReq SetReq
func (a *AuthCodeAuthorize) SetReq(r *http.Request) {
	a.Req = r
}

//SetRes SetRes
func (a *AuthCodeAuthorize) SetRes(w http.ResponseWriter) {
	a.Res = w
}

//GetNew GetNew
func (a *AuthCodeAuthorize) GetNew() AuthCodeUser {
	var ac AuthCodeUser
	ac = a
	return ac
}

//AuthCodeAuthorizeUser authorize a user with grant type code
func (a *AuthCodeAuthorize) AuthCodeAuthorizeUser() bool {
	var rtn = false
	var uri string
	if a.OverrideURI != "" {
		uri = a.OverrideURI
		rtn = true
	} else {
		uri = a.OauthHost + oauthAuthCodeAuthorizeURI1 + a.ClientID + oauthAuthCodeAuthorizeURI2 +
			a.RedirectURI + oauthAuthCodeAuthorizeURI3 + a.Scope + oauthAuthCodeAuthorizeURI4 + a.State
		rtn = true
	}
	http.Redirect(a.Res, a.Req, uri, http.StatusFound)
	return rtn
}

//AuthToken AuthCodeToken
type AuthToken interface {
	AuthCodeToken() *Token
	AuthCodeRefreshToken() *Token
	SetOauthHost(host string)
	SetRedirectURI(uri string)
	SetClientID(id string)
	SetSecret(sec string)
	SetCode(code string)
	SetRefreshToken(tkn string)
	SetOverrideURI(uri string)
}

//AuthCodeToken auth code token
type AuthCodeToken struct {
	OauthHost    string
	RedirectURI  string
	ClientID     string
	Secret       string
	Code         string
	RefreshToken string
	OverrideURI  string
}

//SetOauthHost SetOauthHost
func (t *AuthCodeToken) SetOauthHost(host string) {
	t.OauthHost = host
}

//SetRedirectURI SetRedirectURI
func (t *AuthCodeToken) SetRedirectURI(uri string) {
	t.RedirectURI = uri
}

//SetClientID SetClientID
func (t *AuthCodeToken) SetClientID(id string) {
	t.ClientID = id
}

//SetSecret SetSecret
func (t *AuthCodeToken) SetSecret(sec string) {
	t.Secret = sec
}

//SetCode SetCode
func (t *AuthCodeToken) SetCode(code string) {
	t.Code = code
}

//SetRefreshToken SetRefreshToken
func (t *AuthCodeToken) SetRefreshToken(tkn string) {
	t.RefreshToken = tkn
}

//SetOverrideURI SetOverrideURI
func (t *AuthCodeToken) SetOverrideURI(uri string) {
	t.OverrideURI = uri
}

//Token the access token
type Token struct {
	AccessToken   string `json:"access_token"`
	RefreshToken  string `json:"refresh_token"`
	TokenType     string `json:"token_type"`
	ExpiresIn     int    `json:"expires_in"`
	ErrorReturned string `json:"error"`
}

//GetNew GetNew
func (t *AuthCodeToken) GetNew() AuthToken {
	var at AuthToken
	at = t
	return at
}

//AuthCodeToken auth code token
func (t *AuthCodeToken) AuthCodeToken() *Token {
	var uri string
	if t.OverrideURI != "" {
		uri = t.OverrideURI
	} else {
		uri = t.OauthHost + oauthAuthCodeTokenURI1 + t.ClientID + oauthAuthCodeTokenURI2 + t.Secret +
			oauthAuthCodeTokenURI3 + t.Code + oauthAuthCodeTokenURI4 + t.RedirectURI
	}
	fmt.Print("AuthCode Token URI: ")
	fmt.Println(uri)
	resp, err := http.Post(uri, "", nil)
	if err != nil {
		fmt.Print("Token Post Error: ")
		fmt.Println(err)
		fmt.Print("Token Post Error Resp: ")
		fmt.Println(resp)
	}
	defer resp.Body.Close()
	rtn := new(Token)
	if resp.StatusCode == 200 || resp.StatusCode == 401 {
		decoder := json.NewDecoder(resp.Body)
		error := decoder.Decode(&rtn)
		if error != nil {
			fmt.Print("Token decode Error: ")
			fmt.Println(error)
		}
	}
	return rtn
}

// AuthCodeRefreshToken get refresh token
func (t *AuthCodeToken) AuthCodeRefreshToken() *Token {
	var uri string
	if t.OverrideURI != "" {
		uri = t.OverrideURI
	} else {
		uri = t.OauthHost + oauthAuthCodeRefreshTokenURI1 + t.ClientID + oauthAuthCodeRefreshTokenURI2 + t.Secret +
			oauthAuthCodeRefreshTokenURI3 + t.RefreshToken
	}
	//fmt.Print("AuthCode RefreshToken URI: ")
	//fmt.Println(uri)
	resp, err := http.Post(uri, "", nil)
	if err != nil {
		fmt.Print("Auth code refresh Post Error: ")
		fmt.Println(err)
	}
	defer resp.Body.Close()
	rtn := new(Token)
	if resp.StatusCode == 200 || resp.StatusCode == 401 {
		decoder := json.NewDecoder(resp.Body)
		error := decoder.Decode(&rtn)
		if error != nil {
			fmt.Print("Token decode Error: ")
			fmt.Println(error)
		}
	}
	return rtn
}

//Implicit Implicit
type Implicit interface {
	ImplicitAuthorize() bool
	SetOauthHost(host string)
	SetRedirectURI(uri string)
	SetClientID(id string)
	SetScope(scope string)
	SetState(state string)
	SetOverrideURI(uri string)
	SetReq(r *http.Request)
	SetRes(w http.ResponseWriter)
}

//ImplicitAuthorize implicit authorize
type ImplicitAuthorize struct {
	OauthHost   string
	RedirectURI string
	ClientID    string
	Scope       string
	State       string
	OverrideURI string
	Req         *http.Request
	Res         http.ResponseWriter
}

//SetOauthHost SetOauthHost
func (i *ImplicitAuthorize) SetOauthHost(host string) {
	i.OauthHost = host
}

//SetRedirectURI SetRedirectURI
func (i *ImplicitAuthorize) SetRedirectURI(uri string) {
	i.RedirectURI = uri
}

//SetClientID SetClientID
func (i *ImplicitAuthorize) SetClientID(id string) {
	i.ClientID = id
}

//SetScope SetScope
func (i *ImplicitAuthorize) SetScope(scope string) {
	i.Scope = scope
}

//SetState SetState
func (i *ImplicitAuthorize) SetState(state string) {
	i.State = state
}

//SetOverrideURI SetOverrideURI
func (i *ImplicitAuthorize) SetOverrideURI(uri string) {
	i.OverrideURI = uri
}

//SetReq SetReq
func (i *ImplicitAuthorize) SetReq(r *http.Request) {
	i.Req = r
}

//SetRes SetRes
func (i *ImplicitAuthorize) SetRes(w http.ResponseWriter) {
	i.Res = w
}

//GetNew GetNew
func (i *ImplicitAuthorize) GetNew() Implicit {
	var it Implicit
	it = i
	return it
}

//ImplicitAuthorize implicit authorize
func (i *ImplicitAuthorize) ImplicitAuthorize() bool {
	var rtn = false
	var uri string
	if i.OverrideURI != "" {
		uri = i.OverrideURI
	} else {
		uri = i.OauthHost + implicitAuthorizeURI1 + i.ClientID + implicitAuthorizeURI2 +
			i.RedirectURI + implicitAuthorizeURI3 + i.Scope + implicitAuthorizeURI4 + i.State
		rtn = true
	}
	http.Redirect(i.Res, i.Req, uri, http.StatusFound)
	return rtn
}

//Credentials Credentials
type Credentials interface {
	ClientCredentialsToken() *Token
	SetOauthHost(host string)
	SetClientID(id string)
	SetSecret(sec string)
	SetOverrideURI(uri string)
}

// ClientCredentialsToken client credentials token
type ClientCredentialsToken struct {
	OauthHost   string
	ClientID    string
	Secret      string
	OverrideURI string
}

//SetOauthHost SetOauthHost
func (c *ClientCredentialsToken) SetOauthHost(host string) {
	c.OauthHost = host
}

//SetClientID SetClientID
func (c *ClientCredentialsToken) SetClientID(id string) {
	c.ClientID = id
}

//SetSecret SetSecret
func (c *ClientCredentialsToken) SetSecret(sec string) {
	c.Secret = sec
}

//SetOverrideURI SetOverrideURI
func (c *ClientCredentialsToken) SetOverrideURI(uri string) {
	c.OverrideURI = uri
}

//GetNew GetNew
func (c *ClientCredentialsToken) GetNew() Credentials {
	var cc Credentials
	cc = c
	return cc
}

//ClientCredentialsToken get client credentials token
func (c *ClientCredentialsToken) ClientCredentialsToken() *Token {
	var uri string
	if c.OverrideURI != "" {
		uri = c.OverrideURI
	} else {
		uri = c.OauthHost + clientCredentialsURI1 + c.ClientID + clientCredentialsURI2 + c.Secret +
			clientCredentialsURI3
	}
	//fmt.Print("Client Credentials Token URI: ")
	//fmt.Println(uri)
	resp, err := http.Post(uri, "", nil)
	if err != nil {
		fmt.Print("Client cred Post Error: ")
		fmt.Println(err)
	}
	defer resp.Body.Close()
	rtn := new(Token)
	if resp.StatusCode == 200 || resp.StatusCode == 401 {
		decoder := json.NewDecoder(resp.Body)
		error := decoder.Decode(&rtn)
		if error != nil {
			fmt.Print("Token decode Error: ")
			fmt.Println(error)
		}
	}
	return rtn
}

//go mod init github.com/Ulbora/go-oauth2-client
