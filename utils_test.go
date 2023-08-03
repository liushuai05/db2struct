package db2struct

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLintFieldName(t *testing.T) {
	name := lintFieldName("_")
	Convey("Should get underscore as fieldName", t, func() {
		So(name, ShouldEqual, "_")
	})

	name = lintFieldName("foo_id")
	Convey("Should be able to convert field name", t, func() {
		So(name, ShouldEqual, "FooID")
	})

	name = lintFieldName("foo__id")
	Convey("Should be able to convert field name", t, func() {
		So(name, ShouldEqual, "FooID")
	})

	name = lintFieldName("1__2")
	Convey("Should be able to convert field name", t, func() {
		So(name, ShouldEqual, "1_2")
	})

	name = lintFieldName("_id")
	Convey("Should be able to convert field name", t, func() {
		So(name, ShouldEqual, "ID")
	})

	name = lintFieldName("foo")
	Convey("Should be able to convert field name", t, func() {
		So(name, ShouldEqual, "Foo")
	})
}

func TestMysqlStringGenerate(t *testing.T) {
	expectedStruct :=
		`package test

type testStruct struct {
	NullStringColumn sql.NullString
	StringColumn     string
}
`

	columnMap := map[string]map[string]string{
		"stringColumn":     {"nullable": "NO", "value": "varchar"},
		"nullStringColumn": {"nullable": "YES", "value": "varchar"},
	}
	columnsSorted := []string{"stringColumn", "nullStringColumn"}
	bytes, err := Generate(columnMap, columnsSorted, "test_table", "testStruct", "test", false, false, true, false)

	Convey("Should be able to generate map from string column", t, func() {
		So(err, ShouldBeNil)
		So(string(bytes), ShouldEqual, expectedStruct)
	})
}

func TestMysqlBlobGenerate(t *testing.T) {
	expectedStruct :=
		`package test

type testStruct struct {
	BinaryColumn         []byte
	BlobColumn           []byte
	LongBlobColumn       []byte
	MediumBlobColumn     []byte
	NullBinaryColumn     []byte
	NullBlobColumn       []byte
	NullLongBlobColumn   []byte
	NullMediumBlobColumn []byte
	NullVarbinaryColumn  []byte
	VarbinaryColumn      []byte
}
`

	columnMap := map[string]map[string]string{
		"binaryColumn":         {"nullable": "NO", "value": "binary"},
		"nullBinaryColumn":     {"nullable": "YES", "value": "binary"},
		"blobColumn":           {"nullable": "NO", "value": "blob"},
		"nullBlobColumn":       {"nullable": "YES", "value": "blob"},
		"longBlobColumn":       {"nullable": "NO", "value": "longblob"},
		"nullLongBlobColumn":   {"nullable": "YES", "value": "longblob"},
		"mediumBlobColumn":     {"nullable": "NO", "value": "mediumblob"},
		"nullMediumBlobColumn": {"nullable": "YES", "value": "mediumblob"},
		"varbinaryColumn":      {"nullable": "NO", "value": "varbinary"},
		"nullVarbinaryColumn":  {"nullable": "YES", "value": "varbinary"},
	}
	columnsSorted := []string{
		"binaryColumn",
		"nullBinaryColumn",
		"blobColumn",
		"nullBlobColumn",
		"longBlobColumn",
		"nullLongBlobColumn",
		"mediumBlobColumn",
		"nullMediumBlobColumn",
		"varbinaryColumn",
		"nullVarbinaryColumn"}

	bytes, err := Generate(columnMap, columnsSorted, "test_table", "testStruct", "test", false, false, true, false)

	Convey("Should be able to generate map from string column", t, func() {
		So(err, ShouldBeNil)
		So(string(bytes), ShouldEqual, expectedStruct)
	})
}

