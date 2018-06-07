package util

// ContainsInt is a utility function to see
// if value "e" is contained in int array i
func ContainsInt(i []int, e int) bool {
	for _, a := range i {
		if a == e {
			return true
		}
	}
	return false
}

// ContainsString is a utility function to see
// if value "e" is contained in string array s
func ContainsString(s []string, v string) bool {
	for _, str := range s {
		if str == v {
			return true
		}
	}
	return false
}

// Index gives the index of a string "t" in slice "vs"
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}
