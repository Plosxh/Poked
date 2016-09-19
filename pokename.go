package main

import (
  "database/sql"
	"log"
  "fmt"
  _ "github.com/mattn/go-sqlite3"
  "strconv"
)

type pokemon struct{
  id int
  name string
  monType string
  contre []string
  a1 string
  a2 string
}



func main()  {
  var id int
  var idAtt int
  var monPoke pokemon

    getLesPoke()

  fmt.Println("Entrez l'id du pokemon : ")
  _,err := fmt.Scanln(&id)
  if err != nil {
    fmt.Println("Error: ", err)
  }

  monPoke = getPoke(id)
  counter(&monPoke)

  getLesAttaquesPrim(monPoke.monType)
  fmt.Println("Choisissez son attaque de base : ")
  _,err = fmt.Scanln(&idAtt)
  setAttaquePrim(&monPoke, idAtt)

  getLesAttaquesCha(monPoke.monType)
  fmt.Println("Choisissez son attaque chargee : ")
  _,err = fmt.Scanln(&idAtt)
  setAttaqueCha(&monPoke, idAtt)

  fmt.Println("ID : ",monPoke.id," | Nom : ", monPoke.name," | Contre : ",monPoke.contre," | Attaques : ",monPoke.a1," ",monPoke.a2)

  /*fmt.Println("choisissez son attaque primaire : ")
  _,err = fmt.Scanlen(&a1)
  if err != nil {
    fmt.Println("Error: ", err)
  }
  fmt.Println("Choisissez son attaque chargée : ")
  _,err = fmt.Scanlen(&a2)
  if err != nil {
    fmt.Println("Error: ", err)
  }*/



}


func setAttaquePrim(p *pokemon, att int){

  var id int
  var name string
  var attype int
  var id_type string
  db, err := sql.Open("sqlite3","./pokename.db")
  if err != nil{
    log.Fatal(err)
  }
  defer db.Close()

  rows, err := db.Query("select * from attaque where id=\""+strconv.Itoa(att)+"\"")
  defer rows.Close()
  for rows.Next(){
  err = rows.Scan(&id,&attype,&name,&id_type)
  if err != nil {
    log.Fatal(err)
  }
}

  err = rows.Err()
  if err != nil{
    log.Fatal(err)
  }
  p.a1 = name
}


func setAttaqueCha(p *pokemon, att int){

  var id int
  var name string
  var attype int
  var id_type string
  db, err := sql.Open("sqlite3","./pokename.db")
  if err != nil{
    log.Fatal(err)
  }
  defer db.Close()

  rows, err := db.Query("select * from attaque where id=\""+strconv.Itoa(att)+"\"")
  defer rows.Close()
  for rows.Next(){
  err = rows.Scan(&id,&attype,&name,&id_type)
  if err != nil {
    log.Fatal(err)
  }
}

  err = rows.Err()
  if err != nil{
    log.Fatal(err)
  }
  p.a2 = name
}


func getLesAttaquesPrim(nom string)  {

  db, err := sql.Open("sqlite3","./pokename.db")
  if err != nil{
    log.Fatal(err)
  }
  defer db.Close()

  var id int
  var name string
  var attype int
  var id_type string
    rows, err := db.Query("select * from attaque where attype=1 and id_type=\""+nom+"\"")
    defer rows.Close()
    for rows.Next(){
    err = rows.Scan(&id,&attype,&name,&id_type)
    fmt.Println(id," | ",name)
    if err != nil {
      log.Fatal(err)
    }
  }

    err = rows.Err()
    if err != nil{
      log.Fatal(err)
    }
}

func getLesAttaquesCha(nom string)  {

  db, err := sql.Open("sqlite3","./pokename.db")
  if err != nil{
    log.Fatal(err)
  }
  defer db.Close()

  var id int
  var name string
  var attype int
  var id_type string
    rows, err := db.Query("select * from attaque where attype=2 and id_type=\""+nom+"\"")
    defer rows.Close()
    for rows.Next(){
    err = rows.Scan(&id,&attype,&name,&id_type)
    fmt.Println(id," | ",name)
    if err != nil {
      log.Fatal(err)
    }
  }

    err = rows.Err()
    if err != nil{
      log.Fatal(err)
    }
}


func getPoke(monId int) pokemon{
  db, err := sql.Open("sqlite3","./pokename.db")
  if err != nil{
    log.Fatal(err)
  }
  defer db.Close()
  var monPoke pokemon
  var id int
  var name string
  var monType string
  var monContre []string
  var a1 string
  var a2 string
  rows, err := db.Query("select * from pokemon where id=\""+strconv.Itoa(monId)+"\"")
  defer rows.Close()
  for rows.Next(){
  err = rows.Scan(&id,&name,&monType,&a1,&a2)
  if err != nil {
    log.Fatal(err)
  }
}

  err = rows.Err()
  if err != nil{
    log.Fatal(err)
  }

  monPoke.id = id
  monPoke.name = name
  monPoke.monType = monType
  monPoke.contre = monContre
  monPoke.a1 = a1
  monPoke.a2 = a2
  return monPoke
}

func getLesPoke() {

  db, err := sql.Open("sqlite3","./pokename.db")
  if err != nil{
    log.Fatal(err)
  }
  defer db.Close()

  var id int
  var name string
  var monType string
  var monContre []string
  var a1 string
  var a2 string
  for i := 1; i < 151; i++ {

      var monPoke pokemon
    rows, err := db.Query("select * from pokemon where id=\""+strconv.Itoa(i)+"\"")
    defer rows.Close()
    for rows.Next(){
    err = rows.Scan(&id, &name,&monType,&a1,&a2)
    if err != nil {
      log.Fatal(err)
    }
  }

    err = rows.Err()
    if err != nil{
      log.Fatal(err)
    }

    monPoke.id = id
    monPoke.name = name
    monPoke.monType = monType
    monPoke.contre = monContre
    monPoke.a1 = a1
    monPoke.a2 = a2
    fmt.Println(monPoke.id," | ",monPoke.name)
  }
}

func getMeilleurAttaque()  {

  

}
