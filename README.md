WOA application
=
Woa is a simple standalone query tool to query airport information from a cloud hosted database 
(`https://mikerhodes.cloudant.com/airportdb`). To use you only need to know the Latitude/Longitude coordinates of a 
given position on the map, and specify the distance in Km. The application will query and sort all Cities within that 
distance.  

Usage
== 
Since GO language outputs portable, static-compiled binaries, in order to run the application, you just need to download
the binary specific to your Operating System from the `github.com/ysolt/woa/releases` site.

In general to run the application you need to specify the arguments in the following order. 
 ```
woa <distance in Km> <latitude> <longitude> 
```

Example output
```
                      City name |  Distance |      Latitude | Longitude
                   -------------|  ---------|  -------------|  ---------|
                Budapest Keteli |          3|     47.500497 | 19.085484
                       Ferihegy |         17|     47.436933 | 19.255592
                      Kecskemet |         83|       46.9175 | 19.749222
            Szentkiralyszabadja |         93|     47.077861 | 17.968444
 Győr-Pér International Airport |         93|     47.627097 | 17.808347
                        Szolnok |         99|     47.122861 | 20.2355
```

MacOS / Linux specific execution steps
===
You'll need to add execution bits to the downloaded binary file. Then you'll be able to execute the application like 
this 
```
chmod 700 ./woa
./woa 100 47.497913 19.040236 
```

Windows specific execution steps
===
```
woa.exe 100 47.497913 19.040236
```

Development area
==
If you decided to make changes and potentially even contribute it back, here some steps for local setup

Environment setup
=== 
`go` installed on your local machine via homebrew or your favourite package manager
Fetch the WOA source code via `go get github.com/ysolt/woa`. The source code will be populated into your GO source 
directory like `~/go/src/github.com/ysolt/woa/`

How to test local changes
===
Once you made any change to the source code here is a command, that you can use to re-build and run as one single shot
```
go run ~/go/src/github.com/ysolt/woa 47.497913 19.040236 1000 
```

Compilation / Packaging
===
If you want to compile the application to the same operation system as you are running now:
```
go build ~/go/src/github.com/ysolt/woa
```

If you want to cross-compile the application to other operating systems
Windows:
```
GOOS=windows GOARCH=amd64 go build -o woa.exe ~/go/src/github.com/ysolt/woa
```

Linux:
```
GOOS=linux go build -o woa-linux  ~/go/src/github.com/ysolt/woa
```
