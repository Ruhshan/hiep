package hiepcalculator

import (
	. "github.com/Ruhshan/hiep/backend/models/types"
	"github.com/Ruhshan/hiep/backend/pkg/iep"
	"math"
)

type concurrentSequenceHiepCalculator struct{}

func (c concurrentSequenceHiepCalculator) CalculateIeps(seq string, minWindow int, minThreshold float64, maxThreshold float64) IepResult {
	var maxIep = 0.0
	var allSubsequenceIep = getAllIepData(seq, minWindow)
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
	}

}

func getAllIepData(seq string, minWindow int) map[float64][]SubSequenceData {
	var subSequences []SubSequenceData
	iepMap := map[float64][]SubSequenceData{}


	for i := minWindow; i < len(seq); i++ {
		for _, subsequence := range GetSubsequences(seq, i) {
			subSequences = append(subSequences, subsequence)
		}
	}

	results := make(chan SubsequenceIep, len(subSequences))

	for _, subsequence := range subSequences {
		go func(seq SubSequenceData) {
			predictedIep := iep.PredictIsoelectricPoint(seq.Sequence)
			results <- SubsequenceIep{PredictedIep: predictedIep, Subsequence: seq}
		}(subsequence)
	}


	for i := 0; i < len(subSequences); i++ {
		res := <-results
		predictedIep := res.PredictedIep
		subsequence := res.Subsequence
		subsequence.Iep = predictedIep

		iepMap[predictedIep] = append(iepMap[predictedIep], subsequence)
	}

	close(results)


	return iepMap

}

func GetConcurrentSequenceHiepCalculator() SequenceHiepCalculator {

	return &concurrentSequenceHiepCalculator{}

}
