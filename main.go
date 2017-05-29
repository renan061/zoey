package main

import (
	"fmt" // TODO: Remove
	"log"
	"os/exec"
	"io/ioutil"
)

// gcc-5 -std=c99 -Wall -c main.c -o main.o
// gcc-5 -std=c99 -Wall main.o -o main

var (
	cc = "gcc-5"
	cflags = []string{"-std=c99", "-Wall"}
)

// ==================================================
//
//	GCC
//
// ==================================================

type GCC struct {
	CC string
	Flags []string
	Args []string
}

func NewGCC(cc string, flags []string, args ...string) *GCC {
	return &GCC{cc, flags, args}
}

func (gcc GCC) Command() *exec.Cmd {
	return exec.Command(gcc.CC, append(gcc.Flags, gcc.Args...)...)
}

func (gcc GCC) Run() error {
	return gcc.Command().Run()
}

func (gcc GCC) Output() ([]byte, error) {
	return gcc.Command().Output()
}

// ==================================================
//
//	Main
//
// ==================================================

func main() {
	fmt.Println("Start...\n")
	defer fmt.Println("\nEnd...")

	NewGCC(cc, cflags, "-c", "main.c", "-o", "main.o").Run()
	NewGCC(cc, cflags, "main.o", "-o", "main").Run()

	// Test

	cmd := exec.Command("./main")

	errPipe, _ := cmd.StderrPipe()
	outPipe, _ := cmd.StdoutPipe()
	cmd.Start()

	fmt.Println("Errors:")
	errstr, _ := ioutil.ReadAll(errPipe)
	fmt.Println(string(errstr))

	fmt.Println("Output:")
	outstr, _ := ioutil.ReadAll(outPipe)
	fmt.Println(string(outstr))

	cmd.Wait()
}

func main1() {
	fmt.Println("Start...\n")
	defer fmt.Println("\nEnd...")

	var err error

	err = NewGCC(cc, cflags, "-c", "main.c", "-o", "main.o").Run()
	if err != nil {
		log.Fatalln(err)
		return
	}

	err = NewGCC(cc, cflags, "main.o", "-o", "main").Run()
	if err != nil {
		log.Fatalln(err)
		return
	}

	out, err := exec.Command("./main").CombinedOutput()
	fmt.Println("Output...")
	fmt.Println(string(out))
	if err != nil {
		log.Fatalln(err)
		return
	}
}
