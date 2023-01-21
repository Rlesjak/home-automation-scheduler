import {
    createBrowserRouter
} from "react-router-dom";

import Root from "./routes/root";

import Dashboard from "./routes/dashboard/dashboard";
import DasboardRoutes from "./routes/dashboard/dasboard-routes";

export default createBrowserRouter([
    {
        path: "/",
        element: <Root />,
        children: [
            {
                path: "dashboard",
                element: <Dashboard />,
                children: DasboardRoutes
            }
        ]
    },
])