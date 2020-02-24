package main

import (
	"math"
)

const (
	maxNonce = math.MaxInt64
)

func main() {
	cli := CLI{}
	cli.Run()
}
