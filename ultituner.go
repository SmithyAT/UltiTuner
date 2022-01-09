package main

import (
	"context"
	"fmt"
	"github.com/jbowes/whatsnew"
	"time"
	"ultituner/cmd"
	"ultituner/version"
)

func main() {
	ctx := context.Background()
	fut := whatsnew.Check(ctx, &whatsnew.Options{
		Slug:      "SmithyAT/UltiTuner",
		Cache:     "ultituner_upd-cache.json",
		Version:   version.Version,
		Frequency: 24 * time.Hour,
	})

	fmt.Printf("UltiTuner %s - Copyright by Smithy (Christian Schmied)\n", version.Version)

	cmd.Execute()

	if v, _ := fut.Get(); v != "" {
		fmt.Println()
		fmt.Printf("UPDATE-INFO - New UltiTuner release available: %s\n", v)
		fmt.Println("https://github.com/SmithyAT/UltiTuner/releases/latest")
	}
}
