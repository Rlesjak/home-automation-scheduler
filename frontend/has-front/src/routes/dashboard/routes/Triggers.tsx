import axios from "axios";
import { TriggersGroup } from "../../../api/model/triggersGroup";
import ReactiveMenu from "../../../components/ReactiveMenu";
import ReactiveMenuItem from "../../../components/ReactiveMenuItem";
import { useTemplatedFetch } from "../../../hooks/customFetch";
import { Icon } from '@iconify/react';
import rectangleGroup from '@iconify/icons-heroicons/rectangle-group';
import plus20Solid from '@iconify/icons-heroicons/plus-20-solid';
import tadpoleIcon from '@iconify/icons-svg-spinners/tadpole';
import exclamationCircle from '@iconify/icons-heroicons/exclamation-circle';
import { useRef, useState } from "react";
import cx from "classnames";

export default function() {

    // Build sidemenu with the list of master trigger groups
    let [triggersTemplate, trigUseFetch] = useTemplatedFetch(
        "triggers", 
        TriggersGroup.getMaster,
        (grp) => (
            <ReactiveMenuItem key={grp.uuid}>
                <Icon
                    icon={rectangleGroup}
                    className='h-6 w-6 text-slate-900'
                />
                <p className="w-full ml-2 text-slate-700 truncate">
                    {grp.name}
                </p>
            </ReactiveMenuItem>
        ),
        () => (
            <ReactiveMenuItem>
                <Icon
                    icon={tadpoleIcon}
                    className='h-6 w-6 mx-auto text-slate-900'
                />
            </ReactiveMenuItem>
        ),
        (err, refetch) => (
            <div onClick={() => refetch()}>
                <ReactiveMenuItem>
                    <Icon
                        icon={exclamationCircle}
                        className='h-6 w-6 mx-auto text-red-400'
                    />
                </ReactiveMenuItem>
            </div>
        )
    )

    let [,,, doTriggerFetch] = trigUseFetch;


    let [enteringNewGroup, setEnteringNewGroup] = useState(false)
    const inputEl = useRef<HTMLInputElement>(null)

    async function createNewGroup() {
        const groupName = inputEl.current?.value;
        if (groupName) {
            await TriggersGroup.createMaster({
                name: groupName
            })
            doTriggerFetch(true)
        }
    }

    function addGroupInputRenderer() {

        return (
            <>
                <input autoFocus ref={inputEl} className="w-full mr-2 px-1 appearance-none text-sm bg-slate-50 border-2 border-slate-400 rounded-sm"></input>
                <Icon
                    onClick={() => createNewGroup()}
                    icon={plus20Solid}
                    className='h-6 w-8 mx-auto text-slate-900 rounded-full hover:bg-slate-600 hover:text-slate-100'
                />
            </>
        )
    }
    

    return (
        <div
            className="flex h-full"
        >
            <ReactiveMenu>
                {triggersTemplate}

                <ReactiveMenuItem
                    onClick={() => enteringNewGroup ? null : setEnteringNewGroup(true)} 
                    className={cx({
                        'bg-slate-200': enteringNewGroup
                    })}
                >
                    {enteringNewGroup
                        ? addGroupInputRenderer()
                        :<Icon
                            icon={plus20Solid}
                            className='h-6 w-6 mx-auto text-slate-900'
                        />
                    }
                </ReactiveMenuItem>
            </ReactiveMenu>
            <div className="w-full p-6 bg-slate-100">

            </div>
        </div>
    )
}