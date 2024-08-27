
```sqlite
INSERT INTO clients (client_type, name, email, phone, address, created_at, updated_at)
VALUES ('individual', 'Иван Иванов', 'ivan@example.com', '1234567890', 'Москва, ул. Примерная, д. 1', datetime('now'), datetime('now'));

INSERT INTO individuals (individual_id, first_name, last_name, patronymic, birth_date)
VALUES (last_insert_rowid(), 'Иван', 'Иванов', 'Иванович', '1990-01-01');




INSERT INTO clients (client_type, name, email, phone, address, created_at, updated_at)
VALUES ('company', 'ООО Пример', 'info@primer.com', '0987654321', 'Москва, ул. Примерная, д. 2', datetime('now'), datetime('now'));

INSERT INTO companies (company_id, inn, kpp, registration_date, director)
VALUES (last_insert_rowid(), '1234567890', '123456789', '2020-01-01', 'Петров Петр Петрович');
```

```sqlite
SELECT * FROM clients;
```

```sqlite
SELECT * FROM clients c
JOIN individuals i ON c.client_id = i.individual_id
WHERE c.client_type = 'individual';
```

```sqlite
SELECT * FROM clients c
JOIN companies co ON c.client_id = co.company_id
WHERE c.client_type = 'company';
```
