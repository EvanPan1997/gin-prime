import {Outlet, useNavigate} from "react-router-dom"
import {useState} from "react";
import '@/style/layout.css'
import type {MenuProps} from 'antd';
import {Button, ConfigProvider, Layout, Menu} from "antd";
import {CustomTheme, MainContentBorder} from "@/config/CustomConfigProvider.tsx";

type MenuItem = Required<MenuProps>['items'][number];

const {Header, Content, Footer, Sider} = Layout;

function LayoutView(): JSX.Element {
    const [theme, setTheme] = useState(localStorage.getItem("CUSTOM_THEME") === 'dark' ? 'dark' : 'light');
    localStorage.setItem('CUSTOM_THEME', theme);

    const changeTheme = () => {
        setTheme(theme === 'dark' ? 'light' : 'dark');
        localStorage.setItem('CUSTOM_THEME', theme);
    }

    const items: MenuItem[] = [
        {key: '/a', label: 'Option A'},
        {key: '/b', label: 'Option B'}
    ]

    const navigate = useNavigate();

    const handleMenuClick = (path: string) => {
        navigate(path);
    }

    return (
        <ConfigProvider theme={CustomTheme(theme)}>
            <div id="layout">
                <Layout style={{minHeight: '100vh'}}>
                    <Header style={{
                        boxShadow: "0 6px 16px 0 rgba(0, 0, 0, 0.08), 0 3px 6px -4px rgba(0, 0, 0, 0.12), 0 9px 28px 8px rgba(0, 0, 0, 0.05)",
                        padding: "0px",
                    }}>
                        <div id="gp-logo-box">
                            <div id="gp-inner-logo">
                                <img src="/Gin-Prime.png"
                                     alt="Logo"
                                     style={{width: "100%", height: "100%"}}/>
                            </div>
                        </div>
                    </Header>

                    <Layout>
                        <Sider>
                            <Menu items={items} style={{height: '100%'}}
                                  onClick={({key}) => handleMenuClick(key)}></Menu>
                        </Sider>
                        <Layout style={{minHeight: '100%'}}>
                            <Content style={MainContentBorder(theme)}>
                                <Button onClick={changeTheme}>changeTheme</Button>
                                <Outlet/>
                            </Content>
                            <Footer></Footer>
                        </Layout>
                    </Layout>
                </Layout>
            </div>
        </ConfigProvider>
    )
}

export default LayoutView;
