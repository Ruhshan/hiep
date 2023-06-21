package main

import (
	"github.com/Ruhshan/hiep/backend/api"
	"github.com/Ruhshan/hiep/backend/pkg/service/hiepcalculator"
	"github.com/Ruhshan/hiep/backend/pkg/service/instanthiep"
)

func main() {
	var calculator = hiepcalculator.GetConcurrentSequenceHiepCalculator()
	//var calculator = hiepcalculator.GetSingleThreadedHiepCalculator()
	var processor = instanthiep.GetInstantHiepProcessor(calculator)

	var r = api.NewRoutes()

	api.ConfigureInstantHiepRoutes(r.GetBaseRoute("hiep/instant"), processor)

	r.Run()
}
