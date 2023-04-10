import ApiResult from '../domains/ApiResult';

class CalculateHiepService {
    public static async calculate(sequence: string): Promise<ApiResult> {
        const result = {'highestIep':11.75,
            'results': [
                {
                    'sequence': 'RFRRHRGSPR',
                    'region': '86 - 96 aa'

                },
                {
                    'sequence': 'SRKLPIRSSR',
                    'region': '262 - 272 aa'

                },
                {
                    'sequence': 'RKLPIRSSRI',
                    'region': '263 - 273 aa'

                }
            ]
        } as ApiResult

        return new Promise<ApiResult>(resolve => resolve(result));

    }
}

export default CalculateHiepService