package pkg

import (
	"bufio"
	"os"
)

func ReadLine(fontName string) ([]string, error) {
	if fontName == "" {
		fontName = "standard"
	}

	fileName := "./fonts/" + fontName + ".txt"
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0, 855)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
