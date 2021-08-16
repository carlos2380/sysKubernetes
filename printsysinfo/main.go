package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

type info struct {
	Memory uint64 `json:"memory"`
	Core   int    `json:"core"`
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	path := flag.String("pathfile", "./data.txt", "path and file to write de info")
	flag.Parse()
	infobyte, err := ioutil.ReadFile(*path)
	checkErr(err)
	info := &info{}
	err = json.Unmarshal(infobyte, info)
	checkErr(err)
	log.Printf(fmt.Sprintf("Memory : %dMB", info.Memory))
	log.Printf(fmt.Sprintf("CPUs : %d", info.Core))
}
