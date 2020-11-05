package main

import "fmt"

// 互换方式

func Qsort(left, right int, arr *[10]int) {
	var  tmp int
	l:=left
	r:=right
	middle := arr[(right + left) / 2]


	for; l < r; {
		for ;arr[l] < middle; {
			l +=1
		}
		for ;arr[r] > middle; {
			r -=1
		}
		if l >= r {
			break
		}

		// 交换left和right的值
		tmp = arr[l]
		arr[l] = arr[r]
		arr[r] = tmp

		if arr[l] == middle {
			l+=1
		}
		if arr[r] == middle {
			r-=1
		}
	}
	if l == r {
		l +=1
		r -=1
	}

	if left < r {
		Qsort(left, r, arr)
	}
	if right > l {
		Qsort(l, right, arr)
	}

}

func main() {
	arr := [10]int{-1,9,1,2,10,2,-4,8,5,6}
	Qsort(0,len(arr)-1, &arr)
	fmt.Println(arr)
}