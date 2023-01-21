import { Icon, IconifyIcon } from '@iconify/react';
import cx from 'classnames';

interface t_SidebarIconProps {
    active: boolean
    icon: IconifyIcon
}

export default function(props: t_SidebarIconProps) {
    return (
        <div className={
            cx({
                'w-9 h-9 flex justify-center items-center rounded-xl transition-all': true,
                'border-2 border-gray-600': props.active,
                'hover:bg-gray-100': !props.active
            })
        }>
            <Icon 
                icon={ props.icon } 
                className='w-5 h-5 [&>path]:stroke-2 text-gray-700'
            />
        </div>
    )
}