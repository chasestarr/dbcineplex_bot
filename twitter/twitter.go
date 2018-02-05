package twitter

import (
	"log"
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

func uploadImage(uri string) int64 {
	twitterRes, _, err := client.Media.UploadFile(uri)
	if err != nil {
		log.Fatal(err)
	}
	return twitterRes.MediaID
}

// Tweet sends a tweet
func Tweet(text string, images []string) {
	mediaIds := []int64{}
	for _, image := range images {
		mediaIds = append(mediaIds, uploadImage(image))
	}

	updateParams := &twitter.StatusUpdateParams{MediaIds: mediaIds}
	_, _, err := client.Statuses.Update(text, updateParams)
	if err != nil {
		log.Fatal(err)
	}
}
