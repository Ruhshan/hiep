package errorMessages

import "errors"


var ErrContainsInvalidCharacters = errors.New("Input contains invalid characters")
var ErrContainsMoreThanOneFastaSequence = errors.New("Input contains more than one fasta sequece")
var ErrInvalidScale = errors.New("Scale is not valid")