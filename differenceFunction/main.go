package main

// difference function, which subtracts one list from another and returns the result.

	func ArrayDiff(a, b []int) []int {
		l3:=[]int{}
	   if len(a)==0{
		 return a
	   }else if len(b)==0{
		 return a
	   }else if len(b)==1{
		 for _,l1:=range b{
		 for _,l2:=range a{
		   if l1!=l2{
			 l3=append(l3,l2)
		   }
		 }
	   }
		 }else {
		 for _,l1:=range a{
		   duplicate:=false
		   for _,l2:=range b{
			 if l1==l2{
			   duplicate=true
			   break
			 }
			 }
			if !duplicate{
			   l3=append(l3,l1)
		   }   
		 }
		 }
		 fmt.Println(l3)
	   return l3
	 }

	 func main(){
		ArrayDiff([]int{1,2,3,4,5},[]int{1,2,3})
		ArrayDiff([]int{1,2},[]int{1})
	 }