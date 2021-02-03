package main

import (
	"net/http"
	"log"

	"google.golang.org/api/youtube/v3"
	"google.golang.org/api/googleapi/transport"

	"Youtube-metadata/search"
	"Youtube-metadata/load"
)

func main() {
	//args := new(SearchArgs)
	var args search.SearchArgs
	args.DeveloperKey = load.Read_env_file()
	args.ChannelName, args.VideoName, args.VideoNum, args.Sort, args.IsLiveNow, args.ArchiveOnly = load.Myflag()

	client := &http.Client {
		Transport: &transport.APIKey{Key: args.DeveloperKey},
	}
	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error Creating New Youtube Client: %w", err)
	}

	args.Search_videos(service)
}