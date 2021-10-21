import React, {useEffect, useState} from "react";
import {getOrderDetail, listOrder} from "../../api/order"


function OrderItem(props) {
    return <div className="border-t border-gray-200">
        <dl>
            <div className="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
                <dt className="text-sm font-medium text-gray-500">
                    ID
                </dt>
                <dd className="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2">
                    {props.data.id}
                </dd>
            </div>
        </dl>
    </div>

}

export default function OrderList(props) {
    const [orderList, setOrderList] = useState([]);

    useEffect(() => {
        listOrder().then((res)=>{
            setOrderList(res.data.results)
        });

        setOrderList([

        ])
        }, []);

    return <>
        <div className="container mx-auto bg-white max-w-screen-lg">
            <div className="px-4 py-5 sm:px-6">
                <h3 className="text-lg leading-6 font-medium text-gray-900">
                    Order List
                </h3>
                <p className="mt-1 max-w-2xl text-sm text-gray-500">
                    All orders
                </p>
            </div>
            {orderList.map((x)=><OrderItem data={x}/>)}
        </div>

    </>
}