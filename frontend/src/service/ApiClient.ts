import axios from 'axios';


const API_URL = 'http://0.0.0.0:8989/api/v1/'

const ApiClient = axios.create({
    baseURL: API_URL,
    headers : {
        Accept: 'application/json'
    }
})


ApiClient.interceptors.response.use(
    response => {
        return response
    },
    async error => {
        return Promise.reject(error)
    }
)

export default ApiClient