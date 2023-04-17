interface SequenceAndRegion {
    sequence: string
    position: Array<number>
}

export default interface ApiResult {
    querySequence: string
    maxIep: number
    sequenceAndPositions: Array<SequenceAndRegion>
}

