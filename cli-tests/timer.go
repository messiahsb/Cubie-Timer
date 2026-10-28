import (
	"fmt"
	"bufio"
	"os"
	"time")

func main(){
	fmt.Println("Speedcuber Timer CLI")
	fmt.Println("Press Space to Start Timer")
	
	fmt.Println("Timing...")

	fmt.Println("Press Space to End Timer")
	duration := timeSolve()
	fmt.Println(duration)

}

//func waitForEnter()
//func timeSolve() time.Duration
