package main
import ( "fmt"
         "time" )
var unbuffered = make(chan int)
var   buffered = make(chan int, 5)

func main() {
	count := 5
        go func() {
             for i := 0; i <= count ; i++ {
                  fmt.Println("send message")
                  unbuffered <- i
              }
         }()
       
       time.Sleep(time.Second * 3)
       for i := 0; i <= count; i++ {
          fmt.Println(<- unbuffered)
        }
} // end main
