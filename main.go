package main

import (
	"fmt"
	"go-ssc/gossc"
)

func main() {
	ssc := gossc.InitSsc()
	fmt.Printf("ssc version: %d\nbuild info: %s", ssc.Version, ssc.BuildInfo)
}
