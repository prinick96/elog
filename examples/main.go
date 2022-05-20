package main

import (
	"fmt"
	"os"

	"github.com/prinick96/elog"
)

func main() {
	_, err := fmt.Printf("Hola %v\n", "mundo")
	go elog.New(elog.ERROR, "esto hará nada, porque no hay un error", err)

	_, err = os.OpenFile("./unpathquenoexiste", os.O_RDWR, 0644)
	go elog.New(elog.ERROR, "esto sí me generará un log", err)

	_, err = os.OpenFile("./otropathquenoexiste", os.O_RDWR, 0644)
	elog.New(elog.PANIC, "y esto, será un panic cuando exita error", err)

	fmt.Println("Esto no se ejecutará")
}
