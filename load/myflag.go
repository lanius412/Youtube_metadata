package load

import (
	"flag"
	"fmt"
	"os"
)

//cmdline arguments flags
func Myflag() (string, string, int, string, bool, bool) {
	var (
		channelNameFlag = flag.String("c", "", "channel name flag")
		videoNameFlag = flag.String("v", "", "video name flag")
		videoNumberFlag = flag.Int("n", 5, "number of videos to view")
		sortFlag = flag.String("s", "relevance", "the way to sort(date, rating, relevance, viewCount)")
		liveFlag = flag.Bool("live", false, "whether to search livestreaming at the moment")
		archiveFlag = flag.Bool("a", false, "whether to view only live-archive")
	)
	flag.Usage = func() {
		fmt.Println("yt_meta [options...]")
		fmt.Println("-c or -v flag should be set")
		fmt.Println("Available  Options: ")
		flag.PrintDefaults()
	}
	flag.Parse()
	if len(flag.Args()) < 0 {
		flag.Usage()
		os.Exit(0)
	}
	//log.Println(*channelNameFlag, *videoNameFlag, *videoNumberFlag, *sortFlag, *liveFlag, *archiveFlag)
	return *channelNameFlag, *videoNameFlag, *videoNumberFlag, *sortFlag, *liveFlag, *archiveFlag
}