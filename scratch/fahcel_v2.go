package main

import "fmt"

func main() {


      fmt.Println("Please enter 'C' for Celsius or 'F' for Fahrenheit conversion")
      var typec rune
      n, err := fmt.Scanf("%c\n", &typec)
      if err != nil || n != 1 {
          // handle invalid input
              fmt.Println(n, err)
           return
      }

      if(typec == 'C') {

          fmt.Println("Enter your Celsius temp")
          var cel float32
          n, err := fmt.Scanf("%f\n", &cel)
          if err != nil || n != 1 {
          // handle invalid input
              fmt.Println(n, err)
           return
          }
          fmt.Printf("It's %f Fahrenheit\n", (cel * 9/5) + 32 )

      } else if (typec == 'F'){

          fmt.Println("Enter your Fahrenheit temp")
          var fah float32
          n, err := fmt.Scanf("%f\n", &fah)

          if err != nil || n != 1 {
          // handle invalid input
              fmt.Println(n, err)
           return
          }

          fmt.Printf("It's %f Celsius\n", (fah - 32) * 5/9 )


      } else {
          fmt.Println("Sorry, you didn't follow the instructions! Bye")
      }
}
