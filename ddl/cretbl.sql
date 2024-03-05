CREATE TABLE posts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    subject VARCHAR(255) NOT NULL,
    name VARCHAR(100),
    email VARCHAR(100),
    content TEXT NOT NULL,
    remote_host VARCHAR(100),
    user_agent VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
