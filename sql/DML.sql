--Input User
INSERT INTO users (name, email, password, role, updated_at)
VALUES 
  ('Admin1', 'admin1@example.com', 'admin1_password', 'ADMIN', CURRENT_TIMESTAMP),
  ('Admin2', 'admin2@example.com', 'admin2_password', 'ADMIN', CURRENT_TIMESTAMP),
  ('User1', 'user1@example.com', 'user1_password', 'USER', CURRENT_TIMESTAMP),
  ('User2', 'user2@example.com', 'user2_password', 'USER', CURRENT_TIMESTAMP),
  ('User3', 'user3@example.com', 'user3_password', 'USER', CURRENT_TIMESTAMP),
  ('User4', 'user4@example.com', 'user4_password', 'USER', CURRENT_TIMESTAMP);

--Input Project
INSERT INTO projects (name, manager_id, deadline, updated_at)
VALUES 
    ('Project Web', 'f3bb0a14-b760-45c6-a0bb-e9b2e3e3f178', '2024-02-05', CURRENT_TIMESTAMP),
    ('Mobile App Project', 'f3bb0a14-b760-45c6-a0bb-e9b2e3e3f178', '2024-03-15', CURRENT_TIMESTAMP),
    ('Data Analytics Web', 'e2c7508e-edb6-408f-968f-5a67748b18c3', '2024-04-30', CURRENT_TIMESTAMP),
    ('Infrastructure Upgrade', 'e2c7508e-edb6-408f-968f-5a67748b18c3', '2024-05-10', CURRENT_TIMESTAMP),
    ('E-commerce Platform', 'e2c7508e-edb6-408f-968f-5a67748b18c3', '2024-06-20', CURRENT_TIMESTAMP);

--Input Project Member
INSERT INTO project_members (member_id, project_id)
VALUES
('f3bb0a14-b760-45c6-a0bb-e9b2e3e3f178', '7bdf1090-033f-497b-b161-bb4ce62c422d'),
('c3086749-8326-49bf-bb13-803694731708', '7bdf1090-033f-497b-b161-bb4ce62c422d'),
('70913634-660b-4139-91fd-596d8e00c896', '7bdf1090-033f-497b-b161-bb4ce62c422d');

--Input Task
INSERT INTO tasks (name, status, approval, person_in_charge, deadline, project_Id, updated_at)
VALUES
('fitur-login', 1, false, 'c3086749-8326-49bf-bb13-803694731708', '2024-01-31', '7bdf1090-033f-497b-b161-bb4ce62c422d', CURRENT_TIMESTAMP),
('fitur-search-page', 1, false, 'c3086749-8326-49bf-bb13-803694731708','2024-02-01', '7bdf1090-033f-497b-b161-bb4ce62c422d', CURRENT_TIMESTAMP),
('fitur-create-task', 1, false, 'c3086749-8326-49bf-bb13-803694731708', '2024-02-01', '7bdf1090-033f-497b-b161-bb4ce62c422d', CURRENT_TIMESTAMP),
('fitur-update-task', 1, false, '70913634-660b-4139-91fd-596d8e00c896', '2024-02-02', '7bdf1090-033f-497b-b161-bb4ce62c422d', CURRENT_TIMESTAMP),
('fitur-delete-task', 1, false, '70913634-660b-4139-91fd-596d8e00c896', '2024-02-02', '7bdf1090-033f-497b-b161-bb4ce62c422d', CURRENT_TIMESTAMP);