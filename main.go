package main

import (
	"bufio"
	"fmt"
	"github.com/deckarep/golang-set"
	"os"
	"regexp"
	"strings"
)

var storyMatcher = regexp.MustCompile("#(\\d{8,12})")

func main() {
	set := mapset.NewSet()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		matches := storyMatcher.FindAllStringSubmatch(scanner.Text(), -1)
		for _, val := range matches {
			set.Add(val[1])
		}
	}
	slice := set.ToSlice()
	storyIDs := make([]string, len(slice))
	for i, storyID := range slice {
		storyIDs[i] = storyID.(string)
	}
	fmt.Printf("Parsed story IDs: %s\n", strings.Join(storyIDs, " "))
}
