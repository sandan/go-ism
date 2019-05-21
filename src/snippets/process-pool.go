package main

import (
  "fmt"
  "os/exec"
  "io/ioutil"
  "time"
)

/* doesn't work */
func worker(id int, jobs <-chan int, results chan<- int){

  // wait for jobs to come through channel
  for j:= range jobs {
    time.Sleep(time.Second)
    // command for spawned process to execute
    fmt.Sprintf("[%v] Started", id)
    entrypoint := exec.Command("python", "echo.py")

    // stdin and stdout of spawned process
    in, _ := entrypoint.StdinPipe()
    out, _ := entrypoint.StdoutPipe()

    // execute spawned process
    entrypoint.Start()

    // write to stdin of spawned process
    in.Write([]byte("[" + string(id) + "]" + " got data: " + string(j)))

    // close stdin pipe
    in.Close()

    // read from stdout of spawned process
    byte_result, _ := ioutil.ReadAll(out)

    // wait for process to finish
    entrypoint.Wait()

    // write result to channel
    results <- len(string(byte_result))
    fmt.Sprintf("[%v] Finished", id)

  }
}

func main(){

  jobs := make(chan int, 100)
  results := make(chan int, 100)

  /* start a pool of workers */
  for i:= 1; i <= 3; i++ {

    go worker(i, jobs, results)

  }

  // pass input data to each worker's job channel
  for j := 1; j <= 5; j++ {
    jobs<- j
  }

  close(jobs) // signal no more work

  // read results
  for a := 1; a <= 5; a++ {
    <-results
  }

}
