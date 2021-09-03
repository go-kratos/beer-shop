import service from './index'

export const listCartItem = () => {
    return service.get("/v1/cart")
};

export const addCartItem = (id) => {
    return service.put("/v1/cart", {
        id,
    })
};

export const createOrder = () => {
    return service.post("/v1/orders", {

    })
};

export const listOrder = (pageNum, pageSize) => {
    return service.get("/v1/orders", {
        pageNum, pageSize
    })
};

export const getOrderDetail = (id) => {
    return service.get("/v1/orders/"+id)
};