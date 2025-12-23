import { type FC } from 'react'

import { useParams } from '@tanstack/react-router'
import { Spinner } from "../components/ui/spinner"

import { useQuestionSubmissions } from "../hooks/api"
import QuestionSubmissionsTable from '../components/QuestionSubmissionsTable'

const QuestionMetadataPage: FC = () => {
    const questionId = useParams({
        from: '/questions/$questionId',
        select: (params) => params.questionId
    });

    const { data, isLoading, error } = useQuestionSubmissions(questionId)

    return (
        <div>
            {
                isLoading ? <Spinner /> : null
            }
            {
                data ? <QuestionSubmissionsTable submissions={data.data} /> : null
            }
        </div>
    )
}

export default QuestionMetadataPage
