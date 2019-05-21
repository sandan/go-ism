package main

import (
  "fmt"
  "encoding/json"
  "sort"
)

func main(){

  maps()
  //structs()
  //jsons()

}

func section(header string) {
    pad:= "---------------------------"
    fmt.Printf("\n%s\n\t%s\n%s\n", pad, header, pad)
}

func arrays(){
  section("arrays")

  // fixed size, homogenous type
  // declaration
  // subscript notation
  // array literals
  // overriding array indices with symbols (pattern)
  // order properties
}

func slices(){
  section("slices")

  // dynamic size, homogenous type
  // decl: like arrays but with no size
  // components of a slice:
  //  - cap
  //  - len
  //  - aliasing
  // slice operator
  // slice literal
  // order properties
  //  - zero value of slice type
  // make()
  // append()
  // in place slicing techniques
}

func maps(){
  section("maps")

  // A map is a reference to a hash table
  // a map type is written map[K]V where K and V are the types of its values
  // initialization
  var my_map map[string]int
  fmt.Println(my_map)

  unordered_string_arr := make(map[int]string)
  fmt.Println(unordered_string_arr)

  // the key type must be comparable with == so that the map can test whether the key is already in the map
  // it is not advised to use floating point keys even though they are comparable
  // no restrictions on the value type

  // map literals
  my_literal_map := map[int]int{
    0:1,
    1:2, //<--- note the comma
  }
  fmt.Println(my_literal_map)

  // dynamic size, homogenous type for keys, values but type(key) is not nec. type(value)
  // read

   fmt.Println(my_literal_map[0])
   unordered_string_arr[0] = "hello!"
   fmt.Println(unordered_string_arr[0])

  // delete
  delete(unordered_string_arr, 0) // map, key
  fmt.Println(unordered_string_arr)

  // you can get, set, and delete keys that are not there (uses zero value for type)
  fmt.Println("my_literal_map[100]: ", my_literal_map[100])

  my_literal_map[99]++
  my_literal_map[99]+=1
  my_literal_map[99] = my_literal_map[99] + 1
  fmt.Println("my_literal_map[99]: ", my_literal_map[99])

  delete(unordered_string_arr, 50)
  fmt.Println("delete(unordered_string_arr, 50): ", unordered_string_arr)
  fmt.Println(unordered_string_arr[50]) // prints empty string with newline

  // a map element is not a variable, deref not allowed
  //_ = &unordered_string_arr[0] compile error: cannot take address of map element

  //One reason that we can’t take the address of a map element is that growing a map might cause
  //re hashing of existing elements into new storage locations, thus potentially invalidating the
  //address.

  // iteration
  //enumerating all key value pairs of a map
  fmt.Println("order of iteration of keys in map is non-deterministic")
  for i, j := range my_literal_map{
    fmt.Printf("%d\t%d\n", i, j)
  } // note that even though using a non-existent key from the map returns a zero value, 
    // the key is not in the map so it does not get printed. Those zero-value semantics are only for accessing from a map

  //  - sorted key ranging
  //   - get the keys of the map
  //   - sort the keys
  //   - range over the sorted keys and access the value using the key via the map
  var keys []int = make([]int, 0, len(my_literal_map))
  for key, _ := range my_literal_map {
    keys = append(keys, key)
  }
  sort.Ints(keys)
  fmt.Println("Sorted my_literal_map")
  for _, key := range keys {
    fmt.Printf("%d\t%d\n", key, my_literal_map[key])
  }

  // nil map
  // a map without keys is nil (except if it is allocaetd with make() or a struct literal)
  fmt.Println(unordered_string_arr) // empty
  fmt.Println(unordered_string_arr == nil) //false since it is init'd with make
  fmt.Println(len(unordered_string_arr)) //0

  // var my_map[K]V not a literal nor make'd
  fmt.Println(my_map) // empty
  fmt.Println(my_map == nil) //false since it is init'd with make
  fmt.Println(len(my_map)) //0

  // you can get, delete, range, len a  nil map but not set to a nil map
  //  my_map["panic"] = 0
  //panic: assignment to entry in nil map
  // you must allocate the map before you can store to it 

  // default values for keys not in map are the zero value
  // to distinguish between a non-existent key and the zero value use ok
  // ok is a boolean that reports whether the element is present
  _, ok := my_literal_map[1000]
  if !ok {
    fmt.Println("1000 in my_literal_map? ", ok)
    fmt.Println("1000 is not a key in my_literal_map but it has value ", my_literal_map[1000])
  }

  //short hand for the above is
  if _, ok := my_literal_map[1000]; !ok {
    fmt.Println(ok, my_literal_map[1000])
  }
  /*
  As with slices, maps cannot be compared to each other the only legal comparison is with nil.
  
  To test whether two maps contain the same keys and the same associated values, we must write a loop
  // see func map_equals
  */

  // map as set
  type emptystruct struct{}
  strset := map[string]struct{}{} //struct{} defines a struct with no members struct{}{} instantiates an empty struct
  strset["hello"] = emptystruct{}
  strset["universe"] = struct{}{}
  strset["universe"] = emptystruct{} // <-- adding this again has no effect
  fmt.Println(strset)

  // strset["galaxy"] = nil // cant use nil as it is not an empty struct type

  // slice keys
  // Sometimes we need a map/set whose keys are slices (or other non-comparable types)
  // Do this in two steps:
  // 1) define a function that maps non-comparable type to a comparable type
  //    with the property that f(x) == f(y) iff we consider x and y to be in the same equivalence class
  // 2) create a map whose keys are string (or some type whose == implements the desired behavior) then apply the function to each key before we access the map
  // this technique also allows you to define your semantics over == (equivalence classes) such as case insensitive string comparison

  var m = make(map[string]int)
  add([]string{"haha", "lol"}, m)
  fmt.Println(get([]string{"haha", "lol"}, m))
  fmt.Println(get([]string{"haha", "lmao"}, m))

  // graph data structure as a map
  // keys are string labels
  // value is a set of strings (could just as well be arrays or slices)
  type setstr map[string]emptystruct
  var graph = make(map[string]setstr)
  graph["v"] = map[string]emptystruct {
    "q": emptystruct{},
    "s": emptystruct{},
  }
  fmt.Println(graph)
}

