package main

import (
    "fmt"
    "math/rand"
    "time"
    "strconv"
    "runtime"
    "sync"
    "log"

    "github.com/fatih/color"
)
var wg sync.WaitGroup

var re *color.Color = color.New(color.FgRed)
var red *color.Color = re.Add(color.BgBlack)
var boldRed *color.Color = red.Add(color.Bold)

func main() {

  runtime.GOMAXPROCS(2)
  wg.Add(1)

  var mode, numTables, total, tries int = 0, 0, 0, 0
  var avg float32 = 0
  red.Println("\n\tGame modes:\n1: Pick a random table\n2: Find the avg tries it takes to solve a puzzle\n\nPlease choose 1 or 2: ")
  if _, err := fmt.Scanln(&mode); err != nil {
    log.Fatal(err)
  }

  if mode == 1 {
    picker()
  } else if mode == 2 {
    red.Print("How many tables would you like to calculate for: ")
    if _, err := fmt.Scanln(&numTables); err != nil {
      log.Fatal(err)
    }
    red.Print("How many times would you like to run the experiment: ")
    if _, err := fmt.Scanln(&tries); err != nil {
      log.Fatal(err)
    }
    fmt.Println("\n")
    go calc()
    go printAvg(numTables, total, avg, tries)
    wg.Wait()
  } else {
    red.Println("Goodbye!")
  }
}

func printAvg(numTables int, total int, avg float32, tries int) {
  defer wg.Done()
  defer timeTrack(time.Now(), "Finding the avg")
  for i := 1; i <= tries; i++ {
    total += avgFinder(numTables)
  }
  //a := color.New(color.FgRed, color.Bold)
  avg = float32(total)/float32(tries)
  boldRed.Printf("\rThe avg tries it took to solve %v tables was %v                   \n\n", numTables, avg)
}

func calc() {
  for {
    for i := "Calculating";; i += "." {
      red.Printf("\r%s", i)
      time.Sleep(500 * time.Millisecond)
    }
  }
}

func printWelcome() {
  welcome := []string{"\nWelcome to Colin's Random Table Picker", "Tables are eliminated when number is drawn,", "And put back in the game when their number is drawn again.", "Last number left wins!", "Good Luck.\n"}
  for _, item := range welcome {
    red.Println(item)
    time.Sleep(400 * time.Millisecond)
  }
}

func picker() {
  var tables, comp, wait int
  var compString string
  printWelcome()
  for {
    red.Print("How many tables are playing: ")
    if _, err := fmt.Scanln(&tables); err != nil {
      log.Fatal(err)
    }
    red.Print("How many ms would you like to wait between each iteration: ")
    if _, err := fmt.Scanln(&wait); err != nil {
      log.Fatal(err)
    }
    if tables > 0 {
      break
    }
  }
  tableArray := make([]string, tables)
  for i, _ := range tableArray {
    tableArray[i] = strconv.Itoa(i + 1)
  }
  var total int = 0
  for {
    r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
    comp = r.Intn(tables) + 1;
    compString = strconv.Itoa(comp)
    var numX int = 0
    total++
    red.Println("Random number is:", compString)
    if tableArray[comp - 1] == compString {
      tableArray[comp - 1] = "x"
    } else {
      tableArray[comp - 1] = compString
    }
    red.Println(tableArray)
    for _, element := range tableArray {
      if element == "x" {
        numX++
      }
    }
    if numX >= tables - 1 {
      for _, element := range tableArray {
        if element != "x" {
          boldRed.Println("The lucky winner is", element, "chosen after", total, "rounds.\n")
          break
        }
      }
      break
    }
    time.Sleep(time.Duration(wait) * time.Millisecond)
  }
}

func avgFinder(tables int) int{
  var comp, total int = 0, 0
  var compString string
  tableArray := make([]string, tables)
  for i, _ := range tableArray {
    tableArray[i] = strconv.Itoa(i + 1)
  }
  for {
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

func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    red.Printf("%s took %s", name, elapsed)
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
