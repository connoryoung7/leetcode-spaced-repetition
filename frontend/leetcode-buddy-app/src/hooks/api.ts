import { useEffect, useState } from "react"
import { getQuestionSubmissions, getAllQuestionTags } from "../api"

export function useQuestionSubmissions(questionId?: number) {
  const [data, setData] = useState<any>(null)
  const [error, setError] = useState<Error | null>(null)
  const [isLoading, setIsLoading] = useState(false)

  useEffect(() => {
    if (!questionId) return

    const fetchData = async () => {
      setIsLoading(true)
      setError(null)

      try {
        const result = await getQuestionSubmissions(questionId)
        setData(result)
      } catch (err) {
        setError(err instanceof Error ? err : new Error("Unknown error"))
      } finally {
        setIsLoading(false)
      }
    }

    fetchData()
  }, [questionId])

  return { data, error, isLoading }
}

export function useQuestionTags() {
  const [data, setData] = useState(null)
  const [error, setError] = useState<Error | null>(null)
  const [isLoading, setIsLoading] = useState(false)

  useEffect(() => {
    const fetchData = async () => {
      setIsLoading(true)
      setError(null)

      try {
        const result = await getAllQuestionTags()
        setData(result)
      } catch (err) {
        setError(err instanceof Error ? err : new Error("Unknown error"))
      } finally {
        setIsLoading(false)
      }
    }

    fetchData()
  }, [])

  return { data, error, isLoading }
}
