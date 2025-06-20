PGDMP  6        	            }            inventaris_kantor    16.3    16.3     �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    16751    inventaris_kantor    DATABASE     �   CREATE DATABASE inventaris_kantor WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'Indonesian_Indonesia.1252';
 !   DROP DATABASE inventaris_kantor;
                postgres    false            �            1259    16762    barang    TABLE     �   CREATE TABLE public.barang (
    id integer NOT NULL,
    nama_barang character varying(100) NOT NULL,
    id_kategori integer NOT NULL,
    tanggal_beli date NOT NULL,
    harga numeric(12,2) NOT NULL,
    kode_barang character varying(50)
);
    DROP TABLE public.barang;
       public         heap    postgres    false            �            1259    16761    barang_id_seq    SEQUENCE     �   CREATE SEQUENCE public.barang_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 $   DROP SEQUENCE public.barang_id_seq;
       public          postgres    false    218            �           0    0    barang_id_seq    SEQUENCE OWNED BY     ?   ALTER SEQUENCE public.barang_id_seq OWNED BY public.barang.id;
          public          postgres    false    217            �            1259    16753    kategori    TABLE     �   CREATE TABLE public.kategori (
    id integer NOT NULL,
    nama_kategori character varying(100) NOT NULL,
    deskripsi text
);
    DROP TABLE public.kategori;
       public         heap    postgres    false            �            1259    16752    kategori_id_seq    SEQUENCE     �   CREATE SEQUENCE public.kategori_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public.kategori_id_seq;
       public          postgres    false    216            �           0    0    kategori_id_seq    SEQUENCE OWNED BY     C   ALTER SEQUENCE public.kategori_id_seq OWNED BY public.kategori.id;
          public          postgres    false    215            V           2604    16765 	   barang id    DEFAULT     f   ALTER TABLE ONLY public.barang ALTER COLUMN id SET DEFAULT nextval('public.barang_id_seq'::regclass);
 8   ALTER TABLE public.barang ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    217    218    218            U           2604    16756    kategori id    DEFAULT     j   ALTER TABLE ONLY public.kategori ALTER COLUMN id SET DEFAULT nextval('public.kategori_id_seq'::regclass);
 :   ALTER TABLE public.kategori ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    215    216    216            �          0    16762    barang 
   TABLE DATA           `   COPY public.barang (id, nama_barang, id_kategori, tanggal_beli, harga, kode_barang) FROM stdin;
    public          postgres    false    218   �       �          0    16753    kategori 
   TABLE DATA           @   COPY public.kategori (id, nama_kategori, deskripsi) FROM stdin;
    public          postgres    false    216   v       �           0    0    barang_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public.barang_id_seq', 6, true);
          public          postgres    false    217            �           0    0    kategori_id_seq    SEQUENCE SET     =   SELECT pg_catalog.setval('public.kategori_id_seq', 3, true);
          public          postgres    false    215            Z           2606    16767    barang barang_pkey 
   CONSTRAINT     P   ALTER TABLE ONLY public.barang
    ADD CONSTRAINT barang_pkey PRIMARY KEY (id);
 <   ALTER TABLE ONLY public.barang DROP CONSTRAINT barang_pkey;
       public            postgres    false    218            X           2606    16760    kategori kategori_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.kategori
    ADD CONSTRAINT kategori_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.kategori DROP CONSTRAINT kategori_pkey;
       public            postgres    false    216            [           2606    16768    barang barang_id_kategori_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.barang
    ADD CONSTRAINT barang_id_kategori_fkey FOREIGN KEY (id_kategori) REFERENCES public.kategori(id) ON DELETE CASCADE;
 H   ALTER TABLE ONLY public.barang DROP CONSTRAINT barang_id_kategori_fkey;
       public          postgres    false    4696    218    216            �   j   x�U�!�0@Qݞb��u#���A�0 FR���`�/�$k=��.׽�&ֳ�g���l�K�դ"[V-�m�#e�&�z��RuJ=�l��!��,�      �   R   x�3�t�I�.)�����tJ,J�K�MS
�p	��ļ��".#N�Ң�̒�"N�ԬD��Ң�L�����" �������� '��     