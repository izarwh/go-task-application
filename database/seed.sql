INSERT INTO tasks (id, title, description, status, due_date, created_at, updated_at, is_deleted)
VALUES
    (
        gen_random_uuid(), 
        'Complete Initial Project Setup', 
        '{"details": "Setup project repositories, configure CI/CD pipelines, and prepare boilerplate code.", "tags": ["infrastructure", "setup"]}', 
        'completed', 
        '2026-03-31', 
        CURRENT_TIMESTAMP, 
        CURRENT_TIMESTAMP, 
        false
    ),
    (
        gen_random_uuid(), 
        'Implement Authentication API', 
        '{"details": "Develop JWT-based authentication for the user application.", "priority": "high"}', 
        'pending', 
        '2026-04-05', 
        CURRENT_TIMESTAMP, 
        CURRENT_TIMESTAMP, 
        false
    ),
    (
        gen_random_uuid(), 
        'Review Database Schema', 
        '{"checklist": ["Check index performance", "Ensure cascading rules", "Review naming conventions"]}', 
        'pending', 
        '2026-04-10', 
        CURRENT_TIMESTAMP, 
        CURRENT_TIMESTAMP, 
        false
    );
