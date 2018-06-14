package main

import(
  "fmt"
  // "math"
)

func main() {
  fmt.Print("Enter fahrenheit: ")
  var fahrenheitInput float64
  fmt.Scanf("%f", &fahrenheitInput)

  celsius := (fahrenheitInput - 32) * 5/9
  roundedCelsius := float64(int(celsius * 100)) / 100 //round to 2 decimals

  fmt.Println(fahrenheitInput, " degrees fahrenheit is ", roundedCelsius, " degrees Celsius")
}
