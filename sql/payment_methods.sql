/*
Dear Programmers,

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
*                                                 *
*	This file belongs to Kevin Veros Hamonangan   *
*	and	Fandi Fladimir Dachi and is a part of     *
*	our	last project as the student of Del        *
*	Institute of Technology, Sitoluama.           *
*	Please contact us via Instagram:              *
*	sleepingnext and fandi_dachi                  *
*	before copying this file.                     *
*	Thank you, buddy. 😊                          *
*                                                 *
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
*/

--
-- PostgreSQL database dump
--

-- Dumped from database version 11.7 (Ubuntu 11.7-2.pgdg18.04+1)
-- Dumped by pg_dump version 11.7 (Ubuntu 11.7-2.pgdg18.04+1)

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

SET default_with_oids = false;

--
-- Name: payment_methods; Type: TABLE; Schema: public; Owner: sleepingnext
--

CREATE TABLE public.payment_methods (
    id bigint NOT NULL,
    name character(255) NOT NULL,
    account_number character(255) NOT NULL,
    account_holder_name character(255) NOT NULL,
    status character(255) NOT NULL
);


ALTER TABLE public.payment_methods OWNER TO sleepingnext;

--
-- Name: payment_method_id_seq; Type: SEQUENCE; Schema: public; Owner: sleepingnext
--

CREATE SEQUENCE public.payment_method_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.payment_method_id_seq OWNER TO sleepingnext;

--
-- Name: payment_method_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sleepingnext
--

ALTER SEQUENCE public.payment_method_id_seq OWNED BY public.payment_methods.id;


--
-- Name: payment_methods id; Type: DEFAULT; Schema: public; Owner: sleepingnext
--

ALTER TABLE ONLY public.payment_methods ALTER COLUMN id SET DEFAULT nextval('public.payment_method_id_seq'::regclass);


--
-- Data for Name: payment_methods; Type: TABLE DATA; Schema: public; Owner: sleepingnext
--

COPY public.payment_methods (id, name, account_number, account_holder_name, status) FROM stdin;
\.


--
-- Name: payment_method_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sleepingnext
--

SELECT pg_catalog.setval('public.payment_method_id_seq', 1, false);


--
-- Name: payment_methods payment_method_pkey; Type: CONSTRAINT; Schema: public; Owner: sleepingnext
--

ALTER TABLE ONLY public.payment_methods
    ADD CONSTRAINT payment_method_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

