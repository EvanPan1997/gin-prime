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
        // "loader": () => {
        //     return redirect('/a')
        // },
        children: [
            {
                path: 'a',
                element: loadComponent("./component/ComponentA.tsx")

            },
            {
                path: 'b',
                element: loadComponent("./component/ComponentB.tsx")
            }
        ]
    }
])

class App extends React.Component {
    render() {
        return (
            // <CustomConfigProvider>
                <RouterProvider router={router}></RouterProvider>
            // </CustomConfigProvider>
        )
    }
}

export default App