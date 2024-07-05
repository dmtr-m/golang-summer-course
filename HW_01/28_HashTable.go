package main

import "fmt"

var Prime_1 int = 5
var Prime_2 int = 13

func Hash(str string) int {
	var hash int = 0
	var multiplier = 1

	for _, letter := range str {
		hash += int(letter) * multiplier
		multiplier *= Prime_2
		hash %= Prime_1
	}

	return hash
}

type KeyValue struct {
	key   string
	value int
}

func Add(table [][]KeyValue, key string, value int) {
	hash := Hash(key)

	for _, key_value := range table[hash] {
		if key_value.key == key {
			return
		}
	}

	table[hash] = append(table[hash], KeyValue{key, value})
}

func Get(table [][]KeyValue, key string, if_none int) int {
	hash := Hash(key)

	for _, key_value := range table[hash] {
		if key_value.key == key {
			return key_value.value
		}
	}

	return if_none
}

func Erase(table [][]KeyValue, key string) {
	hash := Hash(key)

	for i, key_value := range table[hash] {
		if key_value.key == key {
			table[hash][i], table[hash][len(table[hash])-1] = table[hash][len(table[hash])-1], table[hash][i]
			table[hash] = table[hash][:len(table[hash])-1]
			return
		}
	}
}

func TestHashTable() {
	var hash_table = make([][]KeyValue, Prime_1)

	Add(hash_table, "baa", 1)
	Add(hash_table, "aab", 1)
	Add(hash_table, "bac", 1)
	Add(hash_table, "aad", 1)
	Add(hash_table, "bca", 1)
	Add(hash_table, "acb", 1)
	Add(hash_table, "bcc", 1)
	Add(hash_table, "acd", 1)
	fmt.Println(hash_table)
	fmt.Println(Get(hash_table, "aab", 0))
	Erase(hash_table, "aab")
	fmt.Println(hash_table)

	fmt.Println(hash_table)
}
