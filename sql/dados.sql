INSERT INTO users (name, nick, email, password)
VALUES
('User 1', 'user_1', 'user1@gmail.com', '$2a$10$0xZIyTIAzKqLlEowZHhyF.eBXtbnhr.BH/JRcTS7kmyyze5b1LTFO'),
('User 2', 'user_2', 'user2@gmail.com', '$2a$10$0xZIyTIAzKqLlEowZHhyF.eBXtbnhr.BH/JRcTS7kmyyze5b1LTFO'),
('User 3', 'user_3', 'user3@gmail.com', '$2a$10$0xZIyTIAzKqLlEowZHhyF.eBXtbnhr.BH/JRcTS7kmyyze5b1LTFO');

INSERT INTO followers (user_id, follower_id)
values
(1, 2),
(3, 1),
(1, 3);