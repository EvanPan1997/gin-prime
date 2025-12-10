import './App.css'
import {createBrowserRouter, RouterProvider} from 'react-router-dom'
import React from "react";
import {loadComponent} from "@/loader.tsx";

let items: {}[] = [
    {
        path: '/',
        componentPath: 'view/LayoutView',
        errorComponentPath: 'view/ErrorPage',
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
            const {componentPath, errorComponentPath, children} = item;

            if (children) {
                if (children.length > 0) {
                    item.children = loadRouter(children);
                }
            }

            if (componentPath) {
                const element: JSX.Element | null = loadComponent(componentPath)
                if (element) {
                    item.element = element;
                }
            }

            if (errorComponentPath) {
                const errorElement: JSX.Element | null = loadComponent(errorComponentPath)
                if (errorElement) {
                    item.errorElement = errorElement
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
                <RouterProvider router={router}></RouterProvider>
        )
    }
}

export default App
