package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Сбой запроса: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()

	return &result, nil
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(result.Items, func(i, j int) bool {
		return result.Items[i].CreatedAt.After(result.Items[j].CreatedAt)
	})

	fmt.Printf("%d тем:\n", result.TotalCount)

	now := time.Now()

	var index int

	fmt.Printf("Last month:\n")

	for indx, item := range result.Items {
		if item.CreatedAt.Before(now.AddDate(0,-1,0)) {
			index = indx
			break
		}

		fmt.Printf("#%-5d %9.9s %s %.55s\n",
			item.Number, item.User.Login, item.CreatedAt, item.Title)
	}

	fmt.Printf("\nLast year:\n")

	for ; index < len(result.Items); index++ {
		if result.Items[index].CreatedAt.Before(now.AddDate(0,-1,0)) &&
			result.Items[index].CreatedAt.After(now.AddDate(-1,0,0)) {
			fmt.Printf("#%-5d %9.9s %s %.55s\n",
				result.Items[index].Number, result.Items[index].User.Login,
				result.Items[index].CreatedAt, result.Items[index].Title)
		} else {
			break
		}
	}

	fmt.Printf("\nAfter year:\n")

	for ; index < len(result.Items); index++ {
		fmt.Printf("#%-5d %9.9s %s %.55s\n", result.Items[index].Number, result.Items[index].User.Login,
			result.Items[index].CreatedAt, result.Items[index].Title)
	}
}
