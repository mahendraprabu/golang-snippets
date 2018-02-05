package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	//Print time in format yyyymmdd_hhmmss.msec
	fmt.Println(startTime.Format("20060102_150405.000"))
}
