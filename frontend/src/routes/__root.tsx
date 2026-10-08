import { createRootRoute, Link, Outlet } from "@tanstack/react-router";

export const Route = createRootRoute({
  component: () => (
    <>
      <nav>
        <Link to='/'>Home</Link>
        <Link to='/create'>Create</Link>
      </nav>
      <Outlet />
    </>
  )
})