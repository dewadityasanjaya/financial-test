package lru

import "container/list"

type entry struct {
	symbol     string
	priceCents int64
}

type Cache struct {
	capacity int
	items    map[string]*list.Element
	order    *list.List
}

func NewCache(capacity int) *Cache {
	return &Cache{
		capacity: capacity,
		items:    make(map[string]*list.Element),
		order:    list.New(),
	}
}

func (c *Cache) Get(symbol string) (int64, bool) {
	element, ok := c.items[symbol]
	if !ok {
		return 0, false
	}

	c.order.MoveToFront(element)
	item := element.Value.(entry)
	return item.priceCents, true
}

func (c *Cache) Put(symbol string, priceCents int64) {
	if c.capacity <= 0 {
		return
	}

	if element, ok := c.items[symbol]; ok {
		element.Value = entry{symbol: symbol, priceCents: priceCents}
		c.order.MoveToFront(element)
		return
	}

	element := c.order.PushFront(entry{symbol: symbol, priceCents: priceCents})
	c.items[symbol] = element

	if len(c.items) > c.capacity {
		oldest := c.order.Back()
		if oldest == nil {
			return
		}

		item := oldest.Value.(entry)
		delete(c.items, item.symbol)
		c.order.Remove(oldest)
	}
}
