package main

type InvertedIndex interface {
	HasNext() bool
	Next() DocId
	GetGE(id DocId) DocId
}
