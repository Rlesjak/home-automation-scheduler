import {
    createHashRouter
} from "react-router-dom";

import Root from "./routes/root";

import Dashboard from "./routes/dashboard/dashboard";
import DasboardRoutes from "./routes/dashboard/dasboard-routes";

export default createHashRouter([
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