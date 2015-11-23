package main

import (
    "fmt"
    "math/rand"
    "time"
    "strconv"
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

  var numTables, total, tries int = 0, 0, 0
  var avg float32 = 0

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
}

func printAvg(numTables int, total int, avg float32, tries int) {
  defer wg.Done()
  defer timeTrack(time.Now(), "Finding the avg")
  for i := 1; i <= tries; i++ {
    total += avgFinder(numTables)
  }
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
    red.Printf("%s took %s\n", name, elapsed)
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
