import React, { useState } from 'react'
import {
  RouterProvider,
} from "react-router-dom";
import router from './router'
import MainSidebar from './routes/dashboard/sections/MainSidebar'
import PageContent from './routes/dashboard/sections/PageContent'

function App() {
  const [count, setCount] = useState(0)

  return (
    <React.Fragment>
      <RouterProvider router={router} />
    </React.Fragment>
  )
}

export default App
