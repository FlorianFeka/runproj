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
	for _, set := range sets {
		if _, exists := find(selectedSets, set.Name); exists == false {
			continue
		}
		fmt.Println(set.Name)
		for _, program := range set.Programs {
			fmt.Println("\t"+program.ProgramPath)
			for _, argument := range program.Arguments {
				fmt.Println("\t\t"+argument)
			}
			cmd := exec.Command(
				program.ProgramPath,
				program.Arguments...
			)
			if err := cmd.Run(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Opened Program succesfully!")
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