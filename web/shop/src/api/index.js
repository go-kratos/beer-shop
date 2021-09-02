import axios from 'axios';
import {getToken} from "../auth";

const service = axios.create({
    baseURL: process.env.NODE_ENV=== "production" ? "//beer.go-kratos.dev" : "//localhost:8000",
    timeout: 1000,
});

service.interceptors.request.use(config => {
    if (getToken()) {
        config.headers['Authorization'] = 'Bearer ' + getToken() || '';
    }
    return config;
});

export default service