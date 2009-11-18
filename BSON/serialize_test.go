package bson

import (
  "testing"
)

func TestSerialize(t *testing.T) {
  m := make(map[string]interface{});
  m["test"] = 45;
  _,error := Serialize(m);
  if error != nil {
    t.Errorf("Fail: %s",error);
  }
}


