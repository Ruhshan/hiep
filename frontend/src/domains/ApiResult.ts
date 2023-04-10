interface SequenceAndRegion {
    sequence: string
    region: string
}

export default interface ApiResult {
    highestIep: number
    results: Array<SequenceAndRegion>
}

