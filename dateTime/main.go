package main

import (
	"fmt"
	"time"
)

func main() {
	currTime := time.Now()
	// fmt.Printf("%T\n",currTime)//time.Time
	fmt.Println(currTime) //2025-10-30 21:21:57.2057443 +0530 IST m=+0.000000001
	// formattedCurrTime:=currTime.Format("02-01-2006, Monday") //02-dd  01-mm  2006-yyyy  Monday-day, note- M should be capital in monday
	formattedCurrTime:=currTime.Format("02/01/2006, 15:04:05 PM") //15- hh  04-mm  05-ss. write PM or AM as needed
	fmt.Println(formattedCurrTime)

	// layoutStr:="2006/01/02"
	// dateStr:="2025-11-25"
	// formatTime,_:=time.Parse(layoutStr,dateStr)
	// fmt.Println(formatTime)//0001-01-01 00:00:00 +0000 UTC , not correct

	layoutStr:="2006/01/02"
	dateStr:="2025/11/25"
	formatTime,_:=time.Parse(layoutStr,dateStr)
	fmt.Println(formatTime)//2025-11-25 00:00:00 +0000 UTC

	//add one more day 
	newDate:=currTime.Add(24*time.Hour)
	fmt.Println(newDate)//2025-10-31 21:31:39.7499292 +0530 IST m=+86400.000508201
	formatNewDate:=newDate.Format("02/01/2006")
	fmt.Println(formatNewDate)//31/10/2025

}
