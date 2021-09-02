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

const logout = () => {
    return service.post("/v1/logout")
};

const listAddress = () => {
    return service.get("/v1/user/addresses")
};

const createAddress = (address) => {
    return service.post("/v1/user/addresses", {
        address,
    })
};

const listCard = (card) => {
    return service.get("/v1/user/cards", {
        card,
    })
};

const createCard = (card) => {
    return service.post("/v1/user/cards", {
        card,
    })
};
