import { useEffect } from "react";
import {
    Outlet,
    useNavigate
  } from "react-router-dom";

export default function() {

  // Redirect to dashboard route
  let navigate = useNavigate()
  useEffect(() => {
      navigate('/dashboard')
  }, [])

    return <Outlet />
}