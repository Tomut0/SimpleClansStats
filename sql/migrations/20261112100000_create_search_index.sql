-- +goose Up
CREATE VIRTUAL TABLE search_documents USING fts5(
  server_id UNINDEXED,
  entity_type UNINDEXED,
  entity_id UNINDEXED,
  name UNINDEXED,
  payload UNINDEXED,
  search_text,
  tokenize = 'unicode61 remove_diacritics 2'
);

INSERT INTO search_documents (server_id, entity_type, entity_id, name, payload, search_text)
SELECT servers.id, 'clan', json_extract(clan.value, '$.externalId'),
       json_extract(clan.value, '$.tag'), clan.value,
       json_extract(clan.value, '$.tag') || ' ' || json_extract(clan.value, '$.name')
FROM servers, json_each(CASE WHEN json_valid(state_payload) THEN state_payload ELSE '{}' END, '$.clans') AS clan;

INSERT INTO search_documents (server_id, entity_type, entity_id, name, payload, search_text)
SELECT servers.id, 'player', json_extract(player.value, '$.uuid'),
       json_extract(player.value, '$.name'), player.value,
       json_extract(player.value, '$.name') || ' ' || COALESCE(json_extract(player.value, '$.clanTag'), '')
FROM servers, json_each(CASE WHEN json_valid(state_payload) THEN state_payload ELSE '{}' END, '$.players') AS player;

-- +goose Down
DROP TABLE IF EXISTS search_documents;
