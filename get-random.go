package main

import (
	"math/rand"
	"time"
)

var nanounix = time.Now().UnixNano()

func init() {
	// init rand
	rand.Seed(nanounix)
}

// GetNumber returns random number
func GetNumber() (num int, err error) {
	num = rand.Int()
	err = nil
	//err = errors.New("My cool error")
	return
}
