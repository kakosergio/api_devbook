INSERT INTO users (name, nick, email, password)
VALUES
('User 1', 'user_1', 'user1@gmail.com', '$2a$10$0xZIyTIAzKqLlEowZHhyF.eBXtbnhr.BH/JRcTS7kmyyze5b1LTFO'),
('User 2', 'user_2', 'user2@gmail.com', '$2a$10$0xZIyTIAzKqLlEowZHhyF.eBXtbnhr.BH/JRcTS7kmyyze5b1LTFO'),
('User 3', 'user_3', 'user3@gmail.com', '$2a$10$0xZIyTIAzKqLlEowZHhyF.eBXtbnhr.BH/JRcTS7kmyyze5b1LTFO');

INSERT INTO followers (user_id, follower_id)
VALUES
(1, 2),
(3, 1),
(1, 3);

INSERT INTO publications (title, body, author_id)
VALUES
('Publicação divertida 1', 'Olha que publicação mais divertida 1', 2),
('Publicação divertida 2', 'Olha que publicação mais divertida 2', 3),
('Publicação divertida 3', 'Olha que publicação mais divertida 3', 1);