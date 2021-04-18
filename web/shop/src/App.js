import React, {Fragment} from 'react'
import {
    BrowserRouter as Router,
    Switch,
    Route,
    Link
} from "react-router-dom"

import Home from "./containers/Home"
import Nav from "./components/Nav"

import './App.css';

function App() {
    return (
        <Fragment>
            <Nav/>
            <Router>
                <Switch>
                    {/*<Route path="/beer/:id">*/}
                        {/*/!*<BeerDetail/>*!/*/}
                    {/*</Route>*/}
                    <Route path="/">
                        <Home/>
                    </Route>
                </Switch>
            </Router>
        </Fragment>
    );
}

export default App;
