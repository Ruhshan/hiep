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
import {InstantHiepRequest} from '../domains/ApiRequest';
import axios, {AxiosResponse} from 'axios';

export function Analyzer() {
    const [seq, setSeq] = React.useState('');
    const [invalidInput, setInvalidInput] = React.useState(false);
    const [isLoading, setIsLoading] = React.useState(false);
    const [apiResult, setApiResult] = React.useState<ApiResult>({} as ApiResult);

    const handleInputChange = (e) => {
        setSeq(e.target.value)
    };

    function isSequenceValid(sequence: string): boolean {
        sequence = sequence.trim().replace(/\n|\r/g, '')
        const regex = /^[ACDEFGHIKLMNPQRSTVWY]+$/i;
        return regex.test(sequence);
    }

    const performSearch = async () => {
        setApiResult({} as ApiResult)
        setIsLoading(true)
        const req: InstantHiepRequest = {sequence: seq, minimumWindowSize:50 } as InstantHiepRequest

        const res:ApiResult = await CalculateHiepService.instantHiep(req)
        setApiResult(res)
        setIsLoading(false)

    }

    return (
        <VStack>
            <Container color="black">
                <FormControl isInvalid={invalidInput}>
                    <FormLabel>Enter Protein Sequence:</FormLabel>
                    <Textarea onChange={handleInputChange}></Textarea>


                    {invalidInput ? (<FormErrorMessage>Invalid input</FormErrorMessage>) :
                        (<FormHelperText>Insert single raw sequence of fasta</FormHelperText>)
                    }

                </FormControl>

                <div style={{'marginTop':'10px'}}>
                    <Button colorScheme='teal' variant='outline' size='sm' onClick={performSearch} isDisabled={invalidInput} isLoading={isLoading}>
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