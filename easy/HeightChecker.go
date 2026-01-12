// условие задачи - https://leetcode.com/problems/height-checker/

package main

// через быструю сортировку
func partition(arr []int, left, right int) ([]int, int) {
    pivot := arr[right]
    i := left
    for j := left; j < right; j++ {
        if arr[j] < pivot {
            arr[j], arr[i] = arr[i], arr[j]
            i++
        }
    }

    arr[right], arr[i] = arr[i], arr[right]
    return arr, i
}

func internalSortFunc(arr []int, left, right int) []int {
    if left < right {
        arr, pivot := partition(arr, left, right)
        arr = internalSortFunc(arr, left, pivot - 1)
        arr = internalSortFunc(arr, pivot + 1, right)
    }

    return arr
}

func qsort(arr []int) []int {
    return internalSortFunc(arr, 0, len(arr) - 1)
}

func heightCheckerQsort(heights []int) int {
	expected := make([]int, len(heights))
	copy(expected, heights)
    
	expected = qsort(expected)

    counter := 0
    for i := 0; i < len(heights); i++ {
        if expected[i] != heights[i] {
            counter++
        }
    }
    return counter
}

// через сортировку подсчетом
func countingSort(arr []int) []int {
    sorted := make([]int, 0, len(arr))

    max := 0;
    for i := 0; i < len(arr); i++ {
        if max < arr[i] {
            max = arr[i]
        }
    }
    nums := make([]int, max + 1)
    for i := 0; i < len(arr); i++ {
        nums[arr[i]]++
    }

    for i := 0; i < len(nums); i++ {
        for nums[i] > 0 {
            sorted = append(sorted, i)
            nums[i]--
        }
    }

    return sorted
}

func heightCheckerCountingSort(heights []int) int {    
	expected := countingSort(heights)

    counter := 0
    for i := 0; i < len(heights); i++ {
        if expected[i] != heights[i] {
            counter++
        }
    }
    return counter  
}
