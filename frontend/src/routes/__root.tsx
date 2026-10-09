import { createRootRoute, Outlet } from "@tanstack/react-router";
import './__root.css'

export const Route = createRootRoute({
  component: () => (
    <>
      <nav>
        <img src='/logo.svg' alt='logo' />
        {/*
        <Link to='/'>Home</Link>
        <Link to='/create'>Create</Link>
        */}
      </nav>
      <Outlet />
    </>
  )
})
