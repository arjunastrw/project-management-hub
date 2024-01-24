--Input User
INSERT INTO users (name, email, password, role, updated_at)
VALUES 
  ('Admin1', 'admin1@example.com', 'admin1_password', 'ADMIN', CURRENT_TIMESTAMP),
  ('Admin2', 'admin2@example.com', 'admin2_password', 'ADMIN', CURRENT_TIMESTAMP),
  ('Manager1', 'manager1@example.com', 'manager1_password', 'MANAGER', CURRENT_TIMESTAMP),
  ('Manager2', 'manager2@example.com', 'manager2_password', 'MANAGER', CURRENT_TIMESTAMP),
  ('Teammember1', 'teammember1@example.com', 'teammember1_password', 'TEAM MEMBER', CURRENT_TIMESTAMP),
  ('Teammember2', 'teammember2@example.com', 'teammember2_password', 'TEAM MEMBER', CURRENT_TIMESTAMP);

--Input Project
INSERT INTO projects (name, manager_id, deadline, updated_at)
VALUES 
    ('Project Web', 'b8719cc2-2117-4e95-a38c-d71fb2bc4926', '2024-02-05', CURRENT_TIMESTAMP),
    ('Mobile App Project', 'b8719cc2-2117-4e95-a38c-d71fb2bc4926', '2024-03-15', CURRENT_TIMESTAMP),
    ('Data Analytics Web', 'b8719cc2-2117-4e95-a38c-d71fb2bc4926', '2024-04-30', CURRENT_TIMESTAMP),
    ('Infrastructure Upgrade', '129a9db0-b7ea-417d-a918-295d9cc64e5a', '2024-05-10', CURRENT_TIMESTAMP),
    ('E-commerce Platform', '129a9db0-b7ea-417d-a918-295d9cc64e5a', '2024-06-20', CURRENT_TIMESTAMP);

--Input Project Member
INSERT INTO project_members (member_id, project_id)
VALUES
('8d782371-f076-4eae-a65a-42339f182ef6', '1f0927ac-492c-414a-9049-573973c10673'),
('8d782371-f076-4eae-a65a-42339f182ef6', '145ca10b-5c47-429b-b840-fc9d8c5feb33'),
('1ffe875d-45c3-4dc7-8d71-e91238e4f1c1', '29c41de7-9a5c-4fc5-a1d4-f026b5163bc2');

--Input Task
INSERT INTO tasks (name, status, approval, person_in_charge, deadline, project_Id, updated_at)
VALUES
('fitur-login', 'In Progress', false, '8d782371-f076-4eae-a65a-42339f182ef6', '2024-01-31', '1f0927ac-492c-414a-9049-573973c10673', CURRENT_TIMESTAMP),
('fitur-search-page','In Progress', false, '8d782371-f076-4eae-a65a-42339f182ef6','2024-02-01', '1f0927ac-492c-414a-9049-573973c10673', CURRENT_TIMESTAMP),
('fitur-create-task', 'In Progress', false, '8d782371-f076-4eae-a65a-42339f182ef6', '2024-02-01', '1f0927ac-492c-414a-9049-573973c10673', CURRENT_TIMESTAMP),
('fitur-update-task', 'In Progress', false, '8d782371-f076-4eae-a65a-42339f182ef6', '2024-02-02', '1f0927ac-492c-414a-9049-573973c10673', CURRENT_TIMESTAMP),
('fitur-delete-task', 'In Progress', false, '8d782371-f076-4eae-a65a-42339f182ef6', '2024-02-02', '1f0927ac-492c-414a-9049-573973c10673', CURRENT_TIMESTAMP);