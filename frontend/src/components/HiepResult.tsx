import ApiResult from '../domains/ApiResult';
import {Box, Container, Table, TableContainer, Tbody, Td, Th, Thead, Tr} from '@chakra-ui/react';

interface Props {
    apiResult: ApiResult;
}

function HiepResult(props: Props) {
    const apiResult: ApiResult = props.apiResult;
    return (
        <Container>
            {apiResult.results ? (<>
                        <Box border="1px" borderRadius="5" borderColor="blackAlpha.500" marginBottom='10px'>
                            <TableContainer>
                                <Table variant="simple">
                                    <Thead>
                                        <Th>Hieghest IEP</Th>
                                        <Th>{apiResult.highestIep}</Th>
                                    </Thead>
                                </Table>
                            </TableContainer>
                        </Box>
                        <Box border="1px" borderRadius="5" borderColor="blackAlpha.500">
                            <TableContainer>
                                <Table variant="simple">
                                    <Thead>
                                        <Tr>
                                            <Th>Sequence</Th>
                                            <Th>Region</Th>
                                        </Tr>
                                    </Thead>
                                    <Tbody>
                                        {apiResult.results.map((sequenceAndRegion, index) => (
                                            <Tr key={index}>
                                                <Td>{sequenceAndRegion.sequence}</Td>
                                                <Td>{sequenceAndRegion.region}</Td>
                                            </Tr>
                                        ))}
                                    </Tbody>

                                </Table>
                            </TableContainer>
                        </Box>
                    </>
                )
                : (<p></p>)}
        </Container>
    );
}

export default HiepResult;