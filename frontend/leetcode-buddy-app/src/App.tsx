import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

import { createRootRoute, createRoute, createRouter, RouterProvider } from '@tanstack/react-router'
import QuestionSubmissionPage from './pages/QuestionSubmissionPage'
import QuestionsPage from './pages/QuestionsPage'

const rootRoute = createRootRoute()
const indexRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: '/',
  component: () => <QuestionSubmissionPage />,
})
const questionsRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: '/questions',
  component: () => <QuestionsPage />
})

const routeTree = rootRoute.addChildren([indexRoute, questionsRoute])
const router = createRouter({ routeTree })

export default function App() {
  return <RouterProvider router={router} />
}
