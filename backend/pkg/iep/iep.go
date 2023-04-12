package iep

import (
	"math"
)

var scales = map[string]map[string]float64{
	"EMBOSS":     {"Cterm": 3.6, "pKAsp": 3.9, "pKGlu": 4.1, "pKCys": 8.5, "pKTyr": 10.1, "pk_his": 6.5, "Nterm": 8.6, "pKLys": 10.8, "pKArg": 12.5},
	"DTASelect":  {"Cterm": 3.1, "pKAsp": 4.4, "pKGlu": 4.4, "pKCys": 8.5, "pKTyr": 10.0, "pk_his": 6.5, "Nterm": 8.0, "pKLys": 10.0, "pKArg": 12.0},
	"Solomon":    {"Cterm": 2.4, "pKAsp": 3.9, "pKGlu": 4.3, "pKCys": 8.3, "pKTyr": 10.1, "pk_his": 6.0, "Nterm": 9.6, "pKLys": 10.5, "pKArg": 12.5},
	"Sillero":    {"Cterm": 3.2, "pKAsp": 4.0, "pKGlu": 4.5, "pKCys": 9.0, "pKTyr": 10.0, "pk_his": 6.4, "Nterm": 8.2, "pKLys": 10.4, "pKArg": 12.0},
	"Rodwell":    {"Cterm": 3.1, "pKAsp": 3.68, "pKGlu": 4.25, "pKCys": 8.33, "pKTyr": 10.07, "pk_his": 6.0, "Nterm": 8.0, "pKLys": 11.5, "pKArg": 11.5},
	"Patrickios": {"Cterm": 4.2, "pKAsp": 4.2, "pKGlu": 4.2, "pKCys": 0.0, "pKTyr": 0.0, "pk_his": 0.0, "Nterm": 11.2, "pKLys": 11.2, "pKArg": 11.2},
	"Wikipedia":  {"Cterm": 3.65, "pKAsp": 3.9, "pKGlu": 4.07, "pKCys": 8.18, "pKTyr": 10.46, "pk_his": 6.04, "Nterm": 8.2, "pKLys": 10.54, "pKArg": 12.48},
	"IPC_peptide": {"Cterm":  2.383, "pKAsp":  3.887, "pKGlu":  4.317, "pKCys":  8.297, "pKTyr":  10.071, "pk_his": 6.018, "Nterm":  9.564, "pKLys":  10.517, "pKArg":  12.503},
	"IPC_protein": {"Cterm":  2.869, "pKAsp":  3.872, "pKGlu":  4.412, "pKCys":  7.555, "pKTyr":  10.85, "pk_his": 5.637, "Nterm":  9.094, "pKLys":  9.052, "pKArg":  11.84,},
	"Bjellqvist": {"Cterm": 3.55, "pKAsp": 4.05, "pKGlu": 4.45, "pKCys": 9.0, "pKTyr": 10.0, "pk_his": 5.98, "Nterm": 7.5, "pKLys": 10.0, "pKArg": 12.0,},
}

var aaDict = map[string]string{
	"Asp": "D",
	"Glu": "E",
	"Cys": "C",
	"Tyr": "Y",
	"His": "H",
	"Lys": "K",
	"Arg": "R",
	"Met": "M",
	"Phe": "F",
	"Leu": "L",
	"Val": "V",
	"Ala": "A",
	"Gly": "G",
	"Gln": "Q",
	"Asn": "N",
	"Ile": "I",
	"Trp": "W",
	"Ser": "S",
	"Thr": "T",
	"Sec": "U",
	"Pro": "P",
	"Xaa": "X",
	"Pyl": "O",
	"Asx": "B",
	"Xle": "J",
}

var acidic = []string{"D", "E", "C", "Y"}
var basic = []string{"K", "R", "H"}

var pKcterminal = map[string]float64{"D": 4.55, "E": 4.75}
var pKnterminal = map[string]float64{"A": 7.59, "M": 7.0, "S": 6.93, "P": 8.36, "T": 6.82, "V": 7.44, "E": 7.7}


func countChar(s string, c rune) float64 {
	count := 0
	for _, r := range s {
		if r == c {
			count++
		}
	}
	return float64(count)
}

func PredictIsoelectricPoint(seq string, scaleOpt ...string) float64 {
	var scale = "IPC_protein"

	if len(scaleOpt) > 0 {
		scale = scaleOpt[0]
	}

	var pKCterm = scales[scale]["Cterm"]
	var pKAsp = scales[scale]["pKAsp"]
	var pKGlu = scales[scale]["pKGlu"]
	var pKCys = scales[scale]["pKCys"]
	var pKTyr = scales[scale]["pKTyr"]
	var pKHis = scales[scale]["pk_his"]
	var pKNterm = scales[scale]["Nterm"]
	var pKLys = scales[scale]["pKLys"]
	var pKArg = scales[scale]["pKArg"]
	var pH = 6.51             //starting po pI = 6.5 - theoretically it should be 7, but average protein pI is 6.5 so we increase the probability of finding the solution
	var pHprev = 0.0
	var pHnext = 14.0
	var E = 0.01             //epsilon means precision [pI = pH +- E]
	var temp = 0.01

	nterm := string(seq[0])
	cterm := string(seq[len(seq)-1])


	if scale == "Bjellqvist" {
		if val, exists := pKnterminal[nterm]; exists {
			pKNterm = val
		}

		if val, exists := pKcterminal[cterm]; exists {
			pKCterm = val
		}
	}

	for {
		QN1 := -1.0/(1.0 + math.Pow(10, (pKCterm-pH)))
		QN2 := -countChar(seq,'D')/(1.0+math.Pow(10,(pKAsp-pH)))
		QN3 := -countChar(seq, 'E')/(1.0+math.Pow(10,(pKGlu-pH)))
		QN4 := -countChar(seq, 'C')/(1.0+math.Pow(10,(pKCys-pH)))
		QN5 := -countChar(seq,'Y')/(1.0+math.Pow(10,(pKTyr-pH)))
		QP1 := countChar(seq,'H')/(1.0+math.Pow(10,(pH-pKHis)))
		QP2 := 1.0/(1.0+math.Pow(10,(pH-pKNterm)))
		QP3 := countChar(seq,'K')/(1.0+math.Pow(10,(pH-pKLys)))
		QP4 := countChar(seq, 'R')/(1.0+math.Pow(10,(pH-pKArg)))

		NQ := QN1+QN2+QN3+QN4+QN5+QP1+QP2+QP3+QP4

		if NQ < 0.0{
			temp = pH
			pH = pH-((pH-pHprev)/2.0)
			pHnext = temp
		} else {
			temp = pH
			pH = pH + ((pHnext-pH)/2.0)
			pHprev = temp
		}

		if (pH-pHprev<E) && (pHnext-pH<E){
			return pH
		} //terminal condition, finding pI with given precision



	}

	
}

