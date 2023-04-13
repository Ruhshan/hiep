import ApiResult from '../domains/ApiResult';

class CalculateHiepService {
    public static async calculate(sequence: string): Promise<ApiResult> {
        const result = {'highestIep':11.75,
            'results': [
                {
                    'sequence': 'RFRRHRGSPR',
                    'position': [86, 96]

                },
                {
                    'sequence': 'SRKLPIRSSR',
                    'position': [262, 252]

                },
                {
                    'sequence': 'RKLPIRSSRI',
                    'position': [290, 370]

                }
            ]
        } as ApiResult

        return new Promise<ApiResult>(resolve => resolve(result));

    }
}

export default CalculateHiepService