package services

import (
	"encoding/json"
	"fmt"

	"net/http"
)

//AuthCodeUser AuthCodeUser
type AuthCodeUser interface {
	AuthCodeAuthorizeUser() bool
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
}

// ClientCredentialsToken client credentials token
type ClientCredentialsToken struct {
	OauthHost   string
	ClientID    string
	Secret      string
	OverrideURI string
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
