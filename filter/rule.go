package filter

import (
	"fmt"
	"regexp"
	"strings"
)

func ConfigToRegexp(config StringMap) ([]*regexp.Regexp, []*regexp.Regexp) {
	regexpList := make([]*regexp.Regexp, 0)
	ignoreRegexpList := make([]*regexp.Regexp, 0)
	for k, v := range config {
		if len(v) == 0 {
			r, _ := regexp.Compile(fmt.Sprintf("%s/.+", k))
			regexpList = append(regexpList, r)
		} else {
			all := true
			for i := 0; i < len(v); i++ {
				item := v[i]
				if strings.HasPrefix(item, "~") {
					r, _ := regexp.Compile(fmt.Sprintf("%s/%s$", k, v[i][1:]))
					ignoreRegexpList = append(ignoreRegexpList, r)
				} else {
					r, _ := regexp.Compile(fmt.Sprintf("%s/%s$", k, v[i]))
					regexpList = append(regexpList, r)
					all = false
				}
			}
			if all {
				r, _ := regexp.Compile(fmt.Sprintf("%s/.+", k))
				regexpList = append(regexpList, r)
			}
		}
	}
	return regexpList, ignoreRegexpList
}
