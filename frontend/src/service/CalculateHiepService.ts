import ApiResult from '../domains/ApiResult';
import {InstantHiepRequest} from '../domains/ApiRequest';
import ApiClient from './ApiClient';


const hiepPath = 'hiep'

class CalculateHiepService {

    public static async instantHiep(request: InstantHiepRequest): Promise<ApiResult> {
        return await ApiClient.post(`${hiepPath}/instant/calculate`, request).then(res => res.data);

    }
}

export default CalculateHiepService