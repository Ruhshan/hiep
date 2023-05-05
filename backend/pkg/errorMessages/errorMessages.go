package errorMessages

import "errors"


var ErrContainsInvalidCharacters = errors.New("input contains invalid characters")
var ErrContainsMoreThanOneFastaSequence = errors.New("input contains more than one fasta sequece")
var ErrInvalidScale = errors.New("scale is not valid")