package main

import "fmt"


func counter(p *pokemon)(){
 //types := [18]string{"Normal","Combat","Vol","Poison","Sol","Roche","Insecte","Fantôme","Acier","Feu","Eau","Plante","Electricité","Psy","Glace","Dragon","Ténèbre","Fée"}
//var contre []string
//var leType = p.monType
 contre :=[]string{}
  switch p.monType {
  case "Normal":
 contre =append(contre,"Combat")

  case "Combat":
    contre=append(contre,"Vol", "Psy", "Fee")

  case "Vol":
 contre =append(contre,"Roche", "Electricte", "Glace")

  case "Poison":
    contre=append(contre,"Sol", "Psy")

  case "Sol":
    contre=append(contre,"Eau", "Plante", "Glace")

  case "Roche":
    contre=append(contre,"Combat", "Sol", "Acier", "Eau", "Plante")

  case "Insecte":
    contre=append(contre,"Vol", "Roche", "Feu")

  case "Fantome":
    contre=append(contre,"Fantome", "Tenebre")

  case "Acier":
    contre=append(contre,"Combat", "Sol", "Feu")

  case "Feu":
    contre=append(contre,"Sol", "Roche", "Eau")

  case "Eau":
    contre=append(contre,"Eau", "Electricite")

  case "Plante":
    contre=append(contre,"Vol", "Poison", "Insecte", "Feu", "Glace")

  case "Electricite":
    contre=append(contre,"Sol")

  case "Psy":
    contre=append(contre,"Insecte", "Fantome", "Tenebre")

  case "Glace":
    contre=append(contre,"Combat", "Roche", "Acier", "Feu")

  case "Dragon":
    contre=append(contre,"Glace", "Dragon", "Fee")

  case "Tenebre":
    contre=append(contre,"Combat", "Insecte", "Fee")

  case "Fee":
    contre=append(contre,"Poison", "Acier")

  default:
    fmt.Printf("Mauvaise saisie !")
}
p.contre =contre
}
