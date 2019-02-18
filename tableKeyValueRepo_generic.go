//go:generate genny -in=$GOFILE -out=./TableKeyValueRepo.go gen "TypeRepo=TableKeyValueRepo TypeTableKey=string TypeItemKey=string TypeValue=string"

package main

import (
	"encoding/json"
	"fmt"

	"github.com/cheekybits/genny/generic"
)

type TypeTableKey generic.Type
type TypeItemKey generic.Type
type TypeValue generic.Type

type TypeRepo struct {
	Data          map[TypeTableKey]map[TypeItemKey]TypeValue
	GenerateNewID func(map[TypeItemKey]TypeValue) TypeItemKey `json:",omitempty"`
}

func NewTypeRepo() TypeRepo {
	r := TypeRepo{
		Data: make(map[TypeTableKey]map[TypeItemKey]TypeValue),
	}
	r.GenerateNewID = r.generateNewItemKey
	return r
}

func (repo TypeRepo) set(tableKey TypeTableKey, itemKey TypeItemKey, value TypeValue) {
	table := repo.getTable(tableKey)
	table[itemKey] = value
}

func (repo TypeRepo) append(tableKey TypeTableKey, value TypeValue) {
	table := repo.getTable(tableKey)
	newItemKey := repo.GenerateNewID(table)
	repo.set(tableKey, newItemKey, value)
}

func (repo TypeRepo) get(tableKey TypeTableKey, itemKey TypeItemKey) (TypeValue, error) {
	table, ok := repo.Data[tableKey]
	if !ok {
		return "", fmt.Errorf("Do not have a table for %v", tableKey)
	}
	item, ok := table[itemKey]
	if !ok {
		return "", fmt.Errorf("No item in %v with index %v", tableKey, itemKey)
	}
	return item, nil
}

func (repo TypeRepo) getAll(tableKey TypeTableKey) (map[TypeItemKey]TypeValue, error) {
	tbl, ok := repo.Data[tableKey]
	if !ok {
		return nil, fmt.Errorf("Do not have a table for %v", tableKey)
	}
	return tbl, nil
}

func (repo TypeRepo) getTable(tableKey TypeItemKey) map[TypeItemKey]TypeValue {
	tbl, ok := repo.Data[tableKey]
	if !ok {
		tbl = repo.newTable()
		repo.Data[tableKey] = tbl
	}
	return tbl
}

func (repo TypeRepo) tables() []TypeTableKey {
	keys := make([]TypeTableKey, len(repo.Data))
	i := 0
	for k := range repo.Data {
		keys[i] = k
		i++
	}
	return keys
}

func (repo TypeRepo) generateNewItemKey(table map[TypeItemKey]TypeValue) TypeItemKey {
	i := 1
	for i = 1; i > 0; i++ {
		_, ok := table[fmt.Sprintf("%v", i)]
		if !ok {
			break
		}
	}
	return fmt.Sprintf("%v", i)
}

func (repo TypeRepo) newTable() map[TypeItemKey]TypeValue {
	return make(map[TypeItemKey]TypeValue)
}

func (repo TypeRepo) String() string {
	return fmt.Sprintf("%v", repo.Data)
}

func (repo TypeRepo) MarshalJSON() ([]byte, error) {
	return json.Marshal(repo.Data)
}
