package main

import (
	esolFmt "esol/example/directory/fmt"
	"fmt"
)

func main() {
	fmt.Println(GreetWorld())
	esolFmt.SPrintf()
}

func GreetWorld() string {
	return fmt.Sprint("Hello, world!")
}

// TODO why are arrays not able to be used as constants
// TODO labels
// TODO TestMain, *testing.M
// TODO why should an error implementation use pointer receiver for Error() string
