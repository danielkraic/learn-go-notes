# golang 

## tools

```bash
go version

go help
go help gopath

godoc -http=":6060"

godoc fmt Println
godoc fmt Println

go build
go build -race
go run main.go
go install
go install github.com/$USER/hello

go test
go test github.com/$USER/stringutil

go get github.com/golang/example/hello
```

```bash
gofmt file.go
go vet //linter
go tool cover
```

### gopath

```bash
export GOPATH="$HOME/work/go"
export PATH="$PATH:$GOPATH/bin"

mkdir -p $GOPATH
mkdir -p $GOPATH/src/github.com/$USER

go env // validate
```

## godep

```bash
export GOPATH=$HOME/example
go get github.com/goinggo/newssearch

cd projetc_dir

godep save -copy=false
godep go build

// make changes
godep save
godep go build
```

## types

```go
uint8, uint16, uint32, uint64
int8, int16, int32, int64
byte // uint8
rune // int32
uint, int, uintptr // machine dependent

float32, float64
complex64, complex128
```

## vars

```go
var v1 string
v1 = "val1"

var v2 sring = "val2"
var v3 = "val3"

v4 := "val4"

var b, c int = 1, 2

```

## const

```go
const (
    a = 5
    b = 10
)
```

## control structures

```go
if i % 2 == 0 {
  // divisible by 2
} else if i % 3 == 0 {
  // divisible by 3
} else if i % 4 == 0 {
  // divisible by 4
}
```

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

i := 1
for i <= 3 {
    fmt.Println(i)
    i = i + 1
}

for {
    fmt.Println("loop")
    break
}
```

```
switch i {
    case 0: fmt.Println("Zero")
    case 1: fmt.Println("One")
    case 2: fmt.Println("Two")
    case 3: fmt.Println("Three")
    case 4: fmt.Println("Four")
    case 5: fmt.Println("Five")
    default: fmt.Println("Unknown Number")
}

switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("it's the weekend")
    default:
        fmt.Println("it's a weekday")
}
```

## array

```go
var x [5]int
x[0] = 10

var y [3]int{1,2,3}
z := [3]int{1,2,3}
```

## slice

```go
x := make([]float64, 5)
y := make([]float64, 5, 10)
```

```go
a := arr[2:5]
b := arr[:5]
c := arr[2:]
```

```go
size := len(x)
x = append(x, item)

slice1 := []int{1,2,3}
slice2 := make([]int, 2)
copy(slice2, slice1)
```

## map

```go
x := make(map[string]int)
x["key"] = 10

if val, ok := x["key"]; ok {
  fmt.Println(val, ok)
}

delete(x, "key")
```

```go
elements := map[string]map[string]string{
    "H": map[string]string{
      "name":"Hydrogen",
      "state":"gas",
    },
    "He": map[string]string{
      "name":"Helium",
      "state":"gas",
    },
}

for k, v := range mymap {
    //
}
```

## functions

```go
func average(xs []float64) float64 {
  panic("Not Implemented")
}

func f() (int, int) {
  return 5, 6
}
```

```go
// Variadic Function
func add(args ...int) int {
  total := 0
  for _, v := range args {
    total += v
  }
  return total
}

func main() {
  fmt.Println(add(1,2,3))
}
```

```go
// closure
func makeEvenGenerator() func() uint {
  i := uint(0)
  return func() (ret uint) {
    ret = i
    i += 2
    return
  }
}
func main() {
  nextEven := makeEvenGenerator()
  fmt.Println(nextEven()) // 0
  fmt.Println(nextEven()) // 2
  fmt.Println(nextEven()) // 4
}
```

## defer, panic, recover

```go
func main() {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("%v", r)
        }
    }()
    panic("PANIC")
}
```

## struct

```go
type Circle struct {
  x float64
  y float64
  r float64
}

var c1 Circle
c1.x = 10.6

c2 := new(Circle)

c3 := Circle{x: 0, y: 0, r: 5}
c4 := Circle{0, 0, 5}

