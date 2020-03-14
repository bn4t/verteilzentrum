package main

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// check if a provided mailing list exists
func ListExists(list string) bool {
	for _, b := range Config.Lists {
		if b.Name == list {
			return true
		}
	}
	return false
}
