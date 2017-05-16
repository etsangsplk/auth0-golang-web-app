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
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"golang.org/x/oauth2"

	jwt "github.com/dgrijalva/jwt-go"
	// "github.com/square/go-jose"
	// "github.com/square/go-jose/jose-util"
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
	errNilClientID         = errors.New("Client Id cannot be nil")
	errAccessTokenNotFound = errors.New("Access token not found")
)

const (
	signSecret = ""
)

type SecurityPrincipal struct {

	// namespace
	// Required: true
	Namespace *string `json:"namespace"`

	// scopes
	// Required: true
	Scopes []string `json:"scopes"`
}

// type tok jwt.Token

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
	fmt.Printf("Response body: %v \n", string(body))

	return nil, err
}

// func ParseAndVerify(access_token *string) (*SecurityPrincipal, error) {
// 	if access_token != nil {

// 	}
// 	return nil, errAccessTokenNotFound
// }

func main() {
	//type tok *jwt.Token
	//var myClaims *jwt.Claims
	//myClaims = &jwt.Claims {}

	access_token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwczovL2F2YW50aWRldi5hdXRoMC5jb20vIiwic3ViIjoiRkh5aUc0Zm1QYVB6SXlsRW01RWJDOFRLNEdnVXRJVWZAY2xpZW50cyIsImF1ZCI6Imh0dHA6Ly9sb2NhbGhvc3QvYXZhbnRpL3YwLjMvIiwiZXhwIjoxNDk0OTgxMjkwLCJpYXQiOjE0OTQ4OTQ4OTAsInNjb3BlIjoiYWxsOmF2YW50aS5jb250YWluZXIubG9ncyByZWFkOmF2YW50aS5jb250YWluZXIuam9iIHdyaXRlOmF2YW50aS5jb250YWluZXIuam9iIHJlYWQ6YXZhbnRpLmNvbnRhaW5lci5zZXJ2aWNlIHdyaXRlOmF2YW50aS5jb250YWluZXIuc2VydmljZSByZWFkOmF2YW50aS5jb250YWluZXIudGFzayB3cml0ZTphdmFudGkuY29udGFpbmVyLnRhc2sgcmVhZDphdmFudGkuY29udGFpbmVyLnRhc2stZGVmaW5pdGlvbiB3cml0ZTphdmFudGkuY29udGFpbmVyLnRhc2stZGVmaW5pdGlvbiByZWFkOmF2YW50aS5jb250YWluZXIubmFtZXNwYWNlIHdyaXRlOmF2YW50aS5jb250YWluZXIubmFtZXNwYWNlIHJlYWQ6YXZhbnRpLmNvbnRhaW5lci5jbHVzdGUgd3JpdGU6YXZhbnRpLmNvbnRhaW5lci5jbHVzdGVyIn0.aa3OiyToMN82G3OOPoS1MzonOaWGUrNNwnQFFqu_WMQ"
	//access_token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwczovL2F2YW50aWRldi5hdXRoMC5jb20vIiwic3ViIjoid3JmdnpCOXk0dEIydDd3S3B4aUM2MlZ6V3ZBWXdwTHhAY2xpZW50cyIsImF1ZCI6Imh0dHA6Ly9sb2NhbGhvc3QvYXZhbnRpL3YwLjMvIiwiZXhwIjoxNDk1MDQwMzUwLCJpYXQiOjE0OTQ5NTM5NTAsInNjb3BlIjoiIn0.ExXYa3tQQnJLRM0XEXNhzR1poICLPHBykVlq-0c71a4"
	tokenString := access_token
	var hmacSecret []byte
	// BAse64 of string??
	hmacSecret = []byte("l9rqay1dsOYc4D7SQqHKDTA1rY0FuvfO")
	//hmacSampleSecret = []byte("c2VjcmV0")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSecret, nil
	})
	if !token.Valid {
		fmt.Println("Token is invalid\n")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if err = claims.Valid(); err != nil {
			fmt.Printf("Claim is invalid error: %s \n", err.Error())
		}
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}

	if token != nil {
		fmt.Fprintf(os.Stderr, "Header:\n%v\n", token.Header)
		fmt.Fprintf(os.Stderr, "Claims:\n%v\n", token.Claims)
	}

	if err != nil {
		fmt.Printf("Couldn't parse token: %s \n", err.Error())
	}
	// Print the token details
	if err := printJSON(token.Claims); err != nil {
		fmt.Printf("Failed to output claims: %s", err.Error())
	}
	fmt.Println("\n\n")
	// principal, err := ParseAndVerify(&access_token)
	// if err != nil {
	// 	fmt.Sprintln("Access Token Verification Error: %v", err)
	// 	return
	// }
	// fmt.Sprintf("Security Principal %v", principal)

	if err := GrantClientAccess(&clientID); err != nil {
		fmt.Println("Error grant client access error=%s", err)
	}

	tokenn, err := GetAccessToken()
	if err != nil {
		fmt.Sprintln("No Token Error: %v", err)
	}
	fmt.Println("Token: %v", tokenn)

	if err := RevokeClientAccess(&clientID); err != nil {
		fmt.Println("Error revoke client access error=%s", err)
	}

	//jwt, err := ExtractJWT()
	// fmt.Println(res)
	// fmt.Println(string(body))
}

// Print a json object in accordance with the prophecy (or the command line options)
func printJSON(j interface{}) error {
	var out []byte
	var err error

	// if *flagCompact == false {
	out, err = json.MarshalIndent(j, "", "    ")
	// } else {
	// 	out, err = json.Marshal(j)
	// }

	if err == nil {
		fmt.Println(string(out))
	}

	return err
}
