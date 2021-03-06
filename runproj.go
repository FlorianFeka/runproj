package main

import (
	"os"
)


func main() {
	args := os.Args[1:]
	
	sets := GetConfigContent()

	ExecuteSelectedSets(sets, args)
}

