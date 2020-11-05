package main

import "fmt"

func main()  {
/*
    切片叫做变长数组，长度不固定
        len（），cap()

    当向切片中添加数据时，如果没有超过容量，直接添加，如果超过容量，自动扩容(成倍增长)
    3-->6-->12-->24-->48
    4-->8-->16-->32
     */
arr := [10]int{1,2,3,4,5,6,7,8,9,10}
fmt.Println(arr)
fmt.Printf("%p\n",&arr) //打印输出数组自己的地址

s1 := arr[:5]
fmt.Println(s1)
fmt.Printf("%p,长度：%d，容量：%d\n",s1,len(s1),cap(s1))
s2 := arr[2:7]
fmt.Println(s2)
//cap为8. 10-开始2的索引
fmt.Printf("%p,长度：%d，容量：%d\n",s2,len(s2),cap(s2))

//修改切片的数据,其实是修改该切片所指向的底层数组的数据。
//导致所有指向该数组的切片的数据，都会改变。
s1[3] = 100
fmt.Println(arr)
//此时arr为[1 2 3 100 5 6 7 8 9 10]
fmt.Println(s1)
//[1 2 3 100 5]
fmt.Println(s2)
//[3 100 5 6 7]


//append(),追加数据，修改了原数组内容[6:10]
s1 = append(s1,1,1,1,1) //s1的长度：5，容量：10
fmt.Println(arr) //[1 2 3 100 5 1 1 1 1 10]
fmt.Println(s1) //[1 2 3 100 5 1 1 1 1]
fmt.Println(cap(s1)) //[1 2 3 100 5 1 1 1 1]
fmt.Println(s2) //[3 100 5 1 1]
fmt.Printf("%p\n",&s2) //

//追加的时候，涉及到了扩容,会更改切片指向的底层数组，意味着不再指向原先底层数组，产生了一个全新的底层数组
//s2:len5,cap8
s2= append(s2, 2,2,2,2,2,3,3,3,3,3)
//更改s2指向的底层数组
fmt.Println(arr)//[1 2 3 100 5 1 1 1 1 10]
fmt.Println(s1) //[1 2 3 100 5 1 1 1 1]
fmt.Println(s2) //[3 100 5 1 1 2 2 2 2 2 3 3 3 3 3]
fmt.Println(cap(s2)) //16, 因为超出了原先自身的cap，原先为8,成倍扩容8*2 -> 16
fmt.Printf("%p\n",&s2) //

//修改底层数组arr[0]为9,验证底层数组指向
arr[0] = 9
//s1切片指向原先数组，因为s1的append并没有导致扩容
fmt.Println(s1) //[9 2 3 100 5 1 1 1 1]
//s2的切片已经不在指向原先数组，因为s2的append导致s2扩容
fmt.Println(s2) //[3 100 5 1 1 2 2 2 2 2 3 3 3 3 3]


}