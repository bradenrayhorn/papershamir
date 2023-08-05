package testutils

import (
	"fmt"
	"regexp"
)

func CreateShareRegexp(expectedLines int, lineSize int, lastLineSize int) *regexp.Regexp {
	checksum := "[[:alnum:]]{8}"
	pairs := "(?:[[:alnum:]]{2} ){%d}"

	reg := "(?m)"

	// add full lines
	for i := 0; i < expectedLines-1; i++ {
		reg += fmt.Sprintf(
			"^%s%s$\n",
			fmt.Sprintf(pairs, lineSize),
			checksum,
		)
	}

	// add last pairs line
	reg += fmt.Sprintf(
		"^%s[ ]{%d}%s$\n",
		fmt.Sprintf(pairs, lastLineSize),
		(lineSize-lastLineSize)*3,
		checksum,
	)

	// add final checksum line
	reg += fmt.Sprintf(
		"^[ ]{%d}%s$",
		lineSize*3,
		checksum,
	)

	return regexp.MustCompile(reg)
}
