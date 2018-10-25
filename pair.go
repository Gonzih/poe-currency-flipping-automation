package main

import "sort"

func getPairs() []Pair {

	pairs := make([]Pair, 0)

	if len(currenciesToScan) == 0 {
		for name := range currencyNames {
			currenciesToScan = append(currenciesToScan, name)
		}
	}

	for _, name := range currenciesToScan {
		switch name {
		case "wisdom", "chaos":
			continue
		default:
		}

		offers1 := searchFor(onlineSearch, "Delve", name, "chaos")
		offers2 := searchFor(onlineSearch, "Delve", "chaos", name)

		for _, offer1 := range offers1 {
			p := Pair{of: offer1, ofs: make([]SearchOffer, 0)}

			for _, offer2 := range offers2 {
				if offer1.IsProfitable(offer2) {
					p.ofs = append(p.ofs, offer2)
				}
			}

			if len(p.ofs) > 0 {
				pairs = append(pairs, p)
			}
		}
	}

	sort.Slice(pairs, func(i, j int) bool { return pairs[i].MaxProfit() > pairs[j].MaxProfit() })

	return pairs
}
