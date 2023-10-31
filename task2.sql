/* Написать SQL-запросы для PostgreSQL добавления трех полей, 
изменения одного поля и добавления двух индексов в базу данных размером свыше 100 ГБ и более 8 миллионов строк. */

-- Добавление трех полей
ALTER TABLE users ADD COLUMN new_field1 VARCHAR(255);
ALTER TABLE users ADD COLUMN new_field2 INT DEFAULT 0;
ALTER TABLE users ADD COLUMN new_field3 DATE;

-- Изменение одного поля (увеличение размера поля)
-- Создаем новое поле с новым типом данных
ALTER TABLE users RENAME COLUMN name TO old_name;
ALTER TABLE users ADD COLUMN name VARCHAR(100);
-- Обновляем новое поле с данными из старого поля
UPDATE users SET name = old_name;
-- Удаляем старое поле
ALTER TABLE users DROP COLUMN old_name;

-- Добавление двух индексов
CREATE INDEX CONCURRENTLY idx_users_new_field1 ON users(new_field1);
CREATE INDEX CONCURRENTLY idx_users_new_field2 ON users(new_field2);