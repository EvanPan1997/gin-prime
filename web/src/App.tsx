import './App.css'
import {createBrowserRouter, RouterProvider} from 'react-router-dom'
import React from "react";
import {loadComponent} from "@/loader.tsx";


const router = createBrowserRouter([
    {
        path: '/',
        element: loadComponent("view", "root"),
        errorElement: loadComponent("view", "error-page"),
        children: [
            {
                path: 'a',
                element: loadComponent("component", "ComponentA")

            },
            {
                path: 'b',
                element: loadComponent("component", "ComponentB")
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