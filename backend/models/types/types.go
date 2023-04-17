package types

type SequenceAndPosition struct {
	Sequence string `json:"sequence"`
	Position [2]int `json:"position"`
}


type MaxIepResult struct {
	QuerySequence string `json:"querySequence"`
	MaxIep float64 `json:"maxIep"`
	SequenceAndPositions []SequenceAndPosition `json:"sequenceAndPositions"`
}