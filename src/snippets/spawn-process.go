package main

import (
  "os/exec"
  "fmt"
  "io/ioutil"
)

func main(){

  // command for spawned process to execute
  entrypoint := exec.Command("python", "echo.py") // replace echo.py with autogenerated script

  // stdin and stdout of spawned process
  in, _ := entrypoint.StdinPipe()
  out, _ := entrypoint.StdoutPipe()

  // execute spawned process
  entrypoint.Start()

  // write to stdin of spawned process
  in.Write([]byte("hello universe!"))

  // close stdin pipe
  in.Close()

  // read from stdout of spawned process
  byte_result, _ := ioutil.ReadAll(out)

  // wait for process to finish
  entrypoint.Wait()

  // write result to stdout
  fmt.Println(string(byte_result))
}