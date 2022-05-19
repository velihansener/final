package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Info struct {
	InfoId   int
	InfoName string
	InfoType string
}

type Cigkofte struct {
	Info
	CgSatici string
	CgGram   int
	CgFiyat  float32
}

type Kebap struct {
	Info
	KbSatici   string
	KbPorsiyon float32
	KbFiyat    float32
}

type Lahmacun struct {
	Info
	LhSatici string
	LhTip    string
	LhAdet   int
	LhFiyat  float32
}

type Pizza struct {
	Info
	PzSatici string
	PzTip    string
	PzBoy    string
	PzAdet   int
	PzFiyat  float32
}

type Susi struct {
	Info
	SsSatici string
	SsTip    string
	SsAdet   int
	SsFiyat  float32
}

type EtBurger struct {
	Info
	BrSatici string
	BrKofte  int
	BrAdet   int
	BrFiyat  float32
}

type TavukBurger struct {
	Info
	BgSatici string
	BgKofte  int
	BgAdet   int
	BgFiyat  float32
}

func bilgiler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Kişisel bilgiler")
	staticPage := staticPages.Lookup("bilgiler.html")
	staticPage.Execute(w, nil)

	cg1 := Cigkofte{
		Info:     Info{InfoId: 1, InfoName: "Çiğköfte", InfoType: "Türk Mutfağı"},
		CgSatici: "Komagene", CgGram: 200, CgFiyat: 37.5,
	}
	cg2 := Cigkofte{
		Info:     Info{InfoId: 1, InfoName: "Çiğköfte", InfoType: "Türk Mutfağı"},
		CgSatici: "Komagene", CgGram: 400, CgFiyat: 67.0,
	}
	cg3 := Cigkofte{
		Info:     Info{InfoId: 1, InfoName: "Çiğköfte", InfoType: "Türk Mutfağı"},
		CgSatici: "Komagene", CgGram: 600, CgFiyat: 83.0,
	}
	cg4 := Cigkofte{
		Info:     Info{InfoId: 1, InfoName: "Çiğköfte", InfoType: "Türk Mutfağı"},
		CgSatici: "Adıyaman Lezzet", CgGram: 200, CgFiyat: 32.5,
	}
	cg5 := Cigkofte{
		Info:     Info{InfoId: 1, InfoName: "Çiğköfte", InfoType: "Türk Mutfağı"},
		CgSatici: "Adıyaman Lezzet", CgGram: 400, CgFiyat: 55,
	}
	cg6 := Cigkofte{
		Info:     Info{InfoId: 1, InfoName: "Çiğköfte", InfoType: "Türk Mutfağı"},
		CgSatici: "Adıyaman Lezzet", CgGram: 600, CgFiyat: 75,
	}
	cg7 := Cigkofte{
		Info:     Info{InfoId: 1, InfoName: "Çiğköfte", InfoType: "Türk Mutfağı"},
		CgSatici: "Battalbey", CgGram: 200, CgFiyat: 35.5,
	}
	cg8 := Cigkofte{
		Info:     Info{InfoId: 1, InfoName: "Çiğköfte", InfoType: "Türk Mutfağı"},
		CgSatici: "Battalbey", CgGram: 400, CgFiyat: 57,
	}
	cg9 := Cigkofte{
		Info:     Info{InfoId: 1, InfoName: "Çiğköfte", InfoType: "Türk Mutfağı"},
		CgSatici: "Battalbey", CgGram: 600, CgFiyat: 77,
	}

	var cgList []Cigkofte
	cgList = append(cgList, cg1, cg2, cg3, cg4, cg5, cg6, cg7, cg8, cg9)
	fmt.Println(cgList)

	kb1 := Kebap{
		Info:     Info{InfoId: 2, InfoName: "Kebap", InfoType: "Türk Mutfağı"},
		KbSatici: "Hacıoğulları Kebap", KbPorsiyon: 1, KbFiyat: 40,
	}
	kb2 := Kebap{
		Info:     Info{InfoId: 2, InfoName: "Kebap", InfoType: "Türk Mutfağı"},
		KbSatici: "Hacıoğulları Kebap", KbPorsiyon: 1.5, KbFiyat: 60,
	}
	kb3 := Kebap{
		Info:     Info{InfoId: 2, InfoName: "Kebap", InfoType: "Türk Mutfağı"},
		KbSatici: "Hacıoğulları Kebap", KbPorsiyon: 2, KbFiyat: 80,
	}
	kb4 := Kebap{
		Info:     Info{InfoId: 2, InfoName: "Kebap", InfoType: "Türk Mutfağı"},
		KbSatici: "Ocakbaşı", KbPorsiyon: 1, KbFiyat: 36,
	}
	kb5 := Kebap{
		Info:     Info{InfoId: 2, InfoName: "Kebap", InfoType: "Türk Mutfağı"},
		KbSatici: "Ocakbaşı", KbPorsiyon: 1.5, KbFiyat: 54,
	}
	kb6 := Kebap{
		Info:     Info{InfoId: 2, InfoName: "Kebap", InfoType: "Türk Mutfağı"},
		KbSatici: "Ocakbaşı", KbPorsiyon: 2, KbFiyat: 75,
	}
	kb7 := Kebap{
		Info:     Info{InfoId: 2, InfoName: "Kebap", InfoType: "Türk Mutfağı"},
		KbSatici: "Kuzey Kebap", KbPorsiyon: 1, KbFiyat: 32,
	}
	kb8 := Kebap{
		Info:     Info{InfoId: 2, InfoName: "Kebap", InfoType: "Türk Mutfağı"},
		KbSatici: "Kuzey Kebap", KbPorsiyon: 1.5, KbFiyat: 45,
	}
	kb9 := Kebap{
		Info:     Info{InfoId: 2, InfoName: "Kebap", InfoType: "Türk Mutfağı"},
		KbSatici: "Kuzey Kebap", KbPorsiyon: 2, KbFiyat: 60,
	}

	var kbList []Kebap
	kbList = append(kbList, kb1, kb2, kb3, kb4, kb5, kb6, kb7, kb8, kb9)
	fmt.Println(kbList)

	lh1 := Lahmacun{
		Info:     Info{InfoId: 3, InfoName: "Lahmacun", InfoType: "Türk Mutfağı"},
		LhSatici: "Ocakbaşı", LhTip: "Acısız", LhAdet: 1, LhFiyat: 15,
	}
	lh2 := Lahmacun{
		Info:     Info{InfoId: 3, InfoName: "Lahmacun", InfoType: "Türk Mutfağı"},
		LhSatici: "Ocakbaşı", LhTip: "Acısız", LhAdet: 4, LhFiyat: 60,
	}
	lh3 := Lahmacun{
		Info:     Info{InfoId: 3, InfoName: "Lahmacun", InfoType: "Türk Mutfağı"},
		LhSatici: "Ocakbaşı", LhTip: "Acılı", LhAdet: 2, LhFiyat: 30,
	}
	lh4 := Lahmacun{
		Info:     Info{InfoId: 3, InfoName: "Lahmacun", InfoType: "Türk Mutfağı"},
		LhSatici: "Lezzet Lahmacun", LhTip: "Acılı", LhAdet: 1, LhFiyat: 12,
	}
	lh5 := Lahmacun{
		Info:     Info{InfoId: 3, InfoName: "Lahmacun", InfoType: "Türk Mutfağı"},
		LhSatici: "Lezzet Lahmacun", LhTip: "Acısız", LhAdet: 2, LhFiyat: 24,
	}
	lh6 := Lahmacun{
		Info:     Info{InfoId: 3, InfoName: "Lahmacun", InfoType: "Türk Mutfağı"},
		LhSatici: "Karadeniz Pide", LhTip: "Acısız", LhAdet: 1, LhFiyat: 17,
	}
	lh7 := Lahmacun{
		Info:     Info{InfoId: 3, InfoName: "Lahmacun", InfoType: "Türk Mutfağı"},
		LhSatici: "Karadeniz Pide", LhTip: "Acılı", LhAdet: 3, LhFiyat: 50,
	}

	var lhList []Lahmacun
	lhList = append(lhList, lh1, lh2, lh3, lh4, lh5, lh6, lh7)
	fmt.Println(lhList)

	pz1 := Pizza{
		Info:     Info{InfoId: 4, InfoName: "Pizza", InfoType: "Dünya Mutfağı"},
		PzSatici: "Dominos", PzTip: "Sosisli", PzBoy: "Küçük Boy", PzAdet: 1, PzFiyat: 35,
	}
	pz2 := Pizza{
		Info:     Info{InfoId: 4, InfoName: "Pizza", InfoType: "Dünya Mutfağı"},
		PzSatici: "Dominos", PzTip: "Sucuklu", PzBoy: "Orta Boy", PzAdet: 1, PzFiyat: 40,
	}
	pz3 := Pizza{
		Info:     Info{InfoId: 4, InfoName: "Pizza", InfoType: "Dünya Mutfağı"},
		PzSatici: "Terra Pizza", PzTip: "Sucuklu", PzBoy: "Büyük Boy", PzAdet: 1, PzFiyat: 75,
	}
	pz4 := Pizza{
		Info:     Info{InfoId: 4, InfoName: "Pizza", InfoType: "Dünya Mutfağı"},
		PzSatici: "Terra Pizza", PzTip: "Karışık", PzBoy: "Orta Boy", PzAdet: 1, PzFiyat: 50,
	}
	pz5 := Pizza{
		Info:     Info{InfoId: 4, InfoName: "Pizza", InfoType: "Dünya Mutfağı"},
		PzSatici: "Sbarro", PzTip: "Sucuklu", PzBoy: "Büyük Boy", PzAdet: 1, PzFiyat: 80,
	}
	pz6 := Pizza{
		Info:     Info{InfoId: 4, InfoName: "Pizza", InfoType: "Dünya Mutfağı"},
		PzSatici: "Sbarro", PzTip: "Karışık", PzBoy: "Küçük Boy", PzAdet: 1, PzFiyat: 40,
	}

	var pzList []Pizza
	pzList = append(pzList, pz1, pz2, pz3, pz4, pz5, pz6)
	fmt.Println(pzList)

	ss1 := Susi{
		Info:     Info{InfoId: 5, InfoName: "Suşi", InfoType: "Dünya Mutfağı"},
		SsSatici: "SushiCo", SsTip: "Somon", SsAdet: 3, SsFiyat: 55,
	}
	ss2 := Susi{
		Info:     Info{InfoId: 5, InfoName: "Suşi", InfoType: "Dünya Mutfağı"},
		SsSatici: "SushiCo", SsTip: "Levrek", SsAdet: 3, SsFiyat: 50,
	}
	ss3 := Susi{
		Info:     Info{InfoId: 5, InfoName: "Suşi", InfoType: "Dünya Mutfağı"},
		SsSatici: "Red Dragon", SsTip: "Somon", SsAdet: 3, SsFiyat: 65,
	}
	ss4 := Susi{
		Info:     Info{InfoId: 5, InfoName: "Suşi", InfoType: "Dünya Mutfağı"},
		SsSatici: "Red Dragon", SsTip: "Levrek", SsAdet: 3, SsFiyat: 60,
	}

	var ssList []Susi
	ssList = append(ssList, ss1, ss2, ss3, ss4)
	fmt.Println(ssList)

	br1 := EtBurger{
		Info:     Info{InfoId: 6, InfoName: "Et Burger", InfoType: "Burger"},
		BrSatici: "Burger King", BrKofte: 1, BrAdet: 1, BrFiyat: 55,
	}
	br2 := EtBurger{
		Info:     Info{InfoId: 6, InfoName: "Et Burger", InfoType: "Burger"},
		BrSatici: "Burger King", BrKofte: 2, BrAdet: 1, BrFiyat: 70,
	}
	br3 := EtBurger{
		Info:     Info{InfoId: 6, InfoName: "Et Burger", InfoType: "Burger"},
		BrSatici: "Atmaca Burger", BrKofte: 1, BrAdet: 1, BrFiyat: 50,
	}
	br4 := EtBurger{
		Info:     Info{InfoId: 6, InfoName: "Et Burger", InfoType: "Burger"},
		BrSatici: "Atmaca Burger", BrKofte: 2, BrAdet: 1, BrFiyat: 67.5,
	}
	br5 := EtBurger{
		Info:     Info{InfoId: 6, InfoName: "Et Burger", InfoType: "Burger"},
		BrSatici: "Doyos Burger", BrKofte: 1, BrAdet: 1, BrFiyat: 40,
	}
	br6 := EtBurger{
		Info:     Info{InfoId: 6, InfoName: "Et Burger", InfoType: "Burger"},
		BrSatici: "Doyos Burger", BrKofte: 2, BrAdet: 1, BrFiyat: 52.5,
	}

	var brList []EtBurger
	brList = append(brList, br1, br2, br3, br4, br5, br6)
	fmt.Println(brList)

	bg1 := TavukBurger{
		Info:     Info{InfoId: 7, InfoName: "Tavuk Burger", InfoType: "Burger"},
		BgSatici: "Burger Kıng", BgKofte: 1, BgAdet: 1, BgFiyat: 45,
	}
	bg2 := TavukBurger{
		Info:     Info{InfoId: 7, InfoName: "Tavuk Burger", InfoType: "Burger"},
		BgSatici: "Burger Kıng", BgKofte: 2, BgAdet: 1, BgFiyat: 65,
	}
	bg3 := TavukBurger{
		Info:     Info{InfoId: 7, InfoName: "Tavuk Burger", InfoType: "Burger"},
		BgSatici: "Popeyes", BgKofte: 1, BgAdet: 1, BgFiyat: 47.5,
	}
	bg4 := TavukBurger{
		Info:     Info{InfoId: 7, InfoName: "Tavuk Burger", InfoType: "Burger"},
		BgSatici: "Popeyes", BgKofte: 2, BgAdet: 1, BgFiyat: 67.5,
	}
	bg5 := TavukBurger{
		Info:     Info{InfoId: 7, InfoName: "Tavuk Burger", InfoType: "Burger"},
		BgSatici: "KFC", BgKofte: 1, BgAdet: 1, BgFiyat: 50,
	}
	bg6 := TavukBurger{
		Info:     Info{InfoId: 7, InfoName: "Tavuk Burger", InfoType: "Burger"},
		BgSatici: "KFC", BgKofte: 2, BgAdet: 1, BgFiyat: 75,
	}

	var bgList []TavukBurger
	bgList = append(bgList, bg1, bg2, bg3, bg4, bg5, bg6)
	fmt.Println(bgList)

	w.Header().Set("Content-Type", "application/json")
	cgjsonString, err := json.Marshal(cgList)
	checkErr(err)

	//w.Header().Set("Content-Type", "application/json")
	kbjsonString, err := json.Marshal(kbList)
	checkErr(err)

	//w.Header().Set("Content-Type", "application/json")
	lhjsonString, err := json.Marshal(lhList)
	checkErr(err)

	//w.Header().Set("Content-Type", "application/json")
	pzjsonString, err := json.Marshal(pzList)
	checkErr(err)

	//w.Header().Set("Content-Type", "application/json")
	ssjsonString, err := json.Marshal(ssList)
	checkErr(err)

	//w.Header().Set("Content-Type", "application/json")
	brjsonString, err := json.Marshal(brList)
	checkErr(err)

	//w.Header().Set("Content-Type", "application/json")
	bgjsonString, err := json.Marshal(bgList)
	checkErr(err)

	fmt.Fprintf(w, string(cgjsonString))
	fmt.Fprintf(w, string(kbjsonString))
	fmt.Fprintf(w, string(lhjsonString))
	fmt.Fprintf(w, string(pzjsonString))
	fmt.Fprintf(w, string(ssjsonString))
	fmt.Fprintf(w, string(brjsonString))
	fmt.Fprintf(w, string(bgjsonString))
}
