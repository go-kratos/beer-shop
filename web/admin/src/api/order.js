import service from './index'

export const createOrder = () => {
    return service.post("/admin/v1/orders", {

    })
};

export const listOrder = (pageNum, pageSize) => {
    return service.get("/admin/v1/orders", {
        pageNum, pageSize
    })
};

export const getOrderDetail = (id) => {
    return service.get("/admin/v1/orders/"+id)
};
