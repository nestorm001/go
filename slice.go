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
}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}