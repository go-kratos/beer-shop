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

export const logout = () => {
    return service.post("/v1/logout")
};

export const listAddress = () => {
    return service.get("/v1/user/addresses")
};

export const createAddress = (address) => {
    return service.post("/v1/user/addresses", {
        address,
    })
};

export const listCard = (card) => {
    return service.get("/v1/user/cards", {
        card,
    })
};

export const createCard = (card) => {
    return service.post("/v1/user/cards", {
        card,
    })
};
