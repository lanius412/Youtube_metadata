package search

import (
	"google.golang.org/api/youtube/v3"

	"log"
	"strconv"
	"os"
	"fmt"

	"Youtube-metadata/check"
	"Youtube-metadata/time"
	"Youtube-metadata/preview_save"
)

type SearchArgs struct {
	DeveloperKey string
	ChannelName string
	VideoName string
	VideoNum int
	Sort string
	IsLiveNow bool
	ArchiveOnly bool
}

//search videos , view metadata
func (args SearchArgs) Search_videos(service *youtube.Service) {
	var call *youtube.SearchListCall
	if args.ChannelName != "" {
		channelId := Search_channelid(service, args.DeveloperKey, args.ChannelName)
		if args.IsLiveNow {
			check.Check_live_now(service, args.DeveloperKey, channelId, "")
			call = service.Search.List([]string{"id, snippet"}).ChannelId(channelId).Type("video").EventType("Live").Order(args.Sort).MaxResults(1)
		} else {
			if args.VideoName != "" {
				if args.ArchiveOnly {
					check.Check_have_archive(service, args.DeveloperKey, channelId, args.VideoName)
					call = service.Search.List([]string{"id, snippet"}).ChannelId(channelId).Type("video").EventType("Completed").Q(args.VideoName).Order(args.Sort).MaxResults(int64(args.VideoNum))
				} else {
					call = service.Search.List([]string{"id, snippet"}).ChannelId(channelId).Type("video").Q(args.VideoName).Order(args.Sort).MaxResults(int64(args.VideoNum))
				}
			} else {
				if args.ArchiveOnly {
					check.Check_have_archive(service, args.DeveloperKey, channelId, "")
					call = service.Search.List([]string{"id, snippet"}).ChannelId(channelId).Type("video").EventType("Completed").Order(args.Sort).MaxResults(int64(args.VideoNum))
				} else {
					call = service.Search.List([]string{"id, snippet"}).ChannelId(channelId).Type("video").Order(args.Sort).MaxResults(int64(args.VideoNum))
				}
			}
		}
	} else {
		if args.VideoName != "" {
			if args.IsLiveNow {
				check.Check_live_now(service, args.DeveloperKey, "", args.VideoName)
				call = service.Search.List([]string{"id, snippet"}).Type("video").EventType("Live").Q(args.VideoName).Order(args.Sort).MaxResults(int64(args.VideoNum))
			} else {
				if args.ArchiveOnly {
					check.Check_have_archive(service, args.DeveloperKey, "", args.VideoName)
					call = service.Search.List([]string{"id, snippet"}).Type("video").EventType("Completed").Q(args.VideoName).Order(args.Sort).MaxResults(int64(args.VideoNum))
				} else {
					call = service.Search.List([]string{"id, snippet"}).Type("video").Q(args.VideoName).Order(args.Sort).MaxResults(int64(args.VideoNum))
				}
			}
		}
	}
	res, err := call.Do()
	if err != nil {
		log.Fatal(err)
	}

	videoNames := make([]string, args.VideoNum)
	thumbnailsUrl := make([]string, args,VideoNum)
	for v := 0; v < args.VideoNum; v++ {
		subCall := service.Videos.List([]string{"contentDetails,statistics"}).Id(res.Items[v].Id.VideoId)
		subRes, err := subCall.Do()
		if err != nil {
			log.Fatal(err)
		}
		contentDetails := subRes.Items[0].ContentDetails
		statistics := subRes.Items[0].Statistics
		snippet := res.Items[v].Snippet
		videoUrl := " ( https://youtu.be/" + res.Items[v].Id.VideoId + " )"
		videoNames[v] = snippet.Title
		thumbnailsUrl[v] = snippet.Thumbnails.High.Url
		fmt.Print("\nTitle : " + snippet.Title + videoUrl)
		if snippet.LiveBroadcastContent == "live" {
			fmt.Print(" [\x1b[31mStreaming Now\x1b[0m]")
		}
		fmt.Println("\nChannel Name : " + snippet.ChannelTitle)
		fmt.Print("ViewCount : " + strconv.FormatUint(statistics.ViewCount, 10))
		fmt.Print("    |    LikeCount : " + strconv.FormatUint(statistics.LikeCount, 10))
		fmt.Println("    |    DislikeCount : " + strconv.FormatUint(statistics.DislikeCount, 10))
		hours, minutes, seconds := time.Convert_time(contentDetails.Duration)
		fmt.Printf("Length : %sh %sm %ss", hours, minutes, seconds)
		fmt.Println("\nUpload Time : " + snippet.PublishedAt)
		fmt.Println("Thumbnail Url : " + snippet.Thumbnails.High.Url)
		if (args.ChannelName != "" && args.IsLiveNow && v == 0) {
			os.Exit(0)
		}
	}
	preview_save.IsPreview_Save(videoNames, thumbnailsUrl, args.VideoNum
}
