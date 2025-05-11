package main

import (
	"strconv"
	"time"
)

func toggleIndex(index string) {
	i := convert(index)
	if validate(i) {
		todos[i].Done = !todos[i].Done
		todos[i].UpdatedAt = time.Now()
	}
}

func convert(index string) int {
	i, _ := strconv.Atoi(index)
	return i
}

func validate(i int) bool {
	return i >= 0 && i < len(todos)
}
