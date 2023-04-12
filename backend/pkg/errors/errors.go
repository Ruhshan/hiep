package errors

import "errors"

var ErrCotainsInvalidCharacters = errors.New("Input contains invalid characters")

var ErrContainsMoreThanOneFastaSequence = errors.New("Input contains more than one fasta sequece")