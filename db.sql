-- Table: public.country

-- DROP TABLE IF EXISTS public.country;

CREATE TABLE IF NOT EXISTS public.country
(
    country_id integer NOT NULL DEFAULT nextval('country_country_id_seq'::regclass),
    country character varying(50) COLLATE pg_catalog."default" NOT NULL,
    last_update timestamp without time zone NOT NULL DEFAULT now(),
    CONSTRAINT country_pkey PRIMARY KEY (country_id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.country
    OWNER to postgres;

-- Trigger: last_updated

-- DROP TRIGGER IF EXISTS last_updated ON public.country;

CREATE OR REPLACE TRIGGER last_updated
    BEFORE UPDATE 
    ON public.country
    FOR EACH ROW
    EXECUTE FUNCTION public.last_updated();

-- Table: public.city

-- DROP TABLE IF EXISTS public.city;

CREATE TABLE IF NOT EXISTS public.city
(
    city_id integer NOT NULL DEFAULT nextval('city_city_id_seq'::regclass),
    city character varying(50) COLLATE pg_catalog."default" NOT NULL,
    country_id smallint NOT NULL,
    last_update timestamp without time zone NOT NULL DEFAULT now(),
    CONSTRAINT city_pkey PRIMARY KEY (city_id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.city
    OWNER to postgres;
-- Index: idx_fk_country_id

-- DROP INDEX IF EXISTS public.idx_fk_country_id;

CREATE INDEX IF NOT EXISTS idx_fk_country_id
    ON public.city USING btree
    (country_id ASC NULLS LAST)
    TABLESPACE pg_default;

-- Trigger: last_updated

-- DROP TRIGGER IF EXISTS last_updated ON public.city;

CREATE OR REPLACE TRIGGER last_updated
    BEFORE UPDATE 
    ON public.city
    FOR EACH ROW
    EXECUTE FUNCTION public.last_updated();

-- Table: public.person

-- DROP TABLE IF EXISTS public.person;

CREATE TABLE IF NOT EXISTS public.person
(
    person_id integer,
    city_id integer,
    identity_id uuid,
    last_update timestamp with time zone
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.person
    OWNER to test;