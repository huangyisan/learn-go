package main

import "fmt"
// 填坑方式实现

func quick_sort(arr *[10]int) [10]int {
	return qSort(0, len(arr)-1, arr )
}

func qSort(left,right int, arr *[10]int) [10]int {
	if left < right {
		pos := get_pos(arr, left, right)
		qSort(left, pos-1, arr)
		qSort(pos+1, right, arr)
	}
	return *arr
}

func get_pos(arr *[10]int, left,right int) int {
	pos := arr[left]
	for; left < right; {
		for; left < right && arr[right] >= pos; {
			right -=1
		}
		arr[left] = arr[right]
		for; left < right && arr[left] <= pos; {
			left +=1
		}
		arr[right] =  arr[left]
	}
	arr[left] = pos
	return left
}


func main() {
	arr := [10]int{-1,9,1,2,2,-3,-4,8,5,6}
	newArray := quick_sort(&arr)
	fmt.Println(newArray)

}