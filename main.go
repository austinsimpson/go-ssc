package main

import (
	"fmt"
	"go-ssc/gossc"
)

func main() {
	ssc := gossc.InitSsc()
	fmt.Printf("ssc version: %d\nbuild info: %s", ssc.Version, ssc.BuildInfo)

	pvWatts, _ := gossc.CreateModule("pvwattsv8")
	defer pvWatts.Free()
	if pvWatts != nil {
		fmt.Print("Created module")
	}

	entries := gossc.GetModuleEntries()
	for _, entry := range entries {
		fmt.Printf("Found entry: %s\n%s\n\n", entry.Name, entry.Description)
	}
}
