package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("date time")
	currentTime := time.Now() //! 2025-01-13 06:58:36.1896058 +0600 +06 m=+0.000523501
	fmt.Println(currentTime)

	formatedTime := currentTime.Format("02-01-2006,Monday,15:04:05,3:04PM") //! 13-01-2025,Monday,07:32:04,7:32AM
	fmt.Println(formatedTime)

	layout := "02-01-2006"
	timeString := "13-01-2025"
	convertedToTime, _ := time.Parse(layout, timeString)
	fmt.Println(convertedToTime)
}
