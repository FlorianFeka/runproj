package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// Set has a name and a number of programs that should execute
type Set struct {
	Name string `json:"name"`
	Programs []Program `json:"programs"`
}

// Program is the program to be executed
type Program struct {
	ProgramPath string `json:"program"`
	Arguments []string `json:"arguments"`
}

// ExecuteSelectedSets executes a set of selected sets
func ExecuteSelectedSets(sets []Set, selectedSets []string) {
	for _, set := range sets {
		if _, exists := Find(selectedSets, set.Name); exists == false {
			continue
		}
		fmt.Println(set.Name)
		for _, program := range set.Programs {
			fmt.Println("\t" + program.ProgramPath)
			for _, argument := range program.Arguments {
				fmt.Println("\t\t" + argument)
			}
			cmd := exec.Command(
				program.ProgramPath,
				program.Arguments...,
			)
			if err := cmd.Run(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Opened Program succesfully!")
			}
		}
	}
}

// GetConfigContent reads the config file and retuns the configured sets
func GetConfigContent() []Set {
	confFile, err := os.Open("./.ignore/runproj.json")

	if err != nil {
		fmt.Println(err)
	}

	defer confFile.Close()

	byteConf, _ := ioutil.ReadAll(confFile)
	var sets []Set

	json.Unmarshal(byteConf, &sets)
	return sets
}
