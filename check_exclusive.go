package main

func checkExclusive(arr []int, inclusive []int) bool {
	mask := makeMask(arr)
	inclusiveMask := makeMask(inclusive)
	if mask&inclusiveMask == mask {
		return true
	}

	if len(arr) != 1 {
		return false
	}

	exclusiveMask := ^inclusiveMask
	return mask&exclusiveMask == mask
}

func makeMask(arr []int) int {
	mask := 0
	for _, i := range arr {
		mask |= 1 << i
	}

	return mask
}
