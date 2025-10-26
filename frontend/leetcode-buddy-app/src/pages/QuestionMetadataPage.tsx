import { type FC } from 'react'

import { useParams } from '@tanstack/react-router'
import { Spinner } from "../components/ui/spinner"

import { useQuestionSubmissions } from "../hooks/api"

import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table"

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
                data ?
                <Table>
                    <TableHeader>
                        <TableRow>
                            <TableHead>Date</TableHead>
                            <TableHead>Confidence</TableHead>
                            <TableHead className="text-right">Time Taken</TableHead>
                        </TableRow>
                    </TableHeader>
                    <TableBody>
                        {
                            data.data.map(submission => (
                                <TableRow>
                                    <TableCell>{submission.date}</TableCell>
                                    <TableCell>{submission.confidenceLevel}</TableCell>
                                    <TableCell>{submission.timeTaken}</TableCell>
                                </TableRow>
                            ))
                        }
                    </TableBody>
                </Table>
                : null
            }
        </div>
    )
}

export default QuestionMetadataPage
