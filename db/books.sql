PRAGMA foreign_keys = ON;

-- =========================
-- AUTHORS
-- =========================
CREATE TABLE authors (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);

-- =========================
-- EDITORIALS
-- =========================
CREATE TABLE editorials (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);

-- =========================
-- TOPICS
-- =========================
CREATE TABLE topics (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE
);

-- =========================
-- LANGUAGES
-- =========================
CREATE TABLE languages (
    code TEXT PRIMARY KEY,
    name TEXT NOT NULL
);

-- =========================
-- BOOKS
-- =========================
CREATE TABLE books (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    year INTEGER,
    language_code TEXT,
    isbn TEXT UNIQUE,
    edition TEXT,
    cover_url TEXT,
    pages INTEGER,
    location TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    editorial_id INTEGER,

    FOREIGN KEY (editorial_id)
        REFERENCES editorials(id)
        ON DELETE SET NULL,

    FOREIGN KEY (language_code)
        REFERENCES languages(code)
        ON DELETE SET NULL
);

-- =========================
-- BOOK_AUTHORS (N:M)
-- =========================
CREATE TABLE book_authors (
    book_id INTEGER NOT NULL,
    author_id INTEGER NOT NULL,

    PRIMARY KEY (book_id, author_id),

    FOREIGN KEY (book_id)
        REFERENCES books(id)
        ON DELETE CASCADE,

    FOREIGN KEY (author_id)
        REFERENCES authors(id)
        ON DELETE CASCADE
);

-- =========================
-- BOOK_TOPICS (N:M)
-- =========================
CREATE TABLE book_topics (
    book_id INTEGER NOT NULL,
    topic_id INTEGER NOT NULL,

    PRIMARY KEY (book_id, topic_id),

    FOREIGN KEY (book_id)
        REFERENCES books(id)
        ON DELETE CASCADE,

    FOREIGN KEY (topic_id)
        REFERENCES topics(id)
        ON DELETE CASCADE
);
