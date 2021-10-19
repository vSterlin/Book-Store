package book

const BookSchema = `
CREATE TABLE books (
	id SERIAL PRIMARY KEY,
	title VARCHAR NOT NULL,
	author VARCHAR NOT NULL,
	pages INTEGER,
	publish_date VARCHAR,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
	
INSERT INTO books (title, author, pages, publish_date)
VALUES ('Brave New World', 'Aldous Huxley', 311, '1932');
`
