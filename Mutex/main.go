package main

import(
	"fmt"
	"sync"
	"time"
)

type safeCounter struct {
	mu sync.Mutex //later we can just call mu.lock() or unlock()
	v map[string]int
}

func main() {
	//so far we have seen how channels can be great ways of communication among go routines
	//so now what if we don want routines to communicate but we want only one routine to access a variable at a time?
	//the concept of mutual exclusion comes here

	//converntional name for the DS is mutex
	//sync.Mutex for the case of go 
	//with its two methods lock and unlock

	//u see what is heppening on the above inc function?
	//we don wnat too incerement a value twice u know. so we use lock for that

	//AND RIGHT OUT OF THE BLUE => WE CAN MAKE A MACHINE REBASE RATHER THAN MERGE EVERYTIME WE PULL, USING THEGIT CONFIG COMMND
	
}