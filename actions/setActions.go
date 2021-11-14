package actions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"sort"
	"syscall"

	"github.com/FlorianFeka/runproj/data"
	"github.com/FlorianFeka/runproj/utils"
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

func ExecuteSet(set *data.Set) error {
	if isInDockerContainer() {
		err := executeFromDockerContainer()
		return err
	}

	json, err := json.Marshal(&set);
	if err != nil {
		return err
	}

	fmt.Printf("Execute set: \n%s", json)
	for _, programSet := range set.ProgramSets {
		cmd := exec.Command(
			programSet.Program.ProgramPath,
			GetArguments(programSet.Arguments)...,
		)
		if err := cmd.Run(); err != nil {
			return err;
		}
	}

	return nil
}

func GetArguments(arguments []*data.Argument) []string {
	var argStrings []string
	sort.SliceStable(arguments, func(i, j int) bool {
		return arguments[i].Order < arguments[j].Order
	})

	for _, argument := range arguments {
		argStrings = append(argStrings, argument.Argument)
	}

	return argStrings;
}

// ExecuteSelectedSets executes a set of selected sets
func ExecuteSelectedSets(sets []Set, selectedSets []string) {
	for _, set := range sets {
		if _, exists := utils.Find(selectedSets, set.Name); !exists {
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

func executeFromDockerContainer() error {
	syscall.Mkfifo("tmp", 0666)
	_, err := os.OpenFile("tmp", os.O_RDWR, os.ModeNamedPipe)
	if err != nil {
		return err
	}
	fmt.Println("Execute from container!")
	return nil
}

func isInDockerContainer() bool {
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}
	return false
}
