package check

import (
	"google.golang.org/api/youtube/v3"

	"fmt"
	"log"
	"os"
)

//check for view of livestreaming
func Check_live_now(service *youtube.Service, developerKey string, channelId string, videoName string) {
	var checkcall *youtube.SearchListCall
	if channelId != "" {
		checkcall = service.Search.List([]string{"snippet"}).ChannelId(channelId).Type("channel").MaxResults(1)
	} else {
		checkcall = service.Search.List([]string{"snippet"}).Q(videoName).Type("video").Order("date").MaxResults(1)
	}
	checkres , err := checkcall.Do()
	if err != nil {
		log.Fatal(err)
	}
	liveState := checkres.Items[0].Snippet.LiveBroadcastContent
	if liveState != "live" {
		fmt.Println("|| Currently, No livestreaming in this search result ||")
		os.Exit(0)
	}
}