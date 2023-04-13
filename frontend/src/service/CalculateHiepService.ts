import ApiResult from '../domains/ApiResult';
import {InstantHiepRequest} from '../domains/ApiRequest';
import ApiClient from './ApiClient';


const hiepPath = 'hiep'

class CalculateHiepService {


    public static async instantHiep(request: InstantHiepRequest): Promise<ApiResult>{
        try{
            const res = await ApiClient.post(`${hiepPath}/instant/calculate`, request)
            console.log('in service '+JSON.stringify(res.data))
            return res.data
        }catch (e){
            throw new Error('Failed to fetch data')
        }
    }
}

export default CalculateHiepService