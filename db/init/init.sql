-- connection.envの設定によりDB作成済みなので不要
-- CREATE DATABASE grpc_stady;

\c grpc_stady

CREATE SCHEMA grpc_stady_schema;

CREATE ROLE dev WITH LOGIN PASSWORD 'pass';

GRANT ALL PRIVILEGES ON SCHEMA grpc_stady_schema TO dev;

CREATE TABLE  grpc_stady_schema.todo (
  id INT,
  title VARCHAR(10) NOT NULL,
  context VARCHAR(10) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY (id)
);

GRANT ALL PRIVILEGES ON grpc_stady_schema.todo TO dev;

INSERT INTO grpc_stady_schema.todo VALUES(1, 'タイトル1', '内容1', now(), now());
INSERT INTO grpc_stady_schema.todo VALUES(2, 'タイトル2', '内容2', now(), now());
INSERT INTO grpc_stady_schema.todo VALUES(3, 'タイトル3', '内容3', now(), now());

