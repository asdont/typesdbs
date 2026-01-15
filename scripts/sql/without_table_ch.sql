SELECT
    -- Integers
    127 AS col_int8,
    32767 AS col_int16,
    2147483647 AS col_int32,
    9223372036854775807 AS col_int64,
    170141183460469231731687303715884105727 AS col_int128,
    57896044618658097711785492504343953926634992332820282019728792003956564819967 AS col_int256,

    -- Unsigned Integers
    255 AS col_uint8,
    65535 AS col_uint16,
    4294967295 AS col_uint32,
    18446744073709551615 AS col_uint64,
    340282366920938463463374607431768211455 AS col_uint128,
    115792089237316195423570985008687907853269984665640564039457584007913129639935 AS col_uint256,

    -- Decimals
    toDecimal32('1234567.89', 2) AS col_decimal,
    toDecimal64('9999999999999.1234', 4) AS col_numeric,

    -- Floats
    4.56 AS col_float32,
    7.123456789123456 AS col_float64,

    -- Nullable Floats
    CAST(8.123 AS Nullable(Float32)) AS col_null_float32,
    CAST(NULL AS Nullable(Float64)) AS col_null_float64,

    -- Special Floats (NaN, Inf)
    CAST('nan' AS Float32) AS col_nan,
    CAST('+inf' AS Float32) AS col_inf_plus,
    CAST('-inf' AS Float32) AS col_inf_minus,

    -- Strings
    CAST('x' AS FixedString(10)) AS col_char,
    'abcdef' AS col_text,

    -- Boolean
    true AS col_boolean,

    -- Nullable Unsigned Integers
    CAST(250 AS Nullable(UInt8)) AS col_null_uint_tiny,
    CAST(NULL AS Nullable(UInt16)) AS col_null_uint_smallint,

    -- UUID
    generateUUIDv4() AS col_uuid,
    CAST('11111111-0000-0000-0000-000000000000' AS Nullable(UUID)) AS col_null_uuid,

    -- LowCardinality
    CAST('computer' AS LowCardinality(String)) AS col_low_cardinality_str,
    CAST(NULL AS LowCardinality(Nullable(String))) AS col_low_cardinality_null_str,

    -- Date/DateTime
    toDate('1971-02-03') AS col_date,
    toDateTime('1972-03-04 05:05:05') AS col_datetime,
    toDateTime('1983-03-04 05:05:05', 'Europe/Moscow') AS col_datetime_target,
    toDateTime64('1984-03-04 05:05:05.123', 3) AS col_datetime_64_without_timezone,
    toDateTime64('1985-03-04 05:05:05.123', 3, 'UTC') AS col_datetime_64,
    toDateTime64('1986-03-04 05:05:05', 0, 'Europe/Moscow') AS col_datetime_64_2,

    -- Nullable Date/DateTime
    CAST(toDate('1971-02-10') AS Nullable(Date)) AS col_null_date,
    CAST(toDateTime('1972-03-10 01:01:01') AS Nullable(DateTime)) AS col_null_datetime,
    CAST(NULL AS Nullable(DateTime('UTC'))) AS col_null_datetime_utc,
    CAST(toDateTime64('1994-03-10 01:01:01.123', 3, 'UTC') AS Nullable(DateTime64(3, 'UTC'))) AS col_null_datetime_64,

    -- Nullable Decimal
    CAST(toDecimal32('0.12345', 5) AS Nullable(Decimal(8, 5))) AS col_null_decimal,

    -- IP
    toIPv4('192.174.0.100') AS col_ipv4,
    toIPv6('2001:44c8:129:2632:33:0:252:2') AS col_ipv6,

    -- Enum
    CAST('A' AS Enum8('A' = 1, 'B' = 2, 'C' = 3)) AS col_enum,

    -- Map
    map('key1', 'value1', 'key2', 'value2') AS col_map_string_string,

    -- Arrays
    [-128, 0, 127] AS col_array_Int8,
    [0, 255] AS col_array_UInt8,
    ['apple', 'banana'] AS col_array_string,
    [true, false] AS col_array_bool,
    [toDate('2000-01-01'), toDate('2001-01-01')] AS col_array_date,
    [toDateTime('2000-01-01 00:00:00'), toDateTime('2001-01-01 00:00:00')] AS col_array_date_time,
    [toDateTime('2000-01-01 00:00:00', 'UTC')] AS col_array_date_time_utc,

    -- Arrays with Nullable
    [CAST(-128 AS Nullable(Int8)), NULL, CAST(127 AS Nullable(Int8))] AS col_array_null_Int8,
    [CAST('apple' AS Nullable(String)), NULL, CAST('cherry' AS Nullable(String))] AS col_array_null_string,
    [CAST(true AS Nullable(Bool)), NULL, CAST(false AS Nullable(Bool))] AS col_array_null_bool,

    -- Nested Arrays
    [['a', 'b'], ['c']] AS col_array_array_string,

    -- Array of Maps
    [map('k1', 'v1'), map('k2', 'v2')] AS col_array_map,

    -- Array of Tuples
    [(1, 2), (3, 4)] AS col_array_tuple;
