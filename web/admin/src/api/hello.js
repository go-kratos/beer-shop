import service from './index'

const getHello = (id) => {
    return service.get("/v1/hello/"+id)
};
