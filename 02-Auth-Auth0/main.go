// package main

// import (
// 	"log"
// 	"os"
// 	"fmt"
// 	"strings"
// 	"net/http"
// 	"io/ioutil"
// 	// "os"
// 	// "crypto"
// 	// "crypto/rand"
// 	// "crypto/rsa"
// 	// "crypto/sha256"
// 	// "crypto/sha512"
// 	// "encoding/base64"
// 	// "encoding/json"

// 	"github.com/joho/godotenv"
// 	"golang.org/x/oauth2"
// )

// func (this *JWT) ParseAndVerify() (*models.SecurityPrincipal, error) {

// }

// // func main() {

// // 	err := godotenv.Load()
// // 	if err != nil {
// // 		log.Fatal("Error loading .env file")
// // 	}

// // 	domain := os.Getenv("LOCALAVANTI_CLIENT_DOMAN")
// // 	conf := &oauth2.Config{
// // 		ClientID:     os.Getenv("LOCALAVANTI__CLIENT_ID"),
// // 		ClientSecret: os.Getenv("LOCALAVANTI__CLIENT_SECRET"),
// // 		RedirectURL:  os.Getenv("LOCALAVANTI__CALLBACK_URL"),
// // 		Scopes:       []string{"openid", "profile"},
// // 		Endpoint: oauth2.Endpoint{
// // 			AuthURL:  "https://" + domain + "/authorize",
// // 			TokenURL: "https://" + domain + "/oauth/token",
// // 		},
// // 	}
// // }

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
)

var (
	domain        string
	clientID      string
	clientSecret  string
	oauthAuthURL  string
	oauthTokenURL string
	audienceURL   string
)

const (
	grantType = "client_credentials"
)

// const (
// 	errTokenError = error.New("Fail to fetch token")
// )

// type accessTokenRequest struct {
// 	clientID     string `json:"client_id"`
// 	clientSecret string `json:"client_secret"`
// 	audience     string `json:"audience"`
// 	grantType    string `json:"grant_type"`
// }

func init() {
	domain = "avantidev.auth0.com"
	clientID = "FHyiG4fmPaPzIylEm5EbC8TK4GgUtIUf"
	clientSecret = "43J9cLwCAmkiMHH1wnvmK6dTt9ejL-pvpgoNzXoALDcFOktonq97SREDJ4juWkhe"
	oauthAuthURL = "https://" + domain + "/authorize"
	oauthTokenURL = "https://" + domain + "/oauth/token"
	audienceURL = "http://localhost/avanti/v0.3/"
}

func GetAccessToken() (token *oauth2.Token, err error) {
	reqBody := fmt.Sprintf("{\"client_id\":\"%s\", \"client_secret\":\"%s\",\"audience\":\"%s\",\"grant_type\":\"%s\"}", clientID, clientSecret, audienceURL, grantType)
	payload := strings.NewReader(reqBody)
	req, err := http.NewRequest("POST", oauthTokenURL, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	fmt.Println("Sent Access Token Reqest")
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println("Received Access Token Response")

	fmt.Printf("Response: %v \n\n", res)
	fmt.Printf("Response body: %v \n", string(body))

	//token = oauth2.Token{}
	return nil, err
}

// func ExtractJWT() (token *oauth2.Token, err error) {
// }

func main() {
	token, err := GetAccessToken()

	if err != nil {
		fmt.Sprintln("No Token Error: %v", err)
		return
	}

	fmt.Println("Token: %v", token)

	//jwt, err := ExtractJWT()
	// fmt.Println(res)
	// fmt.Println(string(body))
}
