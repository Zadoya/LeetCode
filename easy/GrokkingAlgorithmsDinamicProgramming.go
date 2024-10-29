// услвие задачи
//
// Емкость рюкзака 6 фунтов, в моход нужно взять саммые важные предметы. 
// У каждого предмета имеется стоимость; чем она выше, тем важнее предмет:
// вода,   3 фута, 10;
// книга,  1 фута, 3;
// еда,    2 фута, 9;
// куртка, 2 фута, 5;
// камера, 1 фута, 6
// Необходимо найти оптимальный набор предметов.

package main

import (
	"fmt"
)

type Subject struct {
	Name   string
	Weight int
	Price  int
}

type Set struct {
	Name   []string
	Weight int
	Price  int
}

func findOptimalSet(subjects []Subject, capacity int) []string {
	var (
		capacities [2][]Set
	)
	capacities[0] = make([]Set, capacity + 1)
	capacities[1] = make([]Set, capacity + 1)

	curr := 0
	prev := 0
	for _, subject := range subjects {
		for i := subject.Weight; i <= capacity; i++ {
			if capacities[prev][i].Price < capacities[prev][i - subject.Weight].Price + subject.Price {
				capacities[curr][i].Price = capacities[prev][i - subject.Weight].Price + subject.Price
				capacities[curr][i].Weight = capacities[prev][i - subject.Weight].Weight + subject.Weight
				capacities[curr][i].Name = append(capacities[prev][i - subject.Weight].Name, subject.Name)
			} else {
				capacities[curr][i] = capacities[prev][i]
			}
		}
		prev = curr
		curr = curr ^ 1
	}

	return capacities[prev][capacity].Name
}

func main() {
	subjects := []Subject{
		{"water", 3, 10},
		{"book", 1, 3},
		{"food", 2, 9},
		{"jacket", 2, 5},
		{"camera", 1, 6},
	}
	capacity := 6

	fmt.Println(findOptimalSet(subjects, capacity))
}