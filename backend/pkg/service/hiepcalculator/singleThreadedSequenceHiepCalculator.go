package hiepcalculator

import (
	. "github.com/Ruhshan/hiep/backend/models/types"
	"github.com/Ruhshan/hiep/backend/pkg/iep"
	"math"
)


type singleThreadedSequenceHiepCalculator struct {}

func (c singleThreadedSequenceHiepCalculator) CalculateIeps(seq string, minWindow int, minThreshold float64,
	maxThreshold float64, scale string) IepResult {
	var maxIep = 0.0
	var allSubsequenceIep = getAllIepDataSequencial(seq, minWindow, scale)
	var filteredData []SubSequenceData

	var minIepThreshold = math.Inf(-1)
	var maxIepThreshold = math.Inf(1)

	if minThreshold != 0 {
		minIepThreshold = minThreshold
	}

	if maxThreshold != 0 {
		maxIepThreshold = maxThreshold
	}
	if minThreshold == 0.0 && maxThreshold == 0.0 {
		minIepThreshold = 0.0
		maxIepThreshold = 0.0
	}

	for k, v := range allSubsequenceIep {
		if maxIep < k {
			maxIep = k
		}

		for _, e := range v {

			if k > minIepThreshold && k < maxIepThreshold {
				filteredData = append(filteredData, e)
			}
		}
	}

	return IepResult{
		QuerySequence:                seq,
		MaxIep:                       maxIep,
		SequenceAndPositions:         allSubsequenceIep[maxIep],
		FilteredSequenceAndPositions: filteredData,
		WholeSequenceIep: iep.PredictIsoelectricPoint(seq, scale),
	}


}

func getAllIepDataSequencial(seq string, minWindow int, scale string) map[float64][]SubSequenceData {
	var subSequences []SubSequenceData
	iepMap := map[float64][]SubSequenceData{}


	for i := minWindow; i < len(seq); i++ {
		for _, subsequence := range GetSubsequences(seq, i) {
			subSequences = append(subSequences, subsequence)
		}
	}


	for _, subsequence := range subSequences {

		predictedIep := iep.PredictIsoelectricPoint(subsequence.Sequence, scale)
		subsequence.Iep = predictedIep

		iepMap[predictedIep] = append(iepMap[predictedIep], subsequence)


	}



	return iepMap

}

func GetSingleThreadedHiepCalculator() SequenceHiepCalculator {
	return &singleThreadedSequenceHiepCalculator{}
}