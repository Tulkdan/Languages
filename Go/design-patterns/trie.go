package main

type Trie struct {
    children map[rune]*Trie
    isLast   bool
}

func Constructor() Trie {
    return Trie{children: make(map[rune]*Trie), isLast: false}
}

func (this *Trie) Insert(word string) {
    current := this

    for _, c := range word {
        if _, ok := current.children[c]; !ok {
            current.children[c] = &Trie{children: map[rune]*Trie{}}
        }
        current = current.children[c]
    }

    current.isLast = true
}

func (this *Trie) Search(word string) bool {
    current := this

    for _, c := range word {
        if _, ok := current.children[c]; !ok {
            return false
        }
        current = current.children[c]
    }

    return current.isLast
}

func (this *Trie) StartsWith(prefix string) bool {
    current := this

    for _, c := range prefix {
        if _, ok := current.children[c]; !ok {
            return false
        }
        current = current.children[c]
    }

    return true
}

