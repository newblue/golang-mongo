package bson

type BsonTypeId int

const (
	Number	BsonTypeId	= 1;
	String	BsonTypeId	= iota + 1;
	Obj;
	Array;
	Binary;
	Undefined;
	Oid;
	Boolean;
	Date;
	Null;
	Regex;
	Reference;
	Code;
	Symbol;
	CodeWScope;
	Integer;
	Timestamp;
	Long;
	MinKey	BsonTypeId	= -1;
	MaxKey	BsonTypeId	= 127;
)

type BsonType interface {
	TypeNumber() BsonTypeId;
}

type BsonNumber float

func (num BsonNumber) TypeNumber() BsonTypeId	{ return Number }
