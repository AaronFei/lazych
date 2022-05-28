# lazych

Easily create the channel across different module with the channel type and name

```
package main

import (
	"fmt"

	"github.com/AaronFei/lazych"
)

func foo() {
	c, err := lazych.GetCh[string]("foo", 1)
	if err != nil {
		fmt.Println(err)
		panic("")
	} else {
		c <- "Hello World!"
	}
}

func main() {
	foo()
	c, err := lazych.GetCh[string]("foo", 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("from foo:", <-c)
	}
}
```
