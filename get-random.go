package main

import (
	"math/rand"
)

func GetNumber() (err error, num int) {
	num = rand.Int()
	err = nil
	//err = errors.New("My cool error")
	return
}
