import React, {Fragment} from "react"
import './App.css';
import {
    BrowserRouter as Router,
    Switch,
    Route,
    Link,
    useLocation,
    useRouteMatch, NavLink
} from "react-router-dom"

import {Layout, Menu, Avatar, Dropdown} from "antd";
import Dashboard from "./pages/Dashboard";
import BeerList from "./pages/BeerList";
import CustomerList from "./pages/CustomerList";
import OrderList from "./pages/OrderList";

import {
    FileProtectOutlined,
    DashboardOutlined,
    AppstoreOutlined,
    TeamOutlined,
    UserOutlined,
} from '@ant-design/icons';
import BeerDetail from "./pages/BeerDetail";
import CustomerDetail from "./pages/CustomerDetail";
import OrderDetail from "./pages/OrderDetail";
import {renderRoutes} from "react-router-config";

const {Header, Sider, Content} = Layout;

function Main(props) {
    const routes = [
        {
            name: "Dashboard",
            icon: <DashboardOutlined/>,
            exact: true, path: "/", component: Dashboard,
        },
        {
            name: "Beers", icon: <AppstoreOutlined/>,
            exact: true,
            path: "/beers",
            component: BeerList,
            routes: [
                {
                    exact: true,
                    path: "/beers/:id",
                    component: BeerDetail
                },
            ]
        },
        {
            name: "Customers",
            icon: <TeamOutlined/>,
            exact: true,
            path: "/customers", component: CustomerList,
            routes: [
                {
                    exact: true,
                    path: "/customers/:id", component: CustomerDetail
                },
            ]
        },
        {
            name: "Orders", icon: <FileProtectOutlined/>,
            exact: true,
            path: "/orders",
            component: OrderList,
            routes: [
                {
                    path: "/orders/:id",
                    component: OrderDetail
                },
            ]
        },
    ];

    return <Layout>
        <Sider theme="light" collapsible>
            <Menu mode="inline">
                {routes.map((x) =>
                    <Menu.Item key={x.path} icon={x.icon}>
                        <NavLink activeClassName='ant-menu-item-selected' to={x.path}>
                            {x.name}
                        </NavLink>
                    </Menu.Item>
                )}
            </Menu>
        </Sider>
        <Layout style={{
            minHeight: "100vh",
        }}>
            <Header
                style={{
                    backgroundColor: "#fff",
                    display: "flex",
                    justifyContent: "flex-end",
                    alignItems: "center",
                    padding: "0 20px",
                }}>
                <Dropdown placement="bottomRight" arrow overlay={<Menu>
                    <Menu.Item>
                        <a rel="noopener noreferrer" href="#">
                            menu item
                        </a>
                    </Menu.Item>
                </Menu>}>
                    <a onClick={e => e.preventDefault()}>
                        <Avatar size="large" icon={<UserOutlined/>}/>
                    </a>
                </Dropdown>
            </Header>
            <Content className="content">
                <Switch>
                    {renderRoutes(routes)}
                </Switch>
            </Content>
        </Layout>
    </Layout>
}

function App() {
    return (
        <Router>
            <Main/>
        </Router>
    );
}

export default App;
