package sort

import (
	"bufio"
	"os"
)

func SortByMergeFile() {
	sortFile, err := os.OpenFile("test_sort_file", os.O_RDONLY, 0644)
	defer sortFile.Close()
	if err != nil {
		panic(err)
	}
	tmpFile1, err := os.OpenFile("tmp_file_1", os.O_RDWR|os.O_CREATE, 0644)
	defer tmpFile1.Close()
	if err != nil {
		panic(err)
	}
	tmpFile2, err := os.OpenFile("tmp_file_2", os.O_RDWR|os.O_CREATE, 0644)
	defer tmpFile2.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(sortFile)
	var counter int = 0
	for scanner.Scan() {
		if counter%2 == 0 {
			tmpFile2.WriteString(scanner.Text())
			tmpFile2.WriteString("\n")
		} else {
			tmpFile1.WriteString(scanner.Text())
			tmpFile1.WriteString("\n")
		}
		counter++
	}
}
