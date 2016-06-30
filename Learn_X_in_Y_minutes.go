
// Learn X in Y minutes
// Where X=Go
// http://learnxinyminutes.com/docs/go/

package main

import (
  "fmt"
  m "math"
  "strconv"
)

func main() {
  fmt.Println("hello  world")

  beyondHello()
}

func beyondHello() {
  var x int
  x = 4
  y := 5

  sum, prod := learnMultiple(x, y)
  fmt.Println("sum: ", sum, ", prod: ", prod)

  learnTypes()
}

func learnMultiple(x, y int) (sum, prod int) {
  return x + y, x * y
}

func learnTypes() {
  var a4 [4]int
  a3 := [...]int{3, 1, 5}

  s3 := []int{4, 5, 9}
  s4 := make([]int, 4)
  var d2 [][]float64
  bs := []byte("a slice")

  _, _, _, _, _, _ = a4, a3, s3, s4, d2, bs

  s := []int{1, 2, 3}
  s = append(s, 4, 5, 6)
  fmt.Println(s)

  s = append(s, []int{7, 8, 9}...)
  fmt.Println(s)

  p, q := learnMemory()
  fmt.Println(*p, *q)

  m := map[string]int{"three": 3, "four": 4}
  m["one"] = 1
  fmt.Println(m)

  learnFlowControl()
}

func learnMemory() (p, q *int) {
  //p = new(int)
  s := make([]int, 20)
  s[3] = 7
  r := -2

  return &s[3], &r
}

func learnFlowControl() {
  if true {
    fmt.Println("yes")
  } else {
    fmt.Println("no")
  }

  for i := 0; i < 3; i++ {
    fmt.Println("iteration ", i)
  }

  for key, value := range map[string]int{"one":1, "two":2, "three":3} {
    fmt.Printf("%s: %d\n", key, value)
  }

  x := 42.0
  if y := expensiveComputation(); y > x {
    x = y
  }

  // closure
  xBig := func() bool {
    return x > 10000
  }

  fmt.Printf("x: %.2f\n", x)
  fmt.Println("xBig: ", xBig)
  x = 1.3e3
  fmt.Println("xBig: ", xBig)

  fmt.Println("add + double: ",
    func(a, b int) int {
      return (a+b)*2
    }(10, 2))

  learnFuncFactory()
  learnDefer()
  learnInterface()
}

func expensiveComputation() float64 {
  return m.Exp(10)
}

func learnFuncFactory() {
  fmt.Println(funcFactory("middle")("bagin", "end"))

  d := funcFactory("Middle")

  fmt.Println(d("Begin", "End"))
}

func funcFactory(mystring string) func(before, after string) string {
  return func(before, after string) string {
    return fmt.Sprintf("%s:%s:%s", before, mystring, after)
  }
}


func learnDefer() bool {
  defer fmt.Println("Printed second")
  defer fmt.Println("Printed first (because of LIFO order)")

  return true
}


type Stringer interface {
  String() string
}

type pair struct {
  x, y int
}

func (p pair) String() string {
  return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func learnInterface() {
  p := pair{1,5}
  var i Stringer
  i = p

  fmt.Println(p.String())
  fmt.Println(i.String())

  fmt.Println(p)
  fmt.Println(i)

  learnVariadicparams()
}

func learnVariadicparams(vals ...interface{}) {
  for _, val := range vals {
    fmt.Println("param: ", val)
  }

  //fmt.Println("params: ", fmt.Sprintln(vals...))

  learnErrorHandling()
}

func learnErrorHandling() {
  m := map[int]string{1:"one",2:"two", 3:"three"}
  if _, ok := m[1]; !ok {
    fmt.Println("no one there")
  } else {
    fmt.Println("one founded")
  }

  if val, err := strconv.Atoi("no-number"); err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("strconv.Atoi val: ", val)
  }

  learnConcurrency()
}

func inc(i int, c chan int) {
  c <- i + 1  // send to channel
}

func learnConcurrency() {
  c := make(chan int)

  go inc(0, c)
  go inc(10, c)
  go inc(-805, c)

  fmt.Println(<-c, <-c, <-c) // receive operator
}
