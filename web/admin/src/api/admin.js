import service from './index'

export const login = (username, password) => {
    return service.post("/v1/login", {
        username,
        password,
    })
};

export const register = (username, email, password) => {
    return service.post("/v1/register", {
        username,
        email,
        password,
    })
};

