package main

import (
	"encoding/json"
	"fmt"
	al "go-journey/algorithms"
	a "go-journey/arrays"
	f "go-journey/functions"
	"io/ioutil"
	"log"
	"os"
)

type girl interface {
	confess(who string)
}

type Waifu struct {
	Name string
	Age  int
}

func (w *Waifu) setName(name string) {
	w.Name = name
}

func (w Waifu) confess(who string) {
	fmt.Printf("%s, isn't the moon beautiful?)\n", who)
}

type girlfriend struct {
	name string
}

func (g girlfriend) confess(who string) {
	fmt.Printf("%s, sorry...\n", who)
}

func tellIt(g girl, who string) {
	fmt.Println(g)
	g.confess(who)
}

func printType(value interface{}) {
	switch value.(type) {
	case string:
		fmt.Println("type of var is STRING.")
	case int:
		fmt.Println("type of var is INT.")
	default:
		fmt.Println("idk")
	}
}

var funcvar func(int) int

func inc(x int) int {
	return x + 1
}

func dec(x int) int {
	return x - 1
}

func main() {

	//Functions and Packages

	f.Hello()

	x, y := f.Calculate(3, 4)
	fmt.Printf("area: %d, perimeter: %d\n", x, y)

	fmt.Printf("x: %d, y: %d", x, y)
	f.Swap(&x, &y)
	fmt.Printf("\n---Swapped---\n")
	fmt.Printf("x: %d, y: %d", x, y)

	f.Say(func(txt string) string {
		return txt
	}("salom"),
	)

	firstSum := f.PartialSum(5)
	fmt.Printf("\nPARTIAL SUM: %d", firstSum(5))
	fmt.Printf("\nTRIPLEsum func: %d\n", f.TripleSum(3)(4)(7))

	f.VariadicFunc(1, 2, 3, 4, 5, 6)
	f.WithDifferentTypes("salom", 9879087, 4.323)

	funcvar = inc

	fmt.Printf("function as variable -> funcvar(1): %d\n", funcvar(1))

	fmt.Printf("function as argument -> ApplyIt(inc,1): %d\n", f.ApplyIt(inc, 1))
	fmt.Printf("function as argument -> ApplyIt(dec,1): %d\n", f.ApplyIt(dec, 1))
	fmt.Printf("anonymous function -> ApplyIt(anonymous_func,1): %d\n", f.ApplyIt(func(i int) int { return i * 10 }, 1))

	// Arrays

	fmt.Println("\n################## Arrays ###################")
	a.TenZero()

	var array [2]string = [2]string{"hi", "bye"}
	fmt.Println(array)

	var arr = [3]int{1, 2, 3}
	fmt.Println(arr)

	var sl = []int{3, 1, 1, 4}
	fmt.Println(sl[0:2])
	fmt.Println("len: ", len(sl))
	fmt.Println("cap: ", cap(sl))
	fmt.Println("sl from arr: ", arr[0:2])

	// Slices

	fmt.Println("\n################## Slices ###################")

	var msl = make([]int, 2, 8)
	fmt.Println("msl len: ", len(msl))
	fmt.Println("msl cap: ", cap(msl))

	msl = append(msl, 20, 21)
	fmt.Println("after append: ", msl)

	new_msl := []int{12, 123123, 121233}
	msl = append(msl, new_msl...)
	fmt.Println("appending two slices: ", msl)

	neededNums := msl[:len(msl)-3]
	dest := make([]int, len(neededNums))

	copy(neededNums, dest)
	fmt.Println("copied nums: ", dest)

	// Maps

	fmt.Println("\n################## Maps ###################")

	var strMap = map[string]string{"new": "yangi"}
	fmt.Println("strMap: ", strMap)

	var intStrMap = map[int]string{1: "bir", 2: "ikki", 3: "uch"}
	fmt.Println("intStrMap: ", intStrMap)

	var emptyMap = make(map[string]string)
	fmt.Println("emptyMap: ", emptyMap)

	emptyMap["salom"] = "hi"
	fmt.Println("after adding salom: ", emptyMap)

	delete(emptyMap, "salom")
	fmt.Println("again empty: ", emptyMap)

	val, ok := intStrMap[1]
	fmt.Printf("%s in intStrMap: %t\n", val, ok)

	fmt.Println("iterate intStrMap:")
	for k, v := range intStrMap {
		fmt.Printf("%d: %s\n", k, v)
	}

	fmt.Println("\n################## Structs & interfaces ###################")

	megumi := Waifu{Age: 18}
	megumi.setName("Megumi Kato")
	kawori := girlfriend{name: "Kawori Miyazono"}

	tellIt(megumi, "Otabek")
	tellIt(kawori, "Shahruz")

	fmt.Println("\n################## Type Assertions ###################")
	var value interface{} = "salom"
	printType(value)

	fmt.Println("\n################## Scan ###################")
	var age int

	fmt.Print("\nEnter your age: ")

	_, err := fmt.Scan(&age)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("You are %d years old.\n", age)

	fmt.Println("\n################## JSON ###################")

	marsheled, err := json.Marshal(megumi)
	if err != nil {
		fmt.Println("error in marshalling")
	}

	var unmarsheled Waifu
	err = json.Unmarshal(marsheled, &unmarsheled)
	if err != nil {
		fmt.Println("error in UNmarshalling")
	}
	fmt.Println(unmarsheled)
	fmt.Println(megumi)

	fmt.Println("\n################## ioutil & os ###################")
	var filename string = "test.txt"
	txt := "just a test"

	err = ioutil.WriteFile(filename, []byte(txt), 0777)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("test.txt file created.")
	fmt.Println("cat test.txt:")
	text, e := ioutil.ReadFile("test.txt")
	if e != nil {
		log.Fatal(e)
	}
	fmt.Print(string(text))

	fmt.Println("\nChange the content of the file")

	file, error := os.OpenFile(filename, os.O_RDWR, 0777)
	if error != nil {
		fmt.Print(error)
	}

	_, err = file.Write([]byte("Hello World"))
	if err != nil {
		log.Fatal(err)
	}

	file.Close()

	fmt.Printf("cat %s\n", filename)
	text, error_reading := os.ReadFile(filename)
	if error_reading != nil {
		log.Fatal(error_reading)
	}
	fmt.Print(string(text))

	fmt.Println("\n################## algorithms ###################")
	list := []int{1, 2, 3, 4, 5}
	n := 4
	i, err := al.BinarySearch(list, 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("BinarySearch(%v, %d): %d\n", list, n, i)

	fmt.Printf("Factorial(%d): %d", n, al.Factorial(n))
}
