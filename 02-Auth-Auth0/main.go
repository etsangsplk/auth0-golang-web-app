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
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
)

var (
	domain              string
	clientID            string
	clientSecret        string
	oauthAuthURL        string
	oauthTokenURL       string
	oauthClientGrantURL string
	audienceURL         string
)

// Auth0 Managemen API Priviledged Client
var (
	priviledgeClientID     string
	priviledgeclientSecret string
)

var (
	errNilClientID = errors.New("Client Id cannot be nil")
)

func init() {
	domain = "avantidev.auth0.com"
	clientID = "FHyiG4fmPaPzIylEm5EbC8TK4GgUtIUf"
	clientSecret = "43J9cLwCAmkiMHH1wnvmK6dTt9ejL-pvpgoNzXoALDcFOktonq97SREDJ4juWkhe"
	oauthAuthURL = "https://" + domain + "/authorize"
	oauthTokenURL = "https://" + domain + "/oauth/token"
	oauthClientGrantURL = "https://" + domain + "/api/v2/client-grants"
	audienceURL = "http://localhost/avanti/v0.3/"

	priviledgeClientID = ""
	priviledgeclientSecret = ""
}

func GrantClientAccess(client_id *string) error {
	if client_id == nil {
		return errNilClientID
	}
	scopes := []string{
		"create:client_grants",
	}
	// flatten scopes
	reqBody := fmt.Sprintf("{\"client_id\":\"%s\",\"audience\":\"%s\",\"scope\":\"%s\"}", clientID, audienceURL, scopes)
	payload := strings.NewReader(reqBody)
	req, err := http.NewRequest("POST", oauthClientGrantURL, payload)
	if err != nil {
		return err
	}
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	fmt.Println("Sent Grant Client Access Reqest")
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Println("Received Grant Client Access Access Response")

	fmt.Printf("Response: %v \n\n", res)
	fmt.Printf("Response body: %v \n", string(body))

	return nil
}

func RevokeClientAccess(client_id *string) error {
	if client_id == nil {
		return errNilClientID
	}
	scopes := []string{
		"create:client_grants",
	}
	reqBody := fmt.Sprintf("{\"client_id\":\"%s\",\"audience\":\"%s\",\"scope\":\"%s\"}", clientID, audienceURL, scopes)
	payload := strings.NewReader(reqBody)
	req, err := http.NewRequest("POST", oauthClientGrantURL, payload)
	if err != nil {
		return err
	}
	req.Header.Add("content-type", "application/json")
	return nil
}

func UpdateClientAccess(client_id *string) error {
	if client_id == nil {
		return errNilClientID
	}
	scopes := []string{
		"create:client_grants",
	}
	// flatten scopes
	reqBody := fmt.Sprintf("{\"client_id\":\"%s\",\"audience\":\"%s\",\"scope\":\"%s\"}", clientID, audienceURL, scopes)
	payload := strings.NewReader(reqBody)
	req, err := http.NewRequest("POST", oauthClientGrantURL, payload)
	if err != nil {
		return err
	}
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	fmt.Println("Sent Grant Client Access Reqest")
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Println("Received Grant Client Access Access Response")

	fmt.Printf("Response: %v \n\n", res)
	fmt.Printf("Response body: %v \n", string(body))

	return nil
}

func GetAccessToken() (token *oauth2.Token, err error) {
	// you need to grant grantType = "client_credentials" either via dashbaord or via API first or else 403 on below
	grantType := "client_credentials"
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
	fmt.Println("Sent Get Access Token Reqest")
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println("Received Access Token Response")
	fmt.Printf("Response: %v \n\n", res)
	fmt.Printf("Response body: %v \n", string(body))

	return nil, err
}

// func TestRunJob() error{

// }

func main() {

	if err := GrantClientAccess(&clientID); err != nil {
		fmt.Println("Error grant client access error=%s", err)
	}

	token, err := GetAccessToken()
	if err != nil {
		fmt.Sprintln("No Token Error: %v", err)
		return
	}
	fmt.Println("Token: %v", token)

	if err := RevokeClientAccess(&clientID); err != nil {
		fmt.Println("Error revoke client access error=%s", err)
	}

	//jwt, err := ExtractJWT()
	// fmt.Println(res)
	// fmt.Println(string(body))
}
