package WormholeWilliam

import (
	"strings"

	"github.com/psanford/wormhole-william/wordlist"
)

func ValidateCode(code string) bool {
	var split = strings.Split(code, "-")

	if len(split) != 3 {
		return false
	}

	var odd = split[1]
	var even = split[2]
	var evenMatches = false
	var oddMatches = false

	for _, value := range wordlist.RawWords {
		if value.Even == even {
			evenMatches = true
		}
		if value.Odd == odd {
			oddMatches = true
		}
		if evenMatches && oddMatches {
			break
		}
	}

	return evenMatches && oddMatches
}
