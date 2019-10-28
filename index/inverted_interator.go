package main

type InvertedIterator interface {
	HasNext() bool
	Next() interface{}
	GetGE(key interface{}) interface{}
}
