import React, {Fragment} from 'react'
import {
    BrowserRouter as Router,
    Switch,
    Route,
    Link
} from "react-router-dom"

import Nav from "./components/Nav"

import Home from "./pages/Home"
import Login from "./pages/Login";
import Register from "./pages/Register";
import BeerDetail from "./pages/BeerDetail";
import OrderList from "./pages/OrderList";
import Checkout from "./pages/Checkout";

function App() {
    return (
        <Fragment>
            <Nav/>
            <Router>
                <Switch>
                    <Route path="/beer/:id">
                        <BeerDetail/>
                    </Route>
                    <Route path="/login">
                        <Login/>
                    </Route>
                    <Route path="/register">
                        <Register/>
                    </Route>
                    <Route path="/orders">
                        <OrderList/>
                    </Route>
                    <Route path="/checkout">
                        <Checkout/>
                    </Route>
                    <Route path="/">
                        <Home/>
                    </Route>
                </Switch>
            </Router>
        </Fragment>
    );
}

export default App;
