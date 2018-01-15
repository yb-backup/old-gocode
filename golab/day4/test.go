package main

import "fmt"
import "io/ioutil"
import "log"
import "math"
import "net/http"
import "os"
import "sort"
import "strconv"
import "strings"
import "time"

func main() {
	const pi float64 = 3.14159265
	var myName string = "nangbin"
	fmt.Println("pi = ", pi)
	fmt.Println("myName = ", myName)

	var isOver40 bool = true
	fmt.Println("isOver40 : ", isOver40)

	fmt.Printf("%t \n", isOver40)
	fmt.Printf("%T \n", isOver40)

	favNum3 := [5]float64{1, 2, 3, 4, 5}

	for i, v := range favNum3 {
		fmt.Println(i, v)
	}

	numSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("numSlice[:2] = ", numSlice[:2])

	numSlice3 := make([]int, 5, 10)
	copy(numSlice3, numSlice)
	fmt.Println(numSlice3)

	presAge := make(map[string]int)
	presAge["yang"] = 18
	fmt.Println(len(presAge))
	presAge["yang1"] = 18
	fmt.Println(len(presAge))
	fmt.Println("presAge[\"yang\"]= ", presAge["yang"])

	numbers := []float64{0.1, 0.2, 12}
	fmt.Println("sum = ", addThemUp(numbers))

	fmt.Println("sum = ", subtractThem(1, 2, 3))
	fmt.Println("sum = ", subtractThem(1, 2, 3, 4, 5))

	// defer printTwo()
	// printOne()

	fmt.Println("=======================")
	fmt.Println(safeDiv(3, 0))
	fmt.Println(safeDiv(3, 2))

	x := 0
	changeXVal(x)
	fmt.Println("x = ", x)
	changeXValNow(&x)
	fmt.Println("x = ", x)

	fmt.Println("x Memory Address : ", &x)

	fmt.Println("==============================")
	rect1 := Rectangle{1.0, 2.0, 3, 4}
	fmt.Println(rect1.height, rect1.width)
	fmt.Println("area of rect : ", rect1.area())

	fmt.Println("rect area : ", getArea(rect1))
	circle := Circle{2.2}
	fmt.Println("Circle area : ", getArea(circle))
	fmt.Println("==============================")

	csvString := "1, 2, 3, 4, 5"
	fmt.Println(strings.Split(csvString, ","))

	randInts := []int{3, 2, 1, 4, 5}
	sort.Ints(randInts)
	numStrings := []string{"a", "b", "c"}
	fmt.Println(strings.Join(numStrings, ","))

	fmt.Println("==============================")
	file, err := os.Create("/tmp/go.samp.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("This is some random text")
	file.Close()

	stream, err := ioutil.ReadFile("/tmp/go.samp.txt")
	if err != nil {
		log.Fatal(err)
	}
	readString := string(stream)
	fmt.Println(readString)

	fmt.Println("===================================")
	randInt := 5
	randFloat := 10.5
	randString := "100"
	randString2 := "100.5"

	fmt.Println(float64(randInt))
	fmt.Println(int(randFloat))
	newInt, _ := strconv.ParseInt(randString, 0, 64)
	fmt.Println(newInt)
	newFloat, _ := strconv.ParseFloat(randString2, 64)
	fmt.Println(newFloat)
	fmt.Println("===================================")
	// http.HandleFunc("/", handler)
	// http.HandleFunc("/earth", handler2)
	// http.ListenAndServe(":8888", nil)

	fmt.Println("===================================")

	for i := 0; i < 10; i++ {
		go count(i)
	}
	time.Sleep(time.Millisecond * 1000)

	fmt.Println("===================================")
	stringChan := make(chan string)
	for i := 0; i < 3; i++ {
		go makeDough(stringChan)
		go addSauce(stringChan)
		go addToppings(stringChan)
		time.Sleep(time.Microsecond * 5000)
	}

}

var pizzaNum = 0
var pizzaName = ""

func addToppings(stringChan chan string) {
	pizza := <-stringChan
	fmt.Println("Add Toppings to", pizza, "and ship")
	stringChan <- pizzaName
	time.Sleep(time.Microsecond * 10)
}

func addSauce(stringChan chan string) {
	pizza := <-stringChan

	fmt.Println("Add Sauce and Send", pizza, "for Topping")
	stringChan <- pizzaName
	time.Sleep(time.Microsecond * 10)
}

func makeDough(stringChan chan string) {
	pizzaNum++
	pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)

	fmt.Println("Make Dough and Send for Sauce")
	stringChan <- pizzaName
	time.Sleep(time.Microsecond * 10)
}

func count(id int) {
	for i := 0; i < 10; i++ {
		fmt.Println(id, ":", i)

		time.Sleep(time.Millisecond * 1000)
	}
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello earth.")
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world.")
}

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func (cicle Circle) area() float64 {
	return math.Pi * math.Pow(cicle.radius, 2)
}

func (rect Rectangle) area() float64 {
	return rect.height * rect.width
}

type Rectangle struct {
	leftX  float64
	topY   float64
	height float64
	width  float64
}

func changeXVal(x int) {
	x = 2
}

func changeXValNow(x *int) {
	*x = 2
}

func safeDiv(num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()

	solution := num1 / num2
	return solution
}

func printOne() {
	fmt.Println(1)
}

func printTwo() {
	fmt.Println(2)
}

func subtractThem(args ...int) int {
	sum := 0

	for _, val := range args {
		sum += val
	}

	return sum
}

func addThemUp(numbers []float64) float64 {
	sum := 0.0

	for _, val := range numbers {
		sum += val
	}
	return sum
}
