package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/mmcdole/gofeed"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <username>", os.Args[0])
	}

	fp := gofeed.NewParser()
	url := fmt.Sprintf("https://github.com/%s.atom", os.Args[1])
	feed, err := fp.ParseURL(url)
	if err != nil {
		log.Fatalf("Fetch/Parse error for url=%q : %s", url, err)
	}

	sort.Slice(feed.Items, func(i int, j int) bool {
		return feed.Items[i].UpdatedParsed.Before(*feed.Items[j].UpdatedParsed)
	})

	for _, item := range feed.Items {
		fmt.Printf("%s %s\n\t%s\n", item.Updated, item.Title, item.Link)
	}
}
