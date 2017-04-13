package main

import (
	"bufio"
	"fmt"
	"github.com/deckarep/golang-set"
	"gopkg.in/salsita/go-pivotaltracker.v1/v5/pivotal"
	"os"
	"regexp"
	"strconv"
)

var VERSION = "dev"

var storyMatcher = regexp.MustCompile("#(\\d{8,12})")

func main() {

	if len(os.Args) > 1 && (os.Args[1] == "-v" || os.Args[1] == "--version") {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	var token = os.Getenv("PIVOTAL_TOKEN")
	if token == "" {
		fmt.Println("You must provide $PIVOTAL_TOKEN")
		os.Exit(1)
	}

	projectID, err := strconv.Atoi(os.Getenv("PIVOTAL_PROJECT_ID"))
	if err != nil {
		fmt.Println("You must provide $PIVOTAL_PROJECT_ID")
		os.Exit(1)
	}

	set := mapset.NewSet()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		matches := storyMatcher.FindAllStringSubmatch(scanner.Text(), -1)
		for _, val := range matches {
			set.Add(val[1])
		}
	}

	fmt.Printf("Parsed story IDs:")
	for storyID := range set.Iterator().C {
		fmt.Printf(" %s", storyID.(string))
	}
	fmt.Println()

	client := pivotal.NewClient(token)
	stories, err := client.Stories.List(projectID, "state:finished story_type:feature,bug,chore")
	if err != nil {
		fmt.Printf("Error fetching stories from Pivotal: %q\n", err)
		os.Exit(1)
	}

	for _, story := range stories {
		if set.Contains(strconv.Itoa(story.Id)) {
			comment := &pivotal.Comment{
				Text: "Auto-delivered by pivotal-deliver.",
			}
			client.Stories.AddComment(projectID, story.Id, comment)

			req := &pivotal.StoryRequest{
				State: "delivered",
			}
			client.Stories.Update(projectID, story.Id, req)

			fmt.Printf("Delivered %d\n", story.Id)
		}
	}
}