func (c *Circle) area() float64 {
  return math.Pi * c.r*c.r
}
```

```go
type Person struct {
  Name string
}
func (p *Person) Talk() {
  fmt.Println("Hi, my name is", p.Name)
}

// type Android struct {
//   Person Person
//   Model string
// }

type Android struct {
  Person
  Model string
}

a1 := new(Android)
a1.Person.Talk()
// or
a2.Talk()
```

# interface

```go
type Shape interface {
  area() float64
}

func totalArea(shapes ...Shape) float64 {
  var area float64
  for _, s := range shapes {
    area += s.area()
  }
  return area
}

type MultiShape struct {
  shapes []Shape
}

func (m *MultiShape) area() float64 {
  var area float64
  for _, s := range m.shapes {
    area += s.area()
  }
  return area
}
```

## goroutine

```go
go func(msg string) {
    fmt.Println(msg)
}("going")
```

```go
func f(n int) {
  for i := 0; i < 10; i++ {
    fmt.Println(n, ":", i)
  }
}

for i := 0; i < 10; i++ {
    go f(i)
}
```

```go
jobs := make(chan int, 5)
done := make(chan bool)

// Here's the worker goroutine. It repeatedly receives
// from `jobs` with `j, more := <-jobs`. In this
// special 2-value form of receive, the `more` value
// will be `false` if `jobs` has been `close`d and all
// values in the channel have already been received.
// We use this to notify on `done` when we've worked
// all our jobs.
go func() {
    for {
        j, more := <-jobs
        if more {
            fmt.Println("received job", j)
        } else {
            fmt.Println("received all jobs")
            done <- true
            return
        }
    }
}()

// This sends 3 jobs to the worker over the `jobs`
// channel, then closes it.
for j := 1; j <= 3; j++ {
    jobs <- j
    fmt.Println("sent job", j)
}
close(jobs)
fmt.Println("sent all jobs")

// We await the worker using the
// [synchronization](channel-synchronization) approach
// we saw earlier.
<-done
```

```go
queue := make(chan string, 2)
queue <- "one"
queue <- "two"
close(queue)

for elem := range queue {
    fmt.Println(elem)
}
```

```go
func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    // Send a value to notify that we're done.
    done <- true
}

func main() {

    // Start a worker goroutine, giving it the channel to
    // notify on.
    done := make(chan bool, 1)
    go worker(done)

    // Block until we receive a notification from the
    // worker on the channel.
    <-done
}
```

```go
func pinger(c chan<- string)
  for i := 0; ; i++ {
    c <- "ping"
  }
}

func ponger(c chan<- string)
  for i := 0; ; i++ {
    c <- "pong"
  }
}

func printer(c <-chan string)
  for {
    msg := <- c
    fmt.Println(msg)
    time.Sleep(time.Second * 1)
  }
}

func main() {
  var c chan string = make(chan string)

  go pinger(c)
  go ponger(c)
  go printer(c)

  var input string
  fmt.Scanln(&input)
}
```

```go
c1 := make(chan string)
c2 := make(chan string)

go func() {
for {
    c1 <- "from 1"
    time.Sleep(time.Second * 2)
}
}()

go func() {
for {
    c2 <- "from 2"
    time.Sleep(time.Second * 3)
}
}()

go func() {
for {
    select {
    case msg1 := <- c1:
    fmt.Println(msg1)
    case msg2 := <- c2:
    fmt.Println(msg2)
    }
}
}()
```


```go
select {
case msg1 := <- c1:
  fmt.Println("Message 1", msg1)
case msg2 := <- c2:
  fmt.Println("Message 2", msg2)
case <- time.After(time.Second):
  fmt.Println("timeout")
default:
  fmt.Println("nothing ready")
}
```

## core packages

```go
// true
strings.Contains("test", "es")

// 2
strings.Count("test", "t")

// true
strings.HasPrefix("test", "te")

// true
strings.HasSuffix("test", "st")

// 1
strings.Index("test", "e")

// "a-b"
strings.Join([]string{"a","b"}, "-")

// == "aaaaa"
strings.Repeat("a", 5)

// "bbaa"
strings.Replace("aaaa", "a", "b", 2)

