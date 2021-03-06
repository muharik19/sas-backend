PGDMP         1                y            SASDB    13.2    13.2     ?           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            ?           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            ?           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            ?           1262    16403    SASDB    DATABASE     g   CREATE DATABASE "SASDB" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'English_Indonesia.1252';
    DROP DATABASE "SASDB";
                postgres    false            ?            1259    16522    sas_audit_trail    TABLE     %  CREATE TABLE public.sas_audit_trail (
    id integer NOT NULL,
    modules character varying(200) NOT NULL,
    ref_id integer,
    actions character(1) NOT NULL,
    log_activity character varying(4000),
    created_by character varying(100),
    created_at timestamp(0) without time zone
);
 #   DROP TABLE public.sas_audit_trail;
       public         heap    postgres    false            ?            1259    16520    sas_audit_trail_id_seq    SEQUENCE     ?   ALTER TABLE public.sas_audit_trail ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.sas_audit_trail_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    209            ?            1259    16475    sas_item    TABLE     ?  CREATE TABLE public.sas_item (
    id integer NOT NULL,
    item_name character varying(135) NOT NULL,
    price double precision,
    weight_ml double precision,
    weight_mg double precision,
    created_by character varying(100),
    created_at timestamp(0) without time zone,
    modified_by character varying(100),
    modified_at timestamp(0) without time zone,
    is_deleted integer DEFAULT 0 NOT NULL
);
    DROP TABLE public.sas_item;
       public         heap    postgres    false            ?            1259    16473    sas_item_id_seq    SEQUENCE     ?   ALTER TABLE public.sas_item ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.sas_item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    201            ?            1259    16489 
   sas_module    TABLE     f   CREATE TABLE public.sas_module (
    id integer NOT NULL,
    name character varying(100) NOT NULL
);
    DROP TABLE public.sas_module;
       public         heap    postgres    false            ?            1259    16487    sas_module_id_seq    SEQUENCE     ?   ALTER TABLE public.sas_module ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.sas_module_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    203            ?            1259    16503    sas_role    TABLE     ~  CREATE TABLE public.sas_role (
    id integer NOT NULL,
    role_name character varying(100) NOT NULL,
    is_active integer DEFAULT 1 NOT NULL,
    created_by character varying(100) NOT NULL,
    created_at timestamp(0) without time zone NOT NULL,
    modified_by character varying(100),
    modified_at timestamp(0) without time zone,
    is_deleted integer DEFAULT 0 NOT NULL
);
    DROP TABLE public.sas_role;
       public         heap    postgres    false            ?            1259    16494    sas_role_access    TABLE       CREATE TABLE public.sas_role_access (
    id integer NOT NULL,
    role_id integer NOT NULL,
    module_id integer NOT NULL,
    c integer DEFAULT 0 NOT NULL,
    r integer DEFAULT 0 NOT NULL,
    u integer DEFAULT 0 NOT NULL,
    d integer DEFAULT 0 NOT NULL
);
 #   DROP TABLE public.sas_role_access;
       public         heap    postgres    false            ?            1259    16492    sas_role_access_id_seq    SEQUENCE     ?   ALTER TABLE public.sas_role_access ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.sas_role_access_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    205            ?            1259    16501    sas_role_id_seq    SEQUENCE     ?   ALTER TABLE public.sas_role ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.sas_role_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    207            ?            1259    16558    sas_user    TABLE     ?  CREATE TABLE public.sas_user (
    id integer NOT NULL,
    email character varying(75) NOT NULL,
    username character varying(50) NOT NULL,
    password text NOT NULL,
    emp_no character varying(25) NOT NULL,
    fullname character varying(135) NOT NULL,
    grade character varying(135),
    positions character varying(130),
    photo character varying(1000),
    role_id integer NOT NULL,
    created_by character varying(100) NOT NULL,
    created_at timestamp(0) without time zone NOT NULL,
    modified_by character varying(100),
    modified_at timestamp(0) without time zone,
    last_login timestamp(0) without time zone,
    is_deleted integer DEFAULT 0 NOT NULL
);
    DROP TABLE public.sas_user;
       public         heap    postgres    false            ?            1259    16556    sas_user_id_seq    SEQUENCE     ?   ALTER TABLE public.sas_user ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.sas_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    211            ?          0    16522    sas_audit_trail 
   TABLE DATA           m   COPY public.sas_audit_trail (id, modules, ref_id, actions, log_activity, created_by, created_at) FROM stdin;
    public          postgres    false    209   ?"       ?          0    16475    sas_item 
   TABLE DATA           ?   COPY public.sas_item (id, item_name, price, weight_ml, weight_mg, created_by, created_at, modified_by, modified_at, is_deleted) FROM stdin;
    public          postgres    false    201   ?%       ?          0    16489 
   sas_module 
   TABLE DATA           .   COPY public.sas_module (id, name) FROM stdin;
    public          postgres    false    203   ?%       ?          0    16503    sas_role 
   TABLE DATA           z   COPY public.sas_role (id, role_name, is_active, created_by, created_at, modified_by, modified_at, is_deleted) FROM stdin;
    public          postgres    false    207   ?&       ?          0    16494    sas_role_access 
   TABLE DATA           M   COPY public.sas_role_access (id, role_id, module_id, c, r, u, d) FROM stdin;
    public          postgres    false    205   ?&       ?          0    16558    sas_user 
   TABLE DATA           ?   COPY public.sas_user (id, email, username, password, emp_no, fullname, grade, positions, photo, role_id, created_by, created_at, modified_by, modified_at, last_login, is_deleted) FROM stdin;
    public          postgres    false    211   ?&       ?           0    0    sas_audit_trail_id_seq    SEQUENCE SET     E   SELECT pg_catalog.setval('public.sas_audit_trail_id_seq', 57, true);
          public          postgres    false    208            ?           0    0    sas_item_id_seq    SEQUENCE SET     =   SELECT pg_catalog.setval('public.sas_item_id_seq', 1, true);
          public          postgres    false    200            ?           0    0    sas_module_id_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('public.sas_module_id_seq', 3, true);
          public          postgres    false    202            ?           0    0    sas_role_access_id_seq    SEQUENCE SET     D   SELECT pg_catalog.setval('public.sas_role_access_id_seq', 5, true);
          public          postgres    false    204            ?           0    0    sas_role_id_seq    SEQUENCE SET     =   SELECT pg_catalog.setval('public.sas_role_id_seq', 2, true);
          public          postgres    false    206            ?           0    0    sas_user_id_seq    SEQUENCE SET     =   SELECT pg_catalog.setval('public.sas_user_id_seq', 1, true);
          public          postgres    false    210            ?   ?  x??????0?Ϟ??d?f4?lwS衔??CKaq6"??&????;J(eײ???%??????????m???k???4]?X\~I?h??Nz??E???鏡??4???=??B?x??ڵ?Q]?U\V??$???`?,??Z??v3????a?n??B??Na?~???o
??nt+?j?S?????~?=5?f?&???;n?y9?bo,??ò+p?Ţ?W?mw8=?7???<?v?g?<??W?~~??Jn`???n??+@=?ø-?7H???% ~??? iM?]qׅ??<I?%??g4S?e???	&?Y??w??L?وqM}k6m??W)9?T????53?W{??'???y$?/?WyvO?Xy?<??W?!???E?pȫf?b?K ?n??
Y??y?????:m?$N.s"????N????b?Jt?T?I?Hz (? 8?s?z?rb?D?pϞǪJ??{q?Jq???^??L"?S??:??oy??8 eZ?D???s?'Y6q(??⏬??7????<	0ɗ??E??????μ?E??c??;>i??'w??????J?????q?+?s?!12<+`?'?s$6?_??D??????A?yC?6C"ͭK?2O"?`?#??IN??D˩>?????T_$?`s?%??S}?q??T?J؜??5؜?G?8?S}??ׅͩ?Hd??? Ԩs?      ?   V   x?3?t???I?,)?Qp?)M?Q?J-???L?Q04???44 NCSNN?Ĕ?DN##C]]sK+#+??2?????? ?p?      ?   8   x?3??M,.I-RpI,IT-N-?2Br/?/-PpLNN-.?2F??,I??????? ?      ?   M   x?3?p?4??H,???4202?50?50W02?2??2????!.#NG_O??j?@jM??X????qqq y$?      ?   "   x?3?4D@.# i?Ic8?(????qqq ?n      ?   ?   x?U??
?@EןO1??7???jˢ?J#?͠c3?E?R=}Im???\8?R???i?H^?||.?*??ZB?u?-?Sԥ??iԶ3?Vx???Z?{38??6??v?p~?^?fXD?K???|?????d??Vv????߀'?Ȉ?HDsW? ?,S?? ???C?j:p??????????8?Zlj???"=?     