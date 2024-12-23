package sort

func quickSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	i, j, pivot := l-1, r+1, arr[l+(r-l)/2]
	for i < j {
		i++
		for arr[i] < pivot {
			i++
		}
		j--
		for arr[j] > pivot {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	quickSort(arr, l, j)
	quickSort(arr, j+1, r)
}

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}
