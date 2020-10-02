package main

import "fmt"
import "math"
import "strings"
import "errors"

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
   fmt.Println("Pointers:")
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
   fmt.Println("\n")  


   // channel


   var j int
   t:=15
   for j = 0; j < 10; j++ {
      fmt.Printf("value of j: %d\n", j)
   }
   for j <= t {
      fmt.Printf("value of j: %d\n", j)
      j++
   }
   numbers := [7]int{10, 20, 3, 15,1000,60} 
   for n,m := range numbers{
   	fmt.Printf("m=%d at %d\n",m,n)
   }
   fmt.Println("")
   

   number10:=10
   number15:=15
   fmt.Println(max(number15,number10),"\n")


   s1,s2:=swap("ali","reza")
   fmt.Println(s1,s2,"\n")


   // call by value and call by refrence functions


   getSquareRoot := func(x float64) float64 {
      return math.Sqrt(x)
   }
   fmt.Println(getSquareRoot(289),"\n")


   circle:=Circle{x:0,y:0,radius:6}
   fmt.Println(circle.area(),"\n")


   var greeting =  "Hello,Hello!"
   fmt.Printf("normal string: ")
   fmt.Printf("%s", greeting)
   fmt.Printf("\n")
   fmt.Printf("hex bytes: ")
   for i := 0; i < len(greeting); i++ {
       fmt.Printf("%x ", greeting[i])
   }
   newGreeting:=[]string{greeting,"moreHello!"}
   fmt.Println("")
   fmt.Println(newGreeting)
   veryNewGreeting:=strings.Join(newGreeting,"  space between elements  ")
   fmt.Println(veryNewGreeting,"\n")


   var array =[]int {1000, 2, 3, 17, 50}
   for i:=0;i<len(array);i++{
   	array[i]=100+array[i]
   	fmt.Printf("Element[%d]=%d\n",i,array[i])
   }
   var arrayAverage float32
   arrayAverage=getAverage(array,len(array))
   fmt.Println(arrayAverage,"\n")


   x1:=20
   y1:=185
   swap2(&x1,&y1)
   fmt.Println(x1)
   fmt.Println(y1)
   fmt.Println("")


   var Book1 Book   
   var Book2 Book 
   Book1.title = "Go Programming"
   Book1.author = "Mahesh Kumar"
   Book1.subject = "Go Programming Tutorial"
   Book1.book_id = 6495407
   Book2.title = "Telecom Billing"
   Book2.author = "Zara Ali"
   Book2.subject = "Telecom Billing Tutorial"
   Book2.book_id = 6495700
   printBook(&Book1)
   printBook(&Book2)
   pbook:=&Book1
   fmt.Println(Book1.author,"   ", pbook.author, "   ", (&Book1).author)
   fmt.Println("\n")


   nums:=[]int {0,1,2,3,4,5,6,7,8}
   printSlice(nums)
   fmt.Println(nums)
   fmt.Println(nums[1:4])
   fmt.Println(nums[:3])
   fmt.Println(nums[4:])
   nums1:=make([]int,0,5)
   printSlice(nums1)
   nums2:=nums[:2]
   printSlice(nums2)
   nums3:=nums[2:5]
   printSlice(nums3)
   fmt.Println("\n")
   var numberrs []int 
   printSlice(numberrs)
   numberrs=append(numberrs,0)
   printSlice(numberrs)
   numberrs=append(numberrs,1)
   printSlice(numberrs)
   numberrs=append(numberrs,2,3,4)
   printSlice(numberrs)
   numberrs1:=make([]int,len(numberrs),cap(numberrs)*2)
   numberrs1[0]=140
   numberrs1[4]=200
   printSlice(numberrs1)
   copy(numberrs1,numberrs)
   printSlice(numberrs)
   printSlice(numberrs1)
   fmt.Println("\n")


   numbs := []int{0,1,2,3,4,5,6,7,8}    
   for i:= range numbs {
      fmt.Println("Slice item",i,"is",numbs[i])
   }
   countryCapitalMap := map[string] string {"France":"Paris","Italy":"Rome","Japan":"Tokyo"}
   countryCapitalMap["Iran"]="Tehran"
   for country := range countryCapitalMap {
      fmt.Println("Capital of",country,"is",countryCapitalMap[country])
   }
   for country,capital := range countryCapitalMap {
      fmt.Println("Capital of",country,"is",capital)
   }
   delete(countryCapitalMap,"Japan")
   UScapital, exist := countryCapitalMap["United States"]
   fmt.Println(exist)
   fmt.Println("Us capital:",UScapital)
   IRANcapital, exists:=countryCapitalMap["Iran"]
   fmt.Println(exists)
   fmt.Println("Iran capital:",IRANcapital)
   JAPANcapital, existss:=countryCapitalMap["Japan"]
   fmt.Println(existss)
   fmt.Println("Japan capital:",JAPANcapital)


   fmt.Println("0! =",factoriel(0),"  1! =",factoriel(1),"  2! =",factoriel(2),"  5! =",factoriel(5),"  6! =",factoriel(6),"\n")


   fmt.Println("F0 =",fibonacci(0), "  F1 =",fibonacci(1),"  F2 =",fibonacci(2),"  F9 =",fibonacci(9))
   for i := 0; i < 10; i++ {
   		fmt.Printf("%d ",fibonacci(i))
   }
   fmt.Println("\n")


   var sam int = 17
   var cunt int = 5
   var min float32
   min = float32(sam)/float32(cunt)
   fmt.Printf("Value of min : %f\n\n",min)


   circle2:=Circle{x:0,y:0,radius:5}
   var rect Shape
   rect=Rectangle{10,5}
   fmt.Printf("Circle area: %f\n",getArea(circle2))
   fmt.Printf("Rectangle area: %f\n\n",getArea(rect))


   sqrResult, err:=sqrt(-10)
   if err!=nil{
   	fmt.Println(err)
   }else{
   	fmt.Println(sqrResult)
   }
   sqrResult,err=sqrt(169)
   if err!=nil{
   	fmt.Println(err)
   }else{
   	fmt.Println(sqrResult)
   }

}


func max(num1, num2 int) int {
   var result int
   if (num1 > num2) {
      result = num1
   } else {
      result = num2
   }
   return result 
}


func swap(x, y string) (string, string) {
   return y, x
}


type Circle struct{
	x,y,radius float64
}
func(circle Circle) area() float64{
	return math.Pi*circle.radius*circle.radius
}


func getAverage(arr []int, size int) float32 {
	var sum float32
	var avg float32
	for i := 0; i < size; i++ {
		sum+=float32(arr[i])
	}
	avg=sum/float32(size)
	return avg
}


func swap2(x *int , y *int){
	temp:=*x
	*x=*y
	*y=temp
}


type Book struct{
	title string 
	author string
    subject string
    book_id int
}
func printBook( book *Book ) {
   fmt.Printf( "Book title : %s\n", book.title);
   fmt.Printf( "Book author : %s\n", book.author);
   fmt.Printf( "Book subject : %s\n", book.subject);
   fmt.Printf( "Book book_id : %d\n", book.book_id);
}


func printSlice(x []int){
	fmt.Printf("len=%d  cap=%d  slice=%v\n",len(x),cap(x),x)
}


func factoriel(n int) int{
	if(n==1 || n==0){ return 1}
	return n*factoriel(n-1)
}


func fibonacci(n int) int {
	if (n==0){return 0}
	if (n==1){return 1}
	return fibonacci(n-1)+fibonacci(n-2)
}


type Shape interface{
	area() float64
}
type Rectangle struct{
	width,height float64
}
func(rectangle Rectangle) area() float64{
	return rectangle.width * rectangle.height
}
func getArea(shape Shape) float64{
	return shape.area()
}


func sqrt(x float64) (float64,error){
	if(x<0){
		return 0, errors.New("Negative number passed to Sqrt")
	}
	return math.Sqrt(x),nil
}
