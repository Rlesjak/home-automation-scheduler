import { useState } from 'react'

function App() {
  const [count, setCount] = useState(0)

  return (
    <div className="m-6 p-10 min-h-screen bg-gray-100 rounded-xl">
      <h1>HA Scheduler front</h1>
      <div className="mt-12 p-6 bg-white rounded-lg drop-shadow-md">
        <button className="px-4 py-2 bg-gray-300 rounded-md drop-shadow-md" onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
      </div>
    </div>
  )
}

export default App
