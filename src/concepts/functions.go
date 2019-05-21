package main

import (
  "math"
  "fmt"
  "encoding/json"
  "os"
  "time"
  "log"
)

func hypot(x, y float64) (z float64) {
  fmt.Println("result list parameters are initialized with their zero value: ", z)
  // they can be assigned to
  z = math.Sqrt(x*x + y*y) + 1
  return math.Sqrt(x*x + y*y) // but z takes on the value given by the return expression (if given)
}

/* A function without a return list means that it is called purely for its mutating effects
The below function does not compile since it returns a value but specifies no return list:
./functions.go:17: too many arguments to return
func add(x, y float64) {
  return x + y
}
*/

// a function that has a result list must end with a return
// unless it ends with a call to panic() or is in an infinite for{} loop with no break
// then the result list is ignored
func addpanic(x, y int) (s, v int) {
  panic("ahhhh!")
}

// result list doesn't need () if unnamed
func add(x, y int) int {
  return x + y
}

// the result list specifies z
// no need to return z explicitly if we set it in the fn
// this is called a "bare" return
func sub(x, y int) (z int) {
  z = x - y
  return
}
// you can use _ to signal unused parameters
// notice how return b is needed, can't have (b int) and bare return otherwise we get:
//./functions.go:44: duplicate argument b
func skip(_, b int) int {
  return b
}

// you can use unnamed parameters to also signal unused params
func zero(int, int) int { return 0 }

