package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

func GetFiles(dir string) ([]string, error) {
	file, err := os.Open(dir)
	if err != nil {
		return nil, fmt.Errorf("Open directory error=%v\n", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	list, _ := file.Readdirnames(0) // 0 to read all files and folders
	sort.Sort(Alphabetic(list))
	return list, nil
}

func ReadFile(file string) string {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
