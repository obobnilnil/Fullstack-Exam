-- รหัสผ่าน "1234" ที่ hash ด้วย bcrypt (cost=10)
-- ใช้ tool เช่น bcrypt-generator.com หรือ Go bcrypt
-- Hash นี้ใช้ทดสอบ login ได้จริง
-- password_hash = "$2a$10$7EqJtq98hPqEX7fNZaFWoOHi6tQ2gk5z/8iR1c3f8eRBE7NDQW22e"

-- Admin
INSERT INTO users (id, email, password_hash, role_id)
VALUES (
  uuid_generate_v4(),
  'admin@example.com',
  '$2a$10$7EqJtq98hPqEX7fNZaFWoOHi6tQ2gk5z/8iR1c3f8eRBE7NDQW22e',
  1
);

-- Staff
INSERT INTO users (id, email, password_hash, role_id)
VALUES (
  uuid_generate_v4(),
  'staff@example.com',
  '$2a$10$7EqJtq98hPqEX7fNZaFWoOHi6tQ2gk5z/8iR1c3f8eRBE7NDQW22e',
  2
);
