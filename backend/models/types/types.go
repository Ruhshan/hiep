package types

type SubSequenceData struct {
	Sequence string `json:"sequence"`
	Position [2]int `json:"position"`
	Iep float64 `json:"iep"`
}


type MaxIepResult struct {
	QuerySequence string `json:"querySequence"`
	MaxIep float64                         `json:"maxIep"`
	SequenceAndPositions []SubSequenceData `json:"sequenceAndPositions"`
}

type AllIepData struct {
	QuerySequence string `json:"querySequence"`
	IepSequenceMap map[float64][]SubSequenceData
}