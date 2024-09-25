CREATE TABLE accounts (
  id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
  name VARCHAR(20) NOT NULL,
  password VARCHAR(50) NOT NULL
)

INSERT INTO accounts (name, password) VALUES ('a0_name', 'a0_password');
INSERT INTO accounts (name, password) VALUES ('a1_name', 'a1_password');
INSERT INTO accounts (name, password) VALUES ('a2_name', 'a2_password');
INSERT INTO accounts (name, password) VALUES ('a3_name', 'a3_password');
INSERT INTO accounts (name, password) VALUES ('a4_name', 'a4_password');
