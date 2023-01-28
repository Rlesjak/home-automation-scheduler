import axios from "axios";

export function setupAxios() {
    axios.defaults.baseURL = import.meta.env.VITE_API_URL
}