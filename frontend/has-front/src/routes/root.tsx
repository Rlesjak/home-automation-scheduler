import { useEffect } from "react";
import {
    Outlet,
    useNavigate,
    useLocation
  } from "react-router-dom";

export default function() {

  const location = useLocation()
  const navigate = useNavigate()
  
  useEffect(() => {
    // Redirect to dashboard route
    if (location.pathname == '/') navigate('dashboard')
  }, [])

  return <Outlet />
}