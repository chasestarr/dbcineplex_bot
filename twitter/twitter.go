package twitter

import (
	"os"

	"github.com/chasestarr/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

var client *twitter.Client

func init() {
	config := oauth1.NewConfig(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"))
	token := oauth1.NewToken(os.Getenv("TWITTER_TOKEN"), os.Getenv("TWITTER_TOKEN_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)
	client = twitter.NewClient(httpClient)
}

// Tweet sends a tweet
func Tweet(data twitter.Tweet) {
	// client.Statuses.Update()
}
