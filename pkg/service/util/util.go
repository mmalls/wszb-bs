package util

import "strconv"

// CvtDef ...
func CvtDef(str string, defval int) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		i = defval
	}
	return i
}
