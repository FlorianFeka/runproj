package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type set struct {
	Name string `json:"name"`
	Programs []program `json:"programs"`
}

type program struct {
	ProgramPath string `json:"program"`
	Path string `json:"path"`
}

func main() {
	confFile, err := os.Open("./.ignore/runproj.json")

	if err != nil {
		fmt.Println(err)
	}

	defer confFile.Close()
	
	byteConf, _ := ioutil.ReadAll(confFile)

	var set []set

	json.Unmarshal(byteConf, &set)

	for i:=0; i < len(set); i++ {
		fmt.Println(set[i].Name)
		for j:=0; j<len(set[i].Programs) ;j++ {
			fmt.Println(set[i].Programs[j].ProgramPath)
			fmt.Println(set[i].Programs[j].Path)
		}
	}
}