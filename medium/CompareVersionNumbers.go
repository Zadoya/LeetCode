// Условие задачи - https://leetcode.com/problems/compare-version-numbers/description

func validateVersion(version string) []int {
	substr := strings.Split(version, ".")

    valid := make([]int, 0, len(substr))
	for i := range substr {
		num, _ := strconv.Atoi(substr[i])
		valid = append(valid, num)
	}
	return valid
}

func compareNums(version1 []int, version2 []int) int {
	for i := range version1 {
		fmt.Println(version1[i], version2[i])
		switch {
		case version1[i] < version2[i]:
			return -1
		case version1[i] > version2[i]:
			return 1	
		}
	}
	return 0
}

func compareVersion(version1 string, version2 string) int {
    valid1 := validateVersion(version1)
	valid2 := validateVersion(version2)

    for len(valid1) < len(valid2) {
        valid1 = append(valid1, 0)
    }

    for len(valid1) > len(valid2) {
        valid2 = append(valid2, 0)
    }

    return compareNums(valid1, valid2)
}