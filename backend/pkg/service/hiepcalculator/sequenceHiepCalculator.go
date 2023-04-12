package hiepcalculator

import "github.com/Ruhshan/hiep/backend/models/types"

type SequenceHiepCalculator interface {
	CalculateMaxIep(seq string, minWindow int) types.MaxIepResult
}
