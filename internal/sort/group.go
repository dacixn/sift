package filesort

func getGroupKeys(groups map[string][]string) []string {
	groupKeys := make([]string, len(groups))
	for k := range groups {
		groupKeys = append(groupKeys, k)
	}
	return groupKeys
}

func sortGroupKeysByDepth(groups map[string][]string) {
	// slices.SortFunc
}

func depth(path string)

func less(i int, j int) bool
