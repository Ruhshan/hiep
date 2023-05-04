package requests

type InstantHiepRequest struct {
	Sequence string `json:"sequence"  binding:"required"`
	MinimumWindowSize int `json:"minimumWindowSize"  binding:"required"`
	MinIepThreshold float64 `json:"minIepThreshold"`
	MaxIepThreshold float64 `json:"maxIepThreshold"`
	Scale string `json:"scale"`
}


