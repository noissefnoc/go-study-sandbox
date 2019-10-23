package main

import (
	"fmt"
	"regexp"
)

func main() {
	targetStr := "これが置換対象置A換"
	pattern := regexp.MustCompile("置換")

	if pattern.MatchString(targetStr) {
		matchedIdx := pattern.FindAllStringSubmatchIndex(targetStr, -1)

		startIdx := 0

		for _, ary := range matchedIdx {
			matchedStart := ary[0]
			matchedEnd := ary[1]

			fmt.Print(string([]byte(targetStr)[startIdx:matchedStart]))
			fmt.Print("\x1b[31m")
			fmt.Print(string([]byte(targetStr)[matchedStart:matchedEnd]))
			fmt.Print("\x1b[0m")

			startIdx = matchedEnd
		}

		if len(targetStr) > startIdx {
			fmt.Print(string([]byte(targetStr)[startIdx:]))
		}
	}
}
