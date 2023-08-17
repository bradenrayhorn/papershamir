package testutils

import (
	"fmt"
	"regexp"
)

func CreateShareRegexp(expectedLines int, lineSize int, lastLineSize int) *regexp.Regexp {
	fullLine := "(?:[[:alnum:]]{4} ){%d}(?:[[:alnum:]]{4}){%d}"
	partialLine := "(?:[[:alnum:]]{4} ){%d}(?:[[:alnum:]]{2}){%d}"

	reg := "(?m)"

	// add full lines
	for i := 0; i < expectedLines-1; i++ {
		reg += fmt.Sprintf(
			"^%s$\n",
			fmt.Sprintf(fullLine, (lineSize+8)/4-1, 1),
		)
	}

	// add last pairs line
	if (lastLineSize+8)%4 == 0 {
		reg += fmt.Sprintf(
			"^%s$\n",
			fmt.Sprintf(fullLine, ((lastLineSize+8)/4)-1, ((lastLineSize+8)%4)/2+1),
		)
	} else {
		reg += fmt.Sprintf(
			"^%s$\n",
			fmt.Sprintf(partialLine, ((lastLineSize+8)/4), ((lastLineSize+8)%4)/2),
		)
	}

	// add final checksum line
	reg += fmt.Sprintf(
		"^%s$",
		fmt.Sprintf(fullLine, 1, 1),
	)

	return regexp.MustCompile(reg)
}
