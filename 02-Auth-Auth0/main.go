package main

import (
	"log"
	// "io/ioutil"
	// "net/http"
	// "os"
	// "crypto"
	// "crypto/rand"
	// "crypto/rsa"
	// "crypto/sha256"
	// "crypto/sha512"
	// "encoding/base64"
	// "encoding/json"

	// "golang.org/x/oauth2"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
