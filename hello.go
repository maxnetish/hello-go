package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/maxnetish/hello/utils"
)

func main() {

	nowDate := time.Now()
	uid := os.Getuid()
	gid := os.Getgid()
	environments := utils.ReadEnvironmentVariables()
	longJobProgress1 := make(chan string, 5)
	longJobProgress2 := make(chan string, 5)
	longJobLive1 := true
	longJobLive2 := true
	// Показываем в std out
	fmt.Printf("Hello Go world\n")

	fmt.Printf("Start async go routine...\n")
	go utils.LongJob(1, 200, 5, longJobProgress1)
	go utils.LongJob(2, 100, 10, longJobProgress2)

	fmt.Printf("Now is %v\n", nowDate)
	fmt.Printf("Your id is %v\n", uid)
	fmt.Printf("Your gid is %v\n", gid)

	for iter := 0; iter < 3; iter++ {
		if rnd, err := GetNumber(); err == nil {
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

	// block until longJobProgress receives some data
	for longJobLive1 || longJobLive2 {
		select {
		case longJobProgressState1, live1 := <-longJobProgress1:
			longJobLive1 = live1
			if live1 {
				fmt.Println("\n", longJobProgressState1)
			}
		case longJobProgressState2, live2 := <-longJobProgress2:
			longJobLive2 = live2
			if live2 {
				fmt.Println("\n", longJobProgressState2)
			}
		default:
			fmt.Print(" . ")
		}
		time.Sleep(10 * time.Millisecond)
	}

}
