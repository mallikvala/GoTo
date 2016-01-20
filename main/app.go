package main

import (
	"fmt"
	"os"

	"github.com/mallikvala/GoTo/utils"
)

func main() {
	fmt.Println(utils.Hi(os.Args[1]))
	fmt.Println(utils.Hello(os.Args[1]))
}
