package main

import (
	"fmt"
	"time"
)

func castomSleep(t time.Duration) {
	timer := time.NewTimer(t)
	<-timer.C
}

func main() {
	fmt.Println("Start")
	castomSleep(5 * time.Second)
	fmt.Println("End")
}
