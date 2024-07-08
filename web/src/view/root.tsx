import {Outlet} from "react-router-dom"
import React from "react";
import '@/style/layout.css'
import {Button, ConfigProvider, Layout, Menu} from "antd";
import {CustomTheme} from "@/config/CustomConfigProvider.tsx";

const {Header, Content, Footer, Sider} = Layout;

class Root extends React.Component {
    constructor(props: {}) {
        super(props);

        let isDark: string | null = localStorage.getItem("isDark");
        const isDarkValue = isDark === null ? false : JSON.parse(isDark)
        this.state = {
            isDark: isDarkValue
        }
    }

    changeTheme = () => {
        // @ts-ignore
        this.setState({isDark: !this.state.isDark})
    }

    render() {
        // @ts-ignore
        const {isDark} = this.state

        return (
            <ConfigProvider theme={CustomTheme(isDark)}>
                <div id="layout">
                    <Layout style={{minHeight: '100vh'}}>
                        <Header style={{
                            boxShadow: "0 6px 16px 0 rgba(0, 0, 0, 0.08), 0 3px 6px -4px rgba(0, 0, 0, 0.12), 0 9px 28px 8px rgba(0, 0, 0, 0.05)",
                            padding: "0px"
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
                                {/* @ts-ignore */}
                                <Menu items={null}></Menu>
                            </Sider>

                            <Layout>
                                <Content
                                    style={{
                                        margin: "16px 16px 0px 16px",
                                        borderRadius: "8px",
                                        overflowY: "auto",
                                        background: "white"
                                    }}
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

export default Root