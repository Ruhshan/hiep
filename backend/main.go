package main

import (
	"fmt"
	"github.com/Ruhshan/hiep/backend/pkg/service/hiepcalculator"
)

func main() {
	var calculator = hiepcalculator.GetConcurrentSequenceHiepCalculator()
	var res = calculator.CalculateMaxIep("MEASAGLVAGSYRRNELVRIRHESDGGTKPLKNMNGQICQICGDDVGLAET",
		1);

	fmt.Println(res)
	//var processor = instanthiep.GetInstantHiepProcessor(calculator)
	//
	//var r = api.NewRoutes()
	//
	//api.ConfigureInstantHiepRoutes(r.GetBaseRoute("hiep/instant"), processor)
	//
	//r.Run()
}