func TestMysqlDateGenerate(t *testing.T) {
	columnMap := map[string]map[string]string{
		"DateColumn":          {"nullable": "NO", "value": "date"},
		"nullDateColumn":      {"nullable": "YES", "value": "date"},
		"DateTimeColumn":      {"nullable": "NO", "value": "datetime"},
		"nullDateTimeColumn":  {"nullable": "YES", "value": "datetime"},
		"TimeColumn":          {"nullable": "NO", "value": "time"},
		"nullTimeColumn":      {"nullable": "YES", "value": "time"},
		"TimeStampColumn":     {"nullable": "NO", "value": "timestamp"},
		"nullTimeStampColumn": {"nullable": "YES", "value": "timestamp"},
	}

	expectedStruct :=
		`package test

type testStruct struct {
	DateColumn          time.Time
	DateTimeColumn      time.Time
	TimeColumn          time.Time
	TimeStampColumn     time.Time
	NullDateColumn      time.Time
	NullDateTimeColumn  time.Time
	NullTimeColumn      time.Time
	NullTimeStampColumn time.Time
}
`
	columnsSorted := []string{
		"DateColumn",
		"DateTimeColumn",
		"TimeColumn",
		"TimeStampColumn",
		"NullDateColumn",
		"NullDateTimeColumn",
		"NullTimeColumn",
		"NullTimeStampColumn"}

	bytes, err := Generate(columnMap, columnsSorted, "test_table", "testStruct", "test", false, false, true, false)

	Convey("Should be able to generate map from string column", t, func() {
		So(err, ShouldBeNil)
		So(string(bytes), ShouldEqual, expectedStruct)
	})

	expectedStruct =
		`package test

type testStruct struct {
	DateColumn          time.Time
	DateTimeColumn      time.Time
	TimeColumn          time.Time
	TimeStampColumn     time.Time
	NullDateColumn      null.Time
	NullDateTimeColumn  null.Time
	NullTimeColumn      null.Time
	NullTimeStampColumn null.Time
}
`
	bytes, err = Generate(columnMap, columnsSorted, "test_table", "testStruct", "test", false, false, true, true)

	Convey("Should be able to generate map from string column", t, func() {
		So(err, ShouldBeNil)
		So(string(bytes), ShouldEqual, expectedStruct)
	})
}

func TestMysqlFloatGenerate(t *testing.T) {

	columnMap := map[string]map[string]string{
		"floatColumn":       {"nullable": "NO", "value": "float"},
		"nullFloatColumn":   {"nullable": "YES", "value": "float"},
		"doubleColumn":      {"nullable": "NO", "value": "double"},
		"nullDoubleColumn":  {"nullable": "YES", "value": "double"},
		"decimalColumn":     {"nullable": "NO", "value": "decimal"},
		"nullDecimalColumn": {"nullable": "YES", "value": "decimal"},
	}

	expectedStruct :=
		`package test

type testStruct struct {
	DecimalColumn     float64
	DoubleColumn      float64
	FloatColumn       float32
	NullDecimalColumn sql.NullFloat64
	NullDoubleColumn  sql.NullFloat64
	NullFloatColumn   sql.NullFloat64
}
`
	columnsSorted := []string{
		"DecimalColumn",
		"DoubleColumn",
		"FloatColumn",
		"NullDecimalColumn",
		"NullDoubleColumn",
		"NullFloatColumn"}

	bytes, err := Generate(columnMap, columnsSorted, "test_table", "testStruct", "test", false, false, true, false)

	Convey("Should be able to generate map from string column", t, func() {
		So(err, ShouldBeNil)
		So(string(bytes), ShouldEqual, expectedStruct)
	})

	expectedStruct =
		`package test

type testStruct struct {
	DecimalColumn     float64
	DoubleColumn      float64
	FloatColumn       float32
	NullDecimalColumn null.Float
	NullDoubleColumn  null.Float
	NullFloatColumn   null.Float
}
`

	bytes, err = Generate(columnMap, columnsSorted, "test_table", "testStruct", "test", false, false, true, true)

	Convey("Should be able to generate map from string column", t, func() {
		So(err, ShouldBeNil)
		So(string(bytes), ShouldEqual, expectedStruct)
	})
}

