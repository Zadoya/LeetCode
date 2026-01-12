// условие задачи - https://leetcode.com/problems/design-hashset/description/

package main

type MyHashSet struct {
    payload [125001] byte
}


func Constructor() MyHashSet {
    return MyHashSet{}
}


func (this *MyHashSet) Add(key int)  {
	this.payload[key / 8] |= 1 << (key % 8) 
}


func (this *MyHashSet) Remove(key int)  {
    this.payload[key / 8] &= ^(1 << (key % 8))
}


func (this *MyHashSet) Contains(key int) bool {
	return this.payload[key / 8] & (1 << (key % 8)) != 0
}