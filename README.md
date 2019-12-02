WOA application
=

Requirements (dev)
=
`go` installed on your local machine via homebrew or your favourite package manager

Usage
== 
`go run . <distance in Km>`
like
`go run . 2000`

Example output
```
              Kamenz |   551|
       Leipzig Halle |   657|
              Erfurt |   702|
             Hamburg |   933|
              Weelde |  1103|
            Alderney |  1576|
 Manchester Woodford |  1627|
```

Debugging
==
How to get `cludant_response_example.json`? 

Use 
```curl --globoff  'https://mikerhodes.cloudant.com/airportdb/_design/view1/_search/geo?limit=200&q=lon:[-90%20TO%2090]%20AND%20lat:[-90%20TO%2090]' > resources/cludant_response_example.json```