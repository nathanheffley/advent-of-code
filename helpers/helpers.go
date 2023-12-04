package helpers

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func SumSlice(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumMap[K comparable](nums map[K]int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func Permutate[T interface{}](arr []T) [][]T {
	var helper func([]T, int)
	res := [][]T{}

	helper = func(arr []T, n int) {
		if n == 1 {
			tmp := make([]T, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
