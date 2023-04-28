package hiepcalculator

import (
	. "github.com/Ruhshan/hiep/backend/models/types"
	"github.com/Ruhshan/hiep/backend/pkg/iep"
	"math"
)

type concurrentSequenceHiepCalculator struct{}

func (c concurrentSequenceHiepCalculator) CalculateIeps(seq string, minWindow int, minThreshold float64, maxThreshold float64) IepResult {
	var maxIep = 0.0
	var allIepData = getAllIepData(seq, minWindow)
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

	for k, v := range allIepData.IepSequenceMap {
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
		SequenceAndPositions:         allIepData.IepSequenceMap[maxIep],
		FilteredSequenceAndPositions: filteredData,
	}

}

func getAllIepData(seq string, minWindow int) AllIepData {
	var subSequences []SubSequenceData
	iepMap := map[float64][]SubSequenceData{}

	results := make(chan struct {
		predictedIep float64
		subsequence  SubSequenceData
	})

	for i := minWindow; i < len(seq); i++ {
		for _, subsequence := range GetSubsequences(seq, i) {
			subSequences = append(subSequences, subsequence)
		}
	}

	for _, subsequence := range subSequences {
		go func(seq SubSequenceData) {
			predictedIep := iep.PredictIsoelectricPoint(seq.Sequence)
			results <- struct {
				predictedIep float64
				subsequence  SubSequenceData
			}{predictedIep, seq}
		}(subsequence)
	}

	for i := 0; i < len(subSequences); i++ {
		res := <-results
		predictedIep := res.predictedIep
		subsequence := res.subsequence
		subsequence.Iep = predictedIep

		iepMap[predictedIep] = append(iepMap[predictedIep], subsequence)
	}

	close(results)

	return AllIepData{
		QuerySequence:  seq,
		IepSequenceMap: iepMap,
	}

}

func GetConcurrentSequenceHiepCalculator() SequenceHiepCalculator {

	return &concurrentSequenceHiepCalculator{}

}
