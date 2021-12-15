import React, {useState, useEffect} from "react";
import Modal from "../../components/Modal";
import {listCartItem} from "../../api/order"
import {listCard} from "../../api/user"
import {listAddress} from "../../api/user"

export default function Checkout(props) {
    const [addrList, setAddrList] = useState([]);
    const [cardList, setCardList] = useState([]);
    const [cartItemList, setCartItemList] = useState([]);
    const [total, setTotal] = useState(0);


    useEffect(() => {
        listCartItem().then((res)=>{
            setCartItemList(res.data.results)
        });

        listAddress().then((res)=>{
            setAddrList(res.data.results)
        });

        listCard().then((res)=>{
            setCardList(res.data.results)
        });

        setAddrList([
            {
                "id": 1,
                "name": "Eric",
                "mobile": "13012345678",
                "address": "Some Road, Shanghai, China",
                "postCode": 200000,
            },
            {
                "id": 2,
                "name": "Tony",
                "mobile": "13000000000",
                "address": "Some Road, Beijing, China",
                "postCode": 100000,
            },
        ]);

        setCardList([
            {
                "id": 1,
                "name": "Eric",
                "cardNo": "12345678",
            },
            {
                "id": 1,
                "name": "Tony",
                "cardNo": "88888888",
            },
        ]);

        setCartItemList([
            {
                "id": 1,
                "name": "A beer",
                "image": "https://images.unsplash.com/photo-1613254025696-6f80f3172937?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=800&q=80",
                "price": "5.99",
                "quantity": 10,
                "size": "500ml",
            },
            {
                "id": 2,
                "name": "B beer",
                "image": "https://images.unsplash.com/photo-1613254025696-6f80f3172937?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=800&q=80",
                "price": "10.99",
                "quantity": 5,
                "size": "500ml",
            },
        ]);

        setTotal(100);
    }, []);


    return <div className="container mx-auto bg-white max-w-screen-md">
        <h1>Checkout</h1>
        <section>
            <Modal/>
        </section>
        <section className="py-6 border-b">
            <h2>Address</h2>
            <div className="py-2 mx-auto">
                <div className="w-full mx-auto overflow-auto">
                    <table className="table-auto w-full text-left whitespace-no-wrap">
                        <thead>
                        <tr>
                            <th className="px-4 py-3 title-font tracking-wider font-medium text-gray-900 text-sm bg-gray-100 rounded-tl rounded-bl">Name</th>
                            <th className="px-4 py-3 title-font tracking-wider font-medium text-gray-900 text-sm bg-gray-100">Mobile</th>
                            <th className="px-4 py-3 title-font tracking-wider font-medium text-gray-900 text-sm bg-gray-100">Address</th>
                            <th className="px-4 py-3 title-font tracking-wider font-medium text-gray-900 text-sm bg-gray-100">Post
                                Code
                            </th>
                            <th className="w-10 title-font tracking-wider font-medium text-gray-900 text-sm bg-gray-100 rounded-tr rounded-br"/>
                        </tr>
                        </thead>
                        <tbody>
                        {
                            addrList && addrList.map((x) =>
                                <tr>
                                    <td className="px-4 py-3">{x.name}</td>
                                    <td className="px-4 py-3">{x.mobile}</td>
                                    <td className="px-4 py-3">{x.address}</td>
                                    <td className="px-4 py-3 text-lg text-gray-900">{x.postCode}</td>
                                    <td className="w-10 text-center">
                                        <input name="plan" type="radio"/>
                                    </td>
                                </tr>
                            )
                        }
                        </tbody>
                    </table>
                </div>
                <div className="flex pl-4 mt-4 w-full mx-auto">
                    <a className="text-yellow-500 inline-flex items-center md:mb-2 lg:mb-0">Add Address</a>
                </div>
            </div>
        </section>
        <section className="py-6 border-b">
            <h2>Card</h2>
            <div className="py-2 mx-auto">
                <div className="w-full mx-auto overflow-auto">
                    <table className="table-auto w-full text-left whitespace-no-wrap">
                        <thead>
                        <tr>
                            <th className="px-4 py-3 title-font tracking-wider font-medium text-gray-900 text-sm bg-gray-100 rounded-tl rounded-bl">Name</th>
                            <th className="px-4 py-3 title-font tracking-wider font-medium text-gray-900 text-sm bg-gray-100 rounded-tl rounded-bl">CardNo</th>
                            <th className="w-10 title-font tracking-wider font-medium text-gray-900 text-sm bg-gray-100 rounded-tr rounded-br"/>
                        </tr>
                        </thead>
                        <tbody>
                        {
                            cardList && cardList.map((x) =>
                                <tr>
                                    <td className="px-4 py-3">{x.name}</td>
                                    <td className="px-4 py-3">{x.cardNo}</td>
                                    <td className="w-10 text-center">
                                        <input name="plan" type="radio"/>
                                    </td>
                                </tr>
                            )
                        }
                        </tbody>
                    </table>
                </div>
                <div className="flex pl-4 mt-4 w-full mx-auto">
                    <a className="text-yellow-500 inline-flex items-center md:mb-2 lg:mb-0">Add Card</a>
                </div>
            </div>
        </section>
        <section className="py-6 border-b">
            <h2>Review your order</h2>
            {
                cartItemList && cartItemList.map((x) =>
                    <div className="py-2 flex flex-wrap md:flex-nowrap">
                        <div className="md:w-64 md:mb-0 mb-6 flex-shrink-0 flex flex-col">
                            <img alt="ecommerce"
                                 className="lg:w-1/2 w-full lg:h-auto object-cover object-center rounded"
                                 src={x.image}/>
                        </div>
                        <div className="md:flex-grow">
                            <h3 className="font-medium text-gray-900 mb-2">{x.name}</h3>
                            <p>{x.price}</p>
                            <p>Quantity: {x.quantity}</p>
                            <p>Size: {x.size}</p>
                        </div>
                    </div>
                )
            }
        </section>
        <section className="flex py-6">
            <span className="title-font font-medium text-2xl text-gray-900">Order total: {total}</span>
            <button
                className="flex ml-auto text-white bg-yellow-500 border-0 py-2 px-6 focus:outline-none hover:bg-yellow-600 rounded-none">Submit
            </button>
        </section>
    </div>
}