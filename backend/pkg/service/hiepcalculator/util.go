package hiepcalculator

import . "github.com/Ruhshan/hiep/backend/models/types"

func GetSubsequences(seq string, window int) [] SequenceAndPosition {
	var subSequences []SequenceAndPosition

	for i := 0; i<len(seq)-window;i++{

		var sAndP = SequenceAndPosition{
			Sequence: seq[i:i+window],
			Position: [2]int{i, i+window},
		}

		subSequences = append(subSequences, sAndP)
	}

	return subSequences


}
