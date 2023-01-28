import { useLayoutEffect, useState } from "react";
import cx from "classnames";
import { Icon } from '@iconify/react';
import arrowLongLeft from '@iconify/icons-heroicons/arrow-long-left';
import bars3Solid from '@iconify/icons-heroicons/bars-3-solid';

interface t_ReactiveMenuItemProps {
    children?: React.ReactNode
}

export default function(props: t_ReactiveMenuItemProps) {

    return (
        <div
            className='flex items-center w-full h-10 my-auto mt-2 px-3 text-xs truncate rounded-sm text-slate-800 hover:bg-slate-200 hover:cursor-pointer'
        >
            {props.children}
        </div>
    )
}