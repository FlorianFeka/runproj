package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
	args := os.Args[1:]
	
	sets := getConfigContent()

	for i:=0; i < len(sets); i++ {
		if _, exists := find(args, sets[i].Name); exists == false {
			continue
		}
		fmt.Println(sets[i].Name)
		for j:=0; j<len(sets[i].Programs) ;j++ {
			fmt.Println(sets[i].Programs[j].ProgramPath)
			fmt.Println(sets[i].Programs[j].Path)
		}
	}
}

func getConfigContent() []set{
	confFile, err := os.Open("./.ignore/runproj.json") 

	if err != nil {
		fmt.Println(err)
	}

	defer confFile.Close()
	
	byteConf, _ := ioutil.ReadAll(confFile)
	var sets []set

	json.Unmarshal(byteConf, &sets)
	return sets
}

func find(arr []string, str string) (string, bool) {
	for _, a := range arr {
		if strings.EqualFold(a, str) {
			return a, true
		}
	}
	return "", false
}