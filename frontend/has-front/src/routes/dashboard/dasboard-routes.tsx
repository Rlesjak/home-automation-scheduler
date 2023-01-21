import { RouteObject } from "react-router-dom";
import Elements from "./routes/Elements";
import Triggers from "./routes/Triggers";

export default [
    {
        path: 'elements',
        element: <Elements />
    },
    {
        path: 'triggers',
        element: <Triggers />
    }
] as RouteObject[]