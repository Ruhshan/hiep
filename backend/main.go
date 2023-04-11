package main

import (
	"fmt"
	"github.com/Ruhshan/biogo"
	"github.com/Ruhshan/hiep/backend/service"
)



func main() {
	fastas,_ := biogo.ParseFasta("/Users/abir_admin/development/hiep/backend/sequences.txt")

	for _, fasta := range fastas{
		fmt.Println(fasta.Id)
		service.CalculateMaxIep(fasta.Sequence, 10)
		break
	}








}
