import './App.css'

import { createRootRoute, createRoute, createRouter, RouterProvider, Outlet } from '@tanstack/react-router'
import QuestionSubmissionPage from './pages/QuestionSubmissionPage'
import QuestionsPage from './pages/QuestionsPage'
import QuestionMetadataPage from './pages/QuestionMetadataPage'
import ListQuestionSubmissionsPage from './pages/ListQuestionSubmissionsPage'
import { Toaster } from 'sonner'

const rootRoute = createRootRoute()
const indexRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: '/',
  component: () => <QuestionSubmissionPage />,
})
const questionsRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: 'questions',
  component: () => <Outlet />
})
const submissionsRoute = createRoute({
  getParentRoute: () => questionsRoute,
  path: 'submissions',
  component: () => <ListQuestionSubmissionsPage />
})

const questionsListRoute = createRoute({
  getParentRoute: () => questionsRoute,
  path: '/',
  component: () => <QuestionsPage />,
})

const questionMetadataRoute = createRoute({
  getParentRoute: () => questionsRoute,
  path: '$questionId',
  component: () => <QuestionMetadataPage />
})

const routeTree = rootRoute.addChildren([
  indexRoute,
  questionsRoute.addChildren([
    questionsListRoute,
    questionMetadataRoute
  ]),
  submissionsRoute,
])
const router = createRouter({ routeTree })

export default function App() {
  return (
    <>
      <Toaster position="top-right" richColors />
      <RouterProvider router={router} />
    </>
  )
}
