import service from './index'

const listCartItem = () => {
    return service.get("/v1/cart")
};

const addCartItem = (id) => {
    return service.get("/v1/cart", {
        id,
    })
};

const createOrder = () => {
    return service.post("/v1/orders", {

    })
};

const listOrder = (pageNum, pageSize) => {
    return service.get("/v1/orders", {
        pageNum, pageSize
    })
};