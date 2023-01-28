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


export default function() {

    let triggersTemplate = useTemplatedFetch(
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
    

    return (
        <div className="flex h-full">
            <ReactiveMenu>
                {triggersTemplate}
                <ReactiveMenuItem>
                    <Icon
                        icon={plus20Solid}
                        className='h-6 w-6 mx-auto text-slate-900'
                    />
                </ReactiveMenuItem>
            </ReactiveMenu>
            <div className="w-full p-6 bg-slate-100">

            </div>
        </div>
    )
}