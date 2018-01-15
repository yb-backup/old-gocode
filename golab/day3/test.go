package main

import "fmt"

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

// three way to make map
func testBasicMap() {
	var m map[string]int
	m = make(map[string]int)
	m["route"] = 1
	v, ok := m["route"]
	if ok {
		fmt.Println("m.route: ", v)
	}

	mm := make(map[string]int)
	mm["route"] = 11
	v, ok = mm["route"]
	if ok {
		fmt.Println("mm.route: ", v)
	}

	mmm := map[string]int{
		"route": 111,
	}
	v, ok = mmm["route"]
	if ok {
		fmt.Println("mmm.route: ", v)
	}
}

// https://blog.golang.org/go-maps-in-action
func testMap() {
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)

	// insert someone
	personDB["1"] = PersonInfo{"1", "yangbin", "Room 123..."}
	personDB["123"] = PersonInfo{"123", "binnnn", "Room 321..."}

	person, ok := personDB["1"]
	if ok {
		fmt.Println("Found Person", person.Name, "with ID 1")
	} else {
		fmt.Println("Not Found person with ID 1")
	}

	// ":=" 表示声明并赋值，由于在上面person, ok已经声明过了，所以此处只能做
	// 赋值操作，不能重复声
	person, ok = personDB["12"]
	if ok {
		fmt.Println("Found Person", person.Name, "with ID 12")
	} else {
		fmt.Println("Not Found person with ID 12")
	}
}

func add(m map[string]map[string]int, path string, country string) {
	mm, ok := m[path]
	if !ok {
		mm = make(map[string]int)
		m[path] = mm
	}
	mm[country]++
}

func testComplexMap() {
	hits := map[string]map[string]int{}
	// hits := make(map[string]map[string]int)
	add(hits, "/doc/", "china")
	add(hits, "/doc/", "china")
	add(hits, "/doc/", "china")
	add(hits, "/pic/", "china")
	add(hits, "/pic/", "china")

	for path, cs := range hits {
		fmt.Printf("path %s country hit: \n", path)
		for country, value := range cs {
			fmt.Printf("%s : %d\n", country, value)
		}
	}
}

func testMapStructKey() {
	type Key struct {
		Path, Country string
	}

	hits := map[Key]int{}

	hits[Key{"/doc/", "china"}]++
	hits[Key{"/doc/", "china"}]++
	hits[Key{"/doc/", "china"}]++
	hits[Key{"/pic/", "china"}]++
	hits[Key{"/pic/", "china"}]++

	fmt.Printf("/doc/ china : %d\n", hits[Key{"/doc/", "china"}])
	fmt.Printf("/pic/ china : %d\n", hits[Key{"/pic/", "china"}])
}

func testFlowControl() {
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 0
	for {
		sum++
		if sum > 100 {
			break
		}
	}
	fmt.Println(sum)

	i := 3
	switch i {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fallthrough
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("default")
	}

	switch {
	case 0 <= sum && sum < 1000:
		fmt.Println("0 <= sum && sum < 1000 : ", sum)
	case 1000 <= sum && sum < 10000:
		fmt.Println("1000 <= sum && sum < 10000 : ", sum)
	default:
		fmt.Println("default")
	}

}

func main() {
	fmt.Println("vim-go")
	fmt.Println("=========================")
	testBasicMap()
	fmt.Println("=========================")
	testMap()
	fmt.Println("=========================")
	testComplexMap()
	fmt.Println("=========================")
	testMapStructKey()
	fmt.Println("=========================")
	fmt.Println("=========================")
	fmt.Println("=========================")
	testFlowControl()

}
