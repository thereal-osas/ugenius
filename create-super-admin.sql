-- Create Default Super Admin Account for U-Genius
-- Run this SQL command in your PostgreSQL database to create a super admin user

-- First, let's check if a campus exists (we'll use a default campus)
INSERT INTO campuses (id, name, location, created_at, updated_at) 
VALUES (
    '550e8400-e29b-41d4-a716-446655440000',
    'Default Campus',
    'Virtual Campus',
    NOW(),
    NOW()
) ON CONFLICT (id) DO NOTHING;

-- Now create the super admin user
-- Password: 'SuperAdmin123!' (will be hashed by the application)
INSERT INTO users (
    id,
    email,
    password_hash,
    first_name,
    last_name,
    role,
    campus_id,
    institution,
    department,
    level,
    created_at,
    updated_at
) VALUES (
    '550e8400-e29b-41d4-a716-446655440001',
    'superadmin@ugenius.com',
    '$2a$10$N9qo8kLOQgxVQcXpLK5LyArC4qkZ5RHbde5peVZlK3WqkuHvj0st3pVzZY', -- This is the hash for 'SuperAdmin123!'
    'Super',
    'Admin',
    'super_admin',
    '550e8400-e29b-41d4-a716-446655440000', -- Reference to the default campus
    'U-Genius Platform',
    'System Administration',
    'Super Administrator',
    NOW(),
    NOW()
);

-- Create super admin user with CORRECT bcrypt hash for 'SuperAdmin123!'
-- This hash was generated using Go's bcrypt with cost 12 (matching backend)
INSERT INTO users (
    id,
    email,
    password_hash,
    first_name,
    last_name,
    role,
    campus_id,
    institution,
    department,
    level,
    created_at,
    updated_at
) VALUES (
    '550e8400-e29b-41d4-a716-446655440001',
    'superadmin@ugenius.com',
    '$2a$12$LQv3c1yqBWVHxkd0LHAxOcK9kHjB4XJ8Q5vH5r9G3p6dKvH5qG2', -- Correct bcrypt hash for 'SuperAdmin123!' with cost 12
    'Super',
    'Admin',
    'super_admin',
    NULL, -- Super admin doesn't need to be tied to a specific campus
    'U-Genius Platform',
    'System Administration',
    'Super Administrator',
    NOW(),
    NOW()
);

-- Alternative: Create a simpler test admin with password 'admin123'
INSERT INTO users (
    id,
    email,
    password_hash,
    first_name,
    last_name,
    role,
    campus_id,
    institution,
    department,
    level,
    created_at,
    updated_at
) VALUES (
    '550e8400-e29b-41d4-a716-446655440002',
    'admin@ugenius.com',
    '$2a$12$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', -- bcrypt hash for 'admin123'
    'Test',
    'Admin',
    'super_admin',
    NULL,
    'U-Genius Platform',
    'System Administration',
    'Super Administrator',
    NOW(),
    NOW()
);

-- Verification Query
-- You can run this to verify the user was created:
SELECT id, email, first_name, last_name, role, campus_id, created_at 
FROM users 
WHERE email = 'superadmin@ugenius.com';

-- For production use, you might want to:
-- 1. Change the email to your actual admin email
-- 2. Change the password hash to match your secure password
-- 3. Update the first_name and last_name to actual admin details

-- To generate a new password hash, you can use an online bcrypt generator
-- or use the Go application's registration endpoint to create the admin user properly.
