package main

import (
  "os"
  "os/exec"
  "log"
  "fmt"
  "bufio"
)

func main() {
    cmd := exec.Command("python", "echo.py")
    cmd.Stdin = os.Stdin

    cmdReader, _ := cmd.StdoutPipe()
    scanner1 := bufio.NewScanner(cmdReader)

	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}

    cmd.Start()

    go func() {
        for scanner1.Scan() {
            fmt.Printf(scanner1.Text())
        }
    }()


	err = cmd.Wait()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
