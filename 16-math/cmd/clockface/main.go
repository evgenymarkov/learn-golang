package main

import (
	"os"
	"time"

	"github.com/evgenymarkov/learn-golang/16-math/pkg/clockface"
)

func main() {
	currentTime := time.Now()
	clockface.RenderClock(os.Stdout, currentTime)
}
