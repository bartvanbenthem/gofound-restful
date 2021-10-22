--
-- PostgreSQL database dump
--

-- Dumped from database version 12.3
-- Dumped by pg_dump version 12.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    id integer NOT NULL,
    email character varying,
    password character varying,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;

--
-- Name: categories; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.categories (
    id integer NOT NULL,
    category_name character varying,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


--
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.categories_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.categories_id_seq OWNED BY public.categories.id;


--
-- Name: software; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.software (
    id integer NOT NULL,
    name character varying,
    description text,
    year integer,
    release_date date,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


--
-- Name: software_categories; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.software_categories (
    id integer NOT NULL,
    software_id integer,
    category_id integer,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


--
-- Name: software_categories_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.software_categories_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: software_categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.software_categories_id_seq OWNED BY public.software_categories.id;


--
-- Name: software_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.software_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: software_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.software_id_seq OWNED BY public.software.id;

--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);

--
-- Name: categories id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.categories ALTER COLUMN id SET DEFAULT nextval('public.categories_id_seq'::regclass);


--
-- Name: software id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.software ALTER COLUMN id SET DEFAULT nextval('public.software_id_seq'::regclass);


--
-- Name: software_categories id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.software_categories ALTER COLUMN id SET DEFAULT nextval('public.software_categories_id_seq'::regclass);



--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.users (id, email, password, created_at, updated_at) FROM stdin Delimiter ',';
1,admin@domain.com,$2a$12$YZmO3zxVXaKGXORRDxMleOD8COPtz85eSfuxB3ulSwfZmQ6uNzmE2,2021-01-01 00:00:00,2021-01-01 00:00:00
\.


--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.categories (id, category_name, created_at, updated_at) FROM stdin Delimiter ',';
1,Financial,2021-05-17 00:00:00,2021-05-17 00:00:00
2,Web,2021-05-17 00:00:00,2021-05-17 00:00:00
3,Tooling,2021-05-17 00:00:00,2021-05-17 00:00:00
4,Analytical,2021-05-17 00:00:00,2021-05-17 00:00:00
5,Desktop,2021-05-17 00:00:00,2021-05-17 00:00:00
6,OpenSource,2021-05-17 00:00:00,2021-05-17 00:00:00
7,CloudNative,2021-05-17 00:00:00,2021-05-17 00:00:00
\.


--
-- Data for Name: software; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.software (id, name, description, year, release_date, created_at, updated_at) FROM stdin Delimiter ',';
1,VSCode,development environment tooling,2010,2010-01-01,2021-01-01 00:00:00,2021-01-01 00:00:00
2,Excel,Microsoft Office spreadsheets,1995,1995-01-01,2021-01-01 00:00:00,2021-01-01 00:00:00
3,Notepad,lightweight editor,1986,1986-01-01,2021-01-01 00:00:00,2021-01-01 00:00:00
4,FireFox,internet browser,2000,2000-01-01,2021-01-01 00:00:00,2021-01-01 00:00:00
\.

--
-- Data for Name: software_categories; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.software_categories (id, software_id, category_id, created_at, updated_at) FROM stdin Delimiter ',';
1,1,3,2021-01-01 00:00:00,2021-01-01 00:00:00
2,1,6,2021-01-01 00:00:00,2021-01-01 00:00:00
3,2,4,2021-01-01 00:00:00,2021-01-01 00:00:00
4,2,5,2021-01-01 00:00:00,2021-01-01 00:00:00
5,3,6,2021-01-01 00:00:00,2021-01-01 00:00:00
6,4,5,2021-01-01 00:00:00,2021-01-01 00:00:00
7,3,2,2021-01-01 00:00:00,2021-01-01 00:00:00
\.

--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.users_id_seq', 9, true);

--
-- Name: categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.categories_id_seq', 9, true);


--
-- Name: software_categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.software_categories_id_seq', 1, false);


--
-- Name: software_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.software_id_seq', 4, true);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);

--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- Name: software_categories software_categories_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.software_categories
    ADD CONSTRAINT software_categories_pkey PRIMARY KEY (id);


--
-- Name: software software_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.software
    ADD CONSTRAINT software_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

