import ApiResult from '../domains/ApiResult';
import {Box, Container, Table, TableContainer, Tbody, Td, Th, Thead, Tr} from '@chakra-ui/react';

interface Props {
    apiResult: ApiResult;
}

function HiepResult(props: Props) {
    const apiResult: ApiResult = props.apiResult;

    function formatPosition(position: Array<number>): string {
        if (position.length == 2) {
            return `${position[0]} to ${position[1]} aa`;
        } else {
            return '';
        }

    }

    return (
        <Container>
            {apiResult.sequenceAndPositions ? (<>
                        <Box border="1px" borderRadius="5" borderColor="blackAlpha.500" marginBottom="10px">
                            <TableContainer>
                                <Table variant="simple">
                                    <Thead>
                                        <Tr>
                                            <Th>Hieghest IEP</Th>
                                            <Th>{apiResult.maxIep}</Th>
                                        </Tr>
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
                                        {apiResult.sequenceAndPositions.map((sequenceAndRegion, index) => (
                                            <Tr key={index}>
                                                <Td>{sequenceAndRegion.sequence}</Td>
                                                <Td>{formatPosition(sequenceAndRegion.position)}</Td>
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