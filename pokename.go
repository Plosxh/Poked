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
  evo string
  monType string
  contre []string
  a1 string
  a2 string
}

type bm struct{

  id int
  prim1 string
  prim2 string
  cha1 string
  cha2 string
  cha3 string
}



func main()  {
  var id int
  var idAtt int
  var monPoke pokemon
  var monBest bm

    getLesPoke()

  fmt.Println("Entrez l'id du pokemon : ")
  _,err := fmt.Scanln(&id)
  if err != nil {
    fmt.Println("Error: ", err)
  }

  monPoke = getPoke(id)
  counter(&monPoke)

  monBest = getMeilleursAttaques(id)


  fmt.Println("Attaques Principales disponibles pour ce pokemon : ")
  getLesAttaquesPrim(monBest.prim1)
  getLesAttaquesPrim(monBest.prim2)



  fmt.Println("Choisissez son attaque de base : ")
  _,err = fmt.Scanln(&idAtt)
  setAttaquePrim(&monPoke, idAtt)


  fmt.Println("Attaques Chargees disponibles pour ce pokemon : ")
  getLesAttaquesCha(monBest.cha1)
  getLesAttaquesCha(monBest.cha2)
  getLesAttaquesCha(monBest.cha3)

  fmt.Println("Choisissez son attaque chargee : ")
  _,err = fmt.Scanln(&idAtt)
  setAttaqueCha(&monPoke, idAtt)


  fmt.Println("ID : ",monPoke.id," | Nom : ", monPoke.name," | Contre : ",monPoke.contre," | Attaques : ",monPoke.a1," ",monPoke.a2)
  ajoutInventaire(monPoke)
  fmt.Println("Ajout du pokémon avec succès !")

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
    rows, err := db.Query("select * from attaque where attype=1 and nom=\""+nom+"\"")
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
    rows, err := db.Query("select * from attaque where attype=2 and nom=\""+nom+"\"")
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
    var evo string
    var name string
    var monType string
    var monContre []string
    var a1 string
    var a2 string
    rows, err := db.Query("select * from pokemon where id=\""+strconv.Itoa(monId)+"\"")
    defer rows.Close()
    for rows.Next(){
    err = rows.Scan(&id,&evo,&name,&monType,&a1,&a2)
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
    var evo string
    var monType string
    var monContre []string
    var a1 string
    var a2 string
    for i := 1; i < 151; i++ {

        var monPoke pokemon
      rows, err := db.Query("select * from pokemon where id=\""+strconv.Itoa(i)+"\"")
      defer rows.Close()
      for rows.Next(){
      err = rows.Scan(&id,&evo, &name,&monType,&a1,&a2)
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
    monPoke.evo = evo
    monPoke.monType = monType
    monPoke.contre = monContre
    monPoke.a1 = a1
    monPoke.a2 = a2
    fmt.Println(monPoke.id," | ",monPoke.name)
  }
}

func getMeilleursAttaques(id int) bm {

    var monid int
    var prim1 string
    var prim2 string
    var cha1 string
    var cha2 string
    var cha3 string
    var bestM bm

    db, err := sql.Open("sqlite3","./pokename.db")
    if err != nil{
      log.Fatal(err)
    }
    defer db.Close()

    rows, err := db.Query("select * from meilleurs where id=\""+strconv.Itoa(id)+"\"")
    defer rows.Close()
    for rows.Next(){
    err = rows.Scan(&monid,&prim1,&prim2,&cha1,&cha2,&cha3)
    if err != nil {
      log.Fatal(err)
    }
  }

    err = rows.Err()
    if err != nil{
      log.Fatal(err)
    }

    bestM.id = monid
    bestM.prim1 = prim1
    bestM.prim2 = prim2
    bestM.cha1 = cha1
    bestM.cha2 = cha2
    bestM.cha3 = cha3

  return bestM

}


func ajoutInventaire(p pokemon){

  db, err := sql.Open("sqlite3","./pokename.db")
  if err != nil{
    log.Fatal(err)
  }
  defer db.Close()

        _,err =db.Exec("INSERT INTO inventaire VALUES (\""+strconv.Itoa(p.id)+"\",\""+p.name+"\",\""+p.evo+"\",\""+p.monType+"\",\""+p.a1+"\",\""+p.a2+"\")")

}
