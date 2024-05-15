//In a small town the population is p0 = 1000 at the beginning of a year.
//The population regularly increases by 2 percent per year and moreover 50 new inhabitants per year come to live in the town.
//How many years does the town need to see its population greater than or equal to p = 1200 inhabitants?

// Don't forget to convert the percent parameter as a percentage in the body of your function: if the parameter percent is 2 you have to convert it to 0.02.
// There are no fractions of people. At the end of each year, the population count is an integer: 252.8 people round down to 252 persons.
package main

import "fmt"

var count int
func NbYear(p0 int, percent float64, aug int, p int) int {
  actualPercent:=percent/100.0
  p1:=float64(p0)*actualPercent+float64(aug)+float64(p0)
  count=1
  for p>=int(p1){
    NbYear(int(p1),percent,aug,p)
    count++
    break
  }
  return count
}
func main(){
	fmt.Println(NbYear(1500, 5, 100, 5000))
	fmt.Println(NbYear(1500000, 2.5, 10000, 2000000))
}