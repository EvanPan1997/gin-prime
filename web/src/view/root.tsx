import {Outlet} from "react-router-dom"
import React from "react";
import '@/style/layout.css'
import {ConfigProvider, Layout, Menu} from "antd";
import {CustomTheme} from "@/config/CustomConfigProvider.tsx";

const {Header, Content, Footer, Sider} = Layout;

class Root extends React.Component {
    constructor(props: {}) {
        super(props);

        let isDark: string | null = localStorage.getItem("isDark");
        if (!isDark) {
            isDark = "false"
            localStorage.setItem("isDark", "false")
        }
        this.state = {
            isDark: isDark == "true"
        }
    }

    changeTheme = () => {
    }

    render() {
        // @ts-ignore
        const {isDark} = this.state

        return (
            <ConfigProvider theme={CustomTheme(isDark)}>
                {/*<ConfigProvider theme={{algorithm: theme.darkAlgorithm}}>*/}
                <div id="layout">
                    <Layout style={{minHeight: '100vh'}}>
                        <Header style={{
                            boxShadow: "0 6px 16px 0 rgba(0, 0, 0, 0.08), 0 3px 6px -4px rgba(0, 0, 0, 0.12), 0 9px 28px 8px rgba(0, 0, 0, 0.05)"
                        }}>
                            <div id="gp-logo-box">
                                <div id="gp-inner-logo">
                                    <img src="" alt="Logo"/>
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
                                    <Outlet/>
                                </Content>
                                <Footer>
                                </Footer>
                            </Layout>

                        </Layout>
                    </Layout>
                </div>
            </ConfigProvider>
        )
    }
}

export default Root