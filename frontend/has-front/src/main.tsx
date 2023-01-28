import React from 'react'
import ReactDOM from 'react-dom/client'
import { setupAxios } from './api/client/axiosConf'
import App from './App'
import './index.css'

// Setup axios client
setupAxios()

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
)
