package support

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"strings"
)

type ConceptSegmenter struct {
}

var strMap map[string]int

func SegmentConceptCamelcase(methodsName []string) map[string]int {
	// get, set
	strMap = make(map[string]int)
	for _, name := range methodsName {
		fmt.Println(name)
		if strings.HasSuffix("set", name) || strings.HasSuffix("get", name) {
			domainName := name[3:]
			if strMap[domainName] == 0 {
				strMap[domainName] = 1
			} else {
				strMap[domainName] = strMap[domainName] + 1
			}
		} else {
			delimited := strcase.ToDelimited(name, '.')
			split := strings.Split(delimited, ".")
			for _, word := range split {
				if strMap[word] == 0 {
					strMap[word] = 1
				} else {
					strMap[word] = strMap[word] + 1
				}
			}
		}
	}

	return strMap
}
