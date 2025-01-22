CREATE TABLE public."user" (
	id uuid NOT NULL,
	firstname varchar NULL,
	lastname varchar NULL,
	email varchar NULL,
	age int NULL,
	created timestamp NULL,
	CONSTRAINT user_pk PRIMARY KEY (id)
);
