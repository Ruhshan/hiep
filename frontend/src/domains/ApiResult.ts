interface SequenceAndRegion {
    sequence: string
    position: Array<number>
}

export default interface ApiResult {
    maxIep: number
    results: Array<SequenceAndRegion>
}

