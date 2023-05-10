package main

import (
	//"log"
	//"os/exec"
	"fmt"
	//"strconv"
	//"strings"
	"testing"
	//"bytes"
	//"io"
	//"bufio"
	//"os"
)

var prog = "./es1-foresta"

func TestMainEsempioConFile(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(t,
		prog,
		"input-es.txt",
		"output-es.txt",
		"NIENTE")
}

func TestStampaAlbero(t *testing.T) {
	mappa := make(map[string]*oggetto)

	tavolo := &oggetto{nome: "tavolo", tipo: "op", dx: "cassetto", op: '+', sx: "mensola"}
	cassetto := &oggetto{nome: "cassetto", tipo: "num", val: 3}
	mensola := &oggetto{nome: "mensola", tipo: "num", val: 13}

	frigorifero := &oggetto{nome: "frigorifero", tipo: "op", dx: "tazza", op: '*', sx: "tavolo"}
	tazza := &oggetto{nome: "tazza", tipo: "num", val: 2}

	mappa["tavolo"] = tavolo
	mappa["cassetto"] = cassetto
	mappa["mensola"] = mensola
	mappa["frigorifero"] = frigorifero
	mappa["tazza"] = tazza

	f := costruisciForesta(mappa)

	output := CaptureOutput(stampaAlbero, f, "frigorifero")
	expected := "mensola (val = 13)\ntavolo\ncassetto (val = 3)\nfrigorifero\ntazza (val = 2)\n"

	if checkOutput(output, expected) == false {
		t.Fail()
	}

}

func TestSxDxUp(t *testing.T) {
	mappa := make(map[string]*oggetto)

	tavolo := &oggetto{nome: "tavolo", tipo: "op", dx: "cassetto", op: '+', sx: "mensola"}
	cassetto := &oggetto{nome: "cassetto", tipo: "num", val: 3}
	mensola := &oggetto{nome: "mensola", tipo: "num", val: 13}
	mappa["tavolo"] = tavolo
	mappa["cassetto"] = cassetto
	mappa["mensola"] = mensola

	f := costruisciForesta(mappa)
	var figliodx, figliosx, padre string
	var res, res2 bool

	figliodx, _ = dx(f, "tavolo")
	figliosx, _ = sx(f, "tavolo")
	padre, _ = up(f, "mensola")
	_, res = up(f, "tavolo")
	_, res2 = sx(f, "mensola")

	output := fmt.Sprintln(figliodx, figliosx, padre, res, res2)

	expected := "cassetto mensola tavolo false false"

	if checkOutput(output[:len(output)-1], expected) == false {
		t.Fail()
	}

}
