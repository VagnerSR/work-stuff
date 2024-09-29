package main

import (
    "fmt"
	"io/ioutil"
	"encoding/json"
	"os"
	"strings"
	"log"
)



type Site struct{
    SiteName string `json: "siteName"`
    OfficerName string `json: "officerName"`
}

type Data struct{
	Sites []Site `json: "sites"`
}

func main() {
	arg := os.Args
	searched := strings.ToUpper(strings.Join(arg[1:], " "))
	
	file, _ := ioutil.ReadFile("sites.json")
	data := Data{}
	_ = json.Unmarshal([]byte(file), &data)
	
	logFile := "sitesLogFile"
	
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	
	if err != nil {
		log.Fatal("Could not log to file: ", logFile)
	}
	
	log.SetOutput(f)
	
	defer f.Close()
	
	found := false
	 
	for _, site := range data.Sites {
		if strings.Contains(site.SiteName, searched) {
			fmt.Printf("-> %s: %s\n", site.OfficerName, site.SiteName)
			log.Print(" => ", site.OfficerName, ": ", site.SiteName, "\n **********")
			found = true
		}
    }
	
	if (!found) {
		fmt.Printf("-> %s: não encontrado\n", searched)
		log.Print(" => ", searched, ": não encontrado", "\n **********")
	}
}