// []string{"a","b","c","d","e"}
strings.Split("a-b-c-d-e", "-")

// "test"
strings.ToLower("TEST")

// "TEST"
strings.ToUpper("test")
```

```go
import "errors"

func main() {
  err := errors.New("error message")
}

type argError struct {
    arg  int
    prob string
}
func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}
func f() error {
    return  &argError{arg, "can't work with it"}
}
```

```go
// open file
import (
  "fmt"
  "os"
)

func main() {
    file, err := os.Open("test.txt")
    if err != nil {
        // handle the error here
        return
    }
    defer file.Close()

    // get the file size
    stat, err := file.Stat()
    if err != nil {
        return
    }
    // read the file
    bs := make([]byte, stat.Size())
    _, err = file.Read(bs)
    if err != nil {
       return
    }
}
```

```go
// read file
import (
  "fmt"
  "io/ioutil"
)

func main() {
  bs, err := ioutil.ReadFile("test.txt")
  if err != nil {
    return
  }
  str := string(bs)
  fmt.Println(str)
}
```

```go
// create file
import (
  "os"
)

func main() {
  file, err := os.Create("test.txt")
  if err != nil {
    // handle the error here
    return
  }
  defer file.Close()

  file.WriteString("test")
}
```

```go
// read dir
import (
  "fmt"
  "os"
)

func main() {
  dir, err := os.Open(".")
  if err != nil {
    return
  }
  defer dir.Close()

  fileInfos, err := dir.Readdir(-1)
  if err != nil {
    return
  }
  for _, fi := range fileInfos {
    fmt.Println(fi.Name())
  }
}
```

```go
// container - list
import ("fmt" ; "container/list")

func main() {
  var x list.List
  x.PushBack(1)
  x.PushBack(2)
  x.PushBack(3)

  for e := x.Front(); e != nil; e=e.Next() {
    fmt.Println(e.Value.(int))
  }
}
```

```go
// sort
import ("fmt" ; "sort")

type Person struct {
  Name string
  Age int
}

type ByName []Person

func (this ByName) Len() int {
  return len(this)
}
func (this ByName) Less(i, j int) bool {
  return this[i].Name < this[j].Name
}
func (this ByName) Swap(i, j int) {
  this[i], this[j] = this[j], this[i]
}

func main() {
  kids := []Person{
    {"Jill",9},
    {"Jack",10},
  }
  sort.Sort(ByName(kids))
  fmt.Println(kids)
}
```

```
resp, err := http.Get(uri)
if err != nil {
    panic(err)
}

defer resp.Body.Close()

// Convert the response to a byte array
rawDocument, err = ioutil.ReadAll(resp.Body)
if err != nil {
    //
}
```

```go
// http server
import ("net/http" ; "io")

func hello(res http.ResponseWriter, req *http.Request) {
  res.Header().Set(
    "Content-Type",
    "text/html",
  )
  io.WriteString(
    res,
    `<DOCTYPE html>
<html>
  <head>
      <title>Hello World</title>
  </head>
  <body>
      Hello World!
  </body>
</html>`,
  )
}
func main() {
  http.HandleFunc("/hello", hello)
  http.Handle(
    "/assets/",
    http.StripPrefix(
        "/assets/",
        http.FileServer(http.Dir("assets")),
    ),
    )
  http.ListenAndServe(":9000", nil)
}
```

```go
// cmd flags
import ("fmt";"flag";"math/rand")

func main() {
  // Define flags
  maxp := flag.Int("max", 6, "the max value")
  // Parse
  flag.Parse()
  // Generate a number between 0 and max
  fmt.Println(rand.Intn(*maxp))
}
```

```go
// mutex
import (
  "fmt"
  "sync"
  "time"
)

func main() {
  m := new(sync.Mutex)

  for i := 0; i < 10; i++ {
    go func(i int) {
      m.Lock()
      fmt.Println(i, "start")
      time.Sleep(time.Second)
      fmt.Println(i, "end")
      m.Unlock()
    }(i)
  }
}
```

## testing

```go
import "testing"

