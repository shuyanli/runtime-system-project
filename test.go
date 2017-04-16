package main

import "fmt"		//import format package 

func main(){
	// fmt.Println("this is the first go")


	// var num1  int = 10
	// //defalt: var num1 = 10;
	// var num2 float64 = 3.1415
	// num3 := 20		//another way to define integer

	// fmt.Println(num1, num2,num3)
	// fmt.Println(num1+num3)

	// var(
	// 	n1 = 11.1;
	// 	n2 = 12.2;
	// 	)
	// fmt.Println(n1 + n2)

	// //con be redefined
	// n1 = 10;
	// n2 = 20;

	// fmt.Println(n1 + n2)


	// //defining string

	// var s1 = "play with string"
	// fmt.Println("\nthe length of the string is :\n");
	// fmt.Println(len(s1))

	// const pi float64 = 3.1415926
	// fmt.Printf("%T \n", pi) //show the datatype
	// const pi float64 = 3.141592658979
	// fmt.Printf("%d \n", 100) //int
	// fmt.Printf("%b \n", 100) //binary representation
	// fmt.Printf("%c \n", 100) //charactor code
	// fmt.Printf("%x \n", 100) //hex code
	// fmt.Printf("%e \n", pi) //scientific notation


	// //========= loop =========

	// i := 0

	// for i < 10{
	// 	fmt.Println(i)
	// 	i++
	// 	}

	// for j:=0; j <10; j++{
	// 	fmt.Println(j)
	// }

	//========== switch =====
	// marker := 8
	// switch marker{
	// 	 case 16 : fmt.Printf("16 case")
	// 	 case 8 : fmt.Println("8 case")
	// 	 default : fmt.Println ("gg")
	// }

//========== arrays =========
	// myArray := [5]float64 {1.1,2.2,3.3,4.4,5.5}
	// // fmt.Println(myArray[3])

	// // for i:=0;i<5;i++{
	// // 	fmt.Println(myArray[i])
	// // }

	// for i, val := range myArray{  //index, value
	// 	fmt.Println(i, ":", val)
	// }

	// more popular: slice

	mySlice :=[]int {1,2,3,4,5} 
	mySlice2 := mySlice[3:] 
	fmt.Println(mySlice2[0])
	fmt.Println("mySlice[2:] is :", mySlice[2: ])
//The length of a slice is the number of elements it contains.

//The capacity of a slice is the number of elements in the underlying array, 
//counting from the first element in the slice.
	mySlice3 := make([]int, 5,10)
	fmt.Println(mySlice3[0:]) //0,0,0,0,0
	copy(mySlice3, mySlice)
	fmt.Println(mySlice3[0:])  //1,2,3,4,5

//keyword append:
	mySlice3  = append(mySlice3, -1). //add -1 to the end of slice
	fmt.Println(mySlice3[:])


}