func main(){

  // Named functions can only be declared at the package level
  // hence, we will use functions defined outside of main
  // except for lambdas

  fmt.Println(hypot(4.2, 4.14))
  //Every function call must provide an argument for each parameter, in the order in which the parameters were declared.
  // there are no default values for parameters in a function signature
  // there is no way to specify an argument by param name
  // parameters are local variables in the body of the function with the initial value set by values supplied by caller
  // arguments are passed by value (except for the following types: pointer, slice, map, channel, function)

  // pass by value ~ copies of data sent to the function's arguments; modification to the copy does not affect the caller's data
  // pass by reference ~ pointer to (address of/reference) the data is sent to function, modification to referenced pointer affects caller's data

  fmt.Println(add(1,1))
  fmt.Println(sub(9, 1))
  fmt.Println(skip(99,100))
  fmt.Println(skip(0 ,zero(11,12))) // skip & zero still needs a first (+ second for zero) argument even if unused

  // You can print functions
  fmt.Printf("%T\n", add)
  fmt.Printf("%T\n", sub)
  fmt.Printf("%T\n", zero)
  fmt.Printf("%T\n", skip)

  // the type of a function is called its signature
  // two functions have the same type if they have the same sequence of types in the parameter list and the same sequence of types in the result list
  // names of parameter list and result list doesn't affect the type, nor does whether _ or the factored form is used (x, y int)

  //Note: Functions declared without a body means that the function is implemented in a language other than Go

  // Recursion
  for x := 0; x <= 10; x++ {
    fmt.Println(countChange(x, []int{1, 5}))
  }

  tree := Node{Value: "+",
            Left:
              &Node{Value: "5", Left:nil, Right:nil},
            Right:
              &Node{Value: "*",
                   Left:  &Node{Value: "11"}, // notice that Left and Right aren't needed since zero value of *Node is nil
                   Right: &Node{Value: "12"},
                  },
              }

  treejson, err := json.MarshalIndent(tree, "", "    ")
  fmt.Printf("%s\n", treejson)

  fmt.Println("preorder")
  fmt.Println(preorder(&tree))

  fmt.Println("postorder")
  fmt.Println(postorder(&tree))

  fmt.Println("inorder")
  fmt.Println(inorder(&tree))

  v := []Node{
    {Value: "42", Right: nil},
    {Value: "43", Left: nil,Right: nil},
  }

  data, err := json.MarshalIndent(v, "", "  ")
  if err != nil {
      fmt.Printf("%s\n", err)
  }

  fmt.Printf("Pointer Value: %s\nType: %T\n", data, v[0])

  // Multiple Return Values
  a, b, c := return3(8)
  fmt.Println(a, b, c)

  // scope of g h i are only in the if else statement
  if g, h, i, err := return4(22); err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(g,h,i)
  }

  // Error handling
  // By convention, errors are returned last in the result list
  // If the failure only has one possible cause, such as a map lookup, the error is a boolean type
  // Otherwise the error is usually returned as an "error" type
  //   - In this case, an error of nil means success (no error), non-nil error means something went wrong
  //   - errors have a string message associated with them (fmt.Println(err) or fmt.Printf("%v\n", err)

  // Go's approach to error handling is not to use exceptions - those are only used for truly unexpected results ("bugs")
  // The motivation is that the exception gives a stack trace with little context of why something went wrong
  // Go prefers developers use normal control flow with ifs and returns to check for errors and report them
  // When a function call returns an error, it's the caller's responsibility to check it
  // Contrast this with python ~ It is easier to ask forgiveness than permission

  // (1) Propagate the error so that an error in the subroutine becomes an error in the calling routine
  //     Build descriptive errors by successively prefixing additional context to the original error message
  //     When the error is finally handled in main, it should provide a clear causal chain from root problem to overall failure
  //     When chaining, avoid using newlines and using capitalized letters - easier to grep
  //
  //     main  -->  error:   (why?)
  //         if val, err := fn1(x); err != nil {fmt.Errorf("error: %v", err)}

  //         fn1  --> could not create table:   (why?)
  //            if val, err := fn2(x); err != nil {return nil, fmt.Errorf("could not create table: %v", err)}

  //            fn2  -->  database threw an error:     (why?)
  //               if val, err := fn3(x); err != nil {return nil, fmt.Errorf("database threw an error: %v", err)}

  //               fn3  -->  "[Database] Memory Full"       (oh...)
  //                  if val, err := db(x); err != nil { return err}
  // each message in the causal chain should answer why the previous error happened
  // In general, the call f(x) is responsible for reporting the attempted operation f and the argument value x as they relate to the context of the error.

  // (2) Retry errors for transient or unpredictable errors
  //     - limit number of tries
  //     - limit time spent trying before giving up
  //     - use a delay between tries (such as exp backoff)
  //
  //     const timeout = 1 * time.Minute
  //     deadline := time.Now().Add(timeout)
  //     for tries := 0; time.Now().Before(deadline) && tries < MAX_TRIES; tries++ {
  //          {try stuff}
  //          time.Sleep(time.Second << uint(tries)) // exp backoff before trying again
  //     }
  //     fmt.Errorf("failed to respond after %v", timeout)
  //
  //
  // (3) If progress is impossible, the caller can print the error and stop the program gracefully 
  //     if val, err := f(x); err != nil { log.Fatalf("service is down: %v", err) } (print to stderr and os.exit(1))
  // (4) Log error and continue with degraded functionality or partial results
  // (5) In rare cases, we can ignore the error entirely (such as when reading EOF from a file, we don't have to log it just continue)
  // 
  // Get into the habit of considering errors after every function call, and when you deliberately ignore one, document your intention clearly. 

  /* Functions are first class values in Go */
  f := even
  if res, err := int_filter(f, []int{1,2,3,4,5,6,7,8}); err != nil {
    fmt.Errorf("main: error in f: %v", err)
  } else {
    fmt.Println(res)
  }

  f = odd // even and odd functions defined below are the same type: func(int) bool, so f can be re-assigned
  odds, _ := int_filter(f, []int{1,2,3,4,5,6,7,8,9,10})
  fmt.Println(odds)

  // if f was assigned to a different function type, we get a compile error

  // The zero value of a function type is nil
  type Object struct {
    Method func(int) bool
    Name string
  }
  var instance = Object{Method: even, Name: "IsEven"}
  fmt.Println(instance.Method(2))

  var methodless = Object{Name: "Methodless"}
  fmt.Println(methodless)

  var strategy = make(map[string]func(int) (int, int, int))
  strategy["python"] = return3
  m, n, o := strategy["python"](42)
  fmt.Println(strategy, m, n, o)

  // calling a nil function throws a panic:
  ///var nilfn func(string) bool
  // nilfn("panic") 
  // panic: runtime error: invalid memory address or nil pointer dereference
  // function values can be compared with nil (see int_filter) but they cannot be compared with other types
  // so they cannot be directly used as keys in a map

  /* Anonymous functions (lambdas) */
  // declared like functions without names and braces for the body
  // function literals
  // expressions whose value is called an anonymous function
  var lambda = func(x int) bool { return x > 7 }
  fmt.Println(lambda == nil)
  greater_than_7, _ := int_filter(lambda, []int{1,2,3,4,5,6,7,8,9,10})
  fmt.Println(greater_than_7)

  // anonymous functions have access to their entire lexical scope
  // inner functions can reference variable in the outer scope
  incr := increment(99)
  fmt.Println(incr())
  fmt.Println(incr())
  fmt.Println(incr())
  fmt.Println(incr())
  fmt.Println(incr())

  // function values can have state as well as code as seen above
  /*
  When an anonymous function requires recursion, we must first declare a variable, 
  and then assign the anonymous function to that variable. 
  */
  var summer func(func(int) int, []int) int
  summer = func(f func(int) int, arr []int) int {
    head := arr[0]
    if len(arr) == 1 {
      return f(head)
    }
    tail := arr[1:]
    return f(head) + summer(f, tail)
  }
  fmt.Println(summer(func(x int) int { return  x*x }, []int{1,3,5,7,9}))

  /*
  Had these two steps been combined in the declaration, the function literal would not be within the scope of the variable
  so it would have no way to call itself recursively
  var summer func(func(int) int, []int) int = func(f func(int) int, arr []int) int {
    . . .
    return f(head) + summer(f, tail) <-- scope not valid, summer variable not defined yet
  }
  */

  //Capturing iteration loop variables in closures
  //The lexical scope of loop variables interact
  var fs []func()
  for _, i:= range []int{1,1,2,3,5,8}{

    var f = func() {
      //trying to capture the value in the function
      fmt.Println(i)
    }
    fs = append(fs, f)
  }
  for _, fun := range fs {
    fun()
  }
  // the above prints all 8's (the last value in the range)
  /*
  the for loop introduces a new lexical block in which the variable i is declared. 
  All function values (closures) created by this loop "capture" and share the same variable — an addressable
  storage location, not its value at that particular moment
  */
  // so basically, when we define the closure f, it captures a pointer? to the loop variable, not the value
  // to capture the value, we have to assign a new variable in the same scope as the closure
  for _, i:= range []int{1,1,2,3,5,8}{

    i := i //get the value from the loop variable i to use in closure
    var f = func() {
      //trying to capture the value in the function
      fmt.Println(i)
    }
    fs = append(fs, f)
  }
  for _, fun := range fs {
    fun()
  }

  /* Variadic Functions */
  // functions that can take a variable number of arguments
  fmt.Println(sum_it(0, 1, 2, 3, 5, 5))
  // implicit allocation of an array happens when values are passed into vals
  // the array is copied and a slice of the array is passed to the function

  var values = []int{1,2,3,5,5} // allocate a slice
  fmt.Println(sum_it(0, values...)) // invoke a variadic function with a slice

  // The type of a variadic function
  fmt.Printf("Type of sum_it(init int, vals ...int): %T\n", sum_it)
  fmt.Printf("type of not_variadic(init int, vals []int): %T\n", not_variadic)

  /* Defer */
  // duplication of cleanup logic becomes a maintenance problem
  // for example, making sure to clenaup network connections on failures

  // The defer stmt is a function or method call prefixed with the defer keyword

  // the function's arguments and the function expression is evaluated 
  // the call is deferred until the enclosing scope of the defer stmt ends 
  // (whether normally as in the end of a function call or abnormally by panic)

  // Any number of calls may be deferred. Execution happens in reverse order in which they were deferred.

  // A defer statement is often used with paired operations like open and close (files), connect and disconnect (requests, HTTP), or lock and unlock
  // to ensure that resources are released in all cases, no matter how complex the control flow. 
  // The right place for a defer statement that releases a resource is immediately after the resource has been successfully acquired.

  read := func() error {
    f, err := os.Open("functions.go")
    if err != nil {
      return err  // open did not succeed
    }
    // open succeeded!
    defer f.Close() // lets make sure to close it when this function ends, call to Close happens after func is done assigning result variables
    fmt.Println(f)  // there are sync bugs? No only for os.Create, not for os.Open: https://www.joeshaw.org/dont-defer-close-on-writable-files/
                    // book talks about this briefly, see below

    return nil
  }
  read()

  // The defer statement can be used to pair "on entry" and "on exit" actions
  // This is kind of like decorators in python, where the function wraps another 
  // function but defer does so by being called in the function it wraps
  // useful for debugging ~ tracing 8)

  bigSlowOp := func() {
      x := 1
      defer tracer("bigSlowOp()", x)() // note extra parens, this evaluates (the function expression )
                                       // so it will call and evaluate tracer (on entry logic) 
                                       // but defers the call to anonymous func returned by tracer till after bigSlowOp (on exit)
                                       // observe what happens when you remove the parens
      time.Sleep(2 * time.Second)
      x = 2
      fmt.Println("bigSlowOp: x = ", x)
      // bigSlowOp returns, defer statements are executed, the anon func is called (which was returned when the defer was evaluated)
      // the anonymous func returned by tracer() still has the old value of x, since the closure was returned before x was updated
  }
  bigSlowOp()

  // deferred functions run after the return statement updates result variables (of the current function that does the defer stmt)
  // anonymous functions can access named variables in their enclosing scope, including result variables
  // so a deferred anonymous function can observe the function's results
  // so you can also use defer to do logic after a function returns, useful for tracing



  dbl := func(x int) (result int) {
      x = x + 1
      defer func() {
        fmt.Printf("Input: %d Output: %d\n", x, result)
      }()
      x = x + 1 // the func literal runs after result has been assigned, picks up modifications to x (on exit semantics only)

      return x + x // note that result is updated with the return stmt, no need to update it in function body
  }
  dbl(2) //input 4 and output 8

  // combining the on entry and on exit example with the deferred anon func
  trpl := func(x int) (result int) {
      fmt.Println("trpl")

      defer func() (func()){ // on entry - evaluated right away
        fmt.Println("on entry 1")
        fmt.Printf("Input: %d\n", x)
        return func() { // on exit - this doesnt get called until after the return
          fmt.Println("on exit 2, x = ", x) // result gets assigned the result in body
          fmt.Println("output: ", result) // should be 12
        }
      }()()
      x = x + 1
      fmt.Println("trpl")
      result = x + x + x
      return
  }
  trpl(4)

  /* Bug? The above works but when we do it below, it does not */
  // https://github.com/golang/go/issues/32175
  // perhaps something with closure semantics is happening here?
  trpl = func(x int) (result int) {
      fmt.Println("trpl")

      defer func() (func()){ // on entry
        fmt.Println("on entry 1")
        fmt.Printf("Input: %d\n", x)

        return func() { // on exit - this doesnt get called until after the return
          fmt.Println("on exit 2, x = ", x) // but since this closure sees result with its initial value
          fmt.Println("output: ", result) // still 0? but it should be 12 since this anon func happens after result variables are assigned by return
        }
      }()()
      x = x + 1
      fmt.Println("trpl")
      return x + x + x
  }
  trpl(4)

  /* defer statements in a loop */
  // a defer statement in a loop deserves extra scrutiny
  // opening a file and defering the close could potentially exhaust file descriptors available for a process
  // instead, if you need to defer, call a separate function in the for loop and defer inside that function
  doFile := func(f string) error {
    if fl, err := os.Open(f); err != nil {
        return err
    } else {
        defer fl.Close()
        // .. process file ..
    }
    return nil
  }
  fmt.Println(doFile) //use it so we can compile...
  _ = func(filenames []string, doFile func(f string) (error)) error{

     // good
     for _, filename := range filenames{
         if err := doFile(filename); err != nil {
             return err
          }
     }

     // bad - files wont be closed till all of them are open
     /*for _, filename := range filenames{
         if f, err := os.Open(filename); err != nil {
             return err
          }
          else{
            defer f.Close()
             //.. process file ..
          }
       }*/
       return nil
  }

  /* Closing Writable Files */
  // do not defer a Close() to a file created with os.Create()
  // Create() opens a file for writing
  // On many file systems, notably NFS, write errors are not reported immediately but may be postponed until the file is closed. 
  // Failure to check the result of the close operation could cause serious data loss to go unnoticed.
  // In general, calling defer doesn't give you a chance to check the output of the function call


  /* See additional references too:
      https://blog.learngoprogramming.com/gotchas-of-defer-in-go-1-8d070894cb01
  */

  /* panic */
  // used for throwing runtime errors (index out of bounds, nil pointer dereference)
  /*
  - During a typical panic, normal execution stops, all deferred function calls in that goroutine are executed, and the program crashes with a log message. 
  - This log message includes:
    - the panic value, which is usually an error message of some sort, and, 
    - for each goroutine, a stack trace showing the stack of function calls that were active at the time of the panic. 
  
  This log message often has enough information to diagnose the root cause of the problem without running the program again, 
  so it should always be included in a bug report about a panicking program.
  
  Not all panics come from the runtime. The built-in panic function may be called directly ; it accepts any value as an argument. 
  A panic is often the best thing to do when some "impossible" situation happens
  */



  /* recover */

}

