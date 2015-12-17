package gozulipbot

import (
	"flag"
	"fmt"
)

func GetConfigFromFlags() (string, string, error) {
	var (
		emailAddress = flag.String("email", "", "bot email address")
		apiKey       = flag.String("apikey", "", "bot api key")
	)
	flag.Parse()

	if *emailAddress == "" {
		return "", "", fmt.Errorf("--email required, but wasn't set")
	}
	if *apiKey == "" {
		return "", "", fmt.Errorf("--apikey required, but wasn't set")
	}

	return *emailAddress, *apiKey, nil
}