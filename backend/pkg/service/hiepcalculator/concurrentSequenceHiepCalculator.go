package hiepcalculator

import (
	. "github.com/Ruhshan/hiep/backend/models/types"
	"github.com/Ruhshan/hiep/backend/pkg/iep"
)

type concurrentSequenceHiepCalculator struct {}


func (c concurrentSequenceHiepCalculator) CalculateMaxIep(seq string, minWindow int) MaxIepResult  {
	maxIep := 0.0

	var subSequences []SequenceAndPosition
	iepMap := map[float64][]SequenceAndPosition{}


	results := make(chan struct {
		predictedIep float64
		subsequence SequenceAndPosition
	})


	for i := minWindow; i< len(seq); i++{
		for _, subsequence := range GetSubsequences(seq, i){
			subSequences = append(subSequences, subsequence)
		}
	}

	for _, subsequence := range subSequences {
		go func(seq SequenceAndPosition) {
			predictedIep := iep.PredictIsoelectricPoint(seq.Sequence)
			results <- struct {
				predictedIep float64
				subsequence SequenceAndPosition
			}{predictedIep, seq}
		}(subsequence)
	}

	for i:=0;i< len(subSequences);i++ {
		res := <-results
		predictedIep := res.predictedIep
		subsequence := res.subsequence

		if predictedIep > maxIep {
			maxIep = predictedIep
		}

		iepMap[predictedIep] = append(iepMap[predictedIep], subsequence)
	}

	close(results)

	return MaxIepResult{
		MaxIep: maxIep,
		SequenceAndPositions: iepMap[maxIep],
	}



}

func GetConcurrentSequenceHiepCalculator() SequenceHiepCalculator {

	return &concurrentSequenceHiepCalculator{}

}
