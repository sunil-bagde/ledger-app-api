INSERT INTO "public"."users" ("id", "first_name", "last_name", "username", "email", "password_hash", "date_created", "date_updated") VALUES
('145999b7-8361-4af7-8530-a59cdb3a1ab5', 'raj', 'raj', 'raj', 'raj@gmail.com', '$2a$10$hGUKpzZJnRuMrjGgyzUgk.5cjIvsvvL9rvF8rL4EiaNkQs97IugZS', '2024-05-12 11:47:11', '2024-05-12 11:47:11'),
('be6019be-0d0f-4d4a-a4e5-fb19777a7b11', 'ani', 'ani', 'ani', 'ani@gmail.com', '$2a$10$hGUKpzZJnRuMrjGgyzUgk.5cjIvsvvL9rvF8rL4EiaNkQs97IugZS', '2024-05-12 11:47:11', '2024-05-12 11:47:11');


INSERT INTO "public"."accounts" ("id", "user_id", "name", "available_amount", "type", "slug", "date_created", "date_updated") VALUES
('31b03c53-5193-4f71-90e6-274d7e3e94bd', '145999b7-8361-4af7-8530-a59cdb3a1ab5', 'rajac', 0.00, 'saving', 'rajac', '2024-05-12 12:51:57', '2024-05-12 18:53:16'),
('dfe3c883-7b6c-496c-8f09-024fe3f7fee7', 'be6019be-0d0f-4d4a-a4e5-fb19777a7b11', 'aniac', 0.00, 'saving', 'aniac', '2024-05-12 12:51:57', '2024-05-12 18:53:16');