func TestMysqlIntGenerate(t *testing.T) {
	columnMap := map[string]map[string]string{
		"intColumn":           {"nullable": "NO", "value": "int"},
		"nullIntColumn":       {"nullable": "YES", "value": "int"},
		"tinyIntColumn":       {"nullable": "NO", "value": "tinyint"},
		"nullTinyIntColumn":   {"nullable": "YES", "value": "tinyint"},
		"smallIntColumn":      {"nullable": "NO", "value": "smallint"},
		"nullSmallIntColumn":  {"nullable": "YES", "value": "smallint"},
		"mediumIntColumn":     {"nullable": "NO", "value": "mediumint"},
		"nullMediumIntColumn": {"nullable": "YES", "value": "mediumint"},
		"bigIntColumn":        {"nullable": "NO", "value": "bigint"},
		"nullBigIntColumn":    {"nullable": "YES", "value": "bigint"},
	}

	expectedStruct :=
		`package test

type testStruct struct {
	BigIntColumn        int64
	IntColumn           int
	MediumIntColumn     int
	NullBigIntColumn    sql.NullInt64
	NullIntColumn       sql.NullInt64
	NullMediumIntColumn sql.NullInt64
	NullSmallIntColumn  sql.NullInt64
	NullTinyIntColumn   sql.NullInt64
	SmallIntColumn      int
	TinyIntColumn       int
}
`

	columnsSorted := []string{
		"BigIntColumn",
		"IntColumn",
		"MediumIntColumn",
		"NullBigIntColumn",
		"NullIntColumn",
		"NullMediumIntColumn",
		"NullSmallIntColumn",
		"NullTinyIntColumn",
		"SmallIntColumn",
		"TinyIntColumn",
	}

	bytes, err := Generate(columnMap, columnsSorted, "test_table", "testStruct", "test", false, false, false, true)

	Convey("Should be able to generate map from string column", t, func() {
		So(err, ShouldBeNil)
		So(string(bytes), ShouldEqual, expectedStruct)
	})

	expectedStruct =
		`package test

type testStruct struct {
	BigIntColumn        int64
	IntColumn           int
	MediumIntColumn     int
	NullBigIntColumn    null.Int
	NullIntColumn       null.Int
	NullMediumIntColumn null.Int
	NullSmallIntColumn  null.Int
	NullTinyIntColumn   null.Int
	SmallIntColumn      int
	TinyIntColumn       int
}
`

	bytes, err = Generate(columnMap, columnsSorted, "test_table", "testStruct", "test", false, false, false, false)

	Convey("Should be able to generate map from string column", t, func() {
		So(err, ShouldBeNil)
		So(string(bytes), ShouldEqual, expectedStruct)
	})
}

func TestMysqlJSONStringGenerate(t *testing.T) {
	columnMap := map[string]map[string]string{
		"stringColumn":     {"nullable": "NO", "value": "varchar"},
		"nullStringColumn": {"nullable": "YES", "value": "varchar"},
	}

	expectedStruct :=
		`package test

type testStruct struct {
	StringColumn     string         ` + "`json:\"stringColumn\" comment:\"\"`     //" + `
	NullStringColumn sql.NullString ` + "`json:\"nullStringColumn\" comment:\"\"` //" + `
}
`
	columnsSorted := []string{
		"stringColumn",
		"nullStringColumn"}

	bytes, err := Generate(columnMap, columnsSorted, "test_table", "testStruct", "test", true, false, true, false)

	Convey("Should be able to generate map from string column", t, func() {
		So(err, ShouldBeNil)
		So(string(bytes), ShouldEqual, expectedStruct)
	})
}

func TestMysqlGormStringGenerate(t *testing.T) {
	expectedStruct :=
		`package test

type testStruct struct {
	StringColumn     string         ` + "`gorm:\"column:stringColumn\" comment:\"\"`     //" + `
	NullStringColumn sql.NullString ` + "`gorm:\"column:nullStringColumn\" comment:\"\"` //" + `
}

// TableName sets the insert table name for this struct type
func (t *testStruct) TableName() string {
	return "test_table"
}
`

	columnMap := map[string]map[string]string{
		"stringColumn":     {"nullable": "NO", "value": "varchar"},
		"nullStringColumn": {"nullable": "YES", "value": "varchar"},
	}

	columnsSorted := []string{
		"stringColumn",
		"nullStringColumn",
	}

	bytes, err := Generate(columnMap, columnsSorted, "test_table", "testStruct", "test", false, true, true, false)

	Convey("Should be able to generate map from string column", t, func() {
		So(err, ShouldBeNil)
		So(string(bytes), ShouldEqual, expectedStruct)
	})
}

