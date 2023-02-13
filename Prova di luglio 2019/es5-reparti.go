package main

type Bit_node struct {
	n_prog     int
	lavoratori int
	dx, sx     *Bit_node
}

type foresta struct {
	impianti []*Bit_node
}

func lavoratoriCollaudo(root *Bit_node) int {
	if root == nil {
		return 0
	}
	if root.dx == nil && root.sx == nil {
		k := root.lavoratori
		return k
	}
	return lavoratoriCollaudo(root.sx) + lavoratoriCollaudo((root.dx))
}

func totCollaudo(f []foresta) int {
	somma := 0
	for i := 0; i < len(f); i++ {
		somma += lavoratoriCollaudo(f.impianti[i])
	}
}
