package main

import (
	"fmt"
    "math"
    "math/cmplx"
    "unicode/utf8"
    "strconv"
)

/*
basic types
- numbers, strings, and booleans

aggregate types
- arrays, structs

reference types (indirect access to variable/program state)
- pointers, functions, maps, slices, channels

interface types

*/
func main() {

	integers()
    floating()
    complex_()
	bits()
    booleans()
    strings_()

}

func section(header string) {
    pad:= "---------------------------"
	fmt.Printf("\n%s\n\t%s\n%s\n", pad, header, pad)
}

func integers() {

	section("integers")
	// integers come in various sizes of bits
	//signed
	var a int8  // 1 byte
	var b int16 // 2 bytes
	var c int32 // 4 bytes
	var d int64 // 8 bytes

	//unsigned
	var ua uint8
	var ub uint16
	var uc uint32
	var ud uint64

	// int and uint
	// could be 32 or 64 bits, can't assume which
	// chosen by compiler

	//rune ~ int32
	var r rune

	//byte ~ uint8
	var by byte

	// uintptr ~ sufficient width to hold an address value
	// size depends on platform
	var uptr uintptr
	fmt.Println(a, b, c, d, ua, ub, uc, ud, r, by, uptr)
}


func floating(){

  section("floating point")
  //floating32
  // limited precision 6 decimal digits (approx)
  fmt.Println("max float32: ", math.MaxFloat32)
  fmt.Println("small positive float32: ", 1 / math.MaxFloat32)

  //floating64
  // precision 15 decimal digits (approx)
  fmt.Println("max float64: ", math.MaxFloat64)
  fmt.Println("small positive float64: ", 1 / math.MaxFloat64)

  //literals
  fmt.Println("floating literals: ", .1, 1.2, 1., 1e2, 1.22324e32)

  // NaN literals
  fmt.Printf("NaN: %v, is NaN a NaN? %v\n", math.NaN(), math.IsNaN(math.NaN()))
  // all comparisons with NaN return false
  fmt.Printf("nan > 1: %v, nan == nan: %v\n", math.NaN() > 1, math.NaN() == math.NaN())

  var z float64
  fmt.Printf("1/0 = %v, 0/0 = %v, -1/0 = %v", 1/z, z/z, -1/z)

  // printf verbs
  // %g chooses the most compact repr
  // %<width>.<digits>f (no exponent form)
  // %e (exponent form)
  for i := 0; i < 8; i++ {
    fmt.Printf("x = %e, e^x = %8.3f\n", float64(i), math.Exp(float64(i)))
  }

}
func complex_(){

  section("complex")
  //complex64 has two floating32 components
  var z1 complex64 = complex(1, 8)

  //complex128 has two floating64 components
  var z2 complex128 = complex(2, 4)
  fmt.Printf("%v %v\n", z1, z2)

  //get the real and imaginary parts by using
  fmt.Printf("real(z1) = %f imag(z1) = %f\n", real(z1), imag(z1))

  //complex literals
  fmt.Printf("%v\n", (1 + 2i) + 3i + 2)
  fmt.Printf("%v\n", cmplx.Sqrt(-1))

}

func bits() {

	section("bits")
	// 42 = 32 + 8 + 2
	var left uint8 = 42
	fmt.Printf("left: %08b\n", left)

	var right uint8 = (255 >> 3)
	fmt.Printf("right: %08b\n", right)

	// & AND
	fmt.Printf("a & b: %08b\n", left&right)

	// | OR
	fmt.Printf("a | b: %08b\n", left|right)

	// ^ XOR
	fmt.Printf("a ^ b: %08b\n", left^right)

	// &^ AND NOT (bit clear)
	fmt.Printf("a &^ b: %08b\n", left&^right)

	// << shift left
	fmt.Printf("a << b: %08b\n", left<<right)

	// >> shift right
	fmt.Printf("a >> b: %08b\n", left>>right)

}