func TestMysqlStringWithIntGenerate(t *testing.T) {
	expectedStruct :=
		`package test

type testStruct struct {
	OneStringColumn string ` + "`comment:\"\"` //" + `
}
`

	columnMap := map[string]map[string]string{
		"1stringColumn": {"nullable": "NO", "value": "varchar"},
	}

	columnsSorted := []string{
		"1stringColumn",
	}
	bytes, err := Generate(columnMap, columnsSorted, "test_table", "testStruct", "test", false, false, true, false)

	Convey("Should be able to generate map from string column", t, func() {
		So(err, ShouldBeNil)
		So(string(bytes), ShouldEqual, expectedStruct)
	})
}

func TestMysqlStringWithUnderscoresGenerate(t *testing.T) {
	expectedStruct :=
		`package test

type testStruct struct {
	OneStringColumn ` + "`comment:\"\"` //" + `
}
`

	columnMap := map[string]map[string]string{
		"string_Column": {"nullable": "NO", "value": "varchar"},
	}
	columnsSorted := []string{
		"1stringColumn",
	}
	bytes, err := Generate(columnMap, columnsSorted, "test_table", "testStruct", "test", false, false, true, false)

	Convey("Should be able to generate map from string column", t, func() {
		So(err, ShouldBeNil)
		So(string(bytes), ShouldEqual, expectedStruct)
	})
}

func TestMysqlStringWithCommonInitialismGenerate(t *testing.T) {
	expectedStruct :=
		`package test

type testStruct struct {
	API string ` + "`comment:\"\"` //" + `
}
`

	columnMap := map[string]map[string]string{
		"API": {"nullable": "NO", "value": "varchar"},
	}
	columnsSorted := []string{
		"API",
	}
	bytes, err := Generate(columnMap, columnsSorted, "test_table", "testStruct", "test", false, false, true, false)

	Convey("Should be able to generate map from string column", t, func() {
		So(err, ShouldBeNil)
		So(string(bytes), ShouldEqual, expectedStruct)
	})
}

// TestMysqlTypeToGureguType generates the struct and outputs nullable columns as guregu null types
func TestMysqlTypeToGureguType(t *testing.T) {
	expectedStruct :=
		`package test

type testStruct struct {
	VarChar   null.String ` + "`comment:\"\"` //" + `
	TinyInt   int         ` + "`comment:\"\"` //" + `
	Int       null.Int    ` + "`comment:\"\"` //" + `
	BigInt    int64       ` + "`comment:\"\"` //" + `
	Decimal   null.Float  ` + "`comment:\"\"` //" + `
	Float     null.Float  ` + "`comment:\"\"` //" + `
	Double    float64     ` + "`comment:\"\"` //" + `
	DateTime  null.Time   ` + "`comment:\"\"` //" + `
	Time      time.Time   ` + "`comment:\"\"` //" + `
	Date      null.Time   ` + "`comment:\"\"` //" + `
	TimeStamp null.Time   ` + "`comment:\"\"` //" + `
}
`

	columnMap := map[string]map[string]string{
		"VarChar":   {"nullable": "YES", "value": "varchar"},
		"TinyInt":   {"nullable": "NO", "value": "tinyint"},
		"Int":       {"nullable": "YES", "value": "int"},
		"BigInt":    {"nullable": "NO", "value": "bigint"},
		"Decimal":   {"nullable": "YES", "value": "decimal"},
		"Float":     {"nullable": "YES", "value": "float"},
		"Double":    {"nullable": "NO", "value": "double"},
		"DateTime":  {"nullable": "YES", "value": "datetime"},
		"Time":      {"nullable": "NO", "value": "time"},
		"Date":      {"nullable": "YES", "value": "date"},
		"TimeStamp": {"nullable": "YES", "value": "timestamp"},
	}

	columnsSorted := []string{
		"VarChar",
		"TinyInt",
		"Int",
		"BigInt",
		"Decimal",
		"Float",
		"Double",
		"DateTime",
		"Time",
		"Date",
		"TimeStamp",
	}
	bytes, err := Generate(columnMap, columnsSorted, "test_table", "testStruct", "test", false, false, true, true)

	Convey("Should be able to generate map for guregu types", t, func() {
		So(err, ShouldBeNil)
		So(string(bytes), ShouldEqual, expectedStruct)
	})
}
