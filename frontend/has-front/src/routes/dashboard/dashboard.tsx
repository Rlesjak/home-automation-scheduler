import React from "react";
import { Outlet } from "react-router-dom";
import MainSidebar from "./sections/MainSidebar";
import PageContent from "./sections/PageContent";

export default function() {

    return (
        <React.Fragment>
            <MainSidebar className="fixed top-0 left-0 right-0 h-16 md:right-auto md:h-auto md:bottom-0 md:w-16" />
            <PageContent className="fixed top-16 left-0 bottom-0 right-0 md:top-0 md:left-16">
                <Outlet />
            </PageContent>
        </React.Fragment>
    )
}