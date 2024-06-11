import './App.css'
import {createBrowserRouter, RouterProvider} from 'react-router-dom'
import React, {Suspense} from "react";

const loadComponent = (componentPath: string) => {
    const LazyComponent= React.lazy(() => import(componentPath));
    if (!LazyComponent) {
        return null
    }
    return (
        <Suspense fallback={<div>Loading...</div>}>
            <LazyComponent/>
        </Suspense>
    )
}

const router = createBrowserRouter([
    {
        path: '/',
        // element: <Layout/>,
        element: loadComponent("./view/root.tsx"),
        errorElement: loadComponent("./view/error-page.tsx"),
        children: [
            {
                path: 'a',
                // element: <ComponentA/>
                element: loadComponent("./component/ComponentA.tsx")

            },
            {
                path: 'b',
                element: loadComponent("./component/ComponentB.tsx")
            }
        ]
    }
])

class App extends React.Component<any, any> {
    render() {
        return (
            <div id="app">
                <RouterProvider router={router}></RouterProvider>
            </div>
        )
    }
}

export default App