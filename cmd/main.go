package main

import (
	"fmt"
	"os"

	"github.com/SteveHan-233/MTS-to-mp4/pkg/traverser"
)

func main() {
	list := traverser.GetConversionList(
		os.DirFS("/Users/stevehan/Desktop"),
		"家传照片",
		"dest",
	)
	fmt.Println(list)
}
