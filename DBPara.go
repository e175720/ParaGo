package main

import (
  "database/sql"
  "fmt"
  "log"
  "strings"
  "sync"
  "time"

  _ "github.com/mattn/go-sqlite3"
)

func Exec(day string, finished chan []float64, wg *sync.WaitGroup){
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

  //データベースのアドレス
  address := [3] string{"./DB_data/d2019/07/", "", ".db"}
  address[1] = day
  data := strings.Join(address[:],"")

  //データベースに接続
  db, err := sql.Open("sqlite3", data)
  if err != nil{
    panic(err)
  }

  //データベースの全テーブル名を取得
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

  //全てのテーブルのデータに順番にアクセス
  for i := 0; i<len(tables); i++{
    name[0] = "select * from "
    name[1] = tables[i]
    exec := strings.Join(name[:], "")

    rows, err := db.Query(exec)
    if err != nil{
      panic(err)
    }

    //1つ1つのテーブルから各学年のデータを入手し、それぞれ配列に順番に格納
    for rows.Next(){
      err = rows.Scan(&index, &k, &e16, &e17, &e18, &e19)
      if err != nil {
        panic(err)
      }
      vk = append(vk,k)
      v16 = append(v16,e16)
      v17 = append(v17,e17)
      v18 = append(v18,e18)
      v19 = append(v19,e19)
    }
  }
  db.Close()
  defer wg.Done()
}

func main(){

  var timer = 0.0

  log.Print("Start.")

  for t:=0;t<10;t++{
    start := time.Now()
    date := [...] string{"05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17"}
    wg := new(sync.WaitGroup)
    finished := make(chan []float64, len(date))

    //全てのデータベースへのアクセスを並列処理
    for i := 0; i<len(date); i++{
      wg.Add(1)
      go Exec(date[i], finished, wg)
    }

    wg.Wait()
    close(finished)

    end := time.Now()
    timer = timer + end.Sub(start).Seconds()
  }

  fmt.Printf("%f秒\n",timer/10)
  log.Print("Finish.")

}
