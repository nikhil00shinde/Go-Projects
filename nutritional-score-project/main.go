package main 

import (
 "fmt"
)

func main(){
  
	ns :=  GetNutritionalScore(
		NutritionalData{
			Energy:               EnergyFromKcal(),
			Sugars:               SugarGram(),
			SaturatedFattyAcids:  SaturatedFattyAcids(),
			Sodium:               SodiumMilligram(),
			Fruits:               FruitsPercents(),
			Fibre:                FibreGram(),
			Protein:              ProteinGram(),
		}, Food)

   fmt.Printf("Nutritional Score:%d\n",ns.Value)
}