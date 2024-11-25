package main

import (
	"fmt"
	"mylog"
	"os"
	"strconv"
)

func main() {
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	mlogger := mylog.NewLogger()
	mlogger.SetDescriptor(logFile)
	for i := 0; i < 5; i++ {
		mlogger.LogINFO("i: " + strconv.Itoa(i))
	}
	mlogger.CloseMyLogger()
}
