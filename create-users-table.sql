DROP TABLE IF EXISTS users;

CREATE TABLE users_test(
  id serial PRIMARY KEY,
  firstname VARCHAR(128),
  lastname VARCHAR(128),
  email VARCHAR(128),
  company VARCHAR(128),
  phone VARCHAR(128)
);

INSERT INTO users_test(firstName, lastName, email, company, phone)
VALUES
('Jane', 'Doe', 'jane@example.com', 'ABC.Inc', '1234567890'),
('John', 'Doe', 'john@example.com', 'XYZ.Corp', '9876543210');