import axios from 'axios';

const instance = axios.create({
    baseURL: "http://localhost:8000",
    timeout: 5_000
})

export const getAllQuestionTags = async () => {
    const response = await instance.get("/questions/tags")
    return response.data
}

export const getQuestionByID = async (questionID: string) => {
    const response = await instance.get(`/questions/${questionID}`)
    return response.data
}

export const getAllQuestions = async (topics: string[]) => {
    const response = await instance.get("/questions", {
        params: {
            topics
        }
    })

    return response.data
}