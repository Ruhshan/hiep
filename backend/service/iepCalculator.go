package service

import (
	"fmt"
	"github.com/Ruhshan/hiep/backend/iep"
)

type sequenceAndPosition struct {
	sequence string
	position [2]int

}

func getSubsequences(seq string, window int) []sequenceAndPosition {
	var subSequences []sequenceAndPosition

	for i := 0; i<len(seq)-window;i++{

		var sAndP = sequenceAndPosition{
			sequence: seq[i:i+window],
			position: [2]int{i, i+window},
		}

		subSequences = append(subSequences, sAndP)
	}

	return subSequences


}

func CalculateMaxIep(seq string, minWindow int)  {
	iepMap := map[float64][]sequenceAndPosition{}
	maxIep := 0.0

	var subSequences []sequenceAndPosition


	for i := minWindow; i< len(seq); i++{
		for _, subsequence := range getSubsequences(seq, i){
			subSequences = append(subSequences, subsequence)
		}
	}


	for _, subsequence := range subSequences{
		predictedIep := iep.PredictIsoelectricPoint(subsequence.sequence)

		if predictedIep > maxIep{
			maxIep = predictedIep
		}


		iepMap[predictedIep] = append(iepMap[predictedIep], subsequence)
	}

	fmt.Println(iepMap[maxIep])




}
