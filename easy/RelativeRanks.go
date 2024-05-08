// условие задачи - https://leetcode.com/problems/relative-ranks/description

func findRelativeRanks(score []int) []string {
    result := make([]int, len(score)) 
	copy(result, score)
	sortQuick(result, 0, len(result) - 1)
	strRes := make([]string, 0, len(score))
	for i := range score {
		place := binarySearch(result, score[i]) + 1
		switch place {
		case 1:
			strRes = append(strRes, "Gold Medal")
		case 2:
			strRes = append(strRes, "Silver Medal")
		case 3:
			strRes = append(strRes, "Bronze Medal")
		default:
			strRes = append(strRes, strconv.Itoa(place))
		}
	}
	return strRes
}

func partition(array []int, low, high int) ([]int, int) {
	pivot := array[high]

	i := low
	for j := low; j < high; j++ {
		if array[j] > pivot {
			array[i], array[j] = array[j], array[i]
			i++
		}
	}
	array[i], array[high] = array[high], array[i]
	return array, i
}

func sortQuick(array []int, low, high int) []int {
	if low < high {
		array, pivot := partition(array, low, high)
		array = sortQuick(array, low, pivot - 1)
		array = sortQuick(array, pivot + 1, high)
	}
	return array
} 

func binarySearch(array []int, target int) int {
	low := 0
	high := len(array) - 1
	
	for low <= high {
		middle := (high + low) / 2
		if target < array[middle] {
			low = middle + 1
		} else if target > array[middle] {
			high = middle - 1
		} else {
			return middle
		}
	}

	return -1
}