package orderbook

type Side string

const (
	Buy  Side = "buy"
	Sell Side = "sell"
)

type Order struct {
	ID         string
	Side       Side
	PriceCents int64
	Quantity   int64
}

type Trade struct {
	BuyOrderID  string
	SellOrderID string
	PriceCents  int64
	Quantity    int64
}

func MatchOrders(orders []Order) []Trade {
	trade := []Trade{}
	buyOrders := []Order{}
	sellOrders := []Order{}

	for _, order := range orders {
		if order.Side == Buy {
			for order.Quantity > 0 {
				bestIndex := -1
				for idx, sellOrder := range sellOrders {
					if sellOrder.PriceCents > order.PriceCents {
						continue
					}

					if bestIndex == -1 || sellOrder.PriceCents < sellOrders[bestIndex].PriceCents {
						bestIndex = idx
					}
				}

				if bestIndex == -1 {
					break
				}

				minQuantity := min(order.Quantity, sellOrders[bestIndex].Quantity)

				trade = append(trade, Trade{
					BuyOrderID:  order.ID,
					SellOrderID: sellOrders[bestIndex].ID,
					PriceCents:  sellOrders[bestIndex].PriceCents,
					Quantity:    minQuantity,
				})
				sellOrders[bestIndex].Quantity -= minQuantity
				order.Quantity -= minQuantity
				if sellOrders[bestIndex].Quantity < 1 {
					sellOrders = append(sellOrders[:bestIndex], sellOrders[bestIndex+1:]...)
				}
			}
			if order.Quantity > 0 {
				buyOrders = append(buyOrders, order)
			}
		}

		if order.Side == Sell {
			for order.Quantity > 0 {
				bestIndex := -1
				for idx, buyOrder := range buyOrders {
					if buyOrder.PriceCents < order.PriceCents {
						continue
					}

					if bestIndex == -1 || buyOrder.PriceCents > buyOrders[bestIndex].PriceCents {
						bestIndex = idx
					}
				}

				if bestIndex == -1 {
					break
				}

				minQuantity := min(order.Quantity, buyOrders[bestIndex].Quantity)

				trade = append(trade, Trade{
					BuyOrderID:  buyOrders[bestIndex].ID,
					SellOrderID: order.ID,
					PriceCents:  buyOrders[bestIndex].PriceCents,
					Quantity:    minQuantity,
				})
				buyOrders[bestIndex].Quantity -= minQuantity
				order.Quantity -= minQuantity
				if buyOrders[bestIndex].Quantity < 1 {
					buyOrders = append(buyOrders[:bestIndex], buyOrders[bestIndex+1:]...)
				}

			}
			if order.Quantity > 0 {
				sellOrders = append(sellOrders, order)
			}
		}
	}

	return trade
}