func booleans(){

  section("boolean")
  //literals: true , false
  fmt.Println(true, false)

  // operations &&, ||
  fmt.Println(true && true, true && false, false && false, true || true, true ||false)

  // ops use short circuiting behavior
  // if the left operand determines the result of the whole expression, the right operand is not evaluated

  // && has higher precedence than ||
}
func strings_(){

  section("string")
/* string
  immutable sequence of bytes
  bytes are interpreted as UTF-8 encoded sequences of unicode points (runes)
*/

  // len()
  fmt.Println("length of \"hello\": ", len("hello"))

  s := "hello"

  // index operation
  fmt.Println("1st and last char of \"hello\": ", s[0], s[len(s) - 1])
  // will print the corresponding ascii number for the characters

  // substring operator
  // s[i:j] gives a substring from i to j (not including j)
  // result is a slice of j-i bytes

  fmt.Println("slices of \"hello\":\n", s[1:])
  fmt.Println(s[1:2])
  fmt.Println(s[1:3])
  fmt.Println(s[1:4])
  // no new memory is allocated for the substring since 
  // the memory where the string is allocated is never mutated
  // bc strings are immutable

  // concatenation with +
  // comparing is done byte by byte so lexicographic order

  // string literals 
  // given using double quotes

  // every string in Go is encoded using utf-8
  // each character has an associated number called its "code point" or "rune"


  /* Unicode */

  // You can specify a unicode character in Go as a string literal by using unicode escapes
  // 16 bit format: \uhhhh
  // 32 bit format: \Uhhhhhhhh
  // where h is a hexadecimal digit (0 thru F ~ 16 values so each digit requires 4 bits, each hex pair is 1 byte)
  // ASCII character set has 256 characters so you only needed 8 bits to encode a character
  // the unicode character set encodes over 120,000 characters, so the number of bits needed to encode a character increase
  // UTF-32 character set uses 32 bits (4 bytes), however most characters (written in english) can fit in 8 bits.
  // So Go uses UTF-8, which is a variable (minimal) length encoding of characters used to not waste so much space
  // To think about how Unicode character sets work:
  // - each byte of a string is no longer necessarily a character
  // - a unicode character may takes anywhere from 1 to 4 bytes
  // - a unicode character in a string is a contiguous sequence of bytes
  // - unicode characters can be specified using unicode escapes
  // - runes are int32, they are a "character" of a unicode string

  uliteral16 := "\u4e16\u754c" // 16 bit value
  uliteral32 := "\U00004e16\U0000754c" // 32 bit value
  fmt.Println("2 16 bit unicode points: ", uliteral16)
  fmt.Println("the same 2 unicode points in 32 bit form: ", uliteral32)

  str := "hello, " + uliteral16
  fmt.Println("\"hello, \" + uliteral16:", str)

  fmt.Println("\nRanging over string literals decodes into UTF-8 to iterate over each rune not byte:")
  for i, v := range str {
    fmt.Printf("rune literal v: %v\tindex: %d\tstring(v): %v\tstr[index]: %v\n", v, i, string(v), str[i])
  }

  fmt.Println("\nNotice how indexing [] instead of range  allows you to get the bytes of the string")
  for i:= 0 ; i < len(str); i++ {
    fmt.Printf("i: %d\tstr[i]: %v\n", i, str[i])
  }

  /* Runes */
  //get the number of runes in a string
  fmt.Println("\nlen (# bytes) of \"hello\" + uliteral16: ", len(str))
  fmt.Println("RuneCountInString (number of runes): ", utf8.RuneCountInString(str))

  // Runes are numeric types of the corresponding unicode character
  rune1 := '\u4e16'
  rune2 := '\U00004e16'
  fmt.Println("\nrune literals (int32 base type): ", rune1, rune2)

  //iterating over runes
  fmt.Println("\nRanging over rune literals decodes UTF-8:")
  for i, v := range uliteral16 {
    fmt.Println("rune literal", v, "index: ", i, "string cast", string(v), "uliteral16[index]: ", uliteral16[i])
  }

  fmt.Println("\nDecodeRune")
  for i := 0; i < len(str); {
    r, size := utf8.DecodeRuneInString(str[i:]) //returns run and size in bytes of rune
    fmt.Printf("index: %d\tsize (bytes): %d\trune: %c\n", i, size, r)
    i += size
  }

  //A rune whose value is less than 256 may be written with a single hexadecimal escape
  //Anything higher needs a \u or \U
  rune3 := '\x41'
  fmt.Println(rune3) // prints out the decimal value ~ 65

  // A rune conversion of a UTF-8 encoded String returns the sequence of code points
  r := []rune(str)
  fmt.Println(r, string(r))

  /* strconv package */

  // convert integer to string using Sprintf or Itoa
  y := fmt.Sprintf("%d", 123)
  z := strconv.Itoa(123)
  fmt.Printf("type(y) = %T\ttype(z) = %T\n", y, z)

  x, _ := strconv.Atoi("123")
  fmt.Printf("type(x) = %T\ttype(z) = %T\n", x, z)

  // formatting integers using Printf verbs or FormatInt to string
  for i := 0; i < 16; i ++ {
    fmt.Println(strconv.FormatInt(int64(i), 2)) // int64 cast is nec.
    fmt.Println(fmt.Sprintf("x = %b", i))
  }

}

func constants(){

  section("constants")
}

















