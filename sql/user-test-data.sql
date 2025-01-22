INSERT INTO public."user"
(id, firstname, lastname, email, age, created)
VALUES(gen_random_uuid(), 'test', 'testovich', 'test@example.com', 18, now());