package translator


import (
	"golang.org/x/exp/utf8string"
	"net/http"
	"math"
	"strings"
	"text/template"
)


type Pikachu struct {
	Plain string
	Blain string
}


func HandleTrans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	txt := r.FormValue("plain")
	ascii := utf8string.NewString(txt)
	pk := translate(txt)
	if !ascii.IsASCII() {
		pk = "ASCII以外の文字が含まれています。"
	}
	pika := Pikachu{txt, pk}
	tmpl, err := template.ParseFiles("views/layout.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, pika)
	if err != nil {
		panic(err)
	}
}


func round(f float64) int {
	return int(math.Floor(f + .5))
}


func translate(str string) string {
	ex_num := 0
	var fin_pk string
	for i, c := range str {
		num := int(c)
		diff := num - ex_num
		ex_num = num
		var sign string
		if diff > 0 {
			sign = "ピカチュー "
		} else {
			sign = "ピッ "
		}
		diff = int(math.Abs(float64(diff)))
		var pk string
		switch {
		case diff > 5:
			divisor := round(math.Sqrt(float64(diff)))
			dividend := diff / divisor
			remainder := diff % divisor
			pk = strings.Repeat("ピカチュー ", divisor) + "ピィ ピカ " + strings.Repeat(sign, dividend)+ "ピ ピッ ピカァ ピカ " + strings.Repeat(sign, remainder) + "ピッカ "
			if i > 0 {
				pk = "ピ " + pk
			}		
		case diff == 0:
			pk = "ピッカ "
		default:
			pk = strings.Repeat(sign, diff) + "ピッカ "
		}
		fin_pk += pk
	}
	return fin_pk
}

