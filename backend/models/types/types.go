package types

type SequenceAndPosition struct {
	Sequence string
	Position [2]int
}


type MaxIepResult struct {
	MaxIep float64 `json:"maxIep"`
	SequenceAndPositions []SequenceAndPosition `json:"sequenceAndPositions"`
}