func TestAverage(t *testing.T) {
  var v float64
  v = Average([]float64{1,2})
  if v != 1.5 {
    t.Error("Expected 1.5, got ", v)
  }
}

// run with `go test`
```

```go
// sync.waitGroup - https://www.goinggo.net/2014/01/concurrency-goroutines-and-gomaxprocs.html
import (
    "fmt"
    "runtime"
    "sync"
)

func main() {
    runtime.GOMAXPROCS(2)

    var wg sync.WaitGroup
    wg.Add(2)

    fmt.Println("Starting Go Routines")
    go func() {
        defer wg.Done()

        for char := ‘a’; char < ‘a’+26; char++ {
            fmt.Printf("%c ", char)
        }
    }()

    go func() {
        defer wg.Done()

        for number := 1; number < 27; number++ {
            fmt.Printf("%d ", number)
        }
    }()

    fmt.Println("Waiting To Finish")
    wg.Wait()

    fmt.Println("\nTerminating Program")
}
```

## mongo 

```go
// https://www.goinggo.net/2014/02/running-queries-concurrently-against.html

package main

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"sync"
	"time"
)

const (
	MongoDBHosts = "ds035428.mongolab.com:35428"
	AuthDatabase = "goinggo"
	AuthUserName = "guest"
	AuthPassword = "welcome"
	TestDatabase = "goinggo"
)

type (
	// BuoyCondition contains information for an individual station.
	BuoyCondition struct {
		WindSpeed     float64 `bson:"wind_speed_milehour"`
		WindDirection int     `bson:"wind_direction_degnorth"`
		WindGust      float64 `bson:"gust_wind_speed_milehour"`
	}

	// BuoyLocation contains the buoy's location.
	BuoyLocation struct {
		Type        string    `bson:"type"`
		Coordinates []float64 `bson:"coordinates"`
	}

	// BuoyStation contains information for an individual station.
	BuoyStation struct {
		ID        bson.ObjectId `bson:"_id,omitempty"`
		StationId string        `bson:"station_id"`
		Name      string        `bson:"name"`
		LocDesc   string        `bson:"location_desc"`
		Condition BuoyCondition `bson:"condition"`
		Location  BuoyLocation  `bson:"location"`
	}
)

// main is the entry point for the application.
func main() {
	// We need this object to establish a session to our MongoDB.
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{MongoDBHosts},
		Timeout:  60 * time.Second,
		Database: AuthDatabase,
		Username: AuthUserName,
		Password: AuthPassword,
	}

	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}

	// Reads may not be entirely up-to-date, but they will always see the
	// history of changes moving forward, the data read will be consistent
	// across sequential queries in the same session, and modifications made
	// within the session will be observed in following queries (read-your-writes).
	// http://godoc.org/labix.org/v2/mgo#Session.SetMode
	mongoSession.SetMode(mgo.Monotonic, true)

	// Create a wait group to manage the goroutines.
	var waitGroup sync.WaitGroup

	// Perform 10 concurrent queries against the database.
	waitGroup.Add(10)
	for query := 0; query < 10; query++ {
		go RunQuery(query, &waitGroup, mongoSession)
	}

	// Wait for all the queries to complete.
	waitGroup.Wait()
	log.Println("All Queries Completed")
}

// RunQuery is a function that is launched as a goroutine to perform
// the MongoDB work.
func RunQuery(query int, waitGroup *sync.WaitGroup, mongoSession *mgo.Session) {
	// Decrement the wait group count so the program knows this
	// has been completed once the goroutine exits.
	defer waitGroup.Done()

	// Request a socket connection from the session to process our query.
	// Close the session when the goroutine exits and put the connection back
	// into the pool.
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(TestDatabase).C("buoy_stations")

	log.Printf("RunQuery : %d : Executing\n", query)

	// Retrieve the list of stations.
	var buoyStations []BuoyStation
	err := collection.Find(nil).All(&buoyStations)
	if err != nil {
		log.Printf("RunQuery : ERROR : %s\n", err)
		return
	}

	log.Printf("RunQuery : %d : Count[%d]\n", query, len(buoyStations))
}
```
