import React, {useState, useEffect} from "react";
import {Avatar, Button, List, PageHeader, Pagination, Skeleton} from "antd";
import {listBeer} from "../../api/catalog";
import {Link, useRouteMatch} from "react-router-dom";

export default function OrderList() {
    const [orderList, setOrderList] = useState([]);
    let {path, url} = useRouteMatch();

    useEffect(() => {
        // listBeer().then((res) => {
        //     setBeerList(res.data.results)
        // });

        setOrderList([
            {
                "id": 1,
                "name": "cool beer1",
                "price": "5.99",
                "images": ["https://images.unsplash.com/photo-1588704487282-e7c55e0448bc?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=668&q=80"],
            },
            {
                "id": 2,
                "name": "cool beer2",
                "price": "6.99",
                "images": ["https://images.unsplash.com/photo-1588704487282-e7c55e0448bc?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=668&q=80"],
            },
            {
                "id": 3,
                "name": "cool beer3",
                "price": "7.99",
                "images": ["https://images.unsplash.com/photo-1588704487282-e7c55e0448bc?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=668&q=80"],
            },
            {
                "id": 4,
                "name": "cool beer4",
                "price": "8.99",
                "images": ["https://images.unsplash.com/photo-1588704487282-e7c55e0448bc?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=668&q=80"],
            },
            {
                "id": 5,
                "name": "cool beer5",
                "price": "9.99",
                "images": ["https://images.unsplash.com/photo-1588704487282-e7c55e0448bc?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=668&q=80"],
            },
            {
                "id": 6,
                "name": "cool beer6",
                "price": "10.99",
                "images": ["https://images.unsplash.com/photo-1588704487282-e7c55e0448bc?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=668&q=80"],
            },
        ]);
    }, []);

    return (
        <div>
            <PageHeader
                ghost={false}
                onBack={() => window.history.back()}
                title="Order"
            />
            <div style={{background: "#fff"}}>
                <List
                    pagination={<Pagination defaultCurrent={1} total={50}/>}
                    dataSource={orderList}
                    renderItem={(item, i) => (
                        <List.Item
                            actions={[<a key={i}>edit</a>]}
                        >
                            <Skeleton avatar title={false} loading={item.loading} active>
                                <List.Item.Meta
                                    avatar={<Avatar shape="square" size={64} src={item.images[0]}/>}
                                    title={<Link to={`${path}/${item.id}`}>{item.name}</Link>}
                                    description={item.price}
                                />
                            </Skeleton>
                        </List.Item>
                    )}
                />
            </div>
        </div>
    )
}
