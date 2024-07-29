import {Outlet} from "react-router-dom"
import React from "react";
import '@/style/layout.css'
import type {MenuProps} from 'antd';
import {Button, ConfigProvider, Layout, Menu} from "antd";
import {CustomTheme, MainContentBorder} from "@/config/CustomConfigProvider.tsx";

type MenuItem = Required<MenuProps>['items'][number];

const {Header, Content, Footer, Sider} = Layout;

class Frame extends React.Component {
    constructor(props: {}) {
        super(props);

        let customTheme: string | null = localStorage.getItem("CUSTOM_THEME");
        const theme = customTheme==='dark' ? 'dark' : 'light';

        const items: MenuItem[] = [
            {key: '/a', label: 'Option A'},
            {key: '/b', label: 'Option B'}
        ]

        localStorage.setItem('CUSTOM_THEME', theme);
        this.state = {
            theme: theme,
            menuItems: items,
        }
    }

    changeTheme = () => {
        // @ts-ignore
        let theme = this.state.theme
        theme = theme==='dark' ? 'light' : 'dark';
        this.setState({theme: theme});
        localStorage.setItem('CUSTOM_THEME', theme);
    }

    handleMenuClick = (path: string) => {
        window.location.href = path;
    }

    render() {
        // @ts-ignore
        const {theme, menuItems} = this.state
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
                                         style={{
                                             width: "100%",
                                             height: "100%"
                                         }}/>
                                </div>
                            </div>
                        </Header>

                        <Layout>
                            <Sider>
                                <Menu items={menuItems} style={{height: '100%'}}
                                      onClick={({key}) => this.handleMenuClick(key)}></Menu>
                            </Sider>

                            <Layout style={{minHeight: '100%'}}>
                                <Content
                                    style={MainContentBorder(theme)}
                                >
                                    <Button onClick={this.changeTheme}>changeTheme</Button>
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
}

export default Frame