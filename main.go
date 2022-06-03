package main

import (
	"fmt"
	"go-ssc/ssc"
)

func main() {
	sscVersionInfo := ssc.InitSsc()
	fmt.Printf("ssc version: %d\nbuild info: %s", sscVersionInfo.Version, sscVersionInfo.BuildInfo)

	pvWatts, _ := ssc.CreateModule("pvwattsv8")
	defer pvWatts.Free()
	if pvWatts != nil {
		fmt.Print("Created module")
	}

	entries := ssc.GetModuleEntries()
	for _, entry := range entries {
		fmt.Printf("Found entry: %s\n%s\n\n", entry.Name, entry.Description)
	}
}
