import axios from 'axios';

const service = axios.create({
    baseURL: process.env.NODE_ENV=== "production" ? "//beer.go-kratos.dev" : "//localhost:8000",
    timeout: 1000,
});

service.interceptors.request.use(config => {
    config.headers['token'] = sessionStorage.getItem('token') || '';
    return config;
});

export default service