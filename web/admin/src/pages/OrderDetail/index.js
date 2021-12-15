import React, {useEffect, useState} from "react";
import {PageHeader} from "antd";
import {useParams} from "react-router-dom";
import {getBeerDetail} from "../../api/catalog";

export default function OrderDetail() {
    let {id} = useParams();
    const [data, setData] = useState({});

    useEffect(() => {
        getBeerDetail(id).then((res) => {
            setData(res.data)
        });

        setData(
            {
                "id": 1,
                "name": "cool beer1",
                "price": "5.99",
                "images": ["https://images.unsplash.com/photo-1588704487282-e7c55e0448bc?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=668&q=80"],
            });
    }, []);

    return (
        <div>
            <PageHeader
                ghost={false}
                onBack={() => window.history.back()}
                title="Order Detail"
            />
            <div>
                {id}
            </div>
        </div>
    )
}
