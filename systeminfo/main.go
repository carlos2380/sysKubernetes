package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"runtime"
	"syscall"
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
	sysinfo := syscall.Sysinfo_t{}
	err := syscall.Sysinfo(&sysinfo)
	checkErr(err)

	info := info{}
	info.Memory = sysinfo.Totalram / (1024 * 1024)
	info.Core = runtime.NumCPU()

	data, err := json.Marshal(info)
	checkErr(err)
	err = ioutil.WriteFile(*path, data, 0666)
	checkErr(err)
}
