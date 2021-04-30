package utils

import (
	"math/rand"
	"time"
)

func GenSlice(len, upper int) []int {
	randomIntArr := []int{}
	s := time.Now().UnixNano()
	rand.Seed(s)
	for i := 0; i < len; i++ {
		randomIntArr = append(randomIntArr, rand.Int()%upper)
	}
	return randomIntArr
}
