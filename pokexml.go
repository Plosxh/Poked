package main
import(
"fmt"
"io/ioutil"
"encoding/xml"
"database/sql"
"log"
_ "github.com/mattn/go-sqlite3"
//"strconv"
)

type resourcePrim struct{
  XMLName xml.Name `xml:"resources"`
  MovePrim []movePrim `xml:"move"`
}

type movePrim struct{
  XMLName xml.Name `xml:"move"`
  Primary []primary `xml:"primary"`
}

type primary struct{
  XMLName xml.Name `xml:primary"`
  Id []string `xml:"id"`
  Nom []string `xml:"nom"`
  Type []string  `xml:"type"`
}

type resourceCha struct{
  XMLName xml.Name `xml:"resources"`
  MoveCha []moveCha `xml:"move"`
}
type moveCha struct{
  XMLName xml.Name `xml:"move"`
  Charged []charged `xml:"charged"`
}

type charged struct{
  XMLName xml.Name `xml:charged"`
  Id []string `xml:"id"`
  Nom []string `xml:"nom"`
  Type []string  `xml:"type"`
}

type Resourcepoke struct{
  XMLName xml.Name `xml:"resources"`
  Poke []poke `xml:"pokemon"`
}
type poke struct{
  XMLName xml.Name `xml:"pokemon"`
  Nom []string `xml:"nom"`
  Type []string `xml:"type"`
}




func main()  {
  var P resourcePrim
  var C resourceCha
  var Po Resourcepoke
  var primlen int
  var chalen int
  var pokelen int

  xmlContent, _ := ioutil.ReadFile("moves.xml")
  err := xml.Unmarshal(xmlContent, &P)
  //err = xml.Unmarshal(xmlContent, &R)
  if err != nil { panic(err) }
  primlen = len(P.MovePrim[0].Primary)
  fmt.Println(primlen)

    db, err := sql.Open("sqlite3", "./pokename.db")
    if err != nil {
      log.Fatal(err)
    }
    defer db.Close()

    for i := 0; i < primlen; i++ {

      _,err =db.Exec("INSERT INTO attaque VALUES (NULL,1,\""+P.MovePrim[0].Primary[i].Nom[0]+"\",\""+P.MovePrim[0].Primary[i].Type[0]+"\")")


    }


  xmlContent, _ = ioutil.ReadFile("moves.xml")
  err = xml.Unmarshal(xmlContent, &C)
  //err = xml.Unmarshal(xmlContent, &R)
  if err != nil { panic(err) }

    chalen = len(C.MoveCha[0].Charged)
fmt.Println(chalen)
  for j := 0; j < chalen; j++ {

    _,err =db.Exec("INSERT INTO attaque VALUES (NULL,2,\""+C.MoveCha[0].Charged[j].Nom[0]+"\",\""+C.MoveCha[0].Charged[j].Type[0]+"\")")


  }

  db, err = sql.Open("sqlite3", "./pokename.db")
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  xmlContent, _ = ioutil.ReadFile("moves.xml")
  err = xml.Unmarshal(xmlContent, &Po)
  //err = xml.Unmarshal(xmlContent, &R)
  if err != nil { panic(err) }

    pokelen = len(Po.Poke)
    fmt.Println(Po.Poke[15])
  for k := 0; k < pokelen-1; k++ {

    _,err =db.Exec("INSERT INTO pokemon VALUES (NULL,\""+Po.Poke[k].Nom[0]+"\",\""+Po.Poke[k].Type[0]+"\",\"\",\"\")")


  }

  fmt.Println("FINI !!")


}
