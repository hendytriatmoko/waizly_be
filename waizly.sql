--
-- PostgreSQL database dump
--

-- Dumped from database version 12.17 (Ubuntu 12.17-0ubuntu0.20.04.1)
-- Dumped by pg_dump version 12.17 (Ubuntu 12.17-0ubuntu0.20.04.1)

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
-- Name: employees; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.employees (
    employee_id integer NOT NULL,
    name character varying,
    job_title character varying,
    salary integer,
    department character varying,
    joined_date date
);


ALTER TABLE public.employees OWNER TO postgres;

--
-- Name: employees_employee_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.employees_employee_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.employees_employee_id_seq OWNER TO postgres;

--
-- Name: employees_employee_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.employees_employee_id_seq OWNED BY public.employees.employee_id;


--
-- Name: sales_data; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sales_data (
    sales_id integer NOT NULL,
    employee_id character varying,
    sales integer
);


ALTER TABLE public.sales_data OWNER TO postgres;

--
-- Name: sales_data_sales_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sales_data_sales_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sales_data_sales_id_seq OWNER TO postgres;

--
-- Name: sales_data_sales_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sales_data_sales_id_seq OWNED BY public.sales_data.sales_id;


--
-- Name: tasks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tasks (
    id_task character varying(36) NOT NULL,
    id_user character varying(36),
    nama character varying(20),
    deskripsi character varying,
    progress integer,
    status character varying,
    pin character varying,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.tasks OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id_user character varying(36) NOT NULL,
    nama character varying(20),
    email character varying(20),
    no_telp character varying(13),
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    token character varying,
    password character varying
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: employees employee_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employees ALTER COLUMN employee_id SET DEFAULT nextval('public.employees_employee_id_seq'::regclass);


--
-- Name: sales_data sales_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sales_data ALTER COLUMN sales_id SET DEFAULT nextval('public.sales_data_sales_id_seq'::regclass);


--
-- Data for Name: employees; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.employees VALUES (1, 'John Smith', 'Manager', 60000, 'Sales', '2022-01-15');
INSERT INTO public.employees VALUES (2, 'Jane Doe', 'Analyst', 45000, 'Marketing', '2022-02-01');
INSERT INTO public.employees VALUES (3, 'Mike Brown', 'Developer', 55000, 'IT', '2022-03-10');
INSERT INTO public.employees VALUES (4, 'Anna Lee', 'Manager', 65000, 'Sales', '2021-12-05');
INSERT INTO public.employees VALUES (5, 'Mark Wong', 'Developer', 50000, 'IT', '2023-05-20');
INSERT INTO public.employees VALUES (6, 'Emily Chen', 'Analyst', 48000, 'Marketing', '2023-06-02');


--
-- Data for Name: sales_data; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.sales_data VALUES (1, '1', 15000);
INSERT INTO public.sales_data VALUES (2, '2', 12000);
INSERT INTO public.sales_data VALUES (3, '3', 18000);
INSERT INTO public.sales_data VALUES (4, '1', 20000);
INSERT INTO public.sales_data VALUES (5, '4', 22000);
INSERT INTO public.sales_data VALUES (6, '5', 19000);
INSERT INTO public.sales_data VALUES (7, '6', 13000);
INSERT INTO public.sales_data VALUES (8, '2', 14000);


--
-- Data for Name: tasks; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.tasks VALUES ('2m7eienr-4z9b-dcdb-hj2a-7tuta14eot06', 's63pdn8h-ojt7-qmaj-klet-d423pe5ubq0q', 'coba baru', 'tar coba du lu sabar yaa', 90, 'progress', 'false', '2023-12-17 15:08:14.217046+07', '2023-12-17 15:47:43.264853+07', NULL);
INSERT INTO public.tasks VALUES ('1muoodyl-oghp-hry7-4qa9-79mldoftvap4', 's63pdn8h-ojt7-qmaj-klet-d423pe5ubq0q', 'nama ub', 'ini adalah test new task terkini nih', 90, 'progress', 'false', '2023-12-17 03:05:55.424124+07', '2023-12-17 15:28:03.358424+07', '2023-12-17 16:01:46.864305+07');


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users VALUES ('m5t7ty3l-dvte-84i2-5phn-s9kabwtmi9up', 'cobadaftar', 'coba@gmail.com', '', '2023-12-17 16:17:51.849859+07', NULL, NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJjbGllbnRfaWQiOiJtNXQ3dHkzbC1kdnRlLTg0aTItNXBobi1zOWthYnd0bWk5dXAiLCJjbGllbnRfcGxhdGZvcm0iOiJub25lIiwiY2xpZW50X3JvbGUiOiJ1c2VyIiwiY2xpZW50X3R5cGUiOiJub25lIiwiZXhwIjoxNzAyODg3NDcyfQ.2C5CzZGoqW6W_mc9101t9wlEDa-zuPbOImYPF88p3bY', 'aKu8jqR5hpAbQmZrUpGadgDpJSkSKEyLiN_4nEzysr8');
INSERT INTO public.users VALUES ('s63pdn8h-ojt7-qmaj-klet-d423pe5ubq0q', 'hendy', 'hendy@gmail.com', '082111239208', '2023-12-17 01:41:38.004376+07', NULL, NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJjbGllbnRfaWQiOiJzNjNwZG44aC1vanQ3LXFtYWota2xldC1kNDIzcGU1dWJxMHEiLCJjbGllbnRfcGxhdGZvcm0iOiJub25lIiwiY2xpZW50X3JvbGUiOiJ1c2VyIiwiY2xpZW50X3R5cGUiOiJub25lIiwiZXhwIjoxNzAyODkwOTU0fQ.5a5I6OYuF7lAPUc-DJUYpo6JUgssofjDTxN4J678du4', 'FBRfmylzQq6OEEClVnCbWgfP6VoYo9QzRnEu6MMFO6g');


--
-- Name: employees_employee_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.employees_employee_id_seq', 6, true);


--
-- Name: sales_data_sales_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sales_data_sales_id_seq', 8, true);


--
-- Name: employees employees_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employees
    ADD CONSTRAINT employees_pkey PRIMARY KEY (employee_id);


--
-- Name: sales_data sales_data_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sales_data
    ADD CONSTRAINT sales_data_pkey PRIMARY KEY (sales_id);


--
-- Name: tasks task_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT task_pkey PRIMARY KEY (id_task);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id_user);


--
-- PostgreSQL database dump complete
--

