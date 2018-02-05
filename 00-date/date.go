package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	//Print time in yyyymmdd_hhmmss.msec format
	//mm-01(month), dd-02, hh-03(15), mm-04, ss-05, yy-06 (yyyy-2006)
	fmt.Println(startTime.Format("20060102_150405.000"))
}
