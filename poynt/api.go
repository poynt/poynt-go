// Package Poynt provides a basic Poynt API  wrapper.

package poynt

import (
  "os"
  "encoding/json"
  "errors"
  "bytes"
	// "encoding/json"
  "net/http"
  "time"
	"fmt"
  jwtGo "github.com/dgrijalva/jwt-go"
  "github.com/satori/go.uuid"
  "io/ioutil"
)

var (
	tokenUrl = "https://services.poynt.net/token"
  apiVersion = "1.2"
)

type PoyntTokenResponse struct {
	ExpiresIn    int64  `json:"expiresIn"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Scope        string `json:"scope"`
	TokenType    string `json:"tokenType"`
}

type PoyntApi struct {
	applicationId string
  pemBytes []byte
  lastResponse PoyntTokenResponse
  accessToken string
  expirationTime int64
}

type JWT struct {
  exp       int64
  iat       int64
  iss       string
  sub       string
  aud       string
  jti       string
}

func (self *PoyntApi) Init(applicationId string, pemFileName string) error {
	pemFile, err := os.Open(pemFileName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	pemBytes, err := ioutil.ReadAll(pemFile)
	if err != nil {
		fmt.Println(err)
		return err
	}
  self.applicationId = applicationId
	self.pemBytes = pemBytes
	return nil
}

func (self *PoyntApi) GetAccessToken() error {
  token := jwtGo.New(jwtGo.SigningMethodRS256)
  claims := make(jwtGo.MapClaims)
	claims["exp"] = time.Now().UTC().Add(time.Hour * 1).Unix()
	claims["iat"] = time.Now().UTC().Unix()
	claims["iss"] = "https://services.poynt.net"
	claims["sub"] = self.applicationId
	claims["aud"] = self.applicationId
	claims["jti"] = uuid.NewV4().String()
  token.Claims = claims
  fmt.Println("Before signing string")
  fmt.Println(self.pemBytes)
	tokenString, err := token.SignedString(self.pemBytes)
	if err != nil {
    fmt.Println("token string parse error", err, tokenString)
		return err
	}

	reqId := uuid.NewV4().String()

	reqBody := fmt.Sprintf("grantType=urn:ietf:params:oauth:grant-type:jwt-bearer&assertion=%s", tokenString)
	req, err := http.NewRequest("POST", tokenUrl, bytes.NewBufferString(reqBody))
	req.Header = make(map[string][]string)
  req.Header.Set("api-version", apiVersion)
	req.Header.Set("Poynt-Request-Id", reqId)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	fmt.Println("url", tokenUrl)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
    fmt.Println(resp)
		return errors.New("Not an OK response")
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var poyntTokenResponse PoyntTokenResponse
	err = json.Unmarshal(respBytes, &poyntTokenResponse)
	if err != nil {
		return err
	}
	fmt.Println(string(respBytes))
	self.lastResponse = poyntTokenResponse
	self.accessToken = poyntTokenResponse.AccessToken
	self.expirationTime = time.Now().UTC().Unix() + poyntTokenResponse.ExpiresIn
	return nil
}


// // Create an API with either code
// func New(clientId string, clientSecret string, accessToken string, enforceSignedRequest bool) *Api {
// 	if clientId == "" && accessToken == "" {
// 		panic("ClientId or AccessToken must be given to create an Api")
// 	}
//
// 	if enforceSignedRequest && clientSecret == "" {
// 		panic("ClientSecret is required for signed request")
// 	}
//
// 	return &Api{
// 		ClientId:             clientId,
// 		ClientSecret:         clientSecret,
// 		AccessToken:          accessToken,
// 		EnforceSignedRequest: enforceSignedRequest,
// 	}
// }
