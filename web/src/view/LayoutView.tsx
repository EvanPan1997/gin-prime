import {Outlet, useNavigate} from "react-router-dom"
import {useState} from "react";
import '@/style/layout.css'
import type {MenuProps} from 'antd';
import {Button, ConfigProvider, Layout, Menu} from "antd";
import {CustomTheme, MainContentBorder, TopMenuCssProperties} from "@/config/CustomConfigProvider.tsx";

type MenuItem = Required<MenuProps>['items'][number];

const {Header, Content, Footer, Sider} = Layout;

function LayoutView(): JSX.Element {
    const navigate = useNavigate();

    const modules = [
        {key: 'mod1', label: 'Group1'},
        {key: 'mod2', label: 'Group2'},
        {key: 'mod3', label: 'Group3'},
    ]
    const moduleMenus: { [key: string]: MenuItem[] } = {
        'mod1': [
            {key: '/a', label: 'Option A'},
            {key: '/b', label: 'Option B'}
        ],
        'mod2': [
            {type: 'group', key: 'c', label: 'Option C', children: [{key: '/b', label: 'Sub Option D'}]}
        ],
        'mod3': [
            {key: '/a', label: 'Option E'}
        ],
    }
    let defaultModuleKey = '';
    let defaultMenuGroup: MenuItem[] = [];
    // let default
    for (let key in moduleMenus) {
        if (moduleMenus.hasOwnProperty(key) && key !== '') {
            defaultModuleKey = key;
            defaultMenuGroup = moduleMenus[key];
            break;
        }
    }
    // 默认组别, 用于计算默认key
    const defaultGroup = moduleMenus[defaultModuleKey];
    let groupItem = null;
    if (defaultGroup.length > 0) {
        groupItem = defaultGroup[0];
        // @ts-ignore
        while ('children' in groupItem) {
            if (groupItem.children && groupItem.children.length > 0) {
                groupItem = groupItem.children[0]
            } else {
                break;
            }
        }
    }
    const [defaultGroupItem] = useState(groupItem);

    // navigate(menus[defaultGroupItem.key].path);

    const [currentGroup, setCurrentGroup] = useState(defaultGroup);


    const [theme, setTheme] = useState(localStorage.getItem("CUSTOM_THEME") === 'dark' ? 'dark' : 'light');
    localStorage.setItem('CUSTOM_THEME', theme);

    const changeTheme = (): void => {
        setTheme(theme === 'dark' ? 'light' : 'dark');
        localStorage.setItem('CUSTOM_THEME', theme);
    }

    const handleTopMenuClick = (key: string): void => {
        if (moduleMenus.hasOwnProperty(key)) {
            setCurrentGroup(moduleMenus[key])
        } else {
            setCurrentGroup([])
        }
    }
    const handleMenuClick = (key: string): void => {
        navigate(key);
    }

    return (
        <ConfigProvider theme={CustomTheme(theme)}>
            <div id="layout">
                <Layout style={{minHeight: '100vh'}}>
                    <Header style={{
                        boxShadow: "0 6px 16px 0 rgba(0, 0, 0, 0.08), 0 3px 6px -4px rgba(0, 0, 0, 0.12), 0 9px 28px 8px rgba(0, 0, 0, 0.05)",
                        padding: "0px", display: "flex", flexDirection: "row", justifyContent: "space-between"
                    }}>
                        {/*Logo组件: 点击事件回到首页*/}
                        {/*TODO 后续修改新增模块部分可能需要调整*/}
                        <div style={{cursor: "pointer"}} onClick={() => {
                            {/*@ts-ignore*/}
                            window.location.href = defaultMenuGroup[0].key
                            // window.location.href ='/'
                        }}>
                            <div id="gp-logo-box">
                                <div id="gp-inner-logo">
                                    <img src="/Gin-Prime.png" alt="Logo" style={{width: "auto", height: "100%"}}/>
                                </div>
                            </div>
                        </div>

                        {/*<div style={{display: "flex", flexDirection: "row"}}>*/}
                        <Menu items={modules} mode='horizontal' style={TopMenuCssProperties(theme)} defaultActiveFirst
                              onClick={({key}) => handleTopMenuClick(key)}/>
                        <div style={{height: "100%", width: "200px", marginLeft: "30px"}}></div>
                        {/*</div>*/}
                    </Header>

                    <Layout>
                        <Sider>
                            {/*@ts-ignore*/}
                            <Menu items={currentGroup} style={{height: '100%'}} mode='vertical' defaultSelectedKeys={[defaultGroupItem.key]}
                                  onClick={({key}) => handleMenuClick(key)}
                            ></Menu>
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
