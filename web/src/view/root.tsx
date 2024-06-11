import {Outlet} from "react-router-dom"
import React from "react";

class Layout extends React.Component {
    render() {
        return (
            <div>
                <Outlet/>
            </div>
        )
    }
}

export default Layout