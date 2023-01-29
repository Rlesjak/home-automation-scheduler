import { useLayoutEffect, useState } from "react";
import { Icon } from '@iconify/react';
import arrowLongLeft from '@iconify/icons-heroicons/arrow-long-left';
import bars3Solid from '@iconify/icons-heroicons/bars-3-solid';
import cx from "classnames";

interface t_ReactiveMenuItemProps {
    children?: React.ReactNode
    className?: string
    onClick?: React.MouseEventHandler<HTMLDivElement> | undefined
}

export default function(props: t_ReactiveMenuItemProps) {

    return (
        <div
            onClick={props.onClick}
            className={cx({
                'flex items-center w-full h-10 my-auto mt-2 px-3 text-xs truncate rounded-sm text-slate-800 hover:bg-slate-200 hover:cursor-pointer': true,
                [props.className ? props.className : '']: true
            })}
        >
            {props.children}
        </div>
    )
}