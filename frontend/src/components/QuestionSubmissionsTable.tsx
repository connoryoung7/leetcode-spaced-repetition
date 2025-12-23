import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table"
type QuestionSubmissionsTableProps = {
    submissions: any[]
};

export default function QuestionSubmissionsTable({ submissions }: QuestionSubmissionsTableProps) {
    return (
        <Table>
            <TableHeader>
                <TableRow>
                    <TableHead>Date</TableHead>
                    <TableHead>Confidence Level</TableHead>
                    <TableHead className="text-right">Time Taken</TableHead>
                </TableRow>
            </TableHeader>
            <TableBody>
                {
                    submissions.map(submission => (
                        <TableRow>
                            <TableCell>{submission.date}</TableCell>
                            <TableCell>{submission.confidenceLevel}</TableCell>
                            <TableCell>{submission.timeTaken}</TableCell>
                        </TableRow>
                    ))
                }
            </TableBody>
        </Table>
    )
}
