package main

import (
    "fmt"
    "time"
	"log"
	"os"
)

func main() {
	layout := "15:04:05"
    now := time.Now().Format(layout)
	
	time1, _ := time.Parse(layout, now)
    time2, _ := time.Parse(layout, "17:00:00")
	
	duration := time2.Sub(time1)
	
	logFile := "logfile"
	
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	
	if err != nil {
		log.Fatal("Could not log to file: ", logFile)
	}
	
	defer f.Close()
	log.SetOutput(f)
	
	if duration < 0 {
		log.Print("Você está livre!", "\n **********")
		fmt.Println("-> Você está livre!")
	} else {
		log.Print("Faltam:", duration, "\n **********")
		fmt.Println("-> Faltam:", duration)
	}
}