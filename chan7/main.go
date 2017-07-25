package main
import ( "fmt"
         "time" )
var unbuffered = make(chan int)
var   buffered = make(chan int, 5)
// example code of buffered vs unbuffered channels, or 
// sync or async channels
func main() {
	count := 5
        go func() {
             for i := 0; i <= count ; i++ {
                  fmt.Println("send message")
                  // unbuffered <- i 
                     buffered <- i
              }
         }()
       
       time.Sleep(time.Second * 3)
       for i := 0; i <= count; i++ {
          // fmt.Println(<- unbuffered)
          fmt.Println(<- buffered)
        }
} // end main
