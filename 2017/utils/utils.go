package utils

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	u "github.com/juamms/go-utils"
)

// Day represents each day in the Advent Calendar
type Day interface {
	Part1() interface{}
	Part2() interface{}
}

// ParseFlags parses the given flags from the command line
func ParseFlags(day, part *int) {
	flag.IntVar(day, "d", 0, "The day to run (0 or omit to run all)")
	flag.IntVar(part, "p", 0, "The part of the day to run (0 or omit to run both parts)")

	flag.Parse()
}

// Handle handles errors
func Handle(err error) {
	if err != nil {
		panic(err)
	}
}

// GetInputFile returns the input file for the given `day`
func GetInputFile(day int) (*os.File, error) {
	path, err := u.GetExecutablePath()
	if err != nil {
		return nil, err
	}

	inputPath := fmt.Sprintf("inputs/%d.txt", day)
	fPath := u.SafeJoinPaths(path, inputPath)

	return os.Open(fPath)
}

// GetInputString returns the input for the given `day` as a string
func GetInputString(day int) (string, error) {
	file, err := GetInputFile(day)
	if err != nil {
		return "", err
	}

	contents := new(bytes.Buffer)
	_, err = io.Copy(contents, file)

	if err != nil {
		return "", err
	}

	return string(contents.Bytes()), nil
}

// GetInputIntSlice returns the input for the given `day` as an int slice, from the original string split by `sep`.
func GetInputIntSlice(day int, sep string) ([]int, error) {
	str, err := GetInputString(day)

	if err != nil {
		return nil, err
	}

	result := make([]int, 0)
	arr := strings.Split(str, sep)

	for _, s := range arr {
		result = append(result, u.StringToInt(s, 0))
	}

	return result, nil
}

// CircularGetInt wraps the given `index` around, making `slice` a circular slice and returning the element in the `index` % len(`slice`) position
func CircularGetInt(slice []int, index int) int {
	return slice[index % len(slice)]
}