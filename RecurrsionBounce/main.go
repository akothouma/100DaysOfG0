/* How many times will the mother see the ball pass in front of her window (including when it's falling and bouncing)?
Three conditions must be met for a valid experiment:

    Float parameter "h" in meters must be greater than 0
    Float parameter "bounce" must be greater than 0 and less than 1
    Float parameter "window" must be less than h.

If all three conditions above are fulfilled, return a positive integer, otherwise return -1.
Note:

The ball can only be seen if the height of the rebounding ball is strictly greater than the window parameter.*/

package main

func BouncingBall(h, bounce, window float64) int {
	count:=1
	if h>0 &&( bounce>0 && bounce<1) && window<h{
	  for h>window{  
		if h*bounce>window{
		  count=count+2
		 }
		 h=bounce*h
	  }
	  }else{
	  return -1
	}
	 return count
  }
func main(){
      BouncingBall(3, 0.66, 1.5,) //3
      BouncingBall(40, 0.4, 10) //3
      BouncingBall(10, 0.6, 10) //-1
      BouncingBall(40, 1, 10) //-1
}