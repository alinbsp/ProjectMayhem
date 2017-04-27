package main

import (
	"os"
	"encoding/json"
	"log"
	"fmt"

	"github.com/ProjectMayhem/client/backend-go/store"
)

const INSTRUCTIONS = "You must provide the path of two project directories like so:\n" +
		"```\n" +
		"go run main.go /home/mayhem/root1 /home/mayhem/root2\n" +
		"```"

func main() {
	projPath1, projPath2 := readFlags()

	projState1 := store.GetLocalState(projPath1)
	projState2 := store.GetLocalState(projPath2)

	diffState := projState1.GetDiffState(&projState2)

	_ = printStateAsJson(diffState)
}

func readFlags() (string, string) {
	if len(os.Args) != 3 {
		panic(INSTRUCTIONS)
	}
	return os.Args[1], os.Args[2]
}

func printStateAsJson(state store.State) string {
	jsonBytes, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	jsonString := string(jsonBytes)
	fmt.Println(jsonString)
	return jsonString
}

