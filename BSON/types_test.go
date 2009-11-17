package bson

import (
	"testing";
)

func TestBsonInt(t *testing.T) {
	var testVar BsonType;
	var num BsonNumber;
	num = 0.0;
	testVar = num;
	typeNumber := testVar.TypeNumber();
	if typeNumber != Number {
		t.Errorf("Int failed, Expecting %d. Got %d", Number, typeNumber)
	}
}
