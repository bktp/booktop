--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.6
-- Dumped by pg_dump version 9.6.6

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

--
-- Name: make_tsvector(character varying); Type: FUNCTION; Schema: public; Owner: buktp
--

CREATE FUNCTION make_tsvector(query character varying) RETURNS tsvector
    LANGUAGE plpgsql IMMUTABLE
    AS $$
begin
return (setweight(to_tsvector('russian', query), 'A'));
end
$$;


ALTER FUNCTION public.make_tsvector(query character varying) OWNER TO buktp;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: authors; Type: TABLE; Schema: public; Owner: buktp
--

CREATE TABLE authors (
    isbn character(13),
    name character varying(255) NOT NULL
);


ALTER TABLE authors OWNER TO buktp;

--
-- Name: books; Type: TABLE; Schema: public; Owner: buktp
--

CREATE TABLE books (
    isbn character(13) NOT NULL,
    name character varying(255) NOT NULL,
    original character varying(255) DEFAULT ''::character varying,
    published character varying(16) DEFAULT ''::character varying,
    description text DEFAULT ''::text,
    cover character varying(255) DEFAULT '/images/placeholder.png'::character varying,
    category_id integer DEFAULT 0
);


ALTER TABLE books OWNER TO buktp;

--
-- Name: categories; Type: TABLE; Schema: public; Owner: buktp
--

CREATE TABLE categories (
    id integer NOT NULL,
    name character varying(255) NOT NULL
);


ALTER TABLE categories OWNER TO buktp;

--
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: buktp
--

CREATE SEQUENCE categories_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE categories_id_seq OWNER TO buktp;

--
-- Name: categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: buktp
--

ALTER SEQUENCE categories_id_seq OWNED BY categories.id;


--
-- Name: favs; Type: TABLE; Schema: public; Owner: buktp
--

CREATE TABLE favs (
    user_token character varying(32),
    isbn character(13)
);


ALTER TABLE favs OWNER TO buktp;

--
-- Name: pages; Type: TABLE; Schema: public; Owner: buktp
--

CREATE TABLE pages (
    isbn character(13),
    pagenum integer NOT NULL,
    text text
);


ALTER TABLE pages OWNER TO buktp;

--
-- Name: user_roles; Type: TABLE; Schema: public; Owner: buktp
--

CREATE TABLE user_roles (
    token character varying(32),
    role character varying(16) DEFAULT 'user'::character varying
);


ALTER TABLE user_roles OWNER TO buktp;

--
-- Name: users; Type: TABLE; Schema: public; Owner: buktp
--

CREATE TABLE users (
    name character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    token character varying(32) NOT NULL
);


ALTER TABLE users OWNER TO buktp;

--
-- Name: categories id; Type: DEFAULT; Schema: public; Owner: buktp
--

ALTER TABLE ONLY categories ALTER COLUMN id SET DEFAULT nextval('categories_id_seq'::regclass);


--
-- Data for Name: authors; Type: TABLE DATA; Schema: public; Owner: buktp
--

COPY authors (isbn, name) FROM stdin;
9785389074354	Джоан Кэтлин Роулинг
9785353004325	Джоан Роулинг
9785389120426	Джоан Кэтлин Роулинг
9785389120426	 Джон Тиффани
9785389120426	 Джек Торн
9785389077812	Джоан Кэтлин Роулинг
\.


--
-- Data for Name: books; Type: TABLE DATA; Schema: public; Owner: buktp
--

COPY books (isbn, name, original, published, description, cover, category_id) FROM stdin;
9785389074354	Гарри Поттер и Философский камень	Harry Potter			/images/placeholder.png	0
9785353004325	Гарри Поттер и узник Азкабана	Harry Potter and the Prisoner of Azkaban	1999	Гарри Поттер и узник Азкабана	https://upload.wikimedia.org/wikipedia/ru/thumb/6/6a/%D0%93%D0%B0%D1%80%D1%80%D0%B8_%D0%9F%D0%BE%D1%82%D1%82%D0%B5%D1%80_3.jpg/300px-%D0%93%D0%B0%D1%80%D1%80%D0%B8_%D0%9F%D0%BE%D1%82%D1%82%D0%B5%D1%80_3.jpg	0
9785389120426	Гарри Поттер и проклятое дитя	Harry Potter			https://ozon-st.cdn.ngenix.net/multimedia/1016151978.jpg	1
9785389077812	Гарри Поттер и Тайная комната	Harry Potter			/images/placeholder.png	1
\.


--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: buktp
--

COPY categories (id, name) FROM stdin;
0	
1	Роман
4	Научпоп
\.


--
-- Name: categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: buktp
--

