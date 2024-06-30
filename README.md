# AT24 -Implementcaija aktorskog radnog okruzenja
predmeti projeakat iz predmeta agentske tehnologije


## Elementi za implementaciju:
 - Osnovni: Aktori, Asinhrono slanje i primanje poruka, Sanduce, Menjanje ponasanja(stanja) aktora, Reagovanje na lifecycle dogadjaje stvaranja i otkazivanja aktora, Udaljena komunikacija aktora
 - Proizvoljni koncepti: Supervizija aktora, klasterovanje(primarno), klasterovanje(sekundarno)

## Članovi tima
  - Nikola Petkovic RA 232/2015



# POKRETANJE
##  PING-PONG test
>1.1 Pokrenuti prvu instancu servera, parametri port i adresa drugog servera(za prvu instancu moze bilo sta)
>```sh
>  go run udpServer.go 9090 ne_poznajem_druge_servere
>```
>1.2 Konzola prve instance ce ispisati adresu i port na kome je server pokrenut
>1.3 Pokrenuti drugu instancu servera, port, adresa iz takce 1.2
>```sh
> go run udpServer.go 9091 192.168.1.101:9090
>```
>1.4 Serveri 1 i 2 pocinju beskonacno slati PING i PONG jedan drugom\
>1.5. Poslati direktivu nekom serveru da otpocne svoje zaustavljanje
>```sh
>go run direktive.go stop 192.168.1.101:9091
>```
>1.6 Server kojem je poslata direktiva stop ce prestati osluskivati dalje poruke, svim operativicima ce biti poslata direktiva da obrade preostale poruke u sanducetu i prestanu sa primanjem novih poruka.\
>1.7 Svaki operativac pri zaustavljanju blokira svoju rutinu na 10s tako da na konzoli nece biti ispisa 10ak sekundi, isPenzionisan() vraca true(novo stanje).\
>1.8 Kada svi operativci zavrse svoje zaustavljanje server ce ukloniti sve agente koji su penzionisani.\
>1.9 Sluzba koja je primila direktivu stop iz 1.5 uspesno zavrsava svoje poslovanje i main() metoda zavrsava svoje izvrsavanje.

##  FAILURE test
>1.1 Pokrenuti prvu instancu servera, parametri port i adresa drugog servera(za prvu instancu moze bilo sta)
>```sh
>  go run udpServer.go 9090 ne_poznajem_druge_servere
>```
>1.2 Konzola prve instance ce ispisati adresu i port na kome je server pokrenut
>1.3. Poslati direktivu na adresu procitanu sa konzole(1.2)
>```sh
>go run direktive.go force_fail 192.168.1.101:9090
>```
>1.4 direktiva force_fail salje 10 poruka na adresu navedenu u drugom parametru, s tim da pri svakoj obradi postoji samo 33% sanse za uspeh, u slucaju neuspeha, poruka i agent inkrementiraju svoje failure brojace.
>1.5 U slucaju neupeha(66% sanse) agent vraca poruku sluzbi na dalju obradu, ako je failure brojace poruke <5 poruka ce opet biti poslata nekom random agentu, za brojac >=5 sluzba odustaje od dalje obrade poruke(TODO: mogla bi se poslati nazad onome ko ju je poslao)


