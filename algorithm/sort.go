package algorithm

func quickSort(arr []int, head, fin int) {
	if head > fin {
		return
	}
	var pivot = arr[fin]
	i, j := head, head
	for ; j <= fin; j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	quickSort(arr, head, i-2)
	quickSort(arr, i, fin)
}
