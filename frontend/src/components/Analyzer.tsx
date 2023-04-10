import {
    Container,
    Textarea,
    VStack,
    FormControl,
    FormLabel,
    FormHelperText,
    FormErrorMessage,
    Button, Spacer
} from '@chakra-ui/react';
import React from 'react';
import CalculateHiepService from '../service/CalculateHiepService';
import ApiResult from '../domains/ApiResult';
import HiepResult from './HiepResult';

export function Analyzer() {
    const [seq, setSeq] = React.useState('');
    const [invalidInput, setInvalidInput] = React.useState(false);
    const [apiResult, setApiResult] = React.useState<ApiResult>({} as ApiResult);

    const handleInputChange = (e) => {
        if (isSequenceValid(e.target.value)) {
            setSeq(e.target.value)
            setInvalidInput(false);
        } else {
            setInvalidInput(true);
        }
    };

    function isSequenceValid(sequence: string): boolean {
        sequence = sequence.trim().replace(/\n|\r/g, '')
        const regex = /^[ACDEFGHIKLMNPQRSTVWY]+$/i;
        return regex.test(sequence);
    }

    const performSearch = async () => {
        const res:ApiResult = await CalculateHiepService.calculate(seq)

        setApiResult(res)
    }

    return (
        <VStack>
            <Container color="black">
                <FormControl isInvalid={invalidInput}>
                    <FormLabel>Enter Protein Sequence:</FormLabel>
                    <Textarea onChange={handleInputChange}></Textarea>


                    {invalidInput ? (<FormErrorMessage>Invalid input</FormErrorMessage>) :
                        (<FormHelperText>Currently we don&apos;t support FASTA</FormHelperText>)
                    }

                </FormControl>

                <div style={{'marginTop':'10px'}}>
                    <Button colorScheme='teal' variant='outline' size='sm' onClick={performSearch} isDisabled={invalidInput}>
                        Search
                    </Button>
                </div>
            </Container>
            <Container>
                <HiepResult apiResult={apiResult}></HiepResult>
            </Container>
        </VStack>
    );
}