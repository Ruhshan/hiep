import axios from 'axios';


const API_URL = 'http://localhost:8989/v1/hiep'

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