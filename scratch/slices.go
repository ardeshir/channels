package main
 
import "fmt"
 
func main(){
             //index 0      1    2      3      4       5      6     7     8       9      10         11
 
// var days = make([]string, 12) 

//days := []string{"1st", "2nd","3rd", "4th", "5th", "6th","7th", "8th", "9th", "10th", "11th", "12th"}
 
arr := []int32{1,2,3,4,5} 
x := arr[:3]
// gifts := [12]string{"partridge in a pear tree","turtle doves","French hens","calling birds","golden rings","geese a-laying","swans a-swimming","maids a-milking","ladies dancing","lords a-leaping","pipers piping","drummers drumming"}
 
//var first int = 0
//var last  int = 11
// len(days)

// fmt.Printf("First %s\n", days[first])
fmt.Printf("arr_slyc:  %d\n", arr)
//fmt.Printf("Last %s\n", days[ len(days)-1 ])

fmt.Printf("x: %d\n", x)
//fmt.Printf("Last %s\n", days[first:last])
/* 
for i := range days {

fmt.Printf("On the %s day of Christmas, my true love gave to me\n", days[i] )

 if(i == 0) {
 
   fmt.Printf("\tA %s.\n", gifts[i])
 
 } else {

      for j := i; j > 0 ; j-- {
      	
          fmt.Printf("\t%d %s, \n", j+1, gifts[j])
      }

      fmt.Printf("\tAnd a %s.\n", gifts[0])
   } 
 }  */
 
} // end of main 
