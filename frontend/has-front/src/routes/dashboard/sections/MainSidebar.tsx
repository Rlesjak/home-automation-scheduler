import { NavLink } from "react-router-dom";
import rectangleGroup from '@iconify/icons-heroicons/rectangle-group';
import clockIcon from '@iconify/icons-heroicons/clock';

import SidebarIcon from "../../../components/SidebarIcon";

interface t_MainSidebarProps {
    className?: string
}

function MainSidebar(props: t_MainSidebarProps) {
    return (
        <aside className={(props.className || "") + ""}>
            <div className='flex w-full h-full items-center justify-center gap-6 overflow-auto border-b border-gray-300 md:justify-start md:flex-col md:py-6 md:border-r'>
                <NavLink to={'elements'}>
                    { ({isActive}) => (
                        <SidebarIcon icon={rectangleGroup} active={isActive} />
                    )}
                </NavLink>
                <NavLink to={'triggers'}>
                    { ({isActive}) => (
                        <SidebarIcon icon={clockIcon} active={isActive} />
                    )}
                </NavLink>
            </div>
        </aside>
    )
}

export default MainSidebar