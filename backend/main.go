package main

import (
	"fmt"
	"github.com/Ruhshan/biogo"
	"github.com/Ruhshan/hiep/backend/iep"
)



func main() {

	fastas,_ := biogo.ParseFasta("/Users/abir_admin/development/hiep/backend/sequences.txt")

	fmt.Println(fastas[0].Sequence)

	fmt.Println(iep.PredictIsoelectricPoint("MEASAGLVAGSYRRNELVRIRHESD"))




}
