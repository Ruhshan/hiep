import {
    Container,
    Textarea,
    VStack,
    FormControl,
    FormLabel,
    FormHelperText,
    FormErrorMessage,
    Button, Spacer, Center, Box
} from '@chakra-ui/react';
import React from 'react';
import CalculateHiepService from '../service/CalculateHiepService';
import ApiResult from '../domains/ApiResult';
import HiepResult from './HiepResult';
import {InstantHiepRequest} from '../domains/ApiRequest';
import {ProteinFeatureViewer} from './ProteinFeatureViewer';


export function Analyzer() {
    const [seq, setSeq] = React.useState('');
    const [invalidInput, setInvalidInput] = React.useState(false);
    const [isLoading, setIsLoading] = React.useState(false);
    const [apiResult, setApiResult] = React.useState<ApiResult>({} as ApiResult);
    const [apiErrorMessage, setApiErrorMessage] = React.useState('')


    const handleInputChange = (e) => {
        if(e.target.value.trim.length==0){
            setInvalidInput(false)
            setApiErrorMessage('')
        }
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
        const req: InstantHiepRequest = {sequence: seq, minimumWindowSize:1 } as InstantHiepRequest
        try{
            const res:ApiResult = await CalculateHiepService.instantHiep(req)
            setApiResult(res)
        }catch (e) {
            setInvalidInput(true)
            setApiErrorMessage(e.response.data.error)
            console.log(e.response)
        }

        setIsLoading(false)

    }

    return (
        <VStack>
            <Container color="black">
                <FormControl isInvalid={invalidInput}>
                    <FormLabel>Enter Protein Sequence:</FormLabel>
                    <Textarea onChange={handleInputChange}></Textarea>


                    {invalidInput ? (<FormErrorMessage>{apiErrorMessage}</FormErrorMessage>) :
                        (<FormHelperText>Insert one fasta sequence</FormHelperText>)
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

                <ProteinFeatureViewer apiResult={apiResult}></ProteinFeatureViewer>

        </VStack>
    );
}