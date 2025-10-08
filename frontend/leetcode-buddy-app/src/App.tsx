import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

import { createRootRoute, createRoute, createRouter, RouterProvider } from '@tanstack/react-router'
import QuestionSubmissionPage from './pages/QuestionSubmissionPage'

const rootRoute = createRootRoute()
const indexRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: '/',
  component: () => <QuestionSubmissionPage />,
})
const routeTree = rootRoute.addChildren([indexRoute])
const router = createRouter({ routeTree })

export default function App() {
  return <RouterProvider router={router} />
}
