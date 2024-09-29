package main

import (
    "fmt"
	"io/ioutil"
	"encoding/json"
	"os"
	"strings"
	"log"
)



type LegalFee struct{
    Name string `json: "name"`
    Value string `json: "value"`
}

type Data struct{
	LegalFees []LegalFee `json: "legalFees"`
}

func main() {
	arg := os.Args
	searched := strings.ToUpper(strings.Join(arg[1:], " "))
	
	file, _ := ioutil.ReadFile("legalFees.json")
	data := Data{}
	_ = json.Unmarshal([]byte(file), &data)
	
	logFile := "legalFeesLogFile"
	
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	
	if err != nil {
		log.Fatal("Could not log to file: ", logFile)
	}
	
	log.SetOutput(f)
	
	defer f.Close()
	
	found := false
	 
	for _, fee := range data.LegalFees {
		if strings.Contains(fee.Name, searched) {
			fmt.Printf("-> %s: %s\n", fee.Name, fee.Value)
			log.Print(" => ", fee.Name, ": ", fee.Value, "\n **********")
			found = true
		}
    }
	
	if (!found) {
		fmt.Printf("-> %s: não encontrado\n", searched)
		log.Print(" => ", searched, ": não encontrado", "\n **********")
	}
}