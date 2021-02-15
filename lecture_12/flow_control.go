package main

import "fmt"

func main() {
         
       fmt.Printf("Enter your age: ")
       var age int
       fmt.Scanf("%d" , &age)
			
       if age < 3 {
               fmt.Println("infant")
               }else if age>2 && age<13  {//2 to 12
               fmt.Println("children")

                  }else if age>12 && age <= 19{
                  fmt.Println("teenage")
                  }else{
                  fmt.Println("Adult")
                  }
}