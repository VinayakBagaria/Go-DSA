// https://leetcode.com/problems/find-in-mountain-array/description/
package random

type MountainArray struct {
}

func (mountainArray *MountainArray) get(index int) int {
	return -1
}

func (mountainArray *MountainArray) length() int {
	return -1
}

func findPeak(mountainArr *MountainArray) int {
	l := 0
	r := mountainArr.length() - 1

	for l < r {
		mid := l + (r-l)/2
		if mountainArr.get(mid) < mountainArr.get(mid+1) {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}

func binSearch(mountainArr *MountainArray, index, target int) int {
	l := 0
	r := index

	for l <= r {
		mid := l + (r-l)/2
		element := mountainArr.get(mid)
		if element == target {
			return mid
		}

		if target > element {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}

func reverseBinSearch(mountainArr *MountainArray, index, target int) int {
	l := index
	r := mountainArr.length() - 1

	for l <= r {
		mid := l + (r-l)/2
		element := mountainArr.get(mid)
		if element == target {
			return mid
		}

		if target < element {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}

func findInMountainArray(target int, mountainArr *MountainArray) int {
	peak := findPeak(mountainArr)

	result := binSearch(mountainArr, peak, target)
	if result != -1 {
		return result
	}

	return reverseBinSearch(mountainArr, peak+1, target)
}
