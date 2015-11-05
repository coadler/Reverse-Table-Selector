package main

import (
    "fmt"
    "math/rand"
    "time"
    "strconv"
    "runtime"
    "sync"
)
var wg sync.WaitGroup

func main() {
  runtime.GOMAXPROCS(2)
  wg.Add(1)
  var mode, numTables, total, tries, avg int = 0, 0, 0, 0, 0
  fmt.Println("\n\tGame modes:")
  fmt.Println("1: Pick a random table")
  fmt.Println("2: Find the avg tries it takes to solve a puzzle\n")
  fmt.Println("Please choose 1 or 2: ")
  fmt.Scanf("%d", &mode)
  fmt.Println("\n")
  if mode == 1 {
    picker()
  } else if mode == 2 {
    fmt.Print("How many tables would you like to calculate for: ")
    fmt.Scanf("%d", &numTables)
    fmt.Print("How many times would you like to run the experiment: ")
    fmt.Scanf("%d", &tries)
    fmt.Println("\n")
    go calc()
    go printAvg(numTables, total, avg, tries, )
    wg.Wait()
  } else {
    fmt.Println("Goodbye!")
  }
}

func printAvg(numTables int, total int, avg int, tries int) {
  defer wg.Done()
  for i := 1; i <= tries; i++ {
    total += avgFinder(numTables)
  }
  avg = total/tries
  fmt.Printf("\rThe avg tries it took to solve %v tables was %v\n\n", numTables, avg)
}

func calc() {
  for {
    for i := "Calculating";; i += "." {
      fmt.Printf("\r%s", i)
      time.Sleep(500 * time.Millisecond)
    }
  }
}

func printWelcome() {
  welcome := []string{"\nWelcome to Colin's Random Table Picker", "Tables are eliminated when number is drawn,", "And put back in the game when their number is drawn again.", "Last number left wins!", "Good Luck.\n"}
  for _, item := range welcome {
    fmt.Println(item)
    time.Sleep(500 * time.Millisecond)
  }
}

func picker() {
  var tables, comp, wait int
  var cont bool = false
  var end bool = false
  var compString string
  printWelcome()
  for cont == false {
    fmt.Print("How many tables are playing: ")
    fmt.Scanf("%d", &tables)
    fmt.Print("How many ms would you like to wait between each iteration: ")
    fmt.Scanf("%d", &wait)
    if tables > 0 {
      cont = true
    }
  }
  tableArray := []string{}
  for i := 1; i <= tables; i++ {
    tableArray = append(tableArray, strconv.Itoa(i))
  }
  var total int = 0
  for end == false {
    r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
    comp = r.Intn(tables) + 1;
    compString = strconv.Itoa(comp)
    var numX int = 0
    total++
    fmt.Println("Random number is:", compString)
    if tableArray[comp - 1] == compString {
      tableArray[comp - 1] = "x"
    } else {
      tableArray[comp - 1] = compString
    }
    fmt.Println(tableArray)
    for _, element := range tableArray {
      if element == "x" {
        numX++
      }
    }
    if numX >= tables - 1 {
      end = true
      for _, element := range tableArray {
        if element != "x" {
          fmt.Println("The lucky winner is", element, "chosen after", total, "rounds.\n")
          break
        }
      }
    }
    time.Sleep(time.Duration(wait) * time.Millisecond)
  }
}

func avgFinder(tables int) int{
  var comp, total int = 0, 0
  var end bool = false
  var compString string
  tableArray := []string{}
  for i := 1; i <= int(tables); i++ {
    tableArray = append(tableArray, strconv.Itoa(i))
  }
  for end == false {
    r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
    comp = r.Intn(tables) + 1;
    compString = strconv.Itoa(comp)
    var numX int = 0
    total++
    if tableArray[comp - 1] == compString {
      tableArray[comp - 1] = "x"
    } else {
      tableArray[comp - 1] = compString
    }
    for _, element := range tableArray {
      if element == "x" {
        numX++
      }
    }
    if numX >= tables - 1 {
      break
    }
  }
  return total
}

/*
AVERAGES:
1:  0
2:  1
3:  2
4:  6
5:  11
6:  19
7:  34
8:  57
9:  95
10: 164
11: 283
12: 494
*/
