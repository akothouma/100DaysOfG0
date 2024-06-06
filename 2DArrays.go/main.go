/* WAP to create N×N multiplication table, of size provided in parameter.

//For example, when given size is 3:

1 2 3
2 4 6
3 6 9

For the given example, the return value should be:

[[1,2,3],[2,4,6],[3,6,9]]
*/
package main

func MultiplicationTable(size int) [][]int {
	var result2d [][]int
  
	for i:=0;i<size;i++{
		var result [] int
	  for column:=0;column<size;column++{
		result=append(result,(i+1)*(column+1))  
	  }
	  result2d=append(result2d,result) 
	}
	return result2d
  }

  func main(){
	MultiplicationTable(1) //[][]int{{1}
	MultiplicationTable(2) //([][]int{{1, 2}, {2, 4}}))
	MultiplicationTable(3) //([][]int{{1, 2, 3}, {2, 4, 6}, {3, 6, 9}}
  }