func tracer(msg string, x int) func() {
  // This section of tracer is executed (evaluated right away) if there are extra parens in the defer statement
  fmt.Println("hi from tracer!")
  start := time.Now()
  log.Printf("enter %s", msg)
  // If there were no extra parens, then only ths body above gets executed after the call the bigSlowOp
  return func() {log.Printf("exit %s (%s), x=%d", msg, time.Since(start), x)} // note log.Printf does not need newline
}

func not_variadic(init int, vals []int) int {
  return sum_it(init, vals...)
}
/* Variadic functions */
func sum_it(init int, vals ...int) int { // vals must be the final parameter
  var runningSum int
  for _, val := range vals {
    runningSum += val
  }
  fmt.Printf("type of variadic param: %T\n", vals) // int slice []int
  return init + runningSum
}
/* Anonymous functions */
func increment(x int) func()int{
  // zero value of x is given by the caller
  var lambda = func() int {
  // inner lambda can access + update local variables of enclosing function increment() x
  // this hidden variable reference is why functions are reference types and are not comparable
  // their body may look the exact same but they may have access to state and the state may be different btw func values
  // function values having state defined outside the body of the function but used in the function are known as closures,
  // function values and state defined in the same lexical scope where the function value uses state defined outside its body
  // closure ~ function values
        x++
        return x
  } // example of how the lifetime of a variable is not tied to its scope
    // x is defined in increment() but gets updated even after increment() returns
  return lambda
}
/* Multiple return values */
func return3(x int) (int, int, int){
  return x, x + 1, x + 2 // multiple return values are comma separated in return stmt
}

