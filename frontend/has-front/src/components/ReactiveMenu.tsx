import { useLayoutEffect, useState } from "react";
import cx from "classnames";
import { Icon } from '@iconify/react';
import arrowLongLeft from '@iconify/icons-heroicons/arrow-long-left';
import bars3Solid from '@iconify/icons-heroicons/bars-3-solid';

interface t_ReactiveMenuProps {
    children?: React.ReactNode
}

export default function(props: t_ReactiveMenuProps) {

    const [isNavOpen, setIsNavOpen] = useState(true);

    // Figure out the initial state of the menu
    // On desktop default is open
    // On mobile default is closed
    useLayoutEffect(()=>{
        // Tailwind md: is at 768px, so here we use the same breakpoint
        if (window.innerWidth < 768) {
            setIsNavOpen(false)
        }
    }, [])

    return (
        <div className={
            cx({
                "relative h-full border-gray-300 transition-all": true,
                "w-80 border-r md:w-52": isNavOpen,
                "w-0": !isNavOpen
            })
        }>
            <div
                onClick={() => setIsNavOpen(true)}
                className={ cx({
                    "absolute top-0 w-16 py-1 border-r border-b border-gray-300 bg-white rounded-br-md": true,
                    "hover:cursor-pointer hover:bg-gray-100": true,
                    "hidden": isNavOpen
                })}
            >
                <Icon 
                    icon={bars3Solid}
                    className='w-7 h-7 mx-auto text-gray-600'
                />
            </div>
            
            <div
                onClick={() => setIsNavOpen(false)}
                className="py-1 border-b border-gray-300 hover:cursor-pointer hover:bg-gray-100"
            >
                <Icon 
                    icon={arrowLongLeft}
                    className='w-7 h-7 ml-auto mr-3 [&>path]:stroke-1'
                />
            </div>
            <div className="px-3">
                {props.children}
            </div>
        </div>
    )
}