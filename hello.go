package golearning

import (
	"fmt"
	"math"
	"strings"

	//"golang.org/x/tour/pic"
	"io"
	"strconv"
)

/* basic types
bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte // alias for uint8
rune // alias for int32
     // represents a Unicode code point
float32 float64
complex64 complex128
*/

// The zero value is:
// 0 for numeric types,
// false for the boolean type, and
// "" (the empty string) for strings.

func powWithLimit(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v after "if" statement scope
	return lim
}

// ErrNegativeSqrt is the error
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// TODO: without float64 type casting, program will go to infinit loop
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// MySqrt exported comment
func MySqrt(x float64) (float64, error) {
	if x < 0 {
		return 0.0, ErrNegativeSqrt(x)
	}
	acc := 0.000001

	z := 1.0
	lastZ := 1.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println("trying:", i, z)
		if math.Abs(lastZ-z) < acc {
			break
		} else {
			lastZ = z
		}
	}
	return z, nil
}

// IncreaseOne will use pointer to add 1
func IncreaseOne(pointer *int) {
	fmt.Println("address:", pointer, ", value: ", *pointer)
	*pointer++
}

func printIntArray(s []int) {
	// TODO: how to define general []T?
	// https://www.freecodecamp.org/news/generics-in-golang/
	// https://blog.golang.org/why-generics#TOC_3.
	// https://github.com/golang/go/wiki/InterfaceSlice
	fmt.Println("print slice custmoized way.")
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	for i, v := range s {
		fmt.Printf("index:[%d] = value:%d\n", i, v)
	}
}

func createDatas(dx, dy int) [][]uint8 {
	// need to import "golang.org/x/tour/pic"
	samples := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		samples[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			//samples[y][x]= uint8((x+y)/2)
			//samples[y][x]= uint8(x*y)
			//samples[y][x]= uint8(x^y)
			samples[y][x] = uint8(float64(x) * math.Log(float64(y)))
			//samples[y][x]= uint8(x%(y+1))
		}
	}
	return samples
}

func wordCount(s string) map[string]int {
	wordCount := make(map[string]int)
	for _, v := range strings.Fields(s) {
		_, ok := wordCount[v]
		if ok {
			wordCount[v]++
		} else {
			wordCount[v] = 1
		}
	}
	return wordCount
}

// function value
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// function closure
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// function closure
func fibonacci() func() int {
	fn, fn1 := 0, 1
	return func() int {
		temp := fn
		fn, fn1 = fn1, fn+fn1
		return temp
	}
}

// Vertex type can have function then work as class
type Vertex struct {
	X, Y float64
}

// Scale method, pointer receiver can be used to update the value directly
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Abs is the implementation of interface Abser
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//
func (v Vertex) String() string {
	return fmt.Sprintf("Vertex.X=%v Vertex.Y=%v", v.X, v.Y)
}

// Abser interface
type Abser interface {
	Abs() float64
}

type myFloat float64

func (f myFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// assert type by interface
func assertType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%v is int\n", v)
	case float64:
		fmt.Printf("%v is float64\n", v)
	case string:
		fmt.Printf("%q is %v bytes long string\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

// IPAddr is the type
type IPAddr [4]byte

// String method for IPAddr
func (ip IPAddr) String() string {
	output := make([]string, 4)
	for i, v := range ip {
		// TODO: more useful about strconv?
		output[i] = strconv.Itoa(int(v))
	}
	a := strings.Join(output, ".")
	//b := fmt.Sprintf("%v.%v.%v.%v", ip[0],ip[1],ip[2],ip[3])
	return a
}

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(b []byte) (int, error) {
	n, err := rot13.r.Read(b)
	for i, v := range b {
		switch {
		case int(v) > 122: // > z
			//fmt.Println("invalid input string")
			break
		case int(v) > 109: // n-z
			b[i] -= 13
		case int(v) > 96: // a-m
			b[i] += 13
		case int(v) > 90: // > Z
			//fmt.Println("invalid input string")
			break
		case int(v) > 77: // N-Z
			b[i] -= 13
		case int(v) > 64: // A-M
			b[i] += 13
		}
	}
	return n, err
}
