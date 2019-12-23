package main

import (
	"context"

	"golang.org/x/oauth2"
	"zz.dev/go-github/v28/github"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "token-here"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	// do something with client
}
