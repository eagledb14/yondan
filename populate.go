package main

import (
	"github.com/eagledb14/shodan-clone/utils"
	"fmt"
	"os"
	"bufio"
	"time"
)

func Populate(db *utils.ConcurrentMap) {
	dummyRanges := []string{}
	dummyScans := []*utils.Scan{}
	intervalWait := 5

	for range 200 {
		dummyRanges = append(dummyRanges, getRandomCidr())
	}

	ranges := readRanges()

	go func() {
		utils.Poll(ranges, db)

		//loads custom flag data
		flagData := readFlagData(db)
		for _, data := range flagData {
			dummyScans = append(dummyScans, data)
		}

		// loads dummy data
		for _, cidr := range dummyRanges {
			dummyData := createDummyData(cidr, db)
			for _, data := range dummyData {
				dummyScans = append(dummyScans, data)
			}
		}

		populateExamples(db)

		fmt.Println("Full Scan Complete: db size", db.Len())
		time.Sleep(time.Duration(intervalWait) * time.Minute)

		for {
			updateDummyRangeTime(dummyScans)

			fmt.Println("Scanning")
			utils.Poll(ranges, db)
			fmt.Println("Full Scan Complete: db size", db.Len())
			time.Sleep(time.Duration(intervalWait) * time.Minute)
		}
	}()
}

func readRanges() []string {
	file, err := os.Open("./resources/ranges.txt")

	if err != nil {
		panic("Missing file \"resources/ranges.txt\"")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ranges := []string{}

	for scanner.Scan() {
		ranges = append(ranges, scanner.Text())
	}

	return ranges
}

func populateExamples(db *utils.ConcurrentMap) {
	createDummyData("8.8.8.8/24", db)
	createDummyData(getRandomCidr(), db, "example.com")
}

