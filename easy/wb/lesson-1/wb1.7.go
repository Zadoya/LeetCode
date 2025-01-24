package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const seats = 450

type party struct {
	name   string
	voutes int64
	seats  int64
}

func readData() ([]string,map[string]*party, []party, int64) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	
	partyList := make([]string, 0)
	partyName := make(map[string]*party, 0)
	partyNumVoutes := make([]party, 0)
	totalVotes := int64(0)
	i := 0
	for scanner.Scan() && i < 3 {
		line := strings.Split(scanner.Text(), " ")
		num, _ := strconv.ParseInt(line[len(line) - 1], 10, 64)
		name := strings.Join(line[:len(line) - 1], " ")

		partyList = append(partyList, name)
		partyNumVoutes = append(partyNumVoutes, party{name: name, voutes: num})
		partyName[name] = &partyNumVoutes[len(partyNumVoutes) - 1]
		totalVotes += num
		i++
	}

	return partyList, partyName, partyNumVoutes, totalVotes/seats
}

func countSumBoxes(arr [][2]int64) map[int64]int64 {
	set := make(map[int64]int64, len(arr))
	for i := 0; i < len(arr); i++ {
		set[arr[i][0]] += arr[i][1]
	}
	return set
}

func main() {
	partyList, partyMap, party, firstVoutesRule := readData()
	
	totalSeats := int64(0)
	for i := range party {
		party[i].seats = party[i].voutes / firstVoutesRule
		party[i].voutes %= firstVoutesRule
		totalSeats += party[i].seats
	}

	if totalSeats < seats {
		sort.Slice(party, func(i, j int) bool {
			if party[i].voutes == party[j].voutes {
				return party[i].voutes > party[j].voutes
			}
			return party[i].voutes > party[j].voutes
		})
		for i := 0; i < len(party) && totalSeats < seats; i++ {
			party[i].seats++
			totalSeats++
		}
	}

	fmt.Println(partyMap["Party One"])
	fmt.Println(partyMap["Party Two"])
	fmt.Println(partyMap["Party Three"])
	fmt.Println()

	for _, name := range partyList {
		fmt.Fprintln(os.Stdout, partyMap[name].name, partyMap[name].seats)
	} 
}