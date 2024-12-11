-- +goose Up
-- +goose StatementBegin
CREATE TABLE devices (
  id  TEXT PRIMARY KEY, 
  name  TEXT,
  "limit" DECIMAL,
  type TEXT
);

CREATE TABLE data (
  id SERIAL PRIMARY KEY,  
  device_id TEXT, 
  value DECIMAL,
  timestamp TIMESTAMP 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE devices;
DROP TABLE data;
-- +goose StatementEnd
