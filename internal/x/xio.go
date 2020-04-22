package x

import (
	"bufio"
	"os"
	"regexp"

	"github.com/sirupsen/logrus"
)

const (
	tagDefFile  = "src/test/java/test/support/Category.java"
	linePattern = `^\s+public\s+static\s+final\s+String\s+\w+\s+=\s+"([^"]+)";\s*$`
)

var (
	lineMatch = regexp.MustCompile(linePattern)
)

func ReadConcreteTags() []string {
	file, err := os.Open(tagDefFile)
	out := make([]string, 0, 17)

	if err != nil {
		logrus.Fatal(err)
	}

	scan := bufio.NewScanner(file)

	for scan.Scan() {
		line := scan.Bytes()
		match := lineMatch.FindSubmatch(line)

		if len(match) == 0 {
			continue
		}

		out = append(out, string(match[1]))
	}

	return out
}
