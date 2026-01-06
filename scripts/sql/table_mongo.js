db.test_types.insertMany([
    {
        _id: ObjectId("507f1f77bcf86cd799439011"),
        booleanField: true,
        intField: 42,
        longField: NumberLong("9223372036854775807"),
        doubleField: 3.14159,
        decimalField: NumberDecimal("123456.789"),
        stringField: "Пример строки",
        dateField: new Date("2025-12-16T10:00:00Z"),
        objectIdField: ObjectId("507f191e810c19729de860ea"),
        arrayField: [1, "два", null, { nested: true }],
        objectField: { key: "value", nestedNull: null },
        binaryField: BinData(0, "aGVsbG8gd29ybGQ="), // base64-encoded "hello world"
        regexField: /test/i,
        javascriptField: { $code: "function() { return 42; }" },
        minKeyField: MinKey,
        maxKeyField: MaxKey,
        nullField: null,
        undefinedField: undefined, // в BSON → null
        timestampField: new Timestamp(0, 1)
    },
    {
        booleanField: null,
        intField: null,
        longField: null,
        doubleField: null,
        decimalField: null,
        stringField: null,
        dateField: null,
        objectIdField: null,
        arrayField: null,
        objectField: null,
        binaryField: null,
        regexField: null,
        javascriptField: null,
        minKeyField: null,
        maxKeyField: null,
        nullField: null,
        undefinedField: null,
        timestampField: null
    }
])