func return4(x int) (a int, b int, c int, d error){
  a = x + 1
  b = x + 2
  c = x + 3
  d = nil
  return // multiple return values that use named parameters
}

/* Function First Class */
func odd(y int) bool {
  return y % 2 == 1
}

func even(x int) bool {
  return x % 2 == 0
}
func int_filter(f func(int) bool, data []int) ([]int, error){

  if len(data) == 0 {
    return nil, fmt.Errorf("error! no dataz or no funcz!!\n")
  }
  var res []int
  for  _, x := range data {
    if f(x) {
      res = append(res, x)
    }
  }
  return res, nil
}


/* Recursion */
// functions may call themselves directly or indirectly
var memo = make(map[string]int)
func f(money int, coins []int) string { return fmt.Sprintf("%d %q",money, coins) }

// count the number of ways to make change given money and slice of denominations
// can use the memo var that we defined in the same scope
func countChange(money int, coins []int) int {

  if money == 0 {
    return 1 // found a way to make change

  } else if money < 0 || len(coins) <= 0 {
    return 0 // cant return change w/o money or no coins

  } else {
    if result, ok := memo[f(money, coins)]; !ok {
      memo[f(money, coins)] = countChange(money - coins[0], coins) + countChange(money, coins[1:])
    } else {
      return result
    }
  }
  return memo[f(money, coins)]
}


type Node struct {
  Value string `json:"value"`
  Left, Right *Node
}

func preorder(v *Node) string {

  if v.Left == nil || v.Right == nil{
    return v.Value
  }
  return fmt.Sprintf("(%s %s %s)", v.Value, preorder(v.Left), preorder(v.Right))
}

func postorder(v *Node) string {

  if v.Left == nil || v.Right == nil{
    return v.Value
  }
  return fmt.Sprintf("(%s %s %s)", postorder(v.Left), postorder(v.Right), v.Value)
}

func inorder(v *Node) string {

  if v.Left == nil || v.Right == nil{
    return v.Value
  }
  return fmt.Sprintf("(%s %s %s)", inorder(v.Left), v.Value, inorder(v.Right))
}
