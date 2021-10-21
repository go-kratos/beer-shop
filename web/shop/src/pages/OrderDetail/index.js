import React, {useEffect, useState} from "react";
import {useParams} from "react-router-dom";
import {getOrderDetail} from "../../api/order"

export default function OrderDetail(props) {
    const {id} = useParams();
    const [data, setData] = useState({});

    useEffect(() => {
        getOrderDetail(id).then((res)=>{
            setData(res.data)
        });
    }, []);

    return <>
        <h1>OrderDetail</h1>
    </>
}