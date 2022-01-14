// 一ヶ月未満、一年未満、一年以上の期間で分類された結果を報告するようにissuesを修正しなさい。
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

const (
	lessThanMonth = "lessThanMonth"
	lessThanYear  = "lessThanYear"
	moreThanYear  = "moreThanYear"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	classified := classifiedIssuesByCreatedDate(&(result.Items))

	fmt.Printf("%d issues:\n", result.TotalCount)
	for class, issues := range classified {
		fmt.Printf("\n%v: %d issues\n", class, len(issues))
		printIssues(&issues)
	}
}

func classifiedIssuesByCreatedDate(issues *[]*github.Issue) map[string][]*github.Issue {
	classified := make(map[string][]*github.Issue)
	classified[lessThanMonth] = make([]*github.Issue, 0)
	classified[lessThanYear] = make([]*github.Issue, 0)
	classified[moreThanYear] = make([]*github.Issue, 0)
	currentTime := time.Now().UTC()
	for _, issue := range *issues {
		switch {
		case issue.CreatedAt.After(currentTime.Add(-(time.Hour * 24 * 30))):
			classified[lessThanMonth] = append(classified[lessThanMonth], issue)
		case issue.CreatedAt.After(currentTime.Add(-(time.Hour * 24 * 365))):
			classified[lessThanYear] = append(classified[lessThanYear], issue)
		default:
			classified[moreThanYear] = append(classified[moreThanYear], issue)
		}
	}

	return classified
}

func printIssues(issues *[]*github.Issue) {
	for _, issue := range *issues {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			issue.Number, issue.User.Login, issue.Title)
	}
}
