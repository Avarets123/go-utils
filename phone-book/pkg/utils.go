package pkg

func JoinArraysToSlice[T any](arr1, arr2 [12]T) []T {
	res := make([]T, len(arr1)+len(arr2))
	res = append(res, arr1[:]...)
	res = append(res, arr2[:]...)
	return res

}

func JoinSlicesToArray[T any](slice1, slice2 []T) [24]T {

	all := append(slice1, slice2...)
	res := [24]T{}

	for i := range 24 {
		res[i] = all[i]
	}

	return res

}

func JoinArraysToArray[T any](arr1, arr2 [12]T) [24]T {
	res := [24]T{}
	i := 0
	for _, v := range arr1 {
		res[i] = v
		i++
	}
	for _, v := range arr2 {
		res[i] = v
		i++
	}

	return res

}
