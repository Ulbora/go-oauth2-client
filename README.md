go-oauth2-client
==============

Mockable Go OAuth2 client for access a OAuth2 Server from a golang application

# Installation

```
$ go get github.com/Ulbora/go-oauth2-client

```

# Usage

## Auth Code Grant Type

### Auth Code Authorize
```
    var a AuthCodeAuthorize
    a.ClientID = "211"
    a.OauthHost = "http://localhost:3000"
    a.RedirectURI = "http:/localhost/token"
    a.Scope = "write"
    a.State = "12345"
    res := a.AuthCodeAuthorizeUser()
```

### Auth Code Token
```
    var tn AuthCodeToken
    tn.OauthHost = "http://localhost:3000"
    tn.ClientID = "403"
    tn.Secret = "554444vfg55ggfff22454sw2fff2dsfd"
    tn.Code = "yfgk5mj481QSl46n2zIZGl"
    tn.RedirectURI = "http://www.google.com"
    token := tn.AuthCodeToken()

```

### Auth Code Refresh

```
    var tn AuthCodeToken
    tn.OauthHost = "http://localhost:3000"
    tn.ClientID = "403"
    tn.Secret = "554444vfg55ggfff22454sw2fff2dsfd"
    tn.RefreshToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJjb2RlIiwidXNlcklkIjoia2VuIiwiY2xpZW50SWQiOjQwMywiaWF0IjoxNTAyNDE4NDQ1LCJ0b2tlblR5cGUiOiJyZWZyZXNoIiwiZXhwIjoxNTAyNDU0NDQ1LCJpc3MiOiJVbGJvcmEgT2F1dGgyIFNlcnZlciJ9.7rJPyXkVppTS_4_b3K8nUdnnrjmZI0R69_F7ii5_ueA"
    token := tn.AuthCodeRefreshToken()

```
## Implicit Grant Type

### Implicit Authorize

```
    var a ImplicitAuthorize
    a.ClientID = "403"
    a.OauthHost = "http://localhost:3000"
    a.RedirectURI = "http://www.google.com"
    a.Scope = "read"
    a.State = "12345"
    res := a.ImplicitAuthorize()

```
## Client Credentials Grant Type

### Client Credentials Token

```
    var tn ClientCredentialsToken
    tn.OauthHost = "http://localhost:3000"
    tn.ClientID = "403"
    tn.Secret = "554444vfg55ggfff22454sw2fff2dsfd"
    token := tn.ClientCredentialsToken()

```

## Password Grant Type
Not supported

## Returned Token for grant types Auth Code, Refresh, and Client Credentials

```
type Token struct {
    AccessToken   string `json:"access_token"`
    RefreshToken  string `json:"refresh_token"`
    TokenType     string `json:"token_type"`
    ExpiresIn     int    `json:"expires_in"`
    ErrorReturned string `json:"error"`
}
```

## Returned by Implicit
Implicit grand type calls the redirect URI and passes the token as a query parm.

