CREATE TABLE test_types
(
    col_tinyint          TINYINT          DEFAULT 0,
    col_smallint         SMALLINT         DEFAULT 0,
    col_integer          INTEGER          DEFAULT 0,
    col_bigint           BIGINT           DEFAULT 0,
    col_double_precision DOUBLE PRECISION DEFAULT 0,
    col_date             DATE,
    col_datetime         DATETIME,
    col_decimal          DECIMAL          DEFAULT 0,
    col_char             CHAR(10)         DEFAULT '',
    col_varchar          VARCHAR(255)     DEFAULT '',
    col_binary           BINARY           DEFAULT 0
);

INSERT INTO test_types(col_tinyint,
                       col_smallint,
                       col_integer,
                       col_bigint,
                       col_double_precision,
                       col_date,
                       col_datetime,
                       col_decimal,
                       col_char,
                       col_varchar)
VALUES (127,
        32767,
        8388607,
        9223372036854775807,
        3.4028,
        '1971-02-03',
        '1972-03-04 05:05:05',
        1.123,
        's',
        'qwerty');