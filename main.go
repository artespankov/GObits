package main

// IMPORT
import (
	"fmt"
	"math/rand"
	"time"
)

// FUNCTION
func getPower(name string) (i int, int bool) {
	if name == "Gotu" {
		return 9000, true
	}
	return 0, false
}

// STRUCTURE
type Person struct {
	Name string
	Age  int
}

// METHOD FOR STRUCTURE: Associate method with structure - *Person is the receiver of PrintIntro method.
func (person *Person) PrintIntro() {
	fmt.Printf("Hey! my name is %s.\n", person.Name)
	fmt.Printf("I'm %d years old \n", person.Age)
}

// CONSTRUCTOR (returns pointer)
func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

// CONSTRUCTOR (returns the instance itself)
//func NewPerson(name string, age int) Person{
//	return Person{
//		Name: name,
//		Age: age,
//	}
//}

// STRUCTURE FIELDS
type Jedi struct {
	Name   string
	Power  int
	Origin string
	Master *Jedi
}

// COMPOSITION
type Pokemon struct {
	Name string
}

func (p *Pokemon) Summon() {
	fmt.Printf("%s! I choose you!\n", p.Name)
}

type TournamentFighter struct {
	Power int
	*Pokemon
}

////INTERFACES
//type Logger interface {
//	Log(message string)
//}
//
//// interface used as function parameter
//func process(logger Logger){
//	logger.Log("Hello")
//}
//
//// interface implementation
//type ConsoleLogger struct {}
//func (l ConsoleLogger) Log(message string) {
//	fmt.Println(message)
//}

func main() {
	// IMPORTS
	//if len(os.Args) >= 2 {
	//	fmt.Println(os.Args[0])
	//	fmt.Println("It's over", os.Args[1])
	//}

	// VARIABLES
	//var power int = 9000
	//fmt.Printf("It's over %d\n", power)
	//// Inferring type+ Multi-assignment
	//power, name := -1, "Santa"
	//fmt.Printf("%s has the power %d\n", name, power)

	// STRUCTURES
	//drChuck := Person{Name:"Chuck",}
	//drChuck.PrintIntro()

	// STRUCTURES FIELDS
	// NEW (allocate memory required by type)
	//profThunder := &Person{
	//	Name: "TH",
	//	Age: 666,
	//}
	// VS
	//profThunder := new(Person)
	//profThunder.Name = "GR"
	//profThunder.Age = 666

	//luke := &Jedi{
	//	Name: "Luke Skywalker",
	//	Power: 2000,
	//	Master: &Jedi{
	//		Name: "Master Yoda",
	//		Power: 9000,
	//		Master: nil,
	//	},
	//}
	//fmt.Printf("Luke: %s\n", luke)

	//// COMPOSITION
	//// use the method of composed instance
	//myPokemon := & TournamentFighter{
	//	Power: 500,
	//	Pokemon: &Pokemon{"Pikachu"},
	//}
	//myPokemon.Summon()

	//// ARRAYS
	//var scores [10] int
	//scores[0] = 39
	//
	//scores_ := [4]int {11, 22, 17, 22}
	//
	//for i, val := range scores_ {
	//	fmt.Printf("%d = %d\n", i, val)
	//}

	// SLICES
	//scores := []int {22, 17, 42, 1}
	//fmt.Println("Scores :", scores)

	//scores := make([]int, 5)
	//scores = append(scores, 999)
	//scores = append(scores, 777)

	// re-slice the slice to set value by desired index
	//scores = scores[0:8]
	//scores[7] = 12
	//fmt.Println("Scores: ", scores)
	//c := cap(scores)
	//fmt.Println(c)
	//
	//for i:=0; i < 25; i++ {
	//	scores = append(scores, i)
	//
	//	if cap(scores) != c {
	//		c = cap(scores)
	//		fmt.Println(c)
	//	}
	//}

	//fmt.Println("Scores: ", scores)
	//scores = append(scores, 909)
	//fmt.Println(scores)

	// MAPS
	//lookup := make(map[string]int) // create map
	//lookup["goku"] = 9001 // set key-value
	//power, exists := lookup["vegeta"] // find by key
	//total := len(lookup) // number of keys
	//delete(lookup, "goku") // delete by key
	//alternative  map initialization
	//lookup := map[string]int{
	//	"goku": 9001,
	//	"gohan": 2000,
	//}
	//
	//// prints 0(default value for an integer), false
	//fmt.Println(lookup)

	//if len(os.Args) != 2{
	//	os.Exit(1)
	//}

	// Error Handling
	//n, err := strconv.Atoi(os.Args[1])
	//
	//if err != nil {
	//	fmt.Println("not a valid number")
	//} else {
	//	if err == io.EOF {
	//		fmt.Println("no more input")
	//	} else {
	//		fmt.Println(n)
	//	}
	//}

	// DEFER
	// close file via defer after return from main, even if error happened
	//file, err := os.Open("some_file")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer file.Close()

	// INITIALIZED IF
	//if err:= process(); err != nil{
	//	return err
	//}
	//
	//// EMPTY INTERFACE
	//var a interface{}
	//// Convert interface variable to concrete type
	//a.(int)

	// conversion between strings and byte arrays
	//stra := "the spice must flow"
	//byts := []byte(stra)
	//strb := string(byts)


	// GOROUTINES
	//fmt.Println("Start")
	//go process(1)
	//go func(num int) {
	//	fmt.Println("Running concurrently", num)
	//}(2)
	//time.Sleep(time.Millisecond * 10) // delay code execution to make main wait
	//fmt.Println("Complete")

	//SYNCHRONIZATION
	//counter:=0
	//var lock sync.Mutex
	//for i := 0; i < 20; i++ {
	//	go func(){
	//		lock.Lock()
	//		defer lock.Unlock()
	//		counter++
	//		fmt.Println(counter)
	//	}()
	//	time.Sleep(time.Millisecond * 100)
	//}

	// CHANNELS
	c := make(chan int, 100)
	for i:=0; i < 10; i++ {
		worker := &Worker{id: i}
		go worker.process(c)
	}


	for {
		select {
			case c <- rand.Int():
			case <- time.After(time.Millisecond * 100):
				fmt.Println("timed out")
			//default:
			//	fmt.Println("dropped") // this can be left empty to silently drop the data
		}
		time.Sleep(time.Millisecond * 50)
	}


}

// CHANNELS
type Worker struct {
	id int
}

func (w *Worker) process(c chan int){
	for {
		data := <- c
		fmt.Printf("worker %d got %d\n", w.id, data)
		time.Sleep(time.Millisecond * 1000)
	}
}

//FUNCTION TYPE
//type Add func(a int, b int) int
//
//func process (adder Add) int{
//	return adder(1, 2)
//}
//fmt.Println(process(func(a int, b int) int{
//	return a + b
//}))