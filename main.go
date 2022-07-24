package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	var n int64
	isAsync := len(os.Args[1:]) != 0 && os.Args[1:][0] == "async"
	fmt.Println("What is the value of N?")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Errorf("something went wrong grabing input")
	}
	fmt.Printf("calculating the posistion of %v in a fib sequence\n", n)
	start := time.Now()
	var val int64;
	fmt.Printf("Async is set to: %v\n", isAsync)
	if(!isAsync){
		val = fibSync(n)
	}else{
		val = fibAsync(n)
	}
	//val := fibAsync(n)
	end := time.Now()

	elapsed := end.UnixMilli() - start.UnixMilli()
	fmt.Printf("the number %v will be in position: %v, this operation took %v miliseconds\n", val, n, elapsed)
}

func fibSync(n int64) int64 {
	//fmt.Printf("   (%v)\n",n)
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	//fmt.Printf("  /   \\\n /    \\\n(%v)   (%v)\n",n-1,n-2)
	return fibSync(n-1) + fibSync(n-2)
}

func fibAsync(n int64) int64 {

	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	leftCh := make(chan int64)
	rightCh := make(chan int64)

	wg.Add(2)
	go func(){
		x := fibSync(n-1);
		leftCh <-x;
		close(leftCh)

	}()
	go func(){
		x := fibSync(n-2);
		rightCh <-x;
		close(rightCh)

	}()
	
	left := <-leftCh;
	right := <-rightCh;
	//fmt.Printf("left: %v right: %v",left, right )
	//

	return left + right;

}
