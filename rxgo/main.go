package main

import (
	"errors"
	"fmt"

	rxgo "github.com/reactivex/rxgo/v2"
)

func main() {
	observable := rxgo.Just(1, 2, errors.New("unknown"), 3, 4, 5)()
	ch := observable.Observe()
	for item := range ch {
		if item.Error() {
			fmt.Println("err: ", item.E)
		} else {
			fmt.Println(item.V)
		}
	}
}
