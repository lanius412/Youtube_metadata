package preview_save

import (
	"strconv"
	"net/http"
	"os"
	"io"
	"os/exec"
	"bufio"
	"fmt"
	"log"
)

type Thumbnail struct {
	VideoNames []string
	ThumbnailsUrl []string
	VideoNum int
}

func IsPreview_Save(videoNames []string, thumbnailsUrl []string, videoNum int) {
	fmt.Print("\nPreview Thumnails ? y or n ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Text() == "y" {
		tmb := new(Thumbnail)
		tmb.VideoNames = videoNames
		tmb.ThumbnailsUrl = thumbnailsUrl
		tmb.VideoNum = videoNum
		tmb.Preview_thumbnail(scanner)
	} else {
		os.Exit(0)
	}
}

func (tmb Thumbnail)Preview_thumbnail(scanner *bufio.Scanner) {
	fmt.Println("Preview which thumbnails ? Numbers from 1 up to " + strconv.Itoa(tmb.VideoNum))
	//scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	selectedVideo, _ := strconv.Atoi(scanner.Text())
	if selectedVideo < 1 || selectedVideo > tmb.VideoNum {
		fmt.Println("Enter the valid value")
		scanner.Scan()
		selectedVideo, _ = strconv.Atoi(scanner.Text())
	}
	res, err := http.Get(tmb.ThumbnailsUrl[selectedVideo-1])
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	fileName := tmb.VideoNames[selectedVideo-1] + "_thumbnail.jpg"
	file, err := os.Create(fileName) //Create File Name
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	wd, _ := os.Getwd() //Get Current Directory
	thumbnailPath := wd + "/" + fileName
	//log.Println(thumbnailPath)

	io.Copy(file, res.Body) //Temporarily Save
	err = exec.Command("qlmanage", "-p", thumbnailPath, ">/dev/null 2>/dev/null").Start() //View by QuickLook
	if err != nil {
		log.Fatal(err)
	}
	tmb.Save_thumbnail(scanner, fileName, thumbnailPath)
}

func (tmb Thumbnail)Save_thumbnail(scanner *bufio.Scanner, fileName string, thumbnailPath string) {
	fmt.Print("Save this Thumbnail ? y or n ") //Select Save or No
	scanner.Scan()
	if scanner.Text() == "y" {
		fmt.Println("Save Thumbnail : " + fileName)
	} else {
		err := os.Remove(thumbnailPath)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Print("\nPreview more thumbnails ? y or n ")
	scanner.Scan()
	if scanner.Text() == "y" {
		tmb.Preview_thumbnail(scanner)
	} else {
		os.Exit(0)
	}
}