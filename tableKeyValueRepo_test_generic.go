//go:generate genny -in=$GOFILE -out=./TableKeyValueRepo_test.go gen "TypeRepo=TableKeyValueRepo TypeTableKey=string TypeItemKey=string TypeValue=string GENR=_"

package main

import (
	"encoding/json"
	"testing"
)

const key0GENR = "key1"
const key2GENR = "key2"
const key3GENR = "key3"
const val1GENR = "val1"
const val2GENR = "val2"
const val3GENR = "val3"

func Test_GetBeforeSet_ReturnsErrGENR(t *testing.T) {
	repo := NewTypeRepo()
	_, err := repo.get(key0GENR, key2GENR)
	if err == nil {
		t.Errorf("Expected getting value before setting to return error, got nil")
		return
	}
}

func Test_SetThenGet_SucceedsGENR(t *testing.T) {
	repo := NewTypeRepo()
	repo.set(key0GENR, key2GENR, val1GENR)
	item, err := repo.get(key0GENR, key2GENR)
	if err != nil {
		t.Errorf("Expected to get a value without error, got: " + err.Error())
	}
	if item != val1GENR {
		t.Errorf("Expected to value to match what was set ('%v'), got: %v", item, val1GENR)
	}
}

func Test_SettingDifferentKeys_StoreSeparateValsGENR(t *testing.T) {
	repo := NewTypeRepo()
	repo.set(key0GENR, key2GENR, val1GENR)
	repo.set(key0GENR, key3GENR, val2GENR)
	repo.set(key2GENR, key2GENR, val3GENR)
	item, err := repo.get(key0GENR, key2GENR)
	if err != nil {
		t.Errorf("Expected to get a value without error, got: " + err.Error())
	}
	if item != val1GENR {
		t.Errorf("Expected to value to match what was set ('%v'), got: %v", item, val1GENR)
	}
	item, err = repo.get(key0GENR, key3GENR)
	if err != nil {
		t.Errorf("Expected to get a value without error, got: " + err.Error())
	}
	if item != val2GENR {
		t.Errorf("Expected to value to match what was set ('%v'), got: %v", item, val2GENR)
	}
	item, err = repo.get(key2GENR, key2GENR)
	if err != nil {
		t.Errorf("Expected to get a value without error, got: " + err.Error())
	}
	if item != val3GENR {
		t.Errorf("Expected to value to match what was set ('%v'), got: %v", item, val3GENR)
	}
}

func Test_SecondSet_ReplacesValueGENR(t *testing.T) {
	repo := NewTypeRepo()
	repo.set(key0GENR, key2GENR, val1GENR)
	repo.set(key0GENR, key2GENR, val2GENR)
	item, err := repo.get(key0GENR, key2GENR)
	if err != nil {
		t.Errorf("Expected to get a value without error, got: " + err.Error())
	}
	if item != val2GENR {
		t.Errorf("Expected to value to match what was set ('%v'), got: %v", item, val2GENR)
	}
}

func Test_Append_AppendsGENR(t *testing.T) {
	//TODO - not very clean with GCONST
	repo := NewTypeRepo()
	repo.append(key0GENR, val1GENR)
	item, err := repo.get(key0GENR, "1")
	if err != nil {
		t.Errorf("Expected to get a value without error, got: " + err.Error())
	}
	if item != val1GENR {
		t.Errorf("Expected to value to match what was set ('%v'), got: %v", item, val1GENR)
	}

	repo.append(key0GENR, val2GENR)
	item, err = repo.get(key0GENR, "2")
	if err != nil {
		t.Errorf("Expected to get a value without error, got: " + err.Error())
	}
	if item != val2GENR {
		t.Errorf("Expected to value to match what was set ('%v'), got: %v", item, val2GENR)
	}
}

func Test_JSONMarshal_SucceedsGENR(t *testing.T) {
	repo := NewTypeRepo()
	repo.append(key0GENR, val1GENR)
	_, err := json.Marshal(repo)
	if err != nil {
		t.Errorf("Could not marshal data structure:: " + err.Error())
	}
}
