package main

type ExecutionProcess struct {
  Command string
  Args    string
  WorkDir string
  OS      *OSOptions
  Runtime map[string]string
}

type OSOptions struct {

}

type Request struct{
  RawData  []byte
  Identity *Sender
  Info  *SQLInfo
}

type SQLInfo struct{
  Script, ScriptCommand, Contract, ScriptArgs string
}

type Sender struct {
  DB, SessionID, QueryId, AMPId string
}


func main(){

// spawn
// redirect io
// set command + workdir

// execute
//set venv
//ensure script in curdir
//execute command
// instrument:
//  - rows read/written
// log:
//  - exceptions
//  - errors/warn: bad sender info
//  - stdout

}
