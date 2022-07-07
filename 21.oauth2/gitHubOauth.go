package main

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"strings"
)

const (
	clientID     = "<your-client-id>"
	clientSecret = "<your-client-secret>"
	CookieName   = "state-token"
)

var GitHubOauthConfig = &oauth2.Config{
	ClientID:     clientID,
	ClientSecret: clientSecret,
	Endpoint:     github.Endpoint,
}

type GitHubResponse struct {
	Data struct {
		Viewer struct {
			Name string `json:"name"`
		} `json:"viewer"`
	} `json:"data"`
}

func fetchUserData(ctx context.Context, token *oauth2.Token) (GitHubResponse, error) {

	tokenSource := GitHubOauthConfig.TokenSource(ctx, token)

	response, err := oauth2.NewClient(ctx, tokenSource).Post(
		"https://api.github.com/graphql",
		"application/json",
		strings.NewReader(`{"query": "query {viewer {name}}"}`))

	if err != nil {
		return GitHubResponse{}, fmt.Errorf("couldn't get user: %w", err)
	}

	defer response.Body.Close()

	var gitHubResponse GitHubResponse
	if err = json.NewDecoder(response.Body).Decode(&gitHubResponse); err != nil {
		return GitHubResponse{}, fmt.Errorf("invalid GitHub Response: %w", err)
	}
	return gitHubResponse, nil
}
