import React, {useState, useEffect} from "react";
import {useParams} from "react-router-dom";
import {getBeerDetail} from "../../api/beer";
import {addCartItem} from "../../api/order"
import {login} from "../../api/user";

export default function BeerDetail(props) {
    const {id} = useParams();
    const [data, setData] = useState({});
    const [qty, setQty] = useState(1);

    useEffect(() => {
        getBeerDetail(id).then((res)=>{
            setData(res.data)
        });
    }, []);

    const handleSubmit = (event) => {
        event.preventDefault();
        addCartItem(id).then(
            ()=>{
                console.log("ok!");
            }
        ).catch(
            ()=>{
                console.log("err!")
            },
        )
    };

    return <div>
        <section className="text-gray-600 body-font overflow-hidden">
            <div className="container px-5 py-2 mx-auto max-w-screen-md">
                <div className="flex flex-wrap">
                    <img alt="ecommerce" className="lg:w-1/2 w-full lg:h-auto h-64 object-cover object-center rounded"
                         src="https://images.unsplash.com/photo-1613254025696-6f80f3172937?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=800&q=80"/>
                    <div className="lg:w-1/2 w-full lg:pl-10 lg:py-6 mt-6 lg:mt-0">
                        <h2 className="text-sm title-font text-gray-500 tracking-widest">PRODUCT NAME</h2>
                        <h1 className="text-gray-900 text-3xl title-font font-medium mb-1">{data.name}</h1>
                        <p className="leading-relaxed">{data.description}</p>
                        <div className="flex mt-6 items-center pb-5 border-b-2 border-gray-100 mb-5">
                            <div className="flex ml-6 items-center">
                                <span className="mr-3">Qty</span>
                                <div className="relative">
                                    <input type="number" onChange={(evt)=>{setQty(evt.target.value)}} value={qty}
                                        className="rounded border appearance-none border-gray-300 py-2 focus:outline-none focus:ring-2 focus:ring-yellow-200 focus:border-yellow-500 text-base pl-3 pr-10">
                                    </input>
                                    <span
                                        className="absolute right-0 top-0 h-full w-10 text-center text-gray-600 pointer-events-none flex items-center justify-center">
                <svg fill="none" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2"
                     className="w-4 h-4" viewBox="0 0 24 24">
                  <path d="M6 9l6 6 6-6"/>
                </svg>
              </span>
                                </div>
                            </div>
                        </div>
                        <div className="flex">
                            <span className="title-font font-medium text-2xl text-gray-900">{data.price}</span>
                            <button
                                className="flex ml-auto text-white bg-yellow-500 border-0 py-2 px-6 focus:outline-none hover:bg-yellow-600 rounded-none"
                            onClick={(evt)=>{handleSubmit(evt)}}
                            >Add to Cart</button>
                        </div>
                    </div>
                </div>
            </div>
        </section>
    </div>
        ;
}