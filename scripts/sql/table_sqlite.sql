CREATE TABLE test_types
(
    col_integer   INTEGER,
    col_real      REAL,
    col_text      TEXT NOT NULL,
    col_boolean   BOOLEAN,
    col_date      DATE,
    col_timestamp TIMESTAMP,
    col_blob      BLOB,
    col_decimal   DECIMAL(10, 2),
    col_enum      TEXT CHECK (col_enum IN ('val_1', 'val_2', 'val_3')),
    col_json      JSON,
    col_unique    TEXT UNIQUE
);

INSERT INTO test_types(col_integer,
                       col_real,
                       col_text,
                       col_boolean,
                       col_date,
                       col_timestamp,
                       col_blob,
                       col_decimal,
                       col_enum,
                       col_json,
                       col_unique)
VALUES (1,
        1.123,
        'qwerty',
        false,
        '1971-02-03',
        '1972-03-04 05:05:05',
        '[21 5 13]',
        1.23,
        'val_1',
        '{"key1": "val1"}',
        '{"key2": "val2"}');
