package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GetUserInfo(accessToken string) (map[string]interface{}, error) {
    userInfoEndpoint := "https://www.googleapis.com/oauth2/v2/userinfo"
    resp, err := http.Get(fmt.Sprintf("%s?access_token=%s", userInfoEndpoint, accessToken))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var userInfo map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
        return nil, err
    }

    return userInfo, nil
}


// Function to sign a JWT with user information
func SignJWT(userInfo map[string]interface{}) (string, error) {
    // Customize the claims as needed
    claims := jwt.MapClaims{
        "sub": userInfo["id"],
		"name": userInfo["name"],
		"email": userInfo["email"],
        "iss": "oauth-app-golang",
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(), // 30 days
        // Add other claims as needed
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signedToken, err := token.SignedString([]byte("your-secret-key")) // Replace with your actual secret key
    if err != nil {
        return "", err
    }

    return signedToken, nil
}