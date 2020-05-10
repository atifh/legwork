CREATE TABLE IF NOT EXISTS users (
       id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
       name VARCHAR(100) not null,
       age INTEGER,
       bio TEXT,
       location VARCHAR(50),
       is_active boolean NOT NULL default true,
       created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
       updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
