package search

import (
	"google.golang.org/api/youtube/v3"

	"log"
)

//return channel id
func Search_channelid(service *youtube.Service, developerKey string, channelName string) string {
	call := service.Search.List([]string{"snippet"}).Type("channel").Q(channelName).MaxResults(1)
	res, err := call.Do()
	if err != nil {
		log.Fatal(err)
	}
	channelId := res.Items[0].Snippet.ChannelId
	return channelId
}