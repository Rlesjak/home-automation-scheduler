import axios from "axios"

export interface m_TriggersGroup {
    uuid: string
    name: string
    description?: string
}

export interface m_CreateMasterTriggersGroup {
    name: string
    description?: string
}

export interface m_CreateChildTriggersGroup {
    parentuuid: string
    name: string
    description?: string
}

interface m_TriggersGroupService {
    getMaster: () => Promise<[m_TriggersGroup]>
    getChildrenOf: (uuid: string) => Promise<[m_TriggersGroup]>
    createMaster: (info: m_CreateMasterTriggersGroup) => Promise<void>
    createChild: (info: m_CreateChildTriggersGroup) => Promise<void>
}

export const TriggersGroup: m_TriggersGroupService = {
    async getMaster() {
        return (await axios.get<[m_TriggersGroup]>("/trigGroup/master")).data
    },
    async getChildrenOf(uuid) {
        return (await axios.get<[m_TriggersGroup]>(`/trigGroup/child/${uuid}`)).data
    },
    async createMaster(info) {
        await axios.post("/trigGroup/master", info)
    },
    async createChild(info) {
        await axios.post("/trigGroup/child", info)
    }
}