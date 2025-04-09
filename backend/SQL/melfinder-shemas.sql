--
-- PostgreSQL database dump
--

-- Dumped from database version 17.4
-- Dumped by pg_dump version 17.4

-- Started on 2025-04-09 23:59:07

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
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
-- TOC entry 223 (class 1259 OID 16425)
-- Name: ingredients; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.ingredients (
    id integer NOT NULL,
    name character varying
);


ALTER TABLE public.ingredients OWNER TO postgres;

--
-- TOC entry 222 (class 1259 OID 16424)
-- Name: ingredients_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.ingredients_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.ingredients_id_seq OWNER TO postgres;

--
-- TOC entry 4947 (class 0 OID 0)
-- Dependencies: 222
-- Name: ingredients_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.ingredients_id_seq OWNED BY public.ingredients.id;


--
-- TOC entry 219 (class 1259 OID 16400)
-- Name: recipes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.recipes (
    id uuid NOT NULL,
    name character varying,
    "time" integer,
    difficulty integer
);


ALTER TABLE public.recipes OWNER TO postgres;

--
-- TOC entry 224 (class 1259 OID 16433)
-- Name: recipes_ingredients; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.recipes_ingredients (
    recipe_id uuid,
    ingredient_id integer,
    amount character varying,
    unit integer DEFAULT 0
);


ALTER TABLE public.recipes_ingredients OWNER TO postgres;

--
-- TOC entry 229 (class 1259 OID 16472)
-- Name: recipes_tags; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.recipes_tags (
    recipe_id uuid,
    tag_id integer
);


ALTER TABLE public.recipes_tags OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 16408)
-- Name: reviews; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.reviews (
    id integer NOT NULL,
    recipe_id uuid,
    user_id integer,
    review integer
);


ALTER TABLE public.reviews OWNER TO postgres;

--
-- TOC entry 220 (class 1259 OID 16407)
-- Name: reviews_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.reviews_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.reviews_id_seq OWNER TO postgres;

--
-- TOC entry 4948 (class 0 OID 0)
-- Dependencies: 220
-- Name: reviews_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.reviews_id_seq OWNED BY public.reviews.id;


--
-- TOC entry 228 (class 1259 OID 16459)
-- Name: tags; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tags (
    id integer NOT NULL,
    name character varying,
    type_id integer
);


ALTER TABLE public.tags OWNER TO postgres;

--
-- TOC entry 227 (class 1259 OID 16458)
-- Name: tags_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tags_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.tags_id_seq OWNER TO postgres;

--
-- TOC entry 4949 (class 0 OID 0)
-- Dependencies: 227
-- Name: tags_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tags_id_seq OWNED BY public.tags.id;


--
-- TOC entry 226 (class 1259 OID 16450)
-- Name: tags_types; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tags_types (
    id integer NOT NULL,
    name character varying
);


ALTER TABLE public.tags_types OWNER TO postgres;

--
-- TOC entry 225 (class 1259 OID 16449)
-- Name: tags_types_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tags_types_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.tags_types_id_seq OWNER TO postgres;

--
-- TOC entry 4950 (class 0 OID 0)
-- Dependencies: 225
-- Name: tags_types_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tags_types_id_seq OWNED BY public.tags_types.id;


--
-- TOC entry 218 (class 1259 OID 16390)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying,
    created_at timestamp without time zone,
    passwd character varying,
    email character varying,
    name character varying,
    surname character varying,
    phone integer,
    age integer,
    sex boolean,
    weight integer,
    height integer,
    bmi integer
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 16389)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- TOC entry 4951 (class 0 OID 0)
-- Dependencies: 217
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 4768 (class 2604 OID 16428)
-- Name: ingredients id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ingredients ALTER COLUMN id SET DEFAULT nextval('public.ingredients_id_seq'::regclass);


--
-- TOC entry 4767 (class 2604 OID 16411)
-- Name: reviews id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reviews ALTER COLUMN id SET DEFAULT nextval('public.reviews_id_seq'::regclass);


--
-- TOC entry 4771 (class 2604 OID 16462)
-- Name: tags id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tags ALTER COLUMN id SET DEFAULT nextval('public.tags_id_seq'::regclass);


--
-- TOC entry 4770 (class 2604 OID 16453)
-- Name: tags_types id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tags_types ALTER COLUMN id SET DEFAULT nextval('public.tags_types_id_seq'::regclass);


--
-- TOC entry 4766 (class 2604 OID 16393)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 4782 (class 2606 OID 16432)
-- Name: ingredients ingredients_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ingredients
    ADD CONSTRAINT ingredients_pkey PRIMARY KEY (id);


--
-- TOC entry 4777 (class 2606 OID 16406)
-- Name: recipes recipes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.recipes
    ADD CONSTRAINT recipes_pkey PRIMARY KEY (id);


--
-- TOC entry 4780 (class 2606 OID 16413)
-- Name: reviews reviews_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reviews
    ADD CONSTRAINT reviews_pkey PRIMARY KEY (id);


--
-- TOC entry 4788 (class 2606 OID 16466)
-- Name: tags tags_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_pkey PRIMARY KEY (id);


--
-- TOC entry 4785 (class 2606 OID 16457)
-- Name: tags_types tags_types_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tags_types
    ADD CONSTRAINT tags_types_pkey PRIMARY KEY (id);


--
-- TOC entry 4773 (class 2606 OID 16397)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 4775 (class 2606 OID 16399)
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- TOC entry 4783 (class 1259 OID 16486)
-- Name: idx_recipes_ingredients_recipe_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_recipes_ingredients_recipe_id ON public.recipes_ingredients USING btree (recipe_id);


--
-- TOC entry 4789 (class 1259 OID 16488)
-- Name: idx_recipes_tags_recipe_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_recipes_tags_recipe_id ON public.recipes_tags USING btree (recipe_id);


--
-- TOC entry 4778 (class 1259 OID 16485)
-- Name: idx_reviews_recipe_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_reviews_recipe_id ON public.reviews USING btree (recipe_id);


--
-- TOC entry 4786 (class 1259 OID 16487)
-- Name: idx_tags_type_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_tags_type_id ON public.tags USING btree (type_id);


--
-- TOC entry 4792 (class 2606 OID 16444)
-- Name: recipes_ingredients recipes_ingredients_ingredient_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.recipes_ingredients
    ADD CONSTRAINT recipes_ingredients_ingredient_id_fkey FOREIGN KEY (ingredient_id) REFERENCES public.ingredients(id);


--
-- TOC entry 4793 (class 2606 OID 16439)
-- Name: recipes_ingredients recipes_ingredients_recipe_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.recipes_ingredients
    ADD CONSTRAINT recipes_ingredients_recipe_id_fkey FOREIGN KEY (recipe_id) REFERENCES public.recipes(id);


--
-- TOC entry 4795 (class 2606 OID 16475)
-- Name: recipes_tags recipes_tags_recipe_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.recipes_tags
    ADD CONSTRAINT recipes_tags_recipe_id_fkey FOREIGN KEY (recipe_id) REFERENCES public.recipes(id);


--
-- TOC entry 4796 (class 2606 OID 16480)
-- Name: recipes_tags recipes_tags_tag_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.recipes_tags
    ADD CONSTRAINT recipes_tags_tag_id_fkey FOREIGN KEY (tag_id) REFERENCES public.tags(id);


--
-- TOC entry 4790 (class 2606 OID 16414)
-- Name: reviews reviews_recipe_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reviews
    ADD CONSTRAINT reviews_recipe_id_fkey FOREIGN KEY (recipe_id) REFERENCES public.recipes(id);


--
-- TOC entry 4791 (class 2606 OID 16419)
-- Name: reviews reviews_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reviews
    ADD CONSTRAINT reviews_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 4794 (class 2606 OID 16467)
-- Name: tags tags_type_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_type_id_fkey FOREIGN KEY (type_id) REFERENCES public.tags_types(id);


-- Completed on 2025-04-09 23:59:07

--
-- PostgreSQL database dump complete
--

