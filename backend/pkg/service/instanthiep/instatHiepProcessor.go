package instanthiep

import (
	"github.com/Ruhshan/hiep/backend/models/requests"
	"github.com/Ruhshan/hiep/backend/models/types"
	"github.com/Ruhshan/hiep/backend/pkg/errorMessages"
	"github.com/Ruhshan/hiep/backend/pkg/fasta"
	"github.com/Ruhshan/hiep/backend/pkg/service/hiepcalculator"
	"regexp"
	"strings"
)

type HiepProcessor interface {
	ProcessPayload(r requests.InstantHiepRequest) (*types.MaxIepResult, error)
}

type instantHiepProcessor struct {
	calculator hiepcalculator.SequenceHiepCalculator
}

func parseFasta(seq string)(string, error)  {
	var fastas, err = fasta.ParseFastaFromString(seq)

	if err != nil{
		return "", err
	}

	if len(fastas)>1{
		return "", errorMessages.ErrContainsMoreThanOneFastaSequence
	}else{
		return fastas[0].Sequence, nil
	}
}

func sanitizeSequence(r requests.InstantHiepRequest) (string, error){
	var seq = strings.TrimSpace(r.Sequence)

	if seq[0]=='>'{
		parsedSeq, err := parseFasta(seq)

		if err !=nil{
			return "", err
		}

		seq = parsedSeq

	}

	regex := regexp.MustCompile("^[ACDEFGHIKLMNPQRSTVWY]+$")

	if regex.MatchString(seq) == false{
		return "", errorMessages.ErrContainsInvalidCharacters
	}


	return seq,nil


}

func (i instantHiepProcessor) ProcessPayload(r requests.InstantHiepRequest) (*types.MaxIepResult, error) {
	var seq, err = sanitizeSequence(r)

	if err!=nil{
		return nil, err
	}

	var res = i.calculator.CalculateMaxIep(seq, r.MinimumWindowSize)

	return &res, nil
}

func GetInstantHiepProcessor(calculator hiepcalculator.SequenceHiepCalculator) HiepProcessor {
	return &instantHiepProcessor{calculator}
}