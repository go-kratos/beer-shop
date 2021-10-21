import axios from 'axios';
import {getToken} from "../auth";

const service = axios.create({
    baseURL: process.env.NODE_ENV=== "production" ? "//localhost:8000" : "//localhost:8000",
    timeout: 1000,
});

service.interceptors.request.use(config => {
    config.headers['token'] = getToken() || '';
    return config;
});

export default service