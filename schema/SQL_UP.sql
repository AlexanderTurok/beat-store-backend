CREATE TABLE IF NOT EXISTS account (
  id BIGSERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  username TEXT UNIQUE NOT NULL,
  email TEXT UNIQUE NOT NULL,
  photo_path TEXT,
  password_hash TEXT NOT NULL,
  confirmed BOOLEAN DEFAULT 0,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS artist (
  id BIGINT REFERENCES account (id) ON DELETE CASCADE UNIQUE NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS beat (
  id BIGSERIAL PRIMARY KEY,
  artist_id BIGINT REFERENCES artist (id) ON DELETE CASCADE NOT NULL,
  name TEXT NOT NULL,
  bpm TEXT NOT NULL,
  key TEXT NOT NULL,
  photo_path TEXT NOT NULL,
  mp3_path TEXT NOT NULL,
  wav_path TEXT,
  likes BIGINT DEFAULT 0,
  genre TEXT DEFAULT 'All',
  mood TEXT DEFAULT 'All',
  standart_price TEXT NOT NULL,
  premium_price TEXT NOT NULL,
  unlimited_price TEXT NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS tag (
  id BIGSERIAL PRIMARY KEY,
  beat_id BIGINT REFERENCES beat (id) ON DELETE CASCADE NOT NULL,
  tag_name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS account_beat (
  id BIGSERIAL PRIMARY KEY,
  account_id BIGINT REFERENCES account (id) ON DELETE CASCADE NOT NULL,
  beat_id BIGINT REFERENCES beat (id) ON DELETE CASCADE NOT NULL,
  access_status TEXT CHECK IN('standart', 'premium', 'unlimited') NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS playlist (
  id BIGSERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS playlist_beat (
  playlist_id BIGINT REFERENCES playlist (id) ON DELETE CASCADE NOT NULL,
  beat_id BIGINT REFERENCES beat (id) ON DELETE CASCADE NOT NULL,
  PRIMARY KEY (playlist_id, beat_id)
);

CREATE TABLE IF NOT EXISTS account_playlist (
  account_id BIGINT REFERENCES account (id) ON DELETE CASCADE NOT NULL,
  playlist_id BIGINT REFERENCES playlist (id) ON DELETE CASCADE NOT NULL,
  PRIMARY KEY (playlist_id, account_id)
);