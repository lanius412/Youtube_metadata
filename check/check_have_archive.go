package check

import (
	"google.golang.org/api/youtube/v3"

	"fmt"
	"log"
	"os"
)

//check for view of archive
func Check_have_archive(service *youtube.Service, developerKey string, channelId string, videoName string) {
	var checkcall *youtube.SearchListCall
	if videoName != "" {
		checkcall = service.Search.List([]string{"id"}).ChannelId(channelId).Type("video").EventType("Completed").Q(videoName).MaxResults(1)
	} else {
		checkcall = service.Search.List([]string{"id"}).Q(videoName).Type("video").EventType("Completed").MaxResults(1)
	}
	checkres, err := checkcall.Do()
	if err != nil {
		log.Fatal(err)
	}
	if checkres.PageInfo.ResultsPerPage == 0 {
		fmt.Println("|| No archive in this search result ||")
		os.Exit(0)
	}
}