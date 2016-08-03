package main

import (
  "fmt"
  "sort"
)

type StatsPrinter struct {
  times []int
}

func (this *StatsPrinter) Print() {
	fmt.Println("Number of Requests: ", len(this.times))
	fmt.Println("Average: ", this.average())
	fmt.Println("Median: ", this.median())
	fmt.Println("Min: ", this.min())
	fmt.Println("Max: ", this.max())
}

func (this *StatsPrinter) average() int {
	if len(this.times) == 0 {
		return 0
	}

	return this.sum() / len(this.times)
}

func (this *StatsPrinter) median() int {
  sort.Ints(this.times)
	if len(this.times) == 0 {
		return 0
	}

	middle := len(this.times) / 2
	result := this.times[middle]
	if len(this.times)%2 == 0 {
		result = (result + this.times[middle-1]) / 2
	}
	return result
}

func (this *StatsPrinter) sum() (total int) {
	for _, number := range this.times {
		total += number
	}
	return total
}

func (this *StatsPrinter) min() int {
	if len(this.times) == 0 {
		return 0
	}

	min := this.times[0]
	for _, number := range this.times {
		if min > number {
			min = number
		}
	}
	return min
}

func (this *StatsPrinter) max() int {
	if len(this.times) == 0 {
		return 0
	}

	max := this.times[0]
	for _, number := range this.times {
		if max < number {
			max = number
		}
	}
	return max
}