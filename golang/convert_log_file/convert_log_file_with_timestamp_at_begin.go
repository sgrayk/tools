package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func convertDateFromTimestampString(timestampString string) string {
	timestamp, err := strconv.ParseInt(timestampString, 10, 64)
	if err != nil {
		panic(err)
	}

	date := time.UnixMilli(timestamp)

	return date.Format("15:04:05.000")
}

func processLogLine(line string) string {
	var re = regexp.MustCompile(`^\d{13}`)
	timestampString := re.FindString(line)

	dateString := convertDateFromTimestampString(timestampString)

	resultLine := strings.Replace(line, timestampString, dateString, 1)

	return resultLine + "\n"
}

func excuteTool(fileName string) {
	// Open file for reading
	readFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	// Creat a file for writing
	resultFileName := fmt.Sprintf("result_%v", fileName)
	resultFile, err := os.Create(resultFileName)
	if err != nil {
		panic(err)
	}
	defer resultFile.Close()

	writer := bufio.NewWriter(resultFile)
	defer writer.Flush()

	// Scan the read file by line and then process that line
	scanner := bufio.NewScanner(readFile)
	for scanner.Scan() {
		writer.WriteString(processLogLine(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	args := os.Args

	if len(args) > 1 {
		excuteTool(args[1])
	} else {
		fmt.Println("You should run this tools with command: /> go run <tools name> <file name>")
	}
}
