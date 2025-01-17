CREATE TABLE public."user" (
	id uuid NOT NULL,
	firstname varchar NULL,
	lastname varchar NULL,
	email varchar NULL,
	age int NULL,
	created timestamp NULL,
	CONSTRAINT user_pk PRIMARY KEY (id)
);


INSERT INTO public."user"
(id, firstname, lastname, email, age, created)
VALUES(gen_random_uuid(), 'test', 'testovich', 'test@example.com', 18, now());