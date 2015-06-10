package main

import "fmt"

func main() {
   var numbers1 = make([]int,3,5)
   printSlice(numbers1)
   
   var numbers2 []int
   printSlice(numbers2)
   if(numbers2 == nil){
      fmt.Printf("slice is nil")
   }
   
   /* create a slice */
   numbers := []int{0,1,2,3,4,5,6,7,8}   
   printSlice(numbers)
   
   /* print the original slice */
   fmt.Println("numbers ==", numbers)
   
   /* print the sub slice starting from index 1(included) to index 4(excluded)*/
   fmt.Println("numbers[1:4] ==", numbers[1:4])
   
   /* missing lower bound implies 0*/
   fmt.Println("numbers[:3] ==", numbers[:3])
   
   /* missing upper bound implies len(s)*/
   fmt.Println("numbers[4:] ==", numbers[4:])
   
   numbers3 := make([]int,0,5)
   printSlice(numbers3)
   
   /* print the sub slice starting from index 0(included) to index 2(excluded) */
   number4 := numbers[:2]
   printSlice(number4)
   
   /* print the sub slice starting from index 2(included) to index 5(excluded) */
   number5 := numbers[2:5]
   printSlice(number5)
   
   var numbers6 []int
   printSlice(numbers)
   
   /* append allows nil slice */
   numbers6 = append(numbers6, 0)
   printSlice(numbers6)
   
   /* add one element to slice*/
   numbers6 = append(numbers6, 1)
   printSlice(numbers6)
   
   /* add more than one element at a time*/
   numbers6 = append(numbers6, 2, 3, 4, 5, 6)
   printSlice(numbers6)
   
   /* create a slice numbers1 with double the capacity of earlier slice*/
   numbers7 := make([]int, len(numbers6), (cap(numbers6))*2)
   
   /* copy content of numbers to numbers1 */
   copy(numbers7,numbers6)
   printSlice(numbers7)
}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}