import service from './index'


export const listUser = (pageNum, pageSize) => {
    return service.get("/admin/v1/users", {
        pageNum, pageSize
    })
};

export const getUserDetail = (id) => {
    return service.get("/admin/v1/users/"+id)
};
