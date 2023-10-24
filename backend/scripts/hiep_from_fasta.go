package main

import (
	"bytes"
	"fmt"
	. "github.com/Ruhshan/hiep/backend/models/types"
	"github.com/Ruhshan/hiep/backend/pkg/fasta"
	"github.com/Ruhshan/hiep/backend/pkg/service/hiepcalculator"
	"github.com/dnlo/struct2csv"
	"os"
)

type CsvRow struct {
	Id        string
	MaxIEP    float64
	Positions []string
}

func main() {
	fastas, _ := fasta.ParseFastaFromFile("/home/ruhshan/Downloads/TAIR10_pep_20101214")
	calculator := hiepcalculator.GetConcurrentSequenceHiepCalculator()

	subset := fastas[0:100]

	var resultChanel = make(chan CsvRow, len(subset))

	var result []CsvRow

	for _, f := range subset {

		go func(s string, id string) {
			res := calculator.CalculateIeps(s, 1, 0, 0, "IPC_protein")
			positions := getPositions(res.SequenceAndPositions)

			resultChanel <- CsvRow{Id: id, MaxIEP: res.MaxIep, Positions: positions}
		}(f.Sequence, f.Id)

	}

	for i := 0; i < len(subset); i++ {
		row := <-resultChanel
		result = append(result, row)
	}

	buff := &bytes.Buffer{}
	w := struct2csv.NewWriter(buff)
	err := w.WriteStructs(result)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.WriteFile("output.csv", buff.Bytes(), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getPositions(dataArray []SubSequenceData) []string {

	var positionsArray []string
	for _, data := range dataArray {
		positionsArray = append(positionsArray, fmt.Sprintf("%d-%d", data.Position[0], data.Position[1]))
	}

	return positionsArray
}
