import axios from 'axios';
import {getToken} from "../auth";

const service = axios.create({
    baseURL: process.env.NODE_ENV=== "production" ? "//localhost:8100" : "//localhost:8100",
    timeout: 1000,
});

service.interceptors.request.use(config => {
    config.headers['Authorization'] = 'Bearer ' + getToken() || '';
    return config;
});

export default service
