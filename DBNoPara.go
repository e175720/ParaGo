package main

import (
  "database/sql"
  "fmt"
  "strings"
  "time"

  _ "github.com/mattn/go-sqlite3"
)

func main(){
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


  var day string
  var name [2]string
  var exec string

  var timer = 0.0

  date := [...] string{"05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17"}
  base := [3] string{"./DB_data/d2019/07/", "", ".db"}

  for t:=0;t<10;t++{
    start := time.Now()

    //データベースに接続
    for m := 0; m<len(date); m++{
      var tables []string
      base[1] = date[m]
      data := strings.Join(base[:],"")
      db, err := sql.Open("sqlite3", data)
      if err != nil{
        panic(err)
      }

      //データベースの全テーブル名を取得
      table, err := db.Query(`select name from sqlite_master where type='table'`)
      if err != nil{
        panic(err)
      }

      //全てのテーブルのデータに順番にアクセス
      for table.Next(){
        err = table.Scan(&day)
        if err != nil {
          panic(err)
        }
        tables = append(tables, day)
      }
      for i := 0; i<len(tables); i++{
        name[0] = "select * from "
        name[1] = tables[i]
        exec = strings.Join(name[:], "")
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
          vk = append(vk, k)
          v16 = append(v16, e16)
          v17 = append(v17, e17)
          v18 = append(v18, e18)
          v19 = append(v19, e19)
        }
      }

    db.Close()
    }
    end := time.Now()
    timer = timer + end.Sub(start).Seconds()
  }
  fmt.Printf("%f秒\n",timer/10)
}
