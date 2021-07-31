package euspec

const diseaseAgentTargeted = `{
	"valueSetId": "disease-agent-targeted",
	"valueSetDate": "2021-04-27",
	"valueSetValues": {
	  "840539006": {
		"display": "COVID-19",
		"lang": "en",
		"active": true,
		"version": "http://snomed.info/sct/900000000000207008/version/20210131",
		"system": "http://snomed.info/sct"
	  }
	}
  }
`

const vaccineProphylaxis = `{
	"valueSetId": "sct-vaccines-covid-19",
	"valueSetDate": "2021-04-27",
	"valueSetValues": {
	  "1119349007": {
		"display": "SARS-CoV-2 mRNA vaccine",
		"lang": "en",
		"active": true,
		"version": "http://snomed.info/sct/900000000000207008/version/20210131",
		"system": "http://snomed.info/sct"
	  },
	  "1119305005": {
		"display": "SARS-CoV-2 antigen vaccine",
		"lang": "en",
		"active": true,
		"version": "http://snomed.info/sct/900000000000207008/version/20210131",
		"system": "http://snomed.info/sct"
	  },
	  "J07BX03": {
		"display": "covid-19 vaccines",
		"lang": "en",
		"active": true,
		"version": "2021-01",
		"system": "http://www.whocc.no/atc"
	  }
	}
  }  
`

const vaccineMedicalProduct = `{
    "valueSetId": "vaccines-covid-19-names",
    "valueSetDate": "2021-04-27",
    "valueSetValues": {
      "EU/1/20/1528": {
        "display": "Comirnaty",
        "lang": "en",
        "active": true,
        "system": "https://ec.europa.eu/health/documents/community-register/html/",
        "version": ""
      },
      "EU/1/20/1507": {
        "display": "COVID-19 Vaccine Moderna",
        "lang": "en",
        "active": true,
        "system": "https://ec.europa.eu/health/documents/community-register/html/",
        "version": ""
      },
      "EU/1/21/1529": {
        "display": "Vaxzevria",
        "lang": "en",
        "active": true,
        "system": "https://ec.europa.eu/health/documents/community-register/html/",
        "version": ""
      },
      "EU/1/20/1525": {
        "display": "COVID-19 Vaccine Janssen",
        "lang": "en",
        "active": true,
        "system": "https://ec.europa.eu/health/documents/community-register/html/",
        "version": ""
      },
      "CVnCoV": {
        "display": "CVnCoV",
        "lang": "en",
        "active": true,
        "system": "http://ec.europa.eu/temp/vaccineproductname",
        "version": "1.0"
      },
      "Sputnik-V": {
        "display": "Sputnik-V",
        "lang": "en",
        "active": true,
        "system": "http://ec.europa.eu/temp/vaccineproductname",
        "version": "1.0"
      },
      "Convidecia": {
        "display": "Convidecia",
        "lang": "en",
        "active": true,
        "system": "http://ec.europa.eu/temp/vaccineproductname",
        "version": "1.0"
      },
      "EpiVacCorona": {
        "display": "EpiVacCorona",
        "lang": "en",
        "active": true,
        "system": "http://ec.europa.eu/temp/vaccineproductname",
        "version": "1.0"
      },
      "BBIBP-CorV": {
        "display": "BBIBP-CorV",
        "lang": "en",
        "active": true,
        "system": "http://ec.europa.eu/temp/vaccineproductname",
        "version": "1.0"
      },
      "Inactivated-SARS-CoV-2-Vero-Cell": {
        "display": "Inactivated SARS-CoV-2 (Vero Cell)",
        "lang": "en",
        "active": true,
        "system": "http://ec.europa.eu/temp/vaccineproductname",
        "version": "1.0"
      },
      "CoronaVac": {
        "display": "CoronaVac",
        "lang": "en",
        "active": true,
        "system": "http://ec.europa.eu/temp/vaccineproductname",
        "version": "1.0"
      },
      "Covaxin": {
        "display": "Covaxin (also known as BBV152 A, B, C)",
        "lang": "en",
        "active": true,
        "system": "http://ec.europa.eu/temp/vaccineproductname",
        "version": "1.0"
      }
    }
  }`

const vaccineMah = `{
	"valueSetId": "vaccines-covid-19-auth-holders",
	"valueSetDate": "2021-04-27",
	"valueSetValues": {
	  "ORG-100001699": {
		"display": "AstraZeneca AB",
		"lang": "en",
		"active": true,
		"system": "https://spor.ema.europa.eu/v1/organisations",
		"version": ""
	  },
	  "ORG-100030215": {
		"display": "Biontech Manufacturing GmbH",
		"lang": "en",
		"active": true,
		"system": "https://spor.ema.europa.eu/v1/organisations",
		"version": ""
	  },
	  "ORG-100001417": {
		"display": "Janssen-Cilag International",
		"lang": "en",
		"active": true,
		"system": "https://spor.ema.europa.eu/v1/organisations",
		"version": ""
	  },
	  "ORG-100031184": {
		"display": "Moderna Biotech Spain S.L.",
		"lang": "en",
		"active": true,
		"system": "https://spor.ema.europa.eu/v1/organisations",
		"version": ""
	  },
	  "ORG-100006270": {
		"display": "Curevac AG",
		"lang": "en",
		"active": true,
		"system": "https://spor.ema.europa.eu/v1/organisations",
		"version": ""
	  },
	  "ORG-100013793": {
		"display": "CanSino Biologics",
		"lang": "en",
		"active": true,
		"system": "https://spor.ema.europa.eu/v1/organisations",
		"version": ""
	  },
	  "ORG-100020693": {
		"display": "China Sinopharm International Corp. - Beijing location",
		"lang": "en",
		"active": true,
		"system": "https://spor.ema.europa.eu/v1/organisations",
		"version": ""
	  },
	  "ORG-100010771": {
		"display": "Sinopharm Weiqida Europe Pharmaceutical s.r.o. - Prague location",
		"lang": "en",
		"active": true,
		"system": "https://spor.ema.europa.eu/v1/organisations",
		"version": ""
	  },
	  "ORG-100024420": {
		"display": "Sinopharm Zhijun (Shenzhen) Pharmaceutical Co. Ltd. - Shenzhen location",
		"lang": "en",
		"active": true,
		"system": "https://spor.ema.europa.eu/v1/organisations",
		"version": ""
	  },
	  "ORG-100032020": {
		"display": "Novavax CZ AS",
		"lang": "en",
		"active": true,
		"system": "https://spor.ema.europa.eu/v1/organisations",
		"version": ""
	  },
	  "Gamaleya-Research-Institute": {
		"display": "Gamaleya Research Institute",
		"lang": "en",
		"active": true,
		"system": "http://ec.europa.eu/temp/vaccinemanufacturer",
		"version": "1.0"
	  },
	  "Vector-Institute": {
		"display": "Vector Institute",
		"lang": "en",
		"active": true,
		"system": "http://ec.europa.eu/temp/vaccinemanufacturer",
		"version": "1.0"
	  },
	  "Sinovac-Biotech": {
		"display": "Sinovac Biotech",
		"lang": "en",
		"active": true,
		"system": "http://ec.europa.eu/temp/vaccinemanufacturer",
		"version": "1.0"
	  },
	  "Bharat-Biotech": {
		"display": "Bharat Biotech",
		"lang": "en",
		"active": true,
		"system": "http://ec.europa.eu/temp/vaccinemanufacturer",
		"version": "1.0"
	  }
	}
  }`
