package hiepcalculator

import "github.com/Ruhshan/hiep/backend/models/types"

type SequenceHiepCalculator interface {
	CalculateIeps(seq string, minWindow int, minThreshold float64, maxThreshold float64, scale string) types.IepResult
}
