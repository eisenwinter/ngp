package euspec

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/eisenwinter/ngp/pkg/template"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

//EU JSON SPEC
//https://ec.europa.eu/health/sites/default/files/ehealth/docs/covid-certificate_json_specification_en.pdf

//valuesets taken from https://github.com/admin-ch/CovidCertificate-Examples/tree/main/valuesets
type EuJsonPayload struct {
	//Dob Date of birth
	Dob string `json:"dob"`
	Nam struct {
		//Fn Surname(s)
		Fn string `json:"fn"`
		//Fnt Standardised surname(s)
		Fnt string `json:"fnt"`
		//Gn Forename(s)
		Gn string `json:"gn"`
		//Gnt  Standardised forename(s)
		Gnt string `json:"gnt"`
	} `json:"nam"`
	//V: Vaccination group  - T (test group) and R (recovery group) are NOT covered
	V []struct {
		// Unique certificate identifier
		Ci string `json:"ci"`
		//Co Member State or third country in which the vaccine was administered
		Co string `json:"co"`
		//Number in a series of doses
		Dn int `json:"dn"`
		//Date of vaccination
		Dt string `json:"dt"`
		//Is Certificate issuer
		Is string `json:"is"`
		//Ma COVID-19 vaccine marketing authorisation holder or manufacturer - vaccine-mah-manf.json.
		Ma string `json:"ma"`
		//COVID-19 vaccine product - vaccine-medicinal-product.json.
		Mp string `json:"mp"`
		//Sd The overall number of doses in the series (1, 2, 3= Booster)
		Sd int `json:"sd"`
		//Tg  Disease or agent targeted: COVID-19 (SARS-CoV or one of its variants) - disease-agent-targeted.json
		Tg string `json:"tg"`
		//Vp COVID-19 vaccine or prophylaxis - vaccine-prophylaxis.json.
		Vp string `json:"vp"`
	} `json:"v"`
	//Ver Schema version - MUST match the identifier of the schema version used for producing the EUDCC. Example "ver": "1.3.0"
	Ver string `json:"ver"`
}

type EuValueSetItem struct {
	Active  bool   `json:"active"`
	Display string `json:"display"`
	Lang    string `json:"lang"`
	System  string `json:"system"`
	Version string `json:"version"`
}

type EuValueSet struct {
	ValueSetDate   string                    `json:"valueSetDate"`
	ValueSetID     string                    `json:"valueSetId"`
	ValueSetValues map[string]EuValueSetItem `json:"valueSetValues"`
}

func init() {
	message.SetString(language.English, "Surname(s), Forename(s)", "Surname(s), Forename(s)")
	message.SetString(language.German, "Surname(s), Forename(s)", "Nachname(n), Vorname(n)")
	message.SetString(language.German, "Date of birth (yyyy-mm-dd)", "Geburtsdatum (JJJJ-MM-TT)")
	message.SetString(language.German, "Disease or agent targeted", "Zielkrankheit oder -erreger")
	message.SetString(language.German, "COVID-19 vaccine or prophylaxis", "COVID-19-Impfstoff oder -Prophylaxe")
	message.SetString(language.German, "COVID-19 vaccine product name", "COVID-19-Impfstoffhandelsname")
	message.SetString(language.German, "Vaccine marketing authorization holder or manufacturer", "Zulassungsinhaber oder Hersteller des Impfstoffs")
	message.SetString(language.German, "Number in a series of vaccinations / doses", "Nummer der Impfung / Anzahl Dosen")
	message.SetString(language.German, "Date of vaccination (yyyy-mm-dd)", "Datum der Impfung (JJJJ-MM-TT)")
	message.SetString(language.German, "Member State or third country", "Mitgliedstaat oder Drittstaat")

}

func loadValueSetInlined(jsonContent string) (*EuValueSet, error) {

	var valueSet EuValueSet
	err := json.Unmarshal([]byte(jsonContent), &valueSet)
	return &valueSet, err
}

func ToRenderView(json EuJsonPayload, qrSource string, lang string) template.RenderView {
	view := template.RenderView{}
	dat, err := loadValueSetInlined(diseaseAgentTargeted)
	if err != nil {
		fmt.Println("Failed to load value set")
	}
	vacPro, err := loadValueSetInlined(vaccineProphylaxis)
	if err != nil {
		fmt.Println("Failed to load value set")
	}
	vacMed, err := loadValueSetInlined(vaccineMedicalProduct)
	if err != nil {
		fmt.Println("Failed to load value set")
	}
	vacMah, err := loadValueSetInlined(vaccineMah)
	if err != nil {
		fmt.Println("Failed to load value set")
	}

	parsedLang := message.MatchLanguage(lang)
	p := message.NewPrinter(parsedLang)
	view.QrSource = qrSource
	view.Infos = make([]template.DetailView, 0)
	view.Infos = append(view.Infos, template.DetailView{
		Display: p.Sprintf("Surname(s), Forename(s)"),
		Value:   json.Nam.Fn + ", " + json.Nam.Gn,
	})
	view.Infos = append(view.Infos, template.DetailView{
		Display: p.Sprintf("Date of birth (yyyy-mm-dd)"),
		Value:   json.Dob,
	})
	for _, v := range json.V {

		view.Infos = append(view.Infos, template.DetailView{
			Display: p.Sprintf("Disease or agent targeted"),
			Value:   dat.ValueSetValues[v.Tg].Display,
			Detail:  v.Tg,
		})
		view.Infos = append(view.Infos, template.DetailView{
			Display: p.Sprintf("COVID-19 vaccine or prophylaxis"),
			Value:   vacPro.ValueSetValues[v.Vp].Display,
			Detail:  v.Vp,
		})
		view.Infos = append(view.Infos, template.DetailView{
			Display: p.Sprintf("COVID-19 vaccine product name"),
			Value:   vacMed.ValueSetValues[v.Mp].Display,
			Detail:  v.Mp,
		})
		view.Infos = append(view.Infos, template.DetailView{
			Display: p.Sprintf("Vaccine marketing authorization holder or manufacturer"),
			Value:   vacMah.ValueSetValues[v.Ma].Display,
			Detail:  v.Ma,
		})
		view.Infos = append(view.Infos, template.DetailView{
			Display: p.Sprintf("Number in a series of vaccinations / doses"),
			Value:   strconv.Itoa(v.Dn) + " / " + strconv.Itoa(v.Sd),
		})
		view.Infos = append(view.Infos, template.DetailView{
			Display: p.Sprintf("Date of vaccination (yyyy-mm-dd)"),
			Value:   v.Dt,
		})
		view.Infos = append(view.Infos, template.DetailView{
			Display: p.Sprintf("Member State or third country"),
			Value:   v.Co,
		})
		view.Id = v.Ci
		view.Issuer = v.Is
	}

	return view
}
