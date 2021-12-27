INSERT INTO users (id, email) VALUES (1, 'user@gmail.com');

INSERT INTO worlds (id, name, active) VALUES (1, 'Fantasy 1', TRUE);
INSERT INTO worlds (id, name, active) VALUES (2, 'Fantasy 2', FALSE);
INSERT INTO worlds (id, name, active) VALUES (3, 'Fantasy 3', TRUE);

INSERT INTO alliances (id, name, tag) VALUES (1, 'Bears', 'BRS');

INSERT INTO players (id, user_id, world_id, alliance_id, name,
        resource_worker_count, resource_gold_count, resource_lumber_count, resource_mana_count)
    VALUES (
        1, 1, 1, 1, 'lightning bear', 100, 500, 300, 50
    );
