package main

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	config := oauth1.NewConfig("2ZSh0GKXT19OHtHmpaAYFquX7", "7c1Fm1ISZhxxj991jyLqo5QEK11D0eb7NlTE6K1R3ls0HfPOFC")
	token := oauth1.NewToken("973207180462710784-kKzBWLDCBmJVFYcKuLXUSZf2rFdpkjq", "cusy0eKHZb9NYl0PsFmGKb6cDWS73WF1cPcQyRAnXNAKj")
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	tweets, resp, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count: 20,
	})

	if err == nil {
		fmt.Printf("There are %d tweets!\n\n", len(tweets))
		fmt.Println(tweets, resp)
	}

	fmt.Println("Done.")

}
