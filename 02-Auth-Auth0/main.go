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
	"os"
	"strings"

	"golang.org/x/oauth2"
)

var (
	domain          string
	client_id       string
	client_secret   string
	oauth_auth_url  string
	oauth_token_url string
	audience_url    string
)

func init() {
	domain := os.Getenv("LOCALAVANTI_CLIENT_DOMAIN")
	client_id := os.Getenv("LOCALAVANTI_CLIENT_ID")
	client_secret := os.Getenv("LOCALAVANTI_CLIENT_SECRET")
	os.Getenv("LOCALAVANTI_CLIENT_SECRET")
	oauth_auth_url := "https://" + domain + "/authorize"
	oauth_token_url := "https://" + domain + "/oauth/token"
	audience_url := os.Getenv("AUTH0_DOMAIN")
}

func AccessTokenRequest() (token *oauth2.Token, err error) {

	url := oauth_token_url

	payload := strings.NewReader("{\"client_id\":\"wrfvzB9y4tB2t7wKpxiC62VzWvAYwpLx\",\"client_secret\":\"1-4X3ltzTe98m5J7FL0wK179ancCWOxz1hbYfo0VgZgCRUEVNgvYNIaa_9tDabWB\",\"audience\":\"http://localhost/avanti/v0.3/\",\"grant_type\":\"client_credentials\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	return nil, err
}

func main() {
	token, err := AccessTokenRequest()
	if err != nil {
		fmt.Sprintln("No Token Error: %v", err)
		return
	}
	fmt.Sprintln("Token: %v", token)
	// fmt.Println(res)
	// fmt.Println(string(body))
}
