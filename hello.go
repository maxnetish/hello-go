package main

import (
	"fmt"
	"github.com/maxnetish/hello/utils"
	"log"
	"math/rand"
	"os"
	"time"
)

var nanounix = time.Now().UnixNano()

func main() {
	// init rand
	rand.Seed(nanounix)

	nowDate := time.Now()
	uid := os.Getuid()
	gid := os.Getgid()
	environments := utils.ReadEnvironmentVariables()

	// Показываем в std out
	fmt.Printf("Hello Go world\n")
	fmt.Printf("Now is %v\n", nowDate)
	fmt.Printf("Your id is %v\n", uid)
	fmt.Printf("Your gid is %v\n", gid)

	for iter := 0; iter < 3; iter++ {
		if err, rnd := GetNumber(); err == nil {
			defer fmt.Printf("Got random number #%v: %v\n", iter, rnd)
		} else {
			log.Print(err)
			break
		}
	}

	fmt.Printf("Platform is: %v\n", utils.PlatformInfo())
	fmt.Printf("Another take of platform info: %v\n\n", utils.AnotherPlatformInfo())

	fmt.Printf("* Environments *\n")
	fmt.Printf("================\n")
	for key, val := range environments {
		fmt.Printf("%v: \"%v\"\n", key, val)
	}
}
