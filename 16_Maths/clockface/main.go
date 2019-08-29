package main

import (
	maths "go-learn/16_Maths"
	"os"
	"time"
)

func main() {
	t := time.Now()
	maths.SVGWriter(os.Stdout, t)
}
