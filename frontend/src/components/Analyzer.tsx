import {Container, Textarea, VStack, FormControl, FormLabel, FormHelperText} from '@chakra-ui/react';
import React from 'react';

export function Analyzer() {
    const [seq, setSeq] = React.useState('')

    const handleInputChange = (e)=>{
        console.log(e.target.value)
    }

    return (
        <VStack>
            <Container color='black'>
                <FormControl>
                    <FormLabel>Enter Protein Sequence:</FormLabel>
                    <Textarea onChange={handleInputChange}></Textarea>
                    <FormHelperText>Currently we don't support FASTA</FormHelperText>
                </FormControl>
            </Container>
        </VStack>
    );
}