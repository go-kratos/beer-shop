import service from './index'

const listBeer = (pageNum, pageSize) => {
    return service.get("/v1/catalog/beers", {
        pageNum, pageSize
    })
};

const getBeerDetail = (id) => {
    return service.get("/v1/catalog/beers/"+id)
};
