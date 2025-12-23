export const ConfidenceLevel = {
    VeryLow: 1,
    Low: 2,
    Medium: 3,
    High: 4,
} as const;
export type ConfidenceLevel = typeof ConfidenceLevel[keyof typeof ConfidenceLevel];
