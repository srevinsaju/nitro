package skiplist

type Iterator struct {
	s          *Skiplist
	prev, curr *Node
	valid      bool
}

func (s *Skiplist) NewIterator() *Iterator {
	return &Iterator{
		s: s,
	}
}

func (it *Iterator) SeekFirst() {
	it.prev = it.s.head
	it.curr, _ = it.s.head.getNext(0)
	it.valid = true
}

func (it *Iterator) Seek(itm Item) {
	it.valid = true
	preds, succs, found := it.s.findPath(itm)
	it.prev = preds[0]
	it.curr = succs[0]
	if !found {
		it.valid = false
	}
}

func (it *Iterator) Valid() bool {
	if it.valid && it.curr == it.s.tail {
		it.valid = false
	}

	return it.valid
}

func (it *Iterator) Get() Item {
	return it.curr.itm
}

func (it *Iterator) Next() {
retry:
	it.valid = true
	next, deleted := it.curr.getNext(0)
	for deleted {
		if !it.s.helpDelete(0, it.prev, it.curr, next) {
			preds, succs, found := it.s.findPath(it.curr.itm)
			last := it.curr
			it.prev = preds[0]
			it.curr = succs[0]
			if found && last == it.curr {
				goto retry
			} else {
				return
			}
		}
		it.curr, _ = it.prev.getNext(0)
		next, deleted = it.curr.getNext(0)
	}

	it.prev = it.curr
	it.curr = next
}