export function generateLeetcodeURL(questionId: number): string {
    return `https://lcid.cc/${questionId}`;
}

export function generateLinkForLeetcode(slug: string): string {
    return `https://leetcode.com/problems/${slug}/`
}
