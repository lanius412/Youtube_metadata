# Youtube_metadata
By using YoutubeDataAPI v3, get Videos' Metadata

## Contents
* Video Title
* Channel Name
* Video Url
* Whehther live streaming at the moment
* View Count
* Like, Dislike Count
* Length of video
* Upload Time
* Thumbnail Url (High Quality)


## Requirement Package
 * google.golang.org/api/youtube/v3
 * google.golang.org/api/googleapi/transport
 * github.com/joho/godotenv
 
## The Things to do before RUN
API Key(Developer Key)

1. Get Key from [Google Developers Console](https://www.google.com/url?sa=t&rct=j&q=&esrc=s&source=web&cd=&cad=rja&uact=8&ved=2ahUKEwiCrZi1n87uAhUxy4sBHemrA3gQFjAAegQIARAD&url=https%3A%2F%2Fconsole.developers.google.com%2F%3Fhl%3DJA&usg=AOvVaw08CfEIcxcA-mwp7e2f3XVK)
2. Rewrite DEVELOPERKEY in .env file





## Feature
* Search by Channel Name
* Search by Any Kewords
* The Numer of Videos to Search
* Search for Live-Streaming at the moment
* Search for Live-Archive
```
$ go run Youtube-metadata.go -h
  Youtube-metadata [options...]
  -c or -v flag should be set
  Available  Options:
    -a	whether to view only live-archive
    -c string
    	  channel name flag
    -live
     	  whether to search livestreaming at the moment
    -n int
      	number of videos to view (default 5)
    -s string
    	  the way to sort(date, rating, relevance, viewCount) (default "relevance")
    -v string
    	  video name flag
```
