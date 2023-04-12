package requests

type InstantHiepRequest struct {
	Sequence string `json:"sequence"  binding:"required"`
	MinimumWindowSize int `json:"minimumWindowSize"  binding:"required"`
}


