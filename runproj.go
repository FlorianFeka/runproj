package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type set struct {
	Name string `json:"name"`
	Programs []program `json:"programs"`
}

type program struct {
	ProgramPath string `json:"program"`
	Arguments []string `json:"arguments"`
}

func main() {
	args := os.Args[1:]
	
	sets := getConfigContent()

	executeSelectedSets(sets, args)
}

func executeSelectedSets(sets []set, selectedSets []string) {
	for i:=0; i < len(sets); i++ {
		if _, exists := find(selectedSets, sets[i].Name); exists == false {
			continue
		}
		fmt.Println(sets[i].Name)
		for j:=0; j<len(sets[i].Programs) ;j++ {
			fmt.Println("\t"+sets[i].Programs[j].ProgramPath)
			for k:=0; k<len(sets[i].Programs[j].Arguments); k++ {
				fmt.Println("\t\t"+sets[i].Programs[j].Arguments[k])
			}
			cmd := exec.Command(
				sets[i].Programs[j].ProgramPath,
				sets[i].Programs[j].Arguments...
			)
			if err := cmd.Run(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Opend Project succesfully!")
			}
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