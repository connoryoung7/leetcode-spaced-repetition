-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS questions (
    id INTEGER PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    difficulty INT NOT NULL CHECK (difficulty > 0 AND difficulty < 4),
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS questionTags (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    questionId INTEGER NOT NULL,
    tag VARCHAR(50) NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (questionId, tag),
    FOREIGN KEY (questionId) REFERENCES questions(id)
);

CREATE TABLE IF NOT EXISTS questionSubmissions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    userId UUID NOT NULL,
    questionId INT NOT NULL,
    submissionDate DATE NOT NULL,
    confidenceLevel INT NOT NULL CHECK (confidenceLevel > 0 AND confidenceLevel < 6),
    timeTaken INTERVAL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (userId, questionId, submissionDate),
    FOREIGN KEY (questionId) REFERENCES questions(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
