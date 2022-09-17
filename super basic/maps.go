package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	houses := map[string][]string{
		"gryffindor": {"weasley", "hagrid", "dumbledore", "lupin"},
		"hufflepuf":  {"wenlock", "scamander", "helga", "diggory"},
		"ravenclaw":  {"flitwick", "bagnold", "wildsmith", "montmorency"},
		"slytherin":  {"horace", "nigellus", "higgs", "scorpius"},
		"bobo":       {"wizardry", "unwanted"},
		"fake":       {},
	}

	delete(houses, "bobo")

	args := os.Args[1:]

	house, students := args[0], houses[args[0]]
	if students == nil {
		fmt.Printf(" cant find %q house", house)
		return
	}

	dups := append([]string(nil), students...)
	sort.Strings(dups)

	fmt.Printf("%s:\n", house)
	for _, student := range dups {
		fmt.Printf("\t+ %s\n", student)
	}
}
