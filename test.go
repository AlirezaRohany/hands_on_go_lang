package main

import "fmt"

func main() {
	/* This is my first sample program. */
   fmt.Println("Hello, World!!!!! \n")


   var x float64
   x = 20.0
   fmt.Println(x)
   fmt.Printf("x is of type %T\n \n", x)


   var z float64 = 20.0
   y := 42 
   fmt.Println(z)
   fmt.Println(y)
   fmt.Printf("z is of type %T\n", z)
   fmt.Printf("y is of type %T\n \n", y)	


   var a, b, c = 3, 4, "foo"  
   fmt.Println(a)
   fmt.Println(b)
   fmt.Println(c)
   fmt.Printf("a is of type %T\n", a)
   fmt.Printf("b is of type %T\n", b)
   fmt.Printf("c is of type %T\n \n", c)


   fmt.Println("Hello\tWorld!","\n")


   const LENGTH int = 10
   const WIDTH int = 5   
   var area int
   area = LENGTH * WIDTH
   fmt.Printf("value of area : %d", area)
   fmt.Printf("\n \n")


// pointers at all
   g:=4
   ptr:= &g
   fmt.Println(g)
   fmt.Println(ptr)
   fmt.Println(*ptr)
   var h int =85
   var hp *int
   hp = &h
   fmt.Println(*hp)
   fmt.Println(" ")


   var s int = 100
   if( s == 10 ) {
      fmt.Printf("Value of s is 10\n" )
   } else if( s == 20 ) {
      fmt.Printf("Value of s is 20\n" )
   } else if( s == 30 ) {
      fmt.Printf("Value of s is 30\n" )
   } else {
      fmt.Printf("None of the values is matching\n" )
   }
   fmt.Printf("Exact value of s is: %d\n\n", s )


   var p int = 100
   var o int = 200
   if( p == 100 ) {
      if( o == 200 )  {
         fmt.Printf("Value of p is 100 and o is 200\n" );
      }
   }
   fmt.Printf("Exact value of p is : %d\n", p );
   fmt.Printf("Exact value of o is : %d\n\n", o );


   var grade string
   var marks int = 90
   switch marks {
      case 90: grade = "A"
      case 80: grade = "B"
      case 50,60,70 : grade = "C"
      default: grade = "D"  
   }
   switch {
      case grade == "A" :
         fmt.Printf("Excellent!\n" )     
      case grade == "B", grade == "C" :
         fmt.Printf("Well done\n" )      
      case grade == "D" :
         fmt.Printf("You passed\n" )      
      case grade == "F":
         fmt.Printf("Better try again\n" )
      default:
         fmt.Printf("Invalid grade\n" );
   }
   fmt.Printf("Your grade is  %s\n\n", grade ); 


   var w string  
   switch i := interface{}(w).(type) {
      case nil:	  
         fmt.Printf("type of w :%T",i)                
      case int:	  
         fmt.Printf("w is int")                       
      case float64:
         fmt.Printf("w is float64")           
      case func(int) float64:
         fmt.Printf("w is func(int)")                      
      case bool, string:
         fmt.Printf("w is bool or string")       
      default:
         fmt.Printf("don't know the type")     
   }
   fmt.Println("")  


   // channel


    

}