--
-- PostgreSQL database dump
--

-- Dumped from database version 14.2
-- Dumped by pg_dump version 14.2

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

--
-- Name: findnumber(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.findnumber() RETURNS integer
    LANGUAGE plpgsql
    AS $$
declare
  _total int;
begin
    --calculate total before update
    select 
    a,sum(seta) INTO _total
    from dailydata WHERE date='2022-01-08' 
    GROUP BY a;

    return _total;
end;
$$;


ALTER FUNCTION public.findnumber() OWNER TO postgres;

--
-- Name: lock_all(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.lock_all() RETURNS integer
    LANGUAGE plpgsql
    AS $$
declare
  _total int;
begin
    --calculate total before update
    SELECT sum(stock_available)
      INTO _total
      FROM stocks;

    UPDATE stocks
       SET stock_locked = stock_locked + stock_available,
           stock_available = 0;

    return _total;
end;
$$;


ALTER FUNCTION public.lock_all() OWNER TO postgres;

--
-- Name: somefuncname(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.somefuncname() RETURNS integer
    LANGUAGE plpgsql
    AS $$
DECLARE
  one int;
  two int;
BEGIN
  one := 1;
  two := 2;
  RETURN one , two;
END
$$;


ALTER FUNCTION public.somefuncname() OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: books; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.books (
    id integer NOT NULL,
    title character varying(100) NOT NULL,
    primary_author character varying(100)
);


ALTER TABLE public.books OWNER TO postgres;

--
-- Name: books_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.books_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.books_id_seq OWNER TO postgres;

--
-- Name: books_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.books_id_seq OWNED BY public.books.id;


--
-- Name: dailyapproval; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.dailyapproval (
    id integer NOT NULL,
    date date NOT NULL,
    showid character varying(100),
    appliedby character varying(100),
    approvedby character varying(100),
    status character varying(100),
    distributorid character varying(100)
);


ALTER TABLE public.dailyapproval OWNER TO postgres;

--
-- Name: dailyapproval_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.dailyapproval_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.dailyapproval_id_seq OWNER TO postgres;

--
-- Name: dailyapproval_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.dailyapproval_id_seq OWNED BY public.dailyapproval.id;


--
-- Name: dailydata; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.dailydata (
    id integer NOT NULL,
    date date DEFAULT CURRENT_DATE NOT NULL,
    showid character varying NOT NULL,
    distributorid character varying NOT NULL,
    a integer,
    seta integer,
    b integer,
    setb integer,
    c integer,
    setc integer,
    ab integer,
    setab integer,
    ac integer,
    setac integer,
    bc integer,
    setbc integer,
    abcfull integer,
    setabcfull integer,
    abchalf integer,
    setabchalf integer,
    boxfull integer,
    setboxfull integer,
    boxhalf integer,
    setboxhalf integer,
    status character varying
);


ALTER TABLE public.dailydata OWNER TO postgres;

--
-- Name: dailydata_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.dailydata_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.dailydata_id_seq OWNER TO postgres;

--
-- Name: dailydata_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.dailydata_id_seq OWNED BY public.dailydata.id;


--
-- Name: distributor; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.distributor (
    id integer NOT NULL,
    firsrname text NOT NULL,
    lastname text,
    address character(100),
    city character(50),
    email character(50),
    distributorid character(50),
    isactive boolean,
    mobile character varying(50),
    joined date DEFAULT CURRENT_DATE NOT NULL
);


ALTER TABLE public.distributor OWNER TO postgres;

--
-- Name: distributor_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.distributor_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.distributor_id_seq OWNER TO postgres;

--
-- Name: distributor_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.distributor_id_seq OWNED BY public.distributor.id;


--
-- Name: roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.roles (
    roalid integer NOT NULL,
    roalval character varying(255),
    roaldisplayname character varying(255),
    isactive boolean
);


ALTER TABLE public.roles OWNER TO postgres;

--
-- Name: roles_roalid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.roles_roalid_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.roles_roalid_seq OWNER TO postgres;

--
-- Name: roles_roalid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.roles_roalid_seq OWNED BY public.roles.roalid;


--
-- Name: sessionmaster; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sessionmaster (
    sessionid integer NOT NULL,
    presenttime character varying(255) NOT NULL,
    expiretime character varying(255) NOT NULL,
    email character varying
);


ALTER TABLE public.sessionmaster OWNER TO postgres;

--
-- Name: shoppingcart; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.shoppingcart (
    id integer NOT NULL,
    username character varying NOT NULL,
    showtype character varying NOT NULL,
    showid integer,
    price integer,
    showhours character varying,
    timeshows character varying,
    ticketnumber integer,
    expired boolean,
    ordertoken character varying,
    paymentstatus character varying
);


ALTER TABLE public.shoppingcart OWNER TO postgres;

--
-- Name: shoppingcart_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.shoppingcart_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.shoppingcart_id_seq OWNER TO postgres;

--
-- Name: shoppingcart_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.shoppingcart_id_seq OWNED BY public.shoppingcart.id;


--
-- Name: shows; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.shows (
    id integer NOT NULL,
    "time" text NOT NULL,
    isactive boolean NOT NULL,
    isspecialevent boolean,
    bannerimg character varying,
    date timestamp without time zone
);


ALTER TABLE public.shows OWNER TO postgres;

--
-- Name: shows_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.shows_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.shows_id_seq OWNER TO postgres;

--
-- Name: shows_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.shows_id_seq OWNED BY public.shows.id;


--
-- Name: transactions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transactions (
    username character varying,
    totalamount integer,
    responsedata character varying,
    showids character varying,
    token character varying,
    timeofdate character varying,
    id integer NOT NULL,
    paymentstatus character varying
);


ALTER TABLE public.transactions OWNER TO postgres;

--
-- Name: transactions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.transactions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.transactions_id_seq OWNER TO postgres;

--
-- Name: transactions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.transactions_id_seq OWNED BY public.transactions.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    lastname character varying(255),
    firstname character varying(255),
    username character varying(255),
    gender character varying(1),
    roalid integer,
    isactive boolean,
    password character varying(255),
    mailtoken integer
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: books id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.books ALTER COLUMN id SET DEFAULT nextval('public.books_id_seq'::regclass);


--
-- Name: dailyapproval id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dailyapproval ALTER COLUMN id SET DEFAULT nextval('public.dailyapproval_id_seq'::regclass);


--
-- Name: dailydata id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dailydata ALTER COLUMN id SET DEFAULT nextval('public.dailydata_id_seq'::regclass);


--
-- Name: distributor id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.distributor ALTER COLUMN id SET DEFAULT nextval('public.distributor_id_seq'::regclass);


--
-- Name: roles roalid; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles ALTER COLUMN roalid SET DEFAULT nextval('public.roles_roalid_seq'::regclass);


--
-- Name: shoppingcart id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shoppingcart ALTER COLUMN id SET DEFAULT nextval('public.shoppingcart_id_seq'::regclass);


--
-- Name: shows id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shows ALTER COLUMN id SET DEFAULT nextval('public.shows_id_seq'::regclass);


--
-- Name: transactions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions ALTER COLUMN id SET DEFAULT nextval('public.transactions_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: books; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.books (id, title, primary_author) FROM stdin;
\.


--
-- Data for Name: dailyapproval; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.dailyapproval (id, date, showid, appliedby, approvedby, status, distributorid) FROM stdin;
2	2021-12-30	12	dataentry@gmail.com	Thomas@gmail.com	Approved	DC-01222                                          
1	2021-12-28	12	dataentry@gmail.com	qc@gmail.com	Approved	DC-0133                                           
3	2021-12-30	12	dataentry@gmail.com	qc@gmail.com	Approved	DC-0133                                           
4	2021-12-01	12	dataentry@gmail.com		User-Submitted	DC-0133                                           
5	2021-12-01	12	dataentry@gmail.com		User-Submitted	DC-0133                                           
6	2022-01-01	12	dataentry@gmail.com	qc@gmail.com	Approved	DC-0133                                           
7	2022-01-01	11	dataentry@gmail.com	qc@gmail.com	Approved	DC-01222                                          
11	2022-01-08	12	dataentry@gmail.com	qc@gmail.com	Approved	DC-01222                                          
10	2022-01-08	12	dataentry@gmail.com	qc@gmail.com	Approved	SSSSS                                             
9	2022-01-08	12	dataentry@gmail.com	qc@gmail.com	Approved	DC-0133                                           
8	2022-01-08	12	dataentry@gmail.com	qc@gmail.com	Approved	DC-0133                                           
12	2022-01-09	12	dataentry@gmail.com	qc@gmail.com	Approved	DC-0133                                           
\.


--
-- Data for Name: dailydata; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.dailydata (id, date, showid, distributorid, a, seta, b, setb, c, setc, ab, setab, ac, setac, bc, setbc, abcfull, setabcfull, abchalf, setabchalf, boxfull, setboxfull, boxhalf, setboxhalf, status) FROM stdin;
5	2022-01-08	12	DC-01222                                          	9	300	3	33	5	55	0	0	0	0	0	0	0	0	0	0	0	0	0	0	User-Submitted
3	2022-01-08	12	SSSSS                                             	2	40	8	2	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	User-Submitted
4	2022-01-08	12	SSSSS                                             	5	60	9	22	3	30	0	0	0	0	0	0	1	1	1	1	1	1	1	1	User-Submitted
1	2022-01-08	12	DC-0133                                           	1	20	1	20	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	User-Submitted
2	2022-01-08	12	DC-0133                                           	2	10	9	20	7	10	10	22	12	2	0	0	0	0	0	0	0	0	0	0	User-Submitted
6	2022-01-09	12	DC-0133                                           	1	10	1	1	1	2	11	1	0	0	0	0	111	1	111	1	111	22	333	2	User-Submitted
7	2022-01-09	12	DC-0133                                           	2	20	2	1	2	2	22	1	0	0	0	0	222	1	222	2	333	333	444	4	User-Submitted
8	2022-01-09	12	DC-0133                                           	3	30	3	1	3	2	33	1	0	0	0	0	333	1	333	3	0	0	0	0	User-Submitted
9	2022-01-09	12	DC-0133                                           	4	40	4	1	4	2	44	1	0	0	0	0	0	0	0	0	0	0	0	0	User-Submitted
10	2022-01-09	12	DC-0133                                           	5	50	5	1	5	2	55	1	0	0	0	0	0	0	0	0	0	0	0	0	User-Submitted
11	2022-01-09	12	DC-0133                                           	6	60	6	1	6	2	66	1	0	0	0	0	0	0	0	0	0	0	0	0	User-Submitted
12	2022-01-09	12	DC-0133                                           	7	70	7	1	7	2	77	1	0	0	0	0	0	0	0	0	0	0	0	0	User-Submitted
13	2022-01-09	12	DC-0133                                           	8	80	8	1	8	2	88	1	0	0	0	0	0	0	0	0	0	0	0	0	User-Submitted
14	2022-01-09	12	DC-0133                                           	9	90	9	1	9	2	99	1	0	0	0	0	0	0	0	0	0	0	0	0	User-Submitted
\.


--
-- Data for Name: distributor; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.distributor (id, firsrname, lastname, address, city, email, distributorid, isactive, mobile, joined) FROM stdin;
19	Thomas	Albert	102, 2nd street , kk nagar                                                                          	Chennai                                           	thomas33@gmail.com                                	DC-0133                                           	t	9677221122	2021-10-01
20	zdc	asd	sss                                                                                                 	sss                                               	asd@sdfsdf.asd                                    	SSSSS                                             	t	9566812123	2022-01-08
18	ad	asd	ewr                                                                                                 	adf                                               	asd                                               	DC-01222                                          	t	9766464646	2021-12-02
\.


--
-- Data for Name: roles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.roles (roalid, roalval, roaldisplayname, isactive) FROM stdin;
1	Data-Entry-Operator	Data-Entry-Operator	t
2	Quality-Checker	Quality-Checker	t
3	Super-Admin	Super-Admin	t
4	user	clientuser	t
\.


--
-- Data for Name: sessionmaster; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sessionmaster (sessionid, presenttime, expiretime, email) FROM stdin;
177918506	2022-04-06 01:26:29	2022-04-06 01:35:43 	good3@gmail.com
\.


--
-- Data for Name: shoppingcart; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.shoppingcart (id, username, showtype, showid, price, showhours, timeshows, ticketnumber, expired, ordertoken, paymentstatus) FROM stdin;
100	good3@gmail.com	SingleDigit	124	25	10:00	06-Apr-2022	1	f	CKJKVB6V	created
101	good3@gmail.com	SingleDigit	124	25	10:00	06-Apr-2022	2	f	CKJKVB6V	created
102	good3@gmail.com	SingleDigit	124	25	10:00	06-Apr-2022	2	f	WL5DQ5Q8	created
103	good3@gmail.com	TripleDigit	124	100	10:00	06-Apr-2022	211	f	WL5DQ5Q8	created
\.


--
-- Data for Name: shows; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.shows (id, "time", isactive, isspecialevent, bannerimg, date) FROM stdin;
118	19:00	t	f	Nothing	2022-04-05 19:00:00
119	20:00	t	f	Nothing	2022-04-05 20:00:00
120	21:00	t	f	Nothing	2022-04-05 21:00:00
121	22:00	t	f	Nothing	2022-04-05 22:00:00
122	23:00	t	f	Nothing	2022-04-05 23:00:00
123	24:00	t	f	Nothing	2022-04-06 00:00:00
124	10:00	t	f	Nothing	2022-04-06 10:00:00
125	11:00	t	f	Nothing	2022-04-06 11:00:00
126	12:00	t	f	Nothing	2022-04-06 12:00:00
127	13:00	t	f	Nothing	2022-04-06 13:00:00
128	14:00	t	f	Nothing	2022-04-06 14:00:00
\.


--
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.transactions (username, totalamount, responsedata, showids, token, timeofdate, id, paymentstatus) FROM stdin;
good3@gmail.com	50	map[data:map[addresses:map[bitcoin:3DENv9CxxYtbymRhtoLLosfbm82k8Cc6zt bitcoincash:qp0uvdkxutr762xfe0p56zzgafaur3tessmng5t7t4 dai:0x9bf77c025e3144c8fb478f5de10a868d02ce6bb3 dogecoin:DB5TnDnWBoeYCyNBo8udfurZsjkx2Hu9kJ ethereum:0x9bf77c025e3144c8fb478f5de10a868d02ce6bb3 litecoin:MLmZPdEZnGPQK9Ak7Sr3yPmyt3qGWQFwm7 usdc:0x9bf77c025e3144c8fb478f5de10a868d02ce6bb3] cancel_url:http://localhost:8080/#/ code:CKJKVB6V created_at:2022-04-05T19:29:27Z description:Place Your Orders exchange_rates:map[BCH-USD:372.115 BTC-USD:46030.455 DAI-USD:0.9998085 DOGE-USD:0.16535 ETH-USD:3465.365 LTC-USD:124.505 USDC-USD:1.0] expires_at:2022-04-05T20:29:27Z fees_settled:true hosted_url:https://commerce.coinbase.com/charges/CKJKVB6V id:c5d19061-a54c-40cd-b80f-d32408892674 local_exchange_rates:map[BCH-USD:372.115 BTC-USD:46030.455 DAI-USD:0.9998085 DOGE-USD:0.16535 ETH-USD:3465.365 LTC-USD:124.505 USDC-USD:1.0] metadata:map[customer_id: customer_name:] name:Buy Ticket offchain_eligible:false payment_threshold:map[overpayment_absolute_threshold:map[amount:5.00 currency:USD] overpayment_relative_threshold:0.005 underpayment_absolute_threshold:map[amount:5.00 currency:USD] underpayment_relative_threshold:0.005] payments:[] pricing:map[bitcoin:map[amount:0.00108624 currency:BTC] bitcoincash:map[amount:0.13436706 currency:BCH] dai:map[amount:50.009576850000000000 currency:DAI] dogecoin:map[amount:302.38887210 currency:DOGE] ethereum:map[amount:0.014428000 currency:ETH] litecoin:map[amount:0.40159030 currency:LTC] local:map[amount:50.00 currency:USD] usdc:map[amount:50.000000 currency:USDC]] pricing_type:fixed_price pwcb_only:false redirect_url:http://localhost:8080/#/ resource:charge support_email:thomas550i@gmail.com timeline:[map[status:NEW time:2022-04-05T19:29:27Z]] utxo:false]]	100,101	CKJKVB6V	2022-04-06 00:59:28	6	created
good3@gmail.com	125	map[data:map[addresses:map[bitcoin:3MExLaGfux2quRmC7qLtYqBaU635SjQcqq bitcoincash:qr4ly9prupeh7u00t5gxwfa5nvukkm7jz568x5qt7x dai:0xf93aa91ba8412dbe9eb522ecc4c7813c7c37b767 dogecoin:D79gxufAYra6M4iYuzkn5qRK9kaHkrouEE ethereum:0xf93aa91ba8412dbe9eb522ecc4c7813c7c37b767 litecoin:MDpDryh3kukZuRLyaANLMamqWQGUnHNbkL usdc:0xf93aa91ba8412dbe9eb522ecc4c7813c7c37b767] cancel_url:http://localhost:8080/#/ code:WL5DQ5Q8 created_at:2022-04-05T19:37:19Z description:Place Your Orders exchange_rates:map[BCH-USD:370.915 BTC-USD:45915.2 DAI-USD:0.999812 DOGE-USD:0.16485 ETH-USD:3453.725 LTC-USD:124.035 USDC-USD:1.0] expires_at:2022-04-05T20:37:19Z fees_settled:true hosted_url:https://commerce.coinbase.com/charges/WL5DQ5Q8 id:512962f0-90ef-444c-9132-154dd80bcfcd local_exchange_rates:map[BCH-USD:370.915 BTC-USD:45915.2 DAI-USD:0.999812 DOGE-USD:0.16485 ETH-USD:3453.725 LTC-USD:124.035 USDC-USD:1.0] metadata:map[customer_id: customer_name:] name:Buy Ticket offchain_eligible:false payment_threshold:map[overpayment_absolute_threshold:map[amount:5.00 currency:USD] overpayment_relative_threshold:0.005 underpayment_absolute_threshold:map[amount:5.00 currency:USD] underpayment_relative_threshold:0.005] payments:[] pricing:map[bitcoin:map[amount:0.00272241 currency:BTC] bitcoincash:map[amount:0.33700443 currency:BCH] dai:map[amount:125.023504375000000000 currency:DAI] dogecoin:map[amount:758.26508950 currency:DOGE] ethereum:map[amount:0.036193000 currency:ETH] litecoin:map[amount:1.00778006 currency:LTC] local:map[amount:125.00 currency:USD] usdc:map[amount:125.000000 currency:USDC]] pricing_type:fixed_price pwcb_only:false redirect_url:http://localhost:8080/#/ resource:charge support_email:thomas550i@gmail.com timeline:[map[status:NEW time:2022-04-05T19:37:19Z]] utxo:false]]	102,103	WL5DQ5Q8	2022-04-06 01:07:19	7	created
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, lastname, firstname, username, gender, roalid, isactive, password, mailtoken) FROM stdin;
6	Albert	Admin	admin@gmail.com	M	3	t	UGFzc3dvcmQuMQ==	\N
11	qctest	qc	qc@gmail.com	M	2	t	UGFzc3dvcmQuMQ==	\N
12	entry	data	dataentry@gmail.com	M	1	t	UGFzc3dvcmQuMQ==	\N
13	raj	kumari	kumar@gmail.com	F	1	t	UGFzc3dvcmQuMQ==	\N
19	user	kumar	Adminds@gmmmm.hjn	M	2	t	MjM0NDIzNA==	\N
20	good	good	good@gmail.com	M	4	f	Z29vZA==	\N
21		goodHORSE	good3@gmail.com	M	4	t	Z29vZA==	5414643
\.


--
-- Name: books_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.books_id_seq', 1, false);


--
-- Name: dailyapproval_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.dailyapproval_id_seq', 12, true);


--
-- Name: dailydata_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.dailydata_id_seq', 14, true);


--
-- Name: distributor_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.distributor_id_seq', 20, true);


--
-- Name: roles_roalid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.roles_roalid_seq', 3, true);


--
-- Name: shoppingcart_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.shoppingcart_id_seq', 103, true);


--
-- Name: shows_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.shows_id_seq', 128, true);


--
-- Name: transactions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.transactions_id_seq', 7, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 21, true);


--
-- Name: books books_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.books
    ADD CONSTRAINT books_pkey PRIMARY KEY (id);


--
-- Name: dailyapproval dailyapproval_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dailyapproval
    ADD CONSTRAINT dailyapproval_pkey PRIMARY KEY (id);


--
-- Name: dailydata dailydata_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dailydata
    ADD CONSTRAINT dailydata_pkey PRIMARY KEY (id);


--
-- Name: distributor distributor_distributorid_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.distributor
    ADD CONSTRAINT distributor_distributorid_key UNIQUE (distributorid);


--
-- Name: distributor distributor_mobile_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.distributor
    ADD CONSTRAINT distributor_mobile_key UNIQUE (mobile);


--
-- Name: distributor distributor_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.distributor
    ADD CONSTRAINT distributor_pkey PRIMARY KEY (id);


--
-- Name: roles roles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (roalid);


--
-- Name: shoppingcart shoppingcart_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shoppingcart
    ADD CONSTRAINT shoppingcart_pkey PRIMARY KEY (id);


--
-- Name: shows shows_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shows
    ADD CONSTRAINT shows_pkey PRIMARY KEY (id);


--
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

