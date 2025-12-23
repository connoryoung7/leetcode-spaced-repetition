export const ConfidenceLevel = {
    VeryLow: 1,
    Low: 2,
    Medium: 3,
    High: 4,
} as const;
export type ConfidenceLevel = typeof ConfidenceLevel[keyof typeof ConfidenceLevel];

export type QuestionSubmissionWithDetails = {
    id: number;
    questionId: number;
    confidenceLevel: ConfidenceLevel;
    timeTaken: number | null;
    question: {
        id: number;
        title: string;
        slug: string;
        difficulty: number;
        tags: string[];
    };
}

export function convertNumToDifficulty(val: number): string {
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
