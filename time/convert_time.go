package time

import (
	"strings"
)

func Convert_time(length string) (string, string, string) {
	hoursPos := strings.Index(length, "H")
	minutesPos := strings.Index(length, "M")
	secondsPos := strings.Index(length, "S")
	var hours, minutes, seconds string
	if hoursPos != -1 {
		hours = length[2:hoursPos]
		minutes = length[hoursPos+1:minutesPos]
		seconds = length[minutesPos+1:secondsPos]
	} else {
		hours = "0"
		if minutesPos != -1 {
			minutes = length[2:minutesPos]
			seconds = length[minutesPos+1:secondsPos]
		} else {
			minutes = "0"
			if secondsPos != -1 {
				seconds = length[minutesPos+1:secondsPos]
			} else {
				seconds = "0"
			}
		}
	}
	return hours, minutes, seconds
}