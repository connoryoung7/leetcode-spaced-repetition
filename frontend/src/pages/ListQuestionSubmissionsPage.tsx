import { useEffect, useState } from "react";

import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table"
import { useNavigate, Link } from "@tanstack/react-router";
// import { useQuestionSubmissions } from "../hooks/api";
import { getQuestionSubmissionsV2 } from "../api";

const convertNumToDifficulty = (val: number) => {
    switch (val) {
        case 1:
            return "Easy"
        case 2:
            return "Medium"
        case 3:
            return "Difficult"
        default:
            return "N/A"
    }
}

const generateLinkForLeetcode = (slug: string): string => {
    return `https://leetcode.com/problems/${slug}/`
}

const ListQuestionSubmissionsPage = () => {
    const [selectedTags, setSelectedTags] = useState<Set<string>>(new Set())
    const [questions, setQuestions] = useState<any[]>([])

    const navigate = useNavigate({ from: '/questions' })

    useEffect(() => {
        (
            async () => {
                const data = await getQuestionSubmissionsV2([])
                console.log("data =", data)
                setQuestions(data?.data || []);
            }
        )()
    }, [])

    return (
        <div className="absolute inset-0 w-9/10 mx-auto">
            <div>Submissions Page</div>
            <div>
                <Table>
                    {
                        (selectedTags.size  === 0) ?
                        <TableCaption>Please select a topic.</TableCaption> :
                        null
                    }
                    <TableHeader>
                        <TableRow>
                            <TableHead className="w-[100px]">Question #</TableHead>
                            <TableHead className="w-[100px] text-left">Title</TableHead>
                            <TableHead className="text-left">Difficulty</TableHead>
                            <TableHead className="text-left">Open Submissions</TableHead>
                            <TableHead className="text-left">Open Link in LeetCode</TableHead>
                        </TableRow>
                    </TableHeader>
                    <TableBody>
                        {
                            questions.map(question => (
                                <TableRow>
                                    <TableCell className="text-left">{question.id}</TableCell>
                                    <TableCell className="font-medium text-left">{question.title}</TableCell>
                                    <TableCell className="text-left">{convertNumToDifficulty(question.difficulty)}</TableCell>
                                    <TableCell className="text-left">
                                        <Link to={`/questions/${question.id}`}>Open Submissions</Link>
                                    </TableCell>
                                    <TableCell className="text-left">
                                        <a href={generateLinkForLeetcode(question.slug)} target="_blank">Open</a>
                                    </TableCell>
                                </TableRow>
                            ))
                        }
                    </TableBody>
                </Table>
            </div>
        </div>
    );
}

export default ListQuestionSubmissionsPage;
