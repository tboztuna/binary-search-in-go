package binarysearch

/*
	Binary search using recursion and iteration implemented in Go
*/

func recursiveBinarySearch(array []int, target int, lowIndex int, highIndex int) int {
	if highIndex < lowIndex {
		return -1
	}

	mid := int((lowIndex + highIndex) / 2)

	if array[mid] > target {
		return recursiveBinarySearch(array, target, lowIndex, mid)
	} else if array[mid] < target {
		return recursiveBinarySearch(array, target, (mid + 1), highIndex)
	} else {
		return mid
	}
}

func iteratedBinarySearch(array []int, target int, lowIndex int, highIndex int) int {
	startIndex := lowIndex
	endIndex := highIndex

	var mid int

	for startIndex < endIndex {
		mid = int((startIndex + endIndex) / 2)

		if array[mid] > target {
			endIndex = mid
		} else if array[mid] < target {

		} else {
			return mid
		}
	}
	return -1
}