SELECT pg_catalog.setval('categories_id_seq', 4, true);


--
-- Data for Name: favs; Type: TABLE DATA; Schema: public; Owner: buktp
--

COPY favs (user_token, isbn) FROM stdin;
361d8ecddfd6fb0b97bb1b3fbd5fdf8b	9785389120426
147893c73a3f11be83ca3c0cfba4dff9	9785389120426
147893c73a3f11be83ca3c0cfba4dff9	9785389074354
147893c73a3f11be83ca3c0cfba4dff9	9785389077812
\.


--
-- Data for Name: pages; Type: TABLE DATA; Schema: public; Owner: buktp
--

COPY pages (isbn, pagenum, text) FROM stdin;
9785389120426	2	Вторая страница
9785389120426	3	Третья страница
9785389120426	1	Первая страница
\.


--
-- Data for Name: user_roles; Type: TABLE DATA; Schema: public; Owner: buktp
--

COPY user_roles (token, role) FROM stdin;
361d8ecddfd6fb0b97bb1b3fbd5fdf8b	user
147893c73a3f11be83ca3c0cfba4dff9	admin
134dc735e0320271308842784047388f	user
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: buktp
--

COPY users (name, password, token) FROM stdin;
root	root	147893c73a3f11be83ca3c0cfba4dff9
buktp	buktp	361d8ecddfd6fb0b97bb1b3fbd5fdf8b
ramil	ramil	134dc735e0320271308842784047388f
\.


--
-- Name: books books_pkey; Type: CONSTRAINT; Schema: public; Owner: buktp
--

ALTER TABLE ONLY books
    ADD CONSTRAINT books_pkey PRIMARY KEY (isbn);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: buktp
--

ALTER TABLE ONLY categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- Name: favs favs_unique; Type: CONSTRAINT; Schema: public; Owner: buktp
--

ALTER TABLE ONLY favs
    ADD CONSTRAINT favs_unique UNIQUE (user_token, isbn);


--
-- Name: pages pages_unique; Type: CONSTRAINT; Schema: public; Owner: buktp
--

ALTER TABLE ONLY pages
    ADD CONSTRAINT pages_unique UNIQUE (isbn, pagenum);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: buktp
--

ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (name);


--
-- Name: users users_token_key; Type: CONSTRAINT; Schema: public; Owner: buktp
--

ALTER TABLE ONLY users
    ADD CONSTRAINT users_token_key UNIQUE (token);


--
-- Name: authors_make_tsvector_idx; Type: INDEX; Schema: public; Owner: buktp
--

CREATE INDEX authors_make_tsvector_idx ON authors USING gin (make_tsvector(name));


--
-- Name: books_make_tsvector_idx; Type: INDEX; Schema: public; Owner: buktp
--

CREATE INDEX books_make_tsvector_idx ON books USING gin (make_tsvector(name));


--
-- Name: books_make_tsvector_idx1; Type: INDEX; Schema: public; Owner: buktp
--

CREATE INDEX books_make_tsvector_idx1 ON books USING gin (make_tsvector(original));


--
-- Name: authors authors_isbn_fkey; Type: FK CONSTRAINT; Schema: public; Owner: buktp
--

ALTER TABLE ONLY authors
    ADD CONSTRAINT authors_isbn_fkey FOREIGN KEY (isbn) REFERENCES books(isbn) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: books books_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: buktp
--

ALTER TABLE ONLY books
    ADD CONSTRAINT books_category_id_fkey FOREIGN KEY (category_id) REFERENCES categories(id) ON UPDATE CASCADE;


--
-- Name: favs favs_isbn_fkey; Type: FK CONSTRAINT; Schema: public; Owner: buktp
--

ALTER TABLE ONLY favs
    ADD CONSTRAINT favs_isbn_fkey FOREIGN KEY (isbn) REFERENCES books(isbn) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: favs favs_user_token_fkey; Type: FK CONSTRAINT; Schema: public; Owner: buktp
--

ALTER TABLE ONLY favs
    ADD CONSTRAINT favs_user_token_fkey FOREIGN KEY (user_token) REFERENCES users(token) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: pages pages_isbn_fkey; Type: FK CONSTRAINT; Schema: public; Owner: buktp
--

ALTER TABLE ONLY pages
    ADD CONSTRAINT pages_isbn_fkey FOREIGN KEY (isbn) REFERENCES books(isbn) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: user_roles user_roles_token_fkey; Type: FK CONSTRAINT; Schema: public; Owner: buktp
--

ALTER TABLE ONLY user_roles
    ADD CONSTRAINT user_roles_token_fkey FOREIGN KEY (token) REFERENCES users(token) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

