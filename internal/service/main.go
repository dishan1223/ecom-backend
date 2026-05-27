// Package service is used to hold the token generation function
package service

// JwtSecret variable holds the jwt secret and we define it's value with the Init() function
var JwtSecret string

// Init() function is used to initialize the service package from main.go file
func Init(s string) {
	JwtSecret = s
}
