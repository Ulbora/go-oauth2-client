package services

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthCodeAuthorizem_AuthCodeAuthorizeUser(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com", nil) // new(http.Request)
	if err != nil {
		log.Fatal(err)
	}
	res := httptest.NewRecorder()
	fmt.Print("req: ")
	fmt.Println(req)
	fmt.Print("res: ")
	fmt.Println(res)
	var a MockAuthCodeAuthorize
	a.MockRtn = true
	a.ClientID = "211"
	a.OauthHost = "http://localhost:3000"
	a.RedirectURI = "http:/localhost/token"
	a.Scope = "write"
	a.State = "12345"
	a.Req = req
	a.Res = res
	aa := a.GetNew()
	resp := aa.AuthCodeAuthorizeUser()
	if resp != true {
		t.Fail()
	}

}

func TestAuthCodeTokenm(t *testing.T) {
	var tn MockAuthCodeToken
	var tnt Token
	tnt.ErrorReturned = "invalid_client"
	tn.MockToken = &tnt
	tn.OauthHost = "http://localhost:3000"
	tn.ClientID = "10"
	tn.Secret = "dfgdfgdgd"
	tn.Code = "yfgk5mj481QSl46n2zIZGl"
	tn.RedirectURI = "http://www.google.com"
	tnn := tn.GetNew()
	token := tnn.AuthCodeToken()
	fmt.Print("Returned Token: ")
	fmt.Println(token.AccessToken)

	fmt.Print("Returned Type: ")
	fmt.Println(token.TokenType)

	fmt.Print("Returned ExpiresIn: ")
	fmt.Println(token.ExpiresIn)

	fmt.Print("Returned RefreshToken: ")
	fmt.Println(token.RefreshToken)
	fmt.Print("Returned Error: ")
	fmt.Println(token.ErrorReturned)
	if token.ErrorReturned != "invalid_client" {
		t.Fail()
	}
}

func TestAuthCodeRefreshTokenm(t *testing.T) {
	var tn MockAuthCodeToken
	var tnt Token
	tnt.ErrorReturned = "invalid_client"
	tn.MockToken = &tnt
	tn.OauthHost = "http://localhost:3000"
	tn.ClientID = "10"
	tn.Secret = "554444vfg55ggfff22454sw2fff2dsfd"
	tn.RefreshToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJjb2RlIiwidXNlcklkIjoia2VuIiwiY2xpZW50SWQiOjQwMywiaWF0IjoxNTAyNDE4NDQ1LCJ0b2tlblR5cGUiOiJyZWZyZXNoIiwiZXhwIjoxNTAyNDU0NDQ1LCJpc3MiOiJVbGJvcmEgT2F1dGgyIFNlcnZlciJ9.7rJPyXkVppTS_4_b3K8nUdnnrjmZI0R69_F7ii5_ueA"
	token := tn.AuthCodeRefreshToken()
	fmt.Print("Returned Token: ")
	fmt.Println(token.AccessToken)

	fmt.Print("Returned Type: ")
	fmt.Println(token.TokenType)

	fmt.Print("Returned ExpiresIn: ")
	fmt.Println(token.ExpiresIn)

	fmt.Print("Returned RefreshToken: ")
	fmt.Println(token.RefreshToken)
	fmt.Print("Returned Error: ")
	fmt.Println(token.ErrorReturned)
	if token.ErrorReturned != "invalid_client" {
		t.Fail()
	}
}

func TestImplicitAuthorizem(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com", nil) // new(http.Request)
	if err != nil {
		log.Fatal(err)
	}
	res := httptest.NewRecorder()
	var a MockImplicitAuthorize
	a.MockRtn = true
	a.ClientID = "10"
	a.OauthHost = "http://localhost:3000"
	a.RedirectURI = "http://www.google.com"
	a.Scope = "read"
	a.State = "12345"
	a.Req = req
	a.Res = res
	aa := a.GetNew()
	resp := aa.ImplicitAuthorize()
	if resp != true {
		t.Fail()
	}
}

func TestClientCredentialsTokenm(t *testing.T) {
	var tn MockClientCredentialsToken
	var tnt Token
	//tnt.ErrorReturned = "invalid_client"
	tnt.TokenType = "bearer"
	tn.MockToken = &tnt
	tn.OauthHost = "http://localhost:3000"
	tn.ClientID = "10"
	tn.Secret = "jhcy2YGrvgDsm4VRVtUESiI96K65gQeXcA2TQCJYZW0J1cYLio"
	tnn := tn.GetNew()
	token := tnn.ClientCredentialsToken()
	fmt.Print("Returned Token: ")
	fmt.Println(token.AccessToken)

	fmt.Print("Returned Type: ")
	fmt.Println(token.TokenType)

	fmt.Print("Returned ExpiresIn: ")
	fmt.Println(token.ExpiresIn)

	fmt.Print("Returned RefreshToken: ")
	fmt.Println(token.RefreshToken)
	fmt.Print("Returned Error: ")
	fmt.Println(token.ErrorReturned)
	if token.ErrorReturned == "invalid_client" || token.TokenType != "bearer" {
		t.Fail()
	}
}
