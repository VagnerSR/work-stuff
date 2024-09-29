package main

import (
    "fmt"
	"io/ioutil"
	"encoding/json"
	"os"
	"strings"
	"log"
)



type System struct{
    Name string `json: "name"`
    WhereTo string `json: "whereTo"`
	CanIDoIt string `json: "canIDoIt"`
}

type Data struct{
	Systems []System `json: "systems"`
}

func main() {
	arg := os.Args
	searched := strings.ToUpper(strings.Join(arg[1:], " "))
	
	file, _ := ioutil.ReadFile("systems.json")
	data := Data{}
	_ = json.Unmarshal([]byte(file), &data)
	
	logFile := "systemsLogFile"
	
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	
	if err != nil {
		log.Fatal("Could not log to file: ", logFile)
	}
	
	log.SetOutput(f)
	
	defer f.Close()
	
	found := false
	 
	for _, system := range data.Systems {
		if strings.Contains(system.Name, searched) {
			fmt.Printf("-> %s: %s\n-> %s\n", system.Name, system.WhereTo, system.CanIDoIt)
			log.Print(" => ", system.Name, ": ", system.WhereTo, " ", system.CanIDoIt, "\n **********")
			found = true
		}
    }
	
	if (!found) {
		fmt.Printf("-> %s: não encontrado\n", searched)
		log.Print(" => ", searched, ": não encontrado", "\n **********")
	}
}