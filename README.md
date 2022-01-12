# ngp (nicer green pass)

ngp is short for nicer green pass. its a utility tool that reads the QR code of vaccination certificates.
It then uses the retrived information to generate a small document containing the QR code and the information
contained within the QR code. 

This can be used as a handy printout to carry in your wallet.

It may also be interesting for some people to actually see the data 
within the QR code.

If you are interested in further information of how those QR codes work I recommend you to check out:

https://github.com/ehn-dcc-development/hcert-spec/blob/main/README.md

and

https://ec.europa.eu/health/sites/default/files/ehealth/docs/covid-certificate_json_specification_en.pdf


The application has only been tested with austrian vaccination certificates,
it _should_ work with any certificate QR code following the EU standard. 

# FAQ

Q: Is this a official?

A: No there is no affiliation with any authorities. This is a hobby project.

Q: Will the generated vaccination certificate (ngp) be accepted?

A: In theory even the QR code alone should be enough. Your mileage may vary.

Q: Why did you build this?

A: Two reasons, curiosity of whats stored in the QR Code and how it is signed and the fact that the
austrian vaccination certificate is a A4 paper containing information that could be fitted on a 
10 euro bill - which  would conveniently fit in my wallet. 

While they state the QR code is just enough most places just check the tabular information containing
and dont bother scanning the QR code so it had to be placed on their as well.

# Running generation

If you just want to generate the HTML file you can use:

```
go run cmd\generate\generate.go -f qr.png -t templates\default.html -o mynpg.html -lang de
``` 

-lang is the language (either de or en)
-f defines the qr code location (as png)
-o sets the output path name

If you are just curious whats inside the qr code you can use:

```
go run cmd\read\read.go -f ..\path-to-qr\qr.png 
``` 

and it will just output the struct to stdin


# Contribution

Contributions are welcome, altough this is just a hobby project.

# Uses

QR Code
https://github.com/makiuchi-d/gozxing - MIT

CBOR
https://github.com/fxamacker/cbor - MIT

Templating
https://github.com/hoisie/mustache - MIT

CovidCertificates Value Sets
https://github.com/ehn-dcc-development/ehn-dcc-schema/tree/main/valuesets

# License

BSD 2-Clause License