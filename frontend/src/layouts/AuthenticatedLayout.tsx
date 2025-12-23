// layouts/authenticated-layout.tsx
import { Outlet, useNavigate } from "@tanstack/react-router"

// import { useAuth } from "@/auth/auth-context"
import { MainNav } from "@/components/NavigationBar"

export function AuthenticatedLayout() {
//   const { isAuthenticated, isLoading } = useAuth()
  const navigate = useNavigate()

//   React.useEffect(() => {
//     if (!isLoading && !isAuthenticated) {
//       navigate({
//         to: "/login",
//         replace: true,
//       })
//     }
//   }, [isAuthenticated, isLoading, navigate])

//   if (isLoading) {
//     return (
//       <div className="flex h-screen items-center justify-center">
//         <span className="text-sm text-muted-foreground">
//           Checking authentication…
//         </span>
//       </div>
//     )
//   }

//   if (!isAuthenticated) {
//     return null
//   }

  return (
    <div className="min-w-full">
      <header className="border-b">
        <div className="mx-auto max-w-7xl px-6 py-4">
          <MainNav />
        </div>
      </header>

      <main className="mx-auto max-w-7xl px-6 py-6">
        <Outlet />
      </main>
    </div>
  )
}
