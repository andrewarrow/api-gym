--
-- PostgreSQL database dump
--

-- Dumped from database version 15.2
-- Dumped by pg_dump version 15.2

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
-- Name: buildings; Type: TABLE; Schema: public; Owner: carla
--

CREATE TABLE public.buildings (
    id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    name character varying(255) DEFAULT ''::character varying,
    guid character varying(255) DEFAULT ''::character varying,
    address character varying(255) DEFAULT ''::character varying,
    city character varying(255) DEFAULT ''::character varying,
    state character varying(255) DEFAULT ''::character varying,
    postal character varying(255) DEFAULT ''::character varying,
    country character varying(255) DEFAULT ''::character varying,
    url character varying(255) DEFAULT ''::character varying,
    units integer DEFAULT 0,
    about text DEFAULT ''::text,
    phone character varying(255) DEFAULT ''::character varying
);


ALTER TABLE public.buildings OWNER TO carla;

--
-- Name: buildings_id_seq; Type: SEQUENCE; Schema: public; Owner: carla
--

CREATE SEQUENCE public.buildings_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.buildings_id_seq OWNER TO carla;

--
-- Name: buildings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: carla
--

ALTER SEQUENCE public.buildings_id_seq OWNED BY public.buildings.id;


--
-- Name: buildings id; Type: DEFAULT; Schema: public; Owner: carla
--

ALTER TABLE ONLY public.buildings ALTER COLUMN id SET DEFAULT nextval('public.buildings_id_seq'::regclass);


--
-- Data for Name: buildings; Type: TABLE DATA; Schema: public; Owner: carla
--

COPY public.buildings (id, created_at, name, guid, address, city, state, postal, country, url, units, about, phone) FROM stdin;
97	2023-03-06 15:29:11.31317	Broadstone Midtown	d42782ce-041b-87d6-2348-7e55ff2414fe		Atlanta	GA		USA	https://broadstonemidtown.com/	90		
96	2023-03-06 15:29:11.312628	The Brady	3108f844-3799-d175-aaa1-8818de1d9165	930 Howell Mill Rd	Atlanta	GA			https://www.thebrady.com/	500		
95	2023-03-06 15:29:11.312056	Atlantic House	9970b6e6-4f5d-a414-75d3-0d78f315c1e2	1163 West Peachtree St NW	Atlanta	GA	30309		https://www.atlantichousemidtown.com/	198		
91	2023-03-06 15:29:11.213288	The Shoreline at Sol Mia	1eba18c4-ae2d-bf81-51c3-99d90f0a4789	5455 N Marginal Rd	Cleveland	OH			https://www.theshorelineapts.com/	600		
80	2023-03-06 15:29:11.018131	The Medici	79930325-4771-7738-a298-3e0fa4eabcc2	725 South Bixel Street	Los Angeles	CA	90017		https://www.themedici.com/	233	Medici is the premier resort apartment community gracing the dynamic new downtown Los Angeles skyline. The twelve unique floor plans assure you will find the right home to fit your lifestyle. The Medici is the best choice for upscale living in Los Angeles.	
79	2023-03-06 15:29:11.017406	The Lorenzo	db00421d-fcb7-ed4b-0792-cc6c5f14ac63	325 W. Adams Blvd	Los Angeles	CA			https://www.thelorenzo.com/	209		
77	2023-03-06 15:29:11.015638	The Orsini	816cefda-a574-1f95-2a04-8ca1a1b52aaa	550 North Figueroa Street	Los Angeles	CA			https://www.theorsini.com/	800		
76	2023-03-06 15:29:11.014675	The Vermont	39c411d0-019c-bcdb-71f6-cfe3c6e2ab59	3150 Wilshire Blvd Los Angeles	Los Angeles	CA	90010		https://www.thevermont.net/	109		(323) 986-3418
133	2023-03-06 15:29:11.813425	Aertson Midtown	8c52d128-1a83-e591-f122-2fffbcf13636	905 20th Avenue S	Nashville	TN	37203	USA	https://www.aertsonmidtown.com/	411	Midtown is more than just a place in Nashville. It's a community of creativity. It's a town on the rise. And, if you're lucky, it's home. The restaurants satisfy your appetite for trendy dining. The shops keep you and your closet looking good. And the area's entertainment brings everyone together. Get comfortable. You're going to want to stay a while.	+1 844-399-5863
134	2023-03-06 15:29:12.284271	Amaray Las Olas by Windsor	dd3ea019-0afb-9a3e-c46b-a37f17be6350	215 SE 8th Ave	Fort Lauderdale	FL	33301	USA	https://www.amaraylasolas.com/	288	Amaray Las Olas by Windsor is a skyward apartment community with the finest of Fort Lauderdale living. A place to slow down and take in the surroundings, people come here to celebrate a life well earned. We offer studios and one, two or three-bedroom luxury Fort Lauderdale apartments with 26 unique designs.	+1 954-462-8282
125	2023-03-06 15:29:11.713201	Eleven North	82e6d234-4d57-0bce-b5bc-86fcfe729f9b	210 11th Ave N	Nashville	TN	37203	USA	https://www.elevennorth.com/	871	Our one, two, and three-bedroom Nashville apartments feature a variety of upscale amenities ranging from granite countertops, low e-windows, and contemporary light fixtures to our oversized pool with an adjoining spa, and a community lounge with a full chef's kitchen. At Eleven North, we offer the full luxury living experience with all the amenities you're looking for in your new apartment home.	+1 615-805-5402
124	2023-03-06 15:29:11.712566	The Cadence	7d255ee7-c48d-c914-7b2b-e2eb0494e52f	1600 McGavock St	Nashville	TN	37203	USA	https://www.cadencenashville.com/	90	Welcome to Midtown, the neighborhood central to Music Row, Downtown, West End, & Hillsboro Village. Placemakr Music Row, nestled in the heart of Nashville, is smack-dab in the middle of a lively nightlife hub. Must-visit restaurants, can’t-miss music venues, gotta-see sights, and a whole lot more are all just steps from your front door. And with Vanderbilt University just a stone's throw away, this neighborhood attracts people from all walks of life, making it a dynamic place to stay.	+1-615-227-1600
132	2023-03-06 15:29:11.812901	Vertis Green Hills	f86c9914-d64a-2062-7a3d-adc8b2e8d976	4000 Hillsboro Pike	Nashville	TN	37215	USA	https://www.vertisgreenhills.com/	400	Vertis Green Hills is Nashville apartment living at its most inspired and indulgent. The 18-story tower rises over a gently rolling and heavily wooded landscape in the enviable Green Hills neighborhood. The richly detailed residences welcome you with singular style, grace and panoramic views – all complemented by a curated collection of exclusive recreational amenities and refined concierge services.	+1 615-819-5357
105	2023-03-06 15:29:11.413658	The Pacific	513acd20-0871-1989-d91a-82b3ba13bfdb	1380 Hornby Street	Vancouver	BC		Canada	https://www.grosvenorpacific.com/	930	For three centuries, Grosvenor has owned and managed over 300 acres of property in London’s desirable Belgravia and Mayfair districts. Today, we are one of the world’s largest privately owned businesses engaged in international property, food & agtech, rural estates and support for charitable initiatives.	+1 604.559.8888
100	2023-03-06 15:29:11.31494	Cyan on Peachtree	be94c1e5-6ae4-0fdd-f244-c87207a5a3fc	3380 Peachtree Rd NE	Atlanta	GA	30326	USA	https://cyanonpeachtree.com/	122	From the moment you step off Peachtree Street and enter our spectacular lobby of our Buckhead Atlanta apartments, you will be certain you've arrived. Elegantly appointed inside and out, Cyan provides everything you desire, from designer finishes to the ample space. Entertain in your expansive residence where floor-to-ceiling windows provide an impressive view of Buckhead.	+1 470-205-5120
\.


--
-- Name: buildings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: carla
--

SELECT pg_catalog.setval('public.buildings_id_seq', 134, true);


--
-- Name: buildings buildings_pkey; Type: CONSTRAINT; Schema: public; Owner: carla
--

ALTER TABLE ONLY public.buildings
    ADD CONSTRAINT buildings_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

