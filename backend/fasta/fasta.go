package fasta

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type fasta struct {
	Index    int
	Id       string
	Sequence string
}

func ParseFastaFromFile(filepath string) ([]*fasta, error) {
	dat, err := ioutil.ReadFile(filepath)


	if len(dat) == 0 {
		return nil, err
	}

	data_string := string(dat)

	return ParseFastaFromString(data_string)
}

func ParseFastaFromString(data_string string) ([]*fasta, error) {
	fastas := []*fasta{}

	splited := strings.Split(data_string, "\n")

	is_start := true

	f_id := ""
	f_sequence := ""
	f_index := 0
	for _, s := range splited {
		if len(s) != 0 {
			if s[0] == 62 && is_start == true {

				id := fmt.Sprintf("%s", s[1:len(s)])
				f_id = id
				f_sequence = ""
				is_start = false

			} else if s[0] == 62 && is_start == false {

				new_fasta := new(fasta)
				new_fasta.Index = f_index
				new_fasta.Id = strings.TrimSpace(f_id)
				new_fasta.Sequence = strings.TrimSpace(f_sequence)

				fastas = append(fastas, new_fasta)

				id := fmt.Sprintf("%s", s[1:len(s)])

				f_id = id
				f_sequence = ""

				f_index += 1
			} else {
				f_sequence = f_sequence + s
			}
		}

	}

	new_fasta := new(fasta)
	new_fasta.Index = f_index
	new_fasta.Id = strings.TrimSpace(f_id)
	new_fasta.Sequence = strings.TrimSpace(f_sequence)

	fastas = append(fastas, new_fasta)

	return fastas, nil
}

