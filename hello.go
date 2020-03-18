package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
	"math/rand"
	"strings"
	//"golang.org/x/tour/pic"
	"strconv"
	"io"
	"os"
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

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return lim
	}
	return lim
}

// ErrNegativeSqrt is the error
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
	// TODO: without float64 type casting, program will go to infinit loop
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt exported comment
func Sqrt(x float64) (float64, error) {
	if x<0 {
		return 0.0, ErrNegativeSqrt(x)
	}
	
	z := 1.0
	lastZ := 1.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println("trying:", i, z)
		if math.Abs(lastZ-z) < 0.0001 {
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


func printSlice(s []int) {
	// TODO: how to define general []T? 
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	for i, v := range s {
		fmt.Printf("index:[%d] = value:%d\n", i, v)
	}
}


func createDatas(dx, dy int) [][]uint8 {
	// need to import "golang.org/x/tour/pic"
	samples := make([][]uint8, dy)
	for y:=0;y<dy;y++{
		samples[y] = make([]uint8,dx)
		for x:=0;x<dx;x++{
			//samples[y][x]= uint8((x+y)/2)
			//samples[y][x]= uint8(x*y)
			//samples[y][x]= uint8(x^y)
			samples[y][x] = uint8(float64(x)*math.Log(float64(y)))
			//samples[y][x]= uint8(x%(y+1))
		}
	}
	return samples
}


func wordCount(s string) map[string]int {
	wordCount := make(map[string]int)
	for _, v := range strings.Fields(s) {
		_, ok := wordCount[v]
		if ok == true {
			wordCount[v]++
		}else{
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
	return func() int{
		temp := fn
		fn, fn1 = fn1, fn + fn1
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
func assertType(i interface{}){
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
func (ip IPAddr) String() string{
	output := make([]string, 4)
	for i, v := range ip{
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

func (rot13 rot13Reader) Read(b []byte) (int, error){
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

func main() {
	// constant cannot be deleared using := syntax
	// variable can be decleared in group
	const (
		Pi = 3.14
		// Create a huge number by shifting a 1 bit left 100 places.
		// In other words, the binary number that is 1 followed by 100 zeroes.
		Big = 1 << 100
		// Shift it right again 99 places, so we end up with 1<<1, or 2.
		Small = Big >> 99
	)

	// final print defer stack
	for i := 0; i < 3; i++ {
		defer fmt.Println("defer stack:", i)
	}

	// random number
	fmt.Println("My favorite number is", rand.Intn(10))

	// declear more variable at the sametime with the same type
	var test1, test2, test3 int = 10, 20, 5

	// local variable can not be used ourside
	// type convension
	fmt.Println(pow(3, 2, float64(test1)), pow(3, 3, float64(test2)))

	// try for loop
	fmt.Println(Sqrt(float64(test3)))

	// try switch case for string
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// switch case for value
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// pointer usage
	v := 1
	IncreaseOne(&v)
	fmt.Println("value after increasing by pointer: ", v)

	// struct
	vex := Vertex{Y:1}
	p := &vex
	fmt.Println("struct vertex:", vex)
	fmt.Println("first value of vertex:", vex.X)
	fmt.Println("access vertex by pointer:", p.Y)

	// slice
	var empty []int
	a := []int{0,1,2,3,4,5,6,7}
	fmt.Print("created slice a:")
	printSlice(a)
	fmt.Print("a[1:5]：")
	printSlice(a[1:5])
	newA := append(a, 8)
	fmt.Print("append slice newA:")
	printSlice(newA)
	// TODO: why not just expand 1 cap?
	b := make([]int, 3)
	fmt.Print("slice created by make:")
	printSlice(b)
	if empty == nil {
		fmt.Println("this slice is nil")
		printSlice(empty)
	}
	// 创建一个井字板（经典游戏）
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// 两个玩家轮流打上 X 和 O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	fmt.Printf("Fields are: %q\n", strings.Fields("  foo bar  baz   "))
	
	// function value
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	// function closure
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	v1 := Vertex{3, 4}
	fmt.Println(v1)
	v1.Scale(10)
	fmt.Println(v1.Abs())

	// interface usage
	var inter Abser
	f := myFloat(-1)
	inter = &v1
	fmt.Println(inter.Abs())
	inter = f
	fmt.Println(inter.Abs())

	// assertType by empty interface
	assertType("this is string")
	assertType(3.1)
	assertType(0)
	assertType(nil)
	assertType(v1)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	result, err := Sqrt(2)
	if err == nil{
		fmt.Println(result)
	} else{
		fmt.Println(err)
	}

	result, err = Sqrt(-2)
	if err == nil{
		fmt.Println(result)
	} else{
		fmt.Println(err)
	}

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
