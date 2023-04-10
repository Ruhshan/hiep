import ApiResult from '../domains/ApiResult';

class CalculateHiepService {
    public static async calculate(sequence: string): Promise<ApiResult> {
        const result = {} as ApiResult

        return new Promise<ApiResult>(resolve => resolve(result));

    }
}