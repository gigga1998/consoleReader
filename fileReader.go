package consoleReader

import (
	"os"
	"bufio"
	"log"
	"strconv"
)

func FileReadStrings(fileName string) ([]string, error){
	listOfStrings := []string{}
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		listOfStrings = append(listOfStrings, scanner.Text())
	}

	err = file.Close()

	if err != nil {
		return listOfStrings, err
	}

	if scanner.Err() != nil {
		return listOfStrings, scanner.Err()
	}

	return listOfStrings, nil
}


func FileReadInt(fileName string) ([]int64, error) {
	listOfInts := []int64{}
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			return listOfInts, err
		}
		listOfInts = append(listOfInts, num)
	}

	err = file.Close()

	if err != nil {
		return listOfInts, err
	}

	if scanner.Err() != nil {
		return listOfInts, scanner.Err()
	}

	return listOfInts, nil
}
