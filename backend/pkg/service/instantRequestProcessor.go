package service

import (
	"github.com/Ruhshan/hiep/backend/models/requests"
	"github.com/Ruhshan/hiep/backend/models/types"
	"github.com/Ruhshan/hiep/backend/pkg/errors"
	"github.com/Ruhshan/hiep/backend/pkg/fasta"
	"regexp"
	"strings"
)

type InstantRequestProcessor interface {
	Process(r requests.InstantHiepRequest) (*types.MaxIepResult, error)
}

type instantRequestProcessor struct {
}

func parseFasta(seq string)(string, error)  {
	var fastas, err = fasta.ParseFastaFromString(seq)

	if err != nil{
		return "", err
	}

	if len(fastas)>1{
		return "", errors.ErrContainsMoreThanOneFastaSequence
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
		return "", errors.ErrCotainsInvalidCharacters
	}


	return seq,nil


}

func (i instantRequestProcessor) Process(r requests.InstantHiepRequest) (*types.MaxIepResult, error) {
	var seq, err = sanitizeSequence(r)

	if err!=nil{
		return nil, err
	}
	var res = CalculateMaxIep(seq, r.MinimumWindowSize)
	return &res, nil
}

func NewInstantRequestProcessor() InstantRequestProcessor {
	return &instantRequestProcessor{}
}
