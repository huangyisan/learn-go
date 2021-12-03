import axios from 'axios'
const service = axios.create({
    baseURL: process.env.VUE_APP_BASE_URL,
    timeout: 1000,
    headers: { 'X-Custom-Header': 'foobar' }
});

export default service