import './App.css'
import {createBrowserRouter, RouterProvider} from 'react-router-dom'
import React from "react";
import {loadComponent} from "@/loader.tsx";

let items: {}[] = [
    {
        path: '/',
        componentPath: 'view/frame',
        children: [
            {
                path: '/a',
                componentPath: 'component/ComponentA'
            },
            {
                path: '/b',
                componentPath: 'component/ComponentB'
            }
        ]
    }
]

const loadRouter = (items: any[]): any[] => {
    let r: {}[] = [];
    if (items) {
        for (let i: number = 0; i < items.length; i++) {
            let item = items[i];
            const {componentPath, children} = item;

            if (children) {
                if (children.length > 0) {
                    item.children = loadRouter(children);
                }
            }

            if (componentPath) {
                let element: JSX.Element | null = loadComponent(componentPath)
                if (element) {
                    item.element = element;
                }
            }

            r[i] = item;
        }
    }
    return r;
}

class App extends React.Component {

    render() {
        const router = createBrowserRouter(loadRouter(items));
        return (
            // <ConfigProvider theme={CustomTheme(isDark)}>
                <RouterProvider router={router}></RouterProvider>
            // </ConfigProvider>
        )
    }
}

export default App

// const router = createBrowserRouter([
//     {
//         path: '/',
//         element: loadComponent("view\/frame"),
//         errorElement: loadComponent("view\/error-page"),
//         children: [
//             {
//                 path: 'a',
//                 element: loadComponent("component\/ComponentA")
//
//             },
//             {
//                 path: 'b',
//                 element: loadComponent("component\/ComponentB")
//             }
//         ]
//     }
// ])