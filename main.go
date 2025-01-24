package main

import (
	"flag"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"io"
	"os"
	"time"
)

func main() {
	iss := flag.String("iss", "www.campus.com", "The token issuer.")
	sub := flag.String("sub", "jack.ryan", "The token subject.")
	flag.Parse()
	keyBytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyBytes)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	now := time.Now()
	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.RegisteredClaims{
		Issuer:    *iss,
		Subject:   *sub,
		IssuedAt:  jwt.NewNumericDate(now),
		NotBefore: jwt.NewNumericDate(now),
		//ExpiresAt: jwt.NewNumericDate(now.Add(time.Minute * 5)),
	}).SignedString(privateKey)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stdout, token)
}
