package instruction

import "strconv"

func checkValueIsOnly1AndIsInt(values []string) (int, string, bool) {
	if len(values) != 1 {
		return 0, "Failed: the value is not only 1", false
	}
	intValue, err := strconv.Atoi(values[0])
	if err != nil {
		return 0, "Failed: the value is not a number", false
	}
	return intValue, "", true
}
