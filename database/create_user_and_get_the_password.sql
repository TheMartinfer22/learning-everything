-- Sistema de Estrutura
DROP DATABASE IF EXISTS users;
CREATE DATABASE users;
USE users;
CREATE TABLE users_table (
  id INT AUTO_INCREMENT PRIMARY KEY,
  email VARCHAR(50),
  password VARCHAR(120)
);

INSERT INTO users_table (email, password) VALUES (
  'insecure_structure@anonmail.com',
  'passwordeasy123'
), (
  'admin@anonmail.com',
  'passwordeasy123'
);

-- Attacker Simulation

show databases;
use users;
show tables;
describe users_table;

SELECT email, password from users_table where email rlike 'admin';
