INSERT INTO users (public_id, username, password_hash, created_at)
VALUES
    ('MhH0ZPc33vxVt-o9HE1FA', 'admin', '$2a$10$BsN1J0oHSVrUbFhimo7aGev7.hhRhmHVCUukZy9rQ5w2cSEW7KHri', '2025-01-14 17:23:43.094319');

INSERT INTO categories (public_id, name, description, parent_id, created_by)
VALUES
    ('hQW0H2q0nXXKDoXvatxwm', 'Academics', 'Discuss about school!', 'PARENT', 1),
    ('TOTaswlgbOj2MEYKHeckk', 'Hobbies', 'Discuss about hobbies!', 'PARENT', 1),
    ('V8a4HFZXGm-KUTvi71lPi', 'CCA', 'Discuss about CCAs!', 'PARENT', 1);