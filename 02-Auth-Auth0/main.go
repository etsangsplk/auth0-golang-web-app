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
	"log"
	"net/http"
	"strings"

	"github.com/joho/godotenv"
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

	fmt.Println(reqBody)

	req, _ := http.NewRequest("POST", oauthTokenURL, payload)

	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	return nil, err
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	token, err := GetAccessToken()
	if err != nil {
		fmt.Sprintln("No Token Error: %v", err)
		return
	}
	fmt.Sprintln("Token: %v", token)
	// fmt.Println(res)
	// fmt.Println(string(body))
}
