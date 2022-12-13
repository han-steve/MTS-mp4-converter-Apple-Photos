package traverser

import (
	"fmt"
	"io/fs"
	"path"
)

type SourceDestPair struct {
	Source string
	Dest   string
}

func GetConversionList(
	fileSystem fs.FS,
	source string,
	dest string,
) []SourceDestPair {
	var conversionList []SourceDestPair
	files, err := fs.ReadDir(fileSystem, source)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		newSource := path.Join(source, file.Name())
		newDest := path.Join(dest, file.Name())
		if file.IsDir() {
			conversionList = append(
				conversionList,
				GetConversionList(fileSystem, newSource, newDest)...,
			)
		} else if len(file.Name()) > 4 && file.Name()[len(file.Name())-4:] == ".MTS" {
			conversionList = append(conversionList, SourceDestPair{Source: newSource, Dest: newDest})
		}
	}
	return conversionList
}