func add(list []string, m map[string]int) {
  m[f(list)]++ //mutate the map
}
func get(list []string, m map[string]int) int {
  return m[f(list)]
}
func f(list []string) string { return fmt.Sprintf("%q", list) }

func map_equals(x, y map[int]int) bool {
  if len(x) != len(y) { return false } //len is always safe even for nil maps

  for key, value := range x {
    if yvalue, ok := y[key]; !ok || value != yvalue {
      return false
    }
  }
  return true
}

/*
   A struct is an aggregate data type that groups together 
   zero or more named values of arbitrary
   types as a single entity. Each value is called a field. 

   All of these fields are collected into a single entity 
   that can be copied as a unit, passed to functions and returned by them, 
   stored in arrays, and so on.
*/
type emp struct {
  ID   int
  Name string // order matters, otherwise defines a different type
  Dept string
} // newline for each field

func id_1(e emp) *emp{
  return &e
}

func id_2(e emp) emp{
  return e
}

func structs(){

  section("structs")

  type emp1 struct { // defines a different type since...
    ID   int
    Name, Dept string
  }
  /*
    Field order is significant to type identity. 
    Combining the declaration of the Position field (also a string), 
    or interchanging Name and Address, we would be defining a different
    struct type
  */

  var x emp // instance of an emp
  fmt.Println(x.Name) // fields are accessed using dot notation

  var y *emp = &x
  y.Name = "y"
  fmt.Println(y.Name) // dot notation can be used with pointers to structs as well

  (&x).Name = "x"
  fmt.Println((&x).Name) // dot notation can be used with pointers to structs

  id_1(*y).Dept ="Austin"
  //id_2(*y).Dept = "Boston"
  /* cannot assign to id_2(*y).Dept since the function 
     returns emp but not a (pointer to emp) *emp 
    (left-hand side does not identify a variable when id_2 is called and returned) 
  */

  /*
  The name of a struct field is exported if it begins with a capital letter; this is Go’s main access
  control mechanism. A struct type may contain a mixture of exported and unexported fields.
  */

  type mything struct{
    Exported int
    unexported int
  }
  f := mything{1, 2}

  fmt.Println(f.Exported) //prints 1
  fmt.Println(f.unexported) //prints 2
  // maybe this has to do with accessing members in another package?

  // struct literals
  // two forms:
  // (1) by order - tends to be used in the package defining the struct or for obvious filed ordering (2DPoint, 3DPoint)
  e := emp{1, "z", "music"}
  fmt.Println(e) // prints {1 z music}

  // (2) by name - you can omit other fields, they just take zero values
  e = emp{ID: 2, Name:"a"}
  fmt.Println(e) // prints {2 a  }

  // go is pass by value so people often pass structs by reference
  pp := &emp{ID: 1, Name: "b", Dept: "c" } //create and initialize a struct variable and obtain its address

  fmt.Println(pp) //&{1 b c}
  fmt.Println(pp.ID) //1
  fmt.Println(pp.Name) //b
  fmt.Println(pp.Dept) //c

  // the above shorthand notation is equivalent to:
  p := new(emp)
  *p = emp{1,"x", "y"}

  fmt.Println(p) //&{1 x y}
  fmt.Println(p.ID) //1
  fmt.Println(p.Name) //x
  fmt.Println(p.Dept) //y

  //anonymous fields and embeddings
  // fields with no name but a type can be specified in a struct
  // this allows us to reference fields in the type without using the full name
  type Point struct{X, Y int}
  //type Circle struct{Point radius float64} fields must be specified on newlines
  type Circle struct {
    Point
    radius float64
  }

  circ := Circle{Point{1,2}, 42.2}
  fmt.Println(circ.X, circ.Y)
  fmt.Println(circ.Point.X, circ.Point.Y)

  // compare with a non-anon struct
  type Square struct{corner1, corner2 Point}
  sq := Square{Point{2,2}, Point{0,0}}
  //fmt.Println(sq.X, sq.Y)  errs
  //fmt.Println(sq.Point.X, sq.Point.Y)  errs
  fmt.Println(sq.corner1.X, sq.corner2.Y)

}


