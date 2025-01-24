package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Truck struct {
	Index    int64
	Start    int64
	End      int64
	Capacity int64
}

type Cargo struct {
	Index    int64
	Arrival  int64
	TruckNum int64
}

type Data struct {
	Trucks []Truck
	Cargo  []Cargo 
}

func scanStrNew(in *bufio.Reader, size int64) ([]int64, error) {
	result := make([]int64, 0, size)

	line, err := in.ReadString('\n')
	if err != nil {
		return nil, err
	}

	line = line[:len(line)-1]
	strs := strings.Split(line, " ")
	if len(strs) != int(size) {
		return nil, fmt.Errorf("incorrect number of arguments in string")
	}
	for _, numStr := range strs {
		num, err := strconv.ParseInt(numStr, 10, 64)
		if err != nil {
			return nil, err
		}
		result = append(result, num)
	}

	return result, nil
}

func scanning() ([]Data, error) {
	var (
		in         *bufio.Reader
		t, n, m    []int64
		err        error
		dataset      []Data
	)

	in = bufio.NewReader(os.Stdin)

	t, err = scanStrNew(in, int64(1))
	if err != nil {
		return nil, err
	}

	dataset = make([]Data, 0, t[0])
	for i := 0; i < int(t[0]); i++ {
		var data Data
		if n, err = scanStrNew(in, int64(1)); err != nil {
			return nil, err
		}

		arrival, err := scanStrNew(in, n[0])
		if err != nil {
			return nil, err
		}
		data.Cargo = make([]Cargo, n[0])
		for j := range arrival {
			data.Cargo[j].Arrival = arrival[j]
			data.Cargo[j].Index = int64(j)
			data.Cargo[j].TruckNum = -1
		}

		if m, err = scanStrNew(in, int64(1)); err != nil {
			return nil, err
		}
		data.Trucks = make([]Truck, m[0])
		for j := 0; j < int(m[0]); j++ {
			truck, err := scanStrNew(in, 3)
			if err != nil {
				return nil, err
			}
			data.Trucks[j].Index = int64(j + 1)
			data.Trucks[j].Start = truck[0]
			data.Trucks[j].End = truck[1]
			data.Trucks[j].Capacity = truck[2]
		}
		dataset = append(dataset, data)
	}

	return dataset, nil
}

type TruckList struct {
	truck Truck
	prev  *TruckList
	next  *TruckList
}

func Remove(tl *TruckList) {
	if tl.prev != nil {
		tl.prev.next = tl.next
	}
	if tl.next != nil {
		tl.next.prev = tl.prev
	}
}

func NewTruckList(trucks []Truck) (*TruckList, *TruckList) {
	var first, curr *TruckList
	
	for i, truck := range trucks {
		if i == 0 {
			first = &TruckList{
				truck: truck,
			}
			curr = first
			continue
		}
		new := &TruckList{
			truck: truck,
			prev: curr,
		}
		curr.next = new
		curr = new
	}
	
	return first, curr
}

func assigneTruck(trucks []Truck, cargo []Cargo) []Cargo {
	startList, endList := NewTruckList(trucks)
	
	for i := range cargo {
		for startList != nil && startList.truck.Start > cargo[i].Arrival {
			if startList == endList {
				startList = nil
				endList = nil
				break
			}
			Remove(startList)
			startList = startList.next
		}

		for endList != nil && endList.truck.End < cargo[i].Arrival {
			if startList == endList {
				startList = nil
				endList = nil
				break
			}
			Remove(endList)
			endList = endList.prev
		}

		if startList != nil {
			startList.truck.Capacity--
			cargo[i].TruckNum = startList.truck.Index
			if startList.truck.Capacity == 0 {
				if startList == endList {
					startList = nil
					endList = nil
					break
				}
				Remove(startList)
				startList = startList.next
			}
		}
	}
	
	return cargo
}

func main() {
	var out  *bufio.Writer

	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	dataset, err := scanning()
	if err != nil {
		fmt.Println("Scanning error:", err)
	}

	for _, data := range dataset {
		slices.SortFunc(data.Cargo, func(a, b Cargo) int {
			return int(a.Arrival - b.Arrival)
		})
		slices.SortFunc(data.Trucks, func(a, b Truck) int {
			if a.Start == b.Start {
				return int(a.Index - b.Index) 
			}
			return int(a.Start - b.Start)
		})
		
		data.Cargo = assigneTruck(data.Trucks, data.Cargo)

		slices.SortFunc(data.Cargo, func(a, b Cargo) int {
			return int(a.Index - b.Index) 
		})
		for _, cargo := range data.Cargo {
			fmt.Fprintf(out,"%v ", cargo.TruckNum)
		}
		fmt.Fprintln(out)
	}
}