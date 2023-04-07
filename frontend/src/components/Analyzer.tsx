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

export function Analyzer() {
    const [seq, setSeq] = React.useState('');
    const [invalidInput, setInvalidInput] = React.useState(false);

    const handleInputChange = (e) => {
        console.log(isSequenceValid(e.target.value));
        if (isSequenceValid(e.target.value)) {
            setInvalidInput(false);
        } else {
            setInvalidInput(true);
        }
    };

    function isSequenceValid(sequence: string): boolean {
        const regex = /^[ACDEFGHIKLMNPQRSTVWY]+$/i;
        return regex.test(sequence);
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
                    <Button colorScheme='teal' variant='outline' size='sm'>
                        Search
                    </Button>
                </div>
            </Container>
        </VStack>
    );
}