func jsons(){
  section("json")
  // JSON is an encoding of JavaScript values— strings, numbers, booleans, arrays, and objects as Unicode text
  // It’s an efficient yet readable representation for the basic data types and the composite types of this chapter—arrays, slices, structs, and maps.

  // The basic JSON types are numbers (in decimal or scientific notation), booleans (true or false), and strings
  // These basic types may be combined recursively using JSON arrays and objects. 

  // A JSON array is an ordered sequence of values, written as a comma-separated list enclosed in square brackets; 
  // JSON arrays are used to encode Go arrays and slices. 
  // A JSON object is a mapping from strings to values, written as a sequence of name:value pairs separated by commas and surrounded by braces; 
  // JSON objects are used to encode Go maps (with string keys) and structs.

  type Song struct {
    Name string
    Year int
    Clean bool
    Artists []string
  }

  var Songs = []Song {
    {Name: "Shinobi", Year: 2019, Clean: false, Artists: []string{"Manila Grey"}},
    {Name: "Juggin", Year: 2019, Clean: false, Artists: []string{"MBNel", "Selfmade Cooly"}}, //<-- note the last comma needed
  }

  // Converting a Go data structure tp JSON is called marshaling. Marshaling is done by json.Marshal
  data, err := json.Marshal(Songs)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(data) // prints a byte slice by default ([]uint8)
  fmt.Printf("type: %T\n%s\n", data, data)

  // you can indent the Marshaled data
  dataindent, _ := json.MarshalIndent(Songs, "", "   ")
  fmt.Printf("%s\n", dataindent)

  // Marshaling uses the Go struct field names as the field names for the JSON objects (through
  // reflection). Only exported fields are marshaled (Fields starting with an uppercase)


  // Field tags
  // string literals after fields in structs; they are strings of metadata associated at compile time with the field of a struct
  type Book struct {
    Title string
    Year  int `json:"released"`
    Edition int `json:"edition,omitempty"`
    Authors []string
  }

  var Books = []Book {
    {Title: "Advanced Calculus", Year:2006, Authors: []string{"Patrick Fitzpatrick"}},
    {Title: "A Companion to Analysis", Year:2004, Edition: 0, Authors: []string{"T.W. Korner"}},
    {Title: "Invitation to Classical Analysis", Year:2012, Edition: 2, Authors: []string{"Peter Duren"}},
  }
  books, _ := json.MarshalIndent(Books, "", "   ")
  fmt.Printf("%s\n", books)

  // by convention, the field tags are interpreted as key:"value" pairs, but they can be any string literal
  // they are usually written as raw string literals (using `) since they contain double quotes
  // The json key controls the behavior of the encoding/json package
  //  - The key of the json field tag specifies an alternative name to use (usually used for idiomatic JSON naming)
  //  - tags have an additional option, omitempty, which indicates that no JSON output should be produced if there s no value (or the zero value) for that field in the struct instance
  // other encoding/... packages follow the same convention

  // Unmarshaling - taking JSON and transforming it into a Go data structure
  // first define a stuct instance whose fields you want from the JSON
  var titles []struct{ Title string }
  if err:= json.Unmarshal(books, &titles); err != nil {
    fmt.Println(err)
  }
  fmt.Println(titles)

}
