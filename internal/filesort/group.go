package filesort

import (
	"slices"
	"strings"
)

func GetGroupKeys(groups map[string][]string) []string {
	groupKeys := make([]string, len(groups))
	for k := range groups {
		groupKeys = append(groupKeys, k)
	}
	return groupKeys
}

func SortGroupKeysByDepth(groupKeys *[]string) {
	slices.SortFunc(*groupKeys, cmp)
	slices.Reverse(*groupKeys)
}

func cmp(i string, j string) int {
	i, _ = strings.CutSuffix(i, "/")
	j, _ = strings.CutSuffix(j, "/")
	i, _ = strings.CutPrefix(i, "/")
	j, _ = strings.CutPrefix(j, "/")
	depthI := strings.Count(i, "/")
	depthJ := strings.Count(j, "/")

	if depthI > depthJ {
		return 1
	} else if depthI < depthJ {
		return -1
	} else {
		return 0
	}
}
