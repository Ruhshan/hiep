package types

type SubSequenceData struct {
	Sequence string `json:"sequence"`
	Position [2]int `json:"position"`
	Iep float64 `json:"iep"`
}

type IepData struct {
	PredictedIep float64
	Subsequence  SubSequenceData
}


type IepResult struct {
	QuerySequence string `json:"querySequence"`
	MaxIep float64                         `json:"maxIep"`
	SequenceAndPositions []SubSequenceData `json:"sequenceAndPositions"`
	FilteredSequenceAndPositions []SubSequenceData `json:"filteredSequenceAndPositions"`
}

type AllIepData struct {
	QuerySequence string `json:"querySequence"`
	IepSequenceMap map[float64][]SubSequenceData
}