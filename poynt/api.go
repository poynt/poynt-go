// Package Poynt provides a basic Poynt API wrapper.

package poynt

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/satori/go.uuid"

	jwtGo "github.com/dgrijalva/jwt-go"
)

const (
	apiUrl     = "https://services.poynt.net"
	tokenUrl   = "https://services.poynt.net/token"
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
	applicationId  string
	pemBytes       []byte
	poyntToken     *PoyntTokenResponse
	accessToken    string
	expirationTime int64
}

type JWT struct {
	exp int64
	iat int64
	iss string
	sub string
	aud string
	jti string
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

// gets an access token and attaches it to the PoyntAPi instance
func (self *PoyntApi) GetAccessToken() error {
	timeNow := time.Now().UTC()

	token := jwtGo.NewWithClaims(jwtGo.SigningMethodRS256, jwtGo.MapClaims{
		"exp": timeNow.Add(time.Minute * 5).Unix(),
		"iat": timeNow.Unix(),
		"iss": self.applicationId,
		"sub": self.applicationId,
		"aud": "https://services.poynt.net",
		"jti": uuid.NewV4().String(),
	})

	privateKey, err := jwtGo.ParseRSAPrivateKeyFromPEM(self.pemBytes)
	if err != nil {
		fmt.Println(err)
		return err
	}

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		fmt.Println("token string parse error", err)
		return err
	}

	reqId := uuid.NewV4().String()

	reqBody := fmt.Sprintf("grantType=urn:ietf:params:oauth:grant-type:jwt-bearer&assertion=%s", tokenString)
	req, err := http.NewRequest("POST", tokenUrl, bytes.NewBufferString(reqBody))
	req.Header = make(map[string][]string)
	req.Header.Set("api-version", apiVersion)
	req.Header.Set("Poynt-Request-Id", reqId)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

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

	var poyntTokenResponse *PoyntTokenResponse
	err = json.Unmarshal(respBytes, &poyntTokenResponse)
	if err != nil {
		return err
	}

	self.poyntToken = poyntTokenResponse
	self.accessToken = poyntTokenResponse.AccessToken
	self.expirationTime = time.Now().UTC().Unix() + poyntTokenResponse.ExpiresIn

	PrettyPrint(self.poyntToken)

	return nil
}

// Makes a GET request
func (self *PoyntApi) Get(path string, query url.Values, r interface{}) error {
	fmt.Println("GET request", path, query)
	req, err := http.NewRequest("GET", apiUrl+path, nil)

	if err != nil {
		return err
	}

	q := req.URL.Query()

	// add the query params to the query
	for key, value := range query {
		fmt.Println(key, value)

		for _, v := range value {
			q.Add(key, v)
		}
	}

	req.URL.RawQuery = q.Encode()

	return self.do(req, r)
}

// Makes a POST request
// body is the payload for the POST request
// r is the response interface
func (self *PoyntApi) Post(path string, body interface{}, r interface{}) error {
	fmt.Println("\n\nPOST request", path)
	PrettyPrint(body)

	bytes := new(bytes.Buffer)
	json.NewEncoder(bytes).Encode(body)

	req, err := http.NewRequest("POST", apiUrl+path, bytes)

	if err != nil {
		return err
	}

	return self.do(req, r)
}

func (self *PoyntApi) do(req *http.Request, r interface{}) error {
	if self.needsRefresh() {
		self.GetAccessToken()
	}

	req.Header.Set("api-version", apiUrl)
	req.Header.Set("Authorization", fmt.Sprintf("%s %s", self.poyntToken.TokenType, self.accessToken))
	req.Header.Set("Poynt-Request-Id", uuid.NewV4().String())
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{Timeout: 10 * time.Second}
	resp, err := httpClient.Do(req)
	fmt.Println("\nResponse:")
	fmt.Println(resp)

	fmt.Println("Response body:")
	fmt.Println(resp.Body)

	httpError := ErrorHandler(resp)
	if httpError != nil {
		return httpError
	}

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return decodeResponse(resp.Body, r)
}

// checks whether a token is aboue to expire or is already expired
func (self *PoyntApi) needsRefresh() bool {
	// if token is going to expire in an hour
	return self.poyntToken == nil || time.Now().UTC().Unix()+(60*60) > self.expirationTime
}

// decodes a response into the interface
func decodeResponse(body io.Reader, to interface{}) error {
	err := json.NewDecoder(body).Decode(to)

	if err != nil {
		return fmt.Errorf("poynt: error decoding body; %s", err.Error())
	}

	return nil
}
