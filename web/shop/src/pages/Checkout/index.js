import React from "react";
import {useState} from 'react'


export default function Checkout(props) {
    return <div className="w-full">
        <div className="md:container md:mx-auto">
            <h1>Checkout</h1>
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
                            <tr>
                                <td className="px-4 py-3">Eric</td>
                                <td className="px-4 py-3">13012345678</td>
                                <td className="px-4 py-3">Some Road, Shanghai, China</td>
                                <td className="px-4 py-3 text-lg text-gray-900">200000</td>
                                <td className="w-10 text-center">
                                    <input name="plan" type="radio"/>
                                </td>
                            </tr>
                            <tr>
                                <td className="px-4 py-3">Tony</td>
                                <td className="px-4 py-3">13955555555</td>
                                <td className="px-4 py-3">Some Road, Beijing, China</td>
                                <td className="px-4 py-3 text-lg text-gray-900">100000</td>
                                <td className="w-10 text-center">
                                    <input name="plan" type="radio"/>
                                </td>
                            </tr>
                            </tbody>
                        </table>
                    </div>
                    <div className="flex pl-4 mt-4 w-full mx-auto">
                        <a className="text-indigo-500 inline-flex items-center md:mb-2 lg:mb-0">Add Address</a>
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
                            <tr>
                                <td className="px-4 py-3">Eric</td>
                                <td className="px-4 py-3">9999999</td>
                                <td className="w-10 text-center">
                                    <input name="plan" type="radio"/>
                                </td>
                            </tr>
                            <tr>
                                <td className="px-4 py-3">Tony</td>
                                <td className="px-4 py-3">88888888</td>
                                <td className="w-10 text-center">
                                    <input name="plan" type="radio"/>
                                </td>
                            </tr>
                            </tbody>
                        </table>
                    </div>
                    <div className="flex pl-4 mt-4 w-full mx-auto">
                        <a className="text-indigo-500 inline-flex items-center md:mb-2 lg:mb-0">Add Card</a>
                    </div>
                </div>
            </section>
            <section className="py-6 border-b">
                <h2>Review your order</h2>
                <div className="py-2 flex flex-wrap md:flex-nowrap">
                    <div className="md:w-64 md:mb-0 mb-6 flex-shrink-0 flex flex-col">
                        <img alt="ecommerce" className="lg:w-1/2 w-full lg:h-auto object-cover object-center rounded"
                             src="https://images.unsplash.com/photo-1613254025696-6f80f3172937?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=800&q=80"/>
                    </div>
                    <div className="md:flex-grow">
                        <h3 className="font-medium text-gray-900 mb-2">Guinness Draught Stout</h3>
                        <p>$6</p>
                        <p>Quantity: 1</p>
                        <p>Size: 500ml</p>
                    </div>
                </div>
                <div className="py-2 flex flex-wrap md:flex-nowrap">
                    <div className="md:w-64 md:mb-0 mb-6 flex-shrink-0 flex flex-col">
                        <img alt="ecommerce" className="lg:w-1/2 w-full lg:h-auto object-cover object-center rounded"
                             src="https://images.unsplash.com/photo-1613254025696-6f80f3172937?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=800&q=80"/>
                    </div>
                    <div className="md:flex-grow">
                        <h3 className="font-medium text-gray-900 mb-2">Guinness Draught Stout</h3>
                        <p>$6</p>
                        <p>Quantity: 1</p>
                        <p>Size: 500ml</p>
                    </div>
                </div>
            </section>
            <section className="flex py-6">
                <span className="title-font font-medium text-2xl text-gray-900">Order total: $58.00</span>
                <button
                    className="flex ml-auto text-white bg-yellow-500 border-0 py-2 px-6 focus:outline-none hover:bg-yellow-600 rounded">Submit
                </button>
            </section>
        </div>
    </div>
}