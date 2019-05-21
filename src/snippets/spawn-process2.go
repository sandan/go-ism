package main

import (
  "os"
  "os/exec"
)

func main(){

  // command for spawned process to execute
  entrypoint := exec.Command("python", "echo.py")

  // redirect input from go process to python process
  entrypoint.Stdin = os.Stdin

  // redirect stdout from python process to go process
  entrypoint.Stdout = os.Stdout

  // execute spawned process
  entrypoint.Start()

  // wait for process to finish
  entrypoint.Wait()
}
