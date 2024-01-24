package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/shurcooL/graphql"
)

type languageStatsQuery struct {
	MatchedUser struct {
		LanguageProblemCount []struct {
			LanguageName  string `graphql:"languageName"`
			ProblemsSolved int    `graphql:"problemsSolved"`
		} `graphql:"languageProblemCount"`
	} `graphql:"matchedUser(username: $username)"`
}

func main() {
	username := "someone"

	apiEndpoint := "https://leetcode.com/graphql"

	client := graphql.NewClient(apiEndpoint, &http.Client{})

	ctx := context.Background()

	var query languageStatsQuery

	variables := map[string]interface{}{
		"username": graphql.String(username),
	}

	err := client.Query(ctx, &query, variables)
	if err != nil {
		log.Fatal(err)
	}

	for _, lang := range query.MatchedUser.LanguageProblemCount {
		fmt.Printf("Language: %s, Problems Solved: %d\n", lang.LanguageName, lang.ProblemsSolved)
	}
}
