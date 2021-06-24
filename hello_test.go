package golearning

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/tour/tree"
)

// constant cannot be deleared using := syntax
// variable can be decleared in group
const (
	Pi = 3.14
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99

	MaxInt uint64 = 1<<64 - 1
)

func TestDeferStack(t *testing.T) {
	fmt.Println("counting")

	// finally print defer stack
	for i := 0; i < 3; i++ {
		defer fmt.Println("defer stack:", i)
	}
	fmt.Println("print something")
}

func TestRandomNumber(t *testing.T) {
	// random number
	seed := 100
	fmt.Println("My favorite number is", rand.Intn(seed))
}

func TestTypeConversion(t *testing.T) {
	// declear more variable at the sametime with the same type
	var test1, test2, test3 int = 10, 20, 5

	// local variable can not be used ourside
	// type convension
	fmt.Println(powWithLimit(3, 2, float64(test1)), powWithLimit(3, 3, float64(test2)))

	// try for loop
	fmt.Println(MySqrt(float64(test3)))
	fmt.Println("math.Sqrt result:", math.Sqrt(5))
	fmt.Println(MySqrt(-1.0))
}

func TestSwitchStatement(t *testing.T) {
	// try switch case for string
	fmt.Print("Go runs on ")
	os := runtime.GOOS
	switch os {
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

	// switch with true
	n := time.Now()
	switch {
	case n.Hour() < 12:
		fmt.Println("Good morning!")
	case n.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func TestPointer(t *testing.T) {
	// pointer usage
	v := 1
	IncreaseOne(&v)
	fmt.Println("value after increasing by pointer: ", v)
}

func TestStruct(t *testing.T) {
	// struct
	vex := Vertex{Y: 1}
	p := &vex
	assert.Equal(t, "Vertex.X=0 Vertex.Y=1", fmt.Sprintf("%v", vex))
	fmt.Println("first value of vertex:", vex.X)
	fmt.Println("access vertex by pointer:", p.Y)

	// create by pointer
	v1 := &Vertex{3, 4}
	fmt.Println(v1)
	v1.Scale(10)
	fmt.Println(v1.Abs())
}

func TestInterface(t *testing.T) {
	// interface usage
	var abser Abser

	v1 := Vertex{3, 4}
	abser = &v1
	fmt.Println(abser.Abs())

	f := myFloat(-1)
	abser = f
	fmt.Println(abser.Abs())

	// assertType by empty interface
	assert.Equal(t, "string", assertType("this is string"))
	assert.Equal(t, "float64", assertType(3.1))
	assert.Equal(t, "int", assertType(0))
	assert.Equal(t, "nil", assertType(nil))
	assert.Equal(t, "Vertex", assertType(v1))
}
func TestArraySlice(t *testing.T) {
	// Array
	var a [2]int
	a[0] = 1
	a[1] = 2
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	printIntArray(a[0:])

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// slice
	var empty []int
	as := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("created slice a:", as)
	printIntArray(as)
	fmt.Println("a[1:5]：", as[1:5])
	newA := append(as, 8)
	fmt.Println("append slice newA:", newA)
	printIntArray(newA)
	// TODO: why not just expand 1 cap?
	b := make([]int, 3)
	fmt.Println("slice created by make:", b)
	printIntArray(b)
	if empty == nil {
		fmt.Println("this slice is nil")
		printIntArray(empty)
	}
}

func TestBoard(t *testing.T) {
	// 创建一个井字板（经典游戏）
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
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
}

func TestStringsFunctions(t *testing.T) {
	fmt.Printf("Fields are: %q\n", strings.Fields("  foo bar  baz   "))
}

func TestStringField(t *testing.T) {
	s := "a b a b c c d a a a"
	expected := map[string]int{"a": 5, "b": 2, "c": 2, "d": 1}
	assert.Equal(t, expected, wordCount(s))
}

func TestFunctionReference(t *testing.T) {
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
		fmt.Println(pos(i), neg(-2*i))
	}
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func TestStringer(t *testing.T) {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	expected := map[string]string{
		"loopback":  "127.0.0.1",
		"googleDNS": "8.8.8.8",
	}
	for name, ip := range hosts {
		assert.Equal(t, expected[name], fmt.Sprintf("%v", ip))
	}
}

func TestError(t *testing.T) {
	result, err := MySqrt(2)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

	result, err = MySqrt(-2)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}

func TestReader(t *testing.T) {
	// ROT-13 encryption and read from reader
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	b := make([]byte, 64)
	n, _ := r.Read(b)
	assert.Equal(t, "You cracked the code!", string(b[:n]))
}

func TestTypeAssertion(t *testing.T) {
	var i interface{} = "hello"
	s := i.(string)
	assert.Equal(t, "string", reflect.TypeOf(s).String())

	var i2 interface{} = 9.2
	s2, ok := i2.(float64)
	assert.Equal(t, "float64", reflect.TypeOf(s2).String())
	assert.Equal(t, true, ok)

	f, ok := i.(float64)
	assert.Equal(t, "float64", reflect.TypeOf(f).String())
	assert.Equal(t, 0.0, f)
	assert.Equal(t, false, ok)

	assert.Panics(t, func() { f = i.(float64) }, "The code did not panic")
}

func TestGoroutines(t *testing.T) {
	go say("folk:world")
	say("main:hello")

	count := int(math.Pow(2, 20))
	s := make([]int, count)
	for i := 0; i < count; i++ {
		s[i] = i
	}
	start := time.Now()
	fmt.Println(sum1(s))
	duration1 := time.Since(start).Nanoseconds()
	fmt.Printf("sum directly takes %v nano seconds\n", duration1)

	start2 := time.Now()
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x + y)
	duration2 := time.Since(start2).Nanoseconds()
	fmt.Printf("sum by 2 routies takes %v nano seconds\n", duration2)

	diff := duration1 - duration2
	perc := float64(diff) / float64(duration1)
	fmt.Printf("you saved %.2f%%, %v nano seconds\n", perc*100, diff)
}

func TestChannelRangeClose(t *testing.T) {
	c := make(chan int, 10)
	go chanFibonacci(cap(c), c)
	f := fibonacci()
	for v := range c {
		assert.Equal(t, v, f())
	}
}

func TestSelect(t *testing.T) {
	c := make(chan int)
	quit := make(chan int)
	count := 100
	f := fibonacci()
	go func() {
		for i := 0; i < count; i++ {
			assert.Equal(t, <-c, f())
		}
		quit <- 0
	}()
	selectFibonacci(c, quit)
}

func TestChannel(t *testing.T) {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := 1; i < 11; i++ {
		assert.Equal(t, i, <-ch)
	}
	assert.Equal(t, true, Same(tree.New(1), tree.New(1)))
	assert.Equal(t, false, Same(tree.New(1), tree.New(2)))
}
