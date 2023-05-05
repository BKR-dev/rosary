package prayers

import (
	"fmt"
	"time"
)

var dailyMysteryEN = map[string]string{
	"lang":   "en",
	"sunday": "glorious",
}

// myMap := make(map[string]string, 8) - creates a fixed size map, maps is still
// dynmaically sized, compiler only knows about it at compile time that the size is fixed

var dailyMysteryDE = map[string]string{
	"lang":       "de",
	"sonntag":    "glorreichen",
	"montag":     "freudenreichen",
	"dienstag":   "schmerzhaften",
	"mittwoch":   "glorreichen",
	"donnerstag": "lichtreichen",
	"freitag":    "schmerzhaften",
	"samstag":    "freudenreichen",
}

type Prayer struct {
	Language     string `json:"Lang"`
	Creed        string `json:"Creed"`
	OurFather    string `json:"OurFather"`
	GloryBe      string `json:"GloryBe"`
	FatimaPrayer string `json:"FatimaPrayer"`
	FirstDecade  string `json:"FirstDecade"`
	SecondDecade string `json:"SecondDecade"`
	ThirdDecade  string `json:"ThirdDecade"`
	FourthDecade string `json:"FourthDecade"`
	FifthDecade  string `json:"FifthDecade"`
	SalveRegina  string `json:"SalveRegina"`
}

type Rosary struct {
	Weekday  string   `json:"Weekday"`
	Language string   `json:"Lang"`
	Decades  int      `json:"Decades"` // always 5
	Prayers  []Prayer `json:"Prayers"`
}

func getWeekday() string {
	return fmt.Sprint(time.Now().Weekday())
}

// Glorreichen Geheimnisse Sonntag / Mittwoch

// Kreuzzeichen

// Glaubensbekenntnis
// Vaterunser

// Ave Maria +Jesus, der in uns den Glauben vermehre
// Ave Maria +Jesus, der in uns die Hoffnung stärke
// Ave Maria +Jesus, der in uns die Liebe entzüde

// Ehre sei dem Vater
// Vaterunser

// 10x Ave Maria + Geheimniss 1

// Ehre sei dem Vater
// Fatimagebet
// Vaterunser

// 10x Ave Maria + Geheimniss 2

// Ehre sei dem Vater
// Fatimagebet
// Vaterunser

// 10x Ave Maria + Geheimniss 3

// Ehre sei dem Vater
// Fatimagebet
// Vaterunser

// 10x Ave Maria + Geheimniss 4

// Ehre sei dem Vater
// Fatimagebet
// Vaterunser

// 10x Ave Maria + Geheimniss 5

// Ehre sei dem Vater
// Fatimagebet
// Salve Regina

// Glorreiche Geheimnisse:
// 1) Der von den Toten auferstanden ist;
// 2) Der in den Himmel aufgefahren ist;
// 3) Der uns den Heiligen Geist gesandt hat;
// 4) Der dich, o Jungfrau, in den Himmel aufgenommen hat;
// 5) Der dich, o Jungfrau, im Himmel gekrönt hat;

// Glaubensbekenntnis
// Ich glaube an Gott, den Vater, den Allmächtigen,
// den Schöpfer des Himmels und der Erde,

// und an Jesus Christus, seinen eingeborenen Sohn, unsern Herrn,
// empfangen durch den Heiligen Geist, geboren von der Jungfrau Maria,

// gelitten unter Pontius Pilatus,
// gekreuzigt, gestorben und begraben,

// hinabgestiegen in das Reich des Todes, am dritten Tage auferstanden von den Toten,
// aufgefahren in den Himmel;
// er sitzt zur Rechten Gottes, des allmächtigen Vaters;
// von dort wird er kommen, zu richten die Lebenden und die Toten.

// Ich glaube an den Heiligen Geist,
// die heilige katholische Kirche,
// Gemeinschaft der Heiligen,
// Vergebung der Sünden,
// Auferstehung der Toten
// und das ewige Leben.

// Vater unser
// Vater unser im Himmel, geheiligt werde dein Name.
// Dein Reich komme.
// Dein Wille geschehe,
// wie im Himmel so auf Erden.
// Unser tägliches Brot gib uns heute.
// Und vergib uns unsere Schuld,
// wie auch wir vergeben unsern Schuldigern.
// Und führe uns nicht in Versuchung,
// sondern erlöse uns von dem Bösen.
// Denn dein ist das Reich und die Kraft
// und die Herrlichkeit in Ewigkeit.

// Ehre sei dem Vater
// Ehre sei dem Vater und dem Sohn und dem Heiligen Geist.
// Wie im Anfang, so auch jetzt und allezeit und in Ewigkeit.

// Fatimagebet (Bittgebet)
// O mein Jesus, verzeih' uns unsere Sünden. Bewahre uns vor dem Feuer der Hölle. Führe alle Seelen in den Himmel, besonders jene, die Deiner Barmherzigkeit am meisten bedürfen.
// Amen.

// Ave Maria
// Gegrüßet seist du, Maria, voll der Gnade,
// der Herr ist mit dir.
// Du bist gebenedeit unter den Frauen,
// und gebenedeit ist die Frucht deines Leibes,
// Jesus.  (Geheimnis)
// Heilige Maria, Mutter Gottes, bitte für uns Sünder jetzt
// und in der Stunde unseres Todes.
// Amen.

// Salve Regina
// Sei gegrüßt, o Königin,
// Mutter der Barmherzigkeit,
// unser Leben, unsre Wonne
// und unsre Hoffnung, sei gegrüßt!
// Zu dir rufen wir
// verbannte Kinder Evas.
// Zu dir seufzen wir trauernd und weinend
// in diesem Tal der Tränen.
// Wohlan denn, unsre Fürsprecherin,
// wende deine barmherzigen Augen
// uns zu,
// und nach diesem Elend zeige uns Jesus,
// die gebenedeite Frucht deines Leibes!
// O gütige, o milde, o süße Jungfrau Maria!
