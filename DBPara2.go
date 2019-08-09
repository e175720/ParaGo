package main

import (
  "database/sql"
  "fmt"
  "log"
  "strings"
  "sync"

  _ "github.com/mattn/go-sqlite3"
)

func Exec(day string, finished chan bool, wg *sync.WaitGroup){
  var index string
  var k float64
  var e16 float64
  var e17 float64
  var e18 float64
  var e19 float64

  var vk  []float64
  var v16 []float64
  var v17 []float64
  var v18 []float64
  var v19 []float64
  var days string
  var tables []string
  var name [2]string

  address := [3] string{"./DB_data/d2019/07/", "", ".db"}
  address[1] = day
  data := strings.Join(address[:],"")

  db, err := sql.Open("sqlite3", data)
  if err != nil{
    panic(err)
  }

  table, err := db.Query(`select name from sqlite_master where type='table'`)
  if err != nil{
    panic(err)
  }

  for table.Next(){
    err = table.Scan(&days)
    if err != nil {
      panic(err)
    }
    tables = append(tables, days)
  }

  for i := 0; i<len(tables); i++{
    name[0] = "select * from "
    name[1] = tables[i]
    exec := strings.Join(name[:], "")

    rows, err := db.Query(exec)
    if err != nil{
      panic(err)
    }

    for rows.Next(){
      err = rows.Scan(&index, &k, &e16, &e17, &e18, &e19)
      if err != nil {
        panic(err)
      }
      vk = append(vk, k)
      v16 = append(v16, e16)
      v17 = append(v17, e17)
      v18 = append(v18, e18)
      v19 = append(v19, e19)
    }
  }
  db.Close()
  defer wg.Done()
  return v19
}

func main(){
  log.Print("Start.")

  date := [...] string{"05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17"}
  wg := new(sync.WaitGroup)
  finished := make(chan bool, len(date))

  for i := 0; i<len(date); i++{
    wg.Add(1)
    go func(i int){
     Exec(date[i], finished, wg)
    }
  }

  wg.Wait()
  close(finished)
  log.Print("Finish.")

}
