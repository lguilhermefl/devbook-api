INSERT INTO users (name, nick, email, password)
values
("User 1", "user1", "user1@test.com", "$2a$10$hXt55E3Oub0ZbHHM0HXjUeaKbX5mVj6ugAYvMZjGPSLVnP6ctROVa"),
("User 2", "user2", "user2@test.com", "$2a$10$hXt55E3Oub0ZbHHM0HXjUeaKbX5mVj6ugAYvMZjGPSLVnP6ctROVa"),
("User 3", "user3", "user3@test.com", "$2a$10$hXt55E3Oub0ZbHHM0HXjUeaKbX5mVj6ugAYvMZjGPSLVnP6ctROVa");

INSERT INTO followers (user_id, follower_id)
values
(1, 2),
(3, 1),
(1, 3);

INSERT INTO posts(title, content, author_id)
values
("Post user 1", "Este é o post do user 1", 1),
("Post user 2", "Este é o post do user 2", 2),
("Post user 3", "Este é o post do user 3", 3);