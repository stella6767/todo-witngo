CREATE TABLE IF NOT EXISTS todo (
                                    id SERIAL PRIMARY KEY,
                                    title TEXT NOT NULL,
                                    completed BOOLEAN DEFAULT false,
                                    created_at TIMESTAMP DEFAULT NOW()
    );


insert into todo values (1, 'test1', false);

CREATE USER postgres WITH PASSWORD '1234';


-- 기본 사용자 목록
SELECT rolname AS username,
       rolsuper AS is_superuser,
       rolcreatedb AS can_create_db,
       rolcreaterole AS can_create_role,
       rolcanlogin AS can_login,
       rolvaliduntil AS password_valid_until,
       rolconnlimit AS connection_limit
FROM pg_catalog.pg_roles
WHERE rolcanlogin = true; -- 로그인 가능한 사용자만 필터링


-- 사용자 권한 조회
SELECT * FROM information_schema.role_table_grants
WHERE grantee = 'postgres';


GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO postgres;

-- 모든 시퀀스에 권한 부여 (스키마 public 기준)
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO postgres;

-- 향후 생성되는 시퀀스에도 자동 권한 부여
ALTER DEFAULT PRIVILEGES IN SCHEMA public
GRANT USAGE, SELECT ON SEQUENCES TO postgres;
