import React from "react";
import ProductCard from "../../components/ProductCard";

export default function Home(props) {
    return <div>
        <section className="w-full mx-auto bg-nordic-gray-light flex pt-12 md:pt-0 md:items-center bg-cover bg-right"
                 style={{maxWidth: "1600px", height: "24rem", backgroundImage: "url('https://images.unsplash.com/photo-1504502350688-00f5d59bbdeb?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1650&q=80')"}}>
            <div className="container mx-auto">
                <div className="flex flex-col w-full lg:w-1/2 justify-center items-start  px-6 tracking-wide text-white">
                    <h1 className="text-2xl my-4">Your Favorite Beer</h1>
                    <a className="text-xl inline-block no-underline border-b border-gray-600 leading-relaxed hover:text-white hover:border-white"
                       href="#">Buy Now !</a>
                </div>
            </div>
        </section>
        <section className="bg-white py-8">
            <div className="container mx-auto flex items-center flex-wrap pt-4 pb-12">
                <nav id="store" className="w-full z-30 top-0 px-6 py-1">
                    <div
                        className="w-full container mx-auto flex flex-wrap items-center justify-between mt-0 px-2 py-3">

                        <a className="uppercase tracking-wide no-underline hover:no-underline font-bold text-gray-800 text-xl "
                           href="#">
                            Store
                        </a>
                        <div className="flex items-center" id="store-nav-content">

                            <a className="pl-3 inline-block no-underline hover:text-black" href="#">
                                <svg className="fill-current hover:text-black" xmlns="http://www.w3.org/2000/svg"
                                     width="24" height="24" viewBox="0 0 24 24">
                                    <path d="M7 11H17V13H7zM4 7H20V9H4zM10 15H14V17H10z"/>
                                </svg>
                            </a>

                            <a className="pl-3 inline-block no-underline hover:text-black" href="#">
                                <svg className="fill-current hover:text-black" xmlns="http://www.w3.org/2000/svg"
                                     width="24" height="24" viewBox="0 0 24 24">
                                    <path
                                        d="M10,18c1.846,0,3.543-0.635,4.897-1.688l4.396,4.396l1.414-1.414l-4.396-4.396C17.365,13.543,18,11.846,18,10 c0-4.411-3.589-8-8-8s-8,3.589-8,8S5.589,18,10,18z M10,4c3.309,0,6,2.691,6,6s-2.691,6-6,6s-6-2.691-6-6S6.691,4,10,4z"/>
                                </svg>
                            </a>
                        </div>
                    </div>
                </nav>

                <ProductCard/>
                <ProductCard/>
                <ProductCard/>
                <ProductCard/>
                <ProductCard/>
                <ProductCard/>

            </div>
        </section>

        <section className="bg-white py-8">
            <div className="container py-8 px-6 mx-auto">
                <a className="uppercase tracking-wide no-underline hover:no-underline font-bold text-gray-800 text-xl mb-8"
                   href="#">
                    About
                </a>

                <p className="mt-8 mb-8">BeerShop is a simple but complete microservices demo for kratos. created by <a
                        className="text-gray-800 underline hover:text-gray-900"
                        href="https://go-kratos.dev">go-kratos.dev</a></p>
                <p className="mb-8">This project describes the microservices project layout in mono-repo,
                    how the services communicate with each other and the deployment of the services.</p>
                <p className="mb-8">All the backend microservices are built with kratos framework. You could get the project's source code from <a
                    className="text-gray-800 underline hover:text-gray-900"
                    href="https://github.com/go-kratos/beer-shop">https://github.com/go-kratos/beer-shop</a></p>
            </div>

        </section>

        <footer className="container mx-auto bg-white py-8 border-t border-gray-400">
            <div className="container flex px-3 py-8 ">
                <div className="w-full mx-auto flex flex-wrap">
                    <div className="flex w-full lg:w-1/2 ">
                        <div className="px-3 md:px-0">
                            <h3 className="font-bold text-gray-900">Footer</h3>
                            <p className="py-4">
                                beer shop
                            </p>
                        </div>
                    </div>
                </div>
            </div>
        </footer>
    </div>;
}