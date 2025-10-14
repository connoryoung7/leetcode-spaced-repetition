import { useEffect, useState } from "react";

import { getAllQuestionTags } from "../api";
import { Badge } from "../components/ui/badge";
import { cn } from "../lib/utils";

const baseBadgeClass = "cursor-pointer"

const QuestionsPage = () => {
    const [tags, setTags] = useState<string[]>([])
    const [selectedTags, setSelectedTags] = useState<Set<string>>(new Set())
    useEffect(() => {
        (async () => {
            const data = await getAllQuestionTags()
            console.log("data =", data)
            setTags(data.tags)
        })()
    }, [])

    console.log("selectedTags:", selectedTags);

    return (
        <div>
            <div>Questions Page</div>
            <div>
                {
                    tags.map(tag => (
                        <Badge
                            key={tag}
                            variant={(selectedTags.has(tag)) ? "secondary" : "outline"}
                            onClick={() => {
                                console.log("tag being clicked is:", tag)
                                if (selectedTags.has(tag)) {
                                    setSelectedTags(
                                        new Set([...selectedTags].filter(t => t !== tag))
                                    )
                                } else {
                                    console.log("adding tag:", tag)
                                    setSelectedTags(
                                        new Set([...selectedTags, tag])
                                    )
                                }
                            }}
                            className={selectedTags.has(tag) ? cn("bg-blue-500", "text-white", baseBadgeClass) : baseBadgeClass}>
                                {tag}
                        </Badge>
                    ))
                }
            </div>
        </div>
    );
}

export default QuestionsPage;
