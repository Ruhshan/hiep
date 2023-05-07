interface SequenceAndRegion {
    sequence: string
    position: Array<number>
    iep: number
}

export default interface ApiResult {
    querySequence: string
    maxIep: number
    sequenceAndPositions: Array<SequenceAndRegion>
    filteredSequenceAndPositions: Array<SequenceAndRegion> | null
    wholeSequenceIep: number
}

