import React, { useState } from 'react';
import { Input } from '../components/ui/input';
import { Slider } from '../components/ui/slider';
import { zodResolver } from "@hookform/resolvers/zod"
import { Form, useForm } from 'react-hook-form';
import { z } from 'zod';
import { FormField, FormItem, FormLabel } from '../components/ui/form';

enum ConfidenceLevel {
    VeryLow = 1,
    Low = 2,
    Medium = 3,
    High = 4,
    VeryHigh = 5
}

const ConfidenceLevelMemes = [
    {
        level: ConfidenceLevel.VeryLow,
        meme: "simpsons_repeat_stuff.gif",
        text: "I have no clue what's going on"
    },
    {
        level: ConfidenceLevel.Low,
        meme: "drake_explaining.gif",
        text: "I see how they did it, but I did not see that coming"
    },
    {
        level: ConfidenceLevel.Medium,
        meme: "chuck_norris.gif",
        text: "Neither confusion nor mastery"
    },
    {
        level: ConfidenceLevel.High,
        meme: "exploding_brain.gif",
        text: "Things are starting to click..."
    },
    {
        level: ConfidenceLevel.VeryHigh,
        meme: "great_gatsy_nod.gif",
        text: "You did it, buddy"
    }
]

const formSchema = z.object({
    questionId: z.string(),
    confidenceLevel: z.number().min(ConfidenceLevel.VeryLow).max(ConfidenceLevel.VeryHigh)
})

const QuestionSubmissionPage: React.FC = () => {
    const [questionId, setQuestionId] = useState<string | number>("")
    const form = useForm<z.infer<typeof formSchema>>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            questionId: "",
            confidenceLevel: ConfidenceLevel.Medium,
        }
    })

    return (
        <div>
            <h1>Which question did you complete today?</h1>
            <Input value={questionId} onChange={ e => setQuestionId(e.target.value)} />
            <h3>The value is: {questionId}</h3>
            <div>
                <Form {...form}>
                    <FormField
                        control={form.control}
                        name="questionId"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>
                                    <FormLabel>Question ID</FormLabel>
                                </FormLabel>
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="confidenceLevel"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Confidence Level</FormLabel>
                                <Slider defaultValue={[ConfidenceLevel.Medium]} step={1} min={ConfidenceLevel.VeryLow} max={ConfidenceLevel.VeryHigh} onChange={e => { field.value  }  />
                            </FormItem>
                        )}
                    />
                </Form>
            </div>
        </div>
    )
}

export default QuestionSubmissionPage;
