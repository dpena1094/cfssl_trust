package model

// NOTE: this file is autogenerated using gen_sql.go.

var rawSQL = []string{
`CREATE TABLE IF NOT EXISTS schema_version (
	revision	INTEGER UNIQUE PRIMARY KEY,
	created_at	INTEGER
);`,
`INSERT INTO schema_version (revision, created_at)
	SELECT 1, 1485991500
	WHERE NOT EXISTS (SELECT 1 FROM schema_version
				WHERE revision = 1);`,
`CREATE TABLE IF NOT EXISTS certificates (
	ski		TEXT NOT NULL,
	aki		TEXT NOT NULL,
	serial		INTEGER NOT NULL,
	not_before	INTEGER NOT NULL,
	not_after	INTEGER NOT NULL,
	raw		BLOB NOT NULL,
	UNIQUE(ski, serial)
);`,
`CREATE TABLE IF NOT EXISTS sources (
	ski	TEXT NOT NULL,
	url	TEXT NOT NULL,
	FOREIGN KEY (ski) REFERENCES certificates(ski)
);`,
`CREATE TABLE IF NOT EXISTS aia (
	ski	TEXT PRIMARY KEY,
	url	TEXT NOT NULL,
	FOREIGN KEY (ski) REFERENCES certificates(aki)
);`,
`CREATE TABLE IF NOT EXISTS revocations (
	ski		TEXT PRIMARY KEY,
	revoked_at	INTEGER NOT NULL,
	mechanism	TEXT NOT NULL,
	reason		TEXT NOT NULL,
	FOREIGN KEY (ski) REFERENCES certificates(ski)
);`,
`CREATE TABLE IF NOT EXISTS roots (
	ski		TEXT NOT NULL,
	serial		INTEGER NOT NULL,
	release		TEXT NOT NULL,
	UNIQUE (ski, serial, release)
	FOREIGN KEY (ski) REFERENCES certificates(ski),
	FOREIGN KEY (release) REFERENCES root_releases(version)`,
`);`,
`CREATE TABLE IF NOT EXISTS root_releases (
	version		TEXT NOT NULL,
	released_at	INTEGER NOT NULL
);`,
`CREATE TABLE IF NOT EXISTS intermediates (
	ski		TEXT NOT NULL,
	serial		INTEGER NOT NULL,
	release		TEXT NOT NULL,
	UNIQUE (ski, serial, release)
	FOREIGN KEY (ski) REFERENCES certificates(ski),
	FOREIGN KEY (release) REFERENCES intermediate_releases(version)
);`,
`CREATE TABLE IF NOT EXISTS intermediate_releases (
	version		TEXT NOT NULL,
	released_at	INTEGER NOT NULL
);`,
}
