import React, {Fragment} from "react"
import './App.css';
import {
    BrowserRouter as Router,
    Switch,
    Route,
    Link,
    useLocation,
} from "react-router-dom"
import {matchRoutes, renderRoutes} from "react-router-config";

import {Layout, Menu, Avatar, Dropdown} from "antd";
import Dashboard from "./pages/Dashboard";
import Beer from "./pages/Beer";
import Customer from "./pages/Customer";
import Order from "./pages/Order";

import {
    FileProtectOutlined,
    DashboardOutlined,
    AppstoreOutlined,
    TeamOutlined,
    UserOutlined,
} from '@ant-design/icons';

const {Header, Sider, Content} = Layout;

function Main(props) {
    const location = useLocation();
    const routes = [
        {name: "Dashboard", exact: true, path: "/", component: Dashboard, icon: <DashboardOutlined/>},
        {name: "Beers", exact: true, path: "/beers", component: Beer, icon: <AppstoreOutlined />},
        {name: "Customers", path: "/customers", component: Customer, icon: <TeamOutlined />},
        {name: "Orders", path: "/orders", component: Order, icon: <FileProtectOutlined />},
    ];
    const branch = matchRoutes(routes, location.pathname);
    const selectedPath = branch[0] ? branch[0].route.path : "/";
    return <Layout>
        <Sider theme="light" collapsible>
            <Menu mode="inline" selectedKeys={[selectedPath]}>
                {routes.map((x) =>
                    <Menu.Item key={x.path} icon={x.icon}>
                        <Link to={x.path}>
                            {x.name}
                        </Link>
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
