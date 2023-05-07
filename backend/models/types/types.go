package types

type SubSequenceData struct {
	Sequence string `json:"sequence"`
	Position [2]int `json:"position"`
	Iep float64 `json:"iep"`
}

type SubsequenceIep struct {
	PredictedIep float64
	Subsequence  SubSequenceData
}


type IepResult struct {
	QuerySequence string `json:"querySequence"`
	MaxIep float64                         `json:"maxIep"`
	SequenceAndPositions []SubSequenceData `json:"sequenceAndPositions"`
	FilteredSequenceAndPositions []SubSequenceData `json:"filteredSequenceAndPositions"`
	WholeSequenceIep float64 `json:"wholeSequenceIep"`
}