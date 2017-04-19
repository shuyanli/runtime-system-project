package main

import "fmt"		//import format package 
import "net/http"   //for server implementation

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

	// ======= more popular: slice ========

// 	mySlice :=[]int {1,2,3,4,5} 
// 	mySlice2 := mySlice[3:] 
// 	fmt.Println(mySlice2[0])
// 	fmt.Println("mySlice[2:] is :", mySlice[2: ])
// //The length of a slice is the number of elements it contains.

// //The capacity of a slice is the number of elements in the underlying array, 
// //counting from the first element in the slice.
// 	mySlice3 := make([]int, 5,10)
// 	fmt.Println(mySlice3[0:]) //0,0,0,0,0
// 	copy(mySlice3, mySlice)
// 	fmt.Println(mySlice3[0:])  //1,2,3,4,5

// //keyword append:
// 	mySlice3  = append(mySlice3, -1). //add -1 to the end of slice
// 	fmt.Println(mySlice3[:])

// maps.  //collection of key value

	// mymap:= make(map[string] int)
	// mymap["a. a"] = 1
	// mymap["b"] = 2

	// fmt.Println(mymap["a. a"])
	// fmt.Println(len(mymap))
	// delete(mymap,"a. a")
	// fmt.Println(len(mymap))
	// fmt.Println(mymap["a. a"])

//==================define a function============
	// myList := []int {3,5,4,2,1}

	// result := addArray(myList);
	//fmt.Println(result)
	//fmt.Println(addArray(myList))


	// num1, num2 := returnTwo(5,7)
	// fmt.Println(num1, num2)


//============ closures ======================

/*define a function in a main without arguments*/
	// num1 := 3

	// doubleIt := func() int{  //doubleIt in this situation is a function
	// 	num1 *= 2
	// 	return num1
	// }

	// fmt.Println(doubleIt())
	// fmt.Println(doubleIt())   //get 6, 12 respectively


//========= recursion =========

// res := fact(5)

// fmt.Println(res)

	//fmt.Println(devideNum(5,0))
	//fmt.Println(devideNum(4,2))
 //a()

//========= pointer and reference =========

	// x := 2

	// passByRef(&x)

	// fmt.Println(x)
	// fmt.Println(&x)
	//fmt.Println(*x)  //error

	// ptr := new(int)

	// passByRef(ptr)

	// fmt.Println(*ptr)




//========== struct===========
//define our own datatype, just like all other languages
 // rect1 := Rectangle{8 ,20}
 // fmt.Println(rect1.getArea())


//======== Interfaces/porlimorphism ========
	// rect :=rectangle{5,6}
	// tran := trangle{3,6}

	// fmt.Println(getArea(rect))
	// fmt.Println(getArea(tran))

//========= start a server ===========
	http.HandleFunc("/", handler1)   // in the main page
	http.HandleFunc("/hello", handler2)  // in /hello directory
	http.ListenAndServe(":8080", nil) //in port 8080
}  //end of main

// func addArray (inputArray []int) int {
// 	sum := 0
// 	for i:=0; i < len(inputArray); i++{
// 		sum += inputArray[i]
// 	}

// 	return sum
// }

//======= return two values in one function!!!! ============(so wierd)

// func returnTwo(n1 int, n2 int)(int, int){
// 	l1 := n1+1;
// 	l2 := n2*2;

// 	return l1, l2
// }
// ========= recursion by doing fact
// func fact(num1 int) int {
// 	if num1 == 1{
// 		return 1
// 	}

// 	return num1 * fact(num1-1)
// }

/*

TODO:
the output for this function is 0 instead of 1
func a() {
    i := 0
    defer fmt.Println(i)
    i++
    return
}

*/


//========= handle exception

// func devideNum(num1 int, num2 int) int{

// 	defer func(){   //handle the exception, in this case, the number 0
// 		fmt.Println(recover())
// 		//recover()
// 	}()		//dont forget this ()
// 	panic (sol := num1 / num2)

	
// 	return sol
// }


//==========pointer and reference ====

// func passByRef(x *int){
// 	*x = 100
// }

// ============ struct ==========
// type Rectangle struct {
// 	width float64 
// 	height float64
// }

// func (rect Rectangle) getArea() float64{
// 	return rect.width * rect.height
// }

//========= Interfaces/polymophism ==========

// type Shape interface{
// 	area() int
// }

// type rectangle struct{
// 	width int
// 	height int
// }

// type trangle struct{
// 	width int
// 	height int
// }

// func (r rectangle) area() int{
// 	return r.width * r.height
// }

// func (t trangle) area() int{
// 	return t.width * t.height / 2
// }

// func getArea(s Shape) int{
// 	return s.area()
// }

//=========== server implementation ====== 

/*
note that there is no exit server function yet. 
update: on mac control+c will shut down the server, not command + c
*/
func handler1(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello my go server")
}

func handler2(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "this goes to welcome you")
}