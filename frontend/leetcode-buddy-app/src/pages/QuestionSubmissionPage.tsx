import React from 'react';

import { z } from 'zod';
import { Controller, useForm } from 'react-hook-form';
import { zodResolver } from "@hookform/resolvers/zod"
import parse from 'parse-duration'

import { Input } from '../components/ui/input';
import { Slider } from '../components/ui/slider';
import { Field, FieldGroup, FieldLabel } from '../components/ui/field';
import { Button } from '../components/ui/button';

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
    confidenceLevel: z.number().min(ConfidenceLevel.VeryLow).max(ConfidenceLevel.VeryHigh),
    timeTaken: z.number().min(0)
})

const QuestionSubmissionPage: React.FC = () => {
    const form = useForm<z.infer<typeof formSchema>>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            questionId: "",
            confidenceLevel: ConfidenceLevel.Medium,
        }
    })

    const onSubmit = (data: z.infer<typeof formSchema>) => {
        console.log("data =", data)
    }

    console.log("formState =", form.formState.isValid)

    return (
        <div>
            <h1>Which question did you complete today?</h1>
            <div className="my-4">
                <form onSubmit={form.handleSubmit(onSubmit)} >
                    <Controller
                        name="questionId"
                        control={form.control}
                        render={({ field, fieldState }) => (
                            <Field data-invalid={fieldState.invalid}>
                                <FieldLabel htmlFor={field.name}>Question #</FieldLabel>
                                <Input
                                    {...field}
                                    aria-invalid={fieldState.invalid}
                                />
                            </Field>
                        )}
                    />
                    <FieldGroup className="grid grid-cols-2 my-4">
                        <Controller
                            name="confidenceLevel"
                            control={form.control}
                            render={({ field }) => (
                                <Field>
                                    <FieldLabel>ConfidenceLevel</FieldLabel>
                                    <Slider
                                        value={[field.value]}
                                        step={1}
                                        min={ConfidenceLevel.VeryLow}
                                        max={ConfidenceLevel.VeryHigh}
                                        onValueChange={val=> field.onChange(val[0])}
                                    />
                                    <p>{ConfidenceLevelMemes[field.value - 1].text}</p>
                                </Field>
                            )}
                        />
                        <Controller
                            name="timeTaken"
                            control={form.control}
                            render={({ field }) => (
                                <Field>
                                    <FieldLabel htmlFor="timeTaken">Time Taken</FieldLabel>
                                    <Input
                                        id="timeTaken"
                                        onBlur={e => {
                                            const value = parse(e.currentTarget.value)
                                            if (value !== null) {
                                                field.onChange(Math.floor(value / 1_000))
                                            } else {
                                                field.onChange(undefined)
                                            }                                            
                                        }}
                                    />
                                </Field>
                            )}
                        />
                    </FieldGroup>
                    <Button disabled={!form.formState.isValid} type="submit" variant="outline">Create Submission</Button>
                </form>
            </div>
        </div>
    )
}

export default QuestionSubmissionPage;
