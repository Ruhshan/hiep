import {
    Container,
    Textarea,
    VStack,
    FormControl,
    FormLabel,
    FormHelperText,
    FormErrorMessage,
    Button,
    Spacer,
    Center,
    Box,
    Select,
    Input,
    NumberInput,
    NumberInputField,
    NumberInputStepper,
    NumberDecrementStepper, NumberIncrementStepper, SimpleGrid, useBreakpointValue, useColorModeValue, Text, useToast
} from '@chakra-ui/react';
import React from 'react';
import CalculateHiepService from '../service/CalculateHiepService';
import ApiResult from '../domains/ApiResult';
import HiepResult from './HiepResult';
import {InstantHiepRequest} from '../domains/ApiRequest';
import {ProteinFeatureViewer} from './ProteinFeatureViewer';


export function Analyzer() {
    const toast = useToast()
    const [seq, setSeq] = React.useState('')
    const [scale, setScale] = React.useState('IPC_protein')
    const [minIep, setMinIep] = React.useState<number>()
    const [maxIep, setMaxIep] = React.useState<number>()
    const [invalidInput, setInvalidInput] = React.useState(false)
    const [isLoading, setIsLoading] = React.useState(false)
    const [apiResult, setApiResult] = React.useState<ApiResult>({} as ApiResult)
    const [apiErrorMessage, setApiErrorMessage] = React.useState('')

    const scales:Array<string> =['EMBOSS','DTASelect','Solomon','Sillero','Rodwell','Patrickios','Wikipedia',
        'IPC_peptide','IPC_protein','Bjellqvist']


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
        const req: InstantHiepRequest = {sequence: seq, minimumWindowSize:1, scale:scale ,
            minIepThreshold: minIep, maxIepThreshold:maxIep} as InstantHiepRequest
        try{
            const res:ApiResult = await CalculateHiepService.instantHiep(req)
            setApiResult(res)
        }catch (e) {
            setInvalidInput(true)
            setApiErrorMessage(e.response.data.error)
            console.log(e.response)

            toast({
                title:'An error occurred',
                description: e.response.data.error,
                status: 'error',
                duration: 9000,
                isClosable: true,
                position: 'top',
            })
        }

        setIsLoading(false)

    }

    return (
        <VStack>
            <Container color="black">
                <FormControl isInvalid={invalidInput} isRequired>
                    <FormLabel color={useColorModeValue('gray.600', 'white')}>

                            Enter Protein Sequence:

                    </FormLabel>
                    <Textarea onChange={handleInputChange}></Textarea>
                    <FormHelperText>Insert one fasta sequence</FormHelperText>

                </FormControl>
                <FormControl isRequired>
                    <FormLabel color={'gray.600'}>Select Scale:</FormLabel>
                    <Select placeholder='Select Scale' onChange={(e)=>setScale(e.target.value)} value={scale}>
                        {
                            scales.map((scale, index)=>
                                <option value={scale} key={index}>{scale}</option>
                            )
                        }
                    </Select>
                </FormControl>

                <SimpleGrid columns={2} spacing={5}>
                    <FormControl style={{'marginTop':'10px'}}>
                        <FormLabel color={'gray.600'}>Enter Minimum Iep:</FormLabel>
                        <NumberInput>
                            <NumberInputField onChange={(e)=>setMinIep(parseFloat(e.target.value))}/>
                        </NumberInput>
                    </FormControl>
                    <FormControl style={{'marginTop':'10px'}}>
                        <FormLabel color={'gray.600'}>Enter Maximum Iep:</FormLabel>
                        <NumberInput>
                            <NumberInputField onChange={(e)=>setMaxIep(parseFloat(e.target.value))}/>
                        </NumberInput>
                    </FormControl>

                </SimpleGrid>
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