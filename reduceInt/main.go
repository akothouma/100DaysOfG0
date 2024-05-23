package main
import ("fmt")
func main(){
	as := []int{500, 2}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)

}

func ReduceInt(a []int, f func(int, int) int) {
  fmt.Println(f(a[0],a[1]))
}
func mul(a,b int)int{
	return a*b
}
func sum(a,b int)int{
	return a+b
}
func div(a,b int)int{
	return a/b
}
