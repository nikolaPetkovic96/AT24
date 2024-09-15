# AT24 -Implementcaija aktorskog radnog okruzenja
predmeti projeakat iz predmeta agentske tehnologije


## Elementi za implementaciju:
 - Osnovni: Aktori, Asinhrono slanje i primanje poruka, Sanduce, Menjanje ponasanja(stanja) aktora, Reagovanje na lifecycle dogadjaje stvaranja i otkazivanja aktora, Udaljena komunikacija aktora
 - Proizvoljni koncepti: Supervizija aktora, klasterovanje(primarno), klasterovanje(sekundarno)

## Članovi tima
  - Nikola Petkovic RA 232/2015



# POKRETANJE
##  test1.go
> Pokrece jedan sistem, 2 agenta razlicitog tipa, Worker tip spawnuje 5 child actora koji primaju po jednu poruku, i na kraju main meteode se sistem gasi "gracefully"
>TODO: nekad puca pri shutdownu jer napravi nekonzistentu mapu aktora pri njihovom pojedinacnom uklanjnaju iz mape
>```sh
>  go run test1.go
>```

## 2. Prost remoting sa slanje poruke preko preze
>TODO: treba razraditi metode za formiranje poruka da da bi bilo olaksana razmena poruka izmedju 2 servera, 
>```sh
>  go run test2.go
>```

#NAPOMENE 
>-poruke su interfejsi, koje se pri razmeni ovaviju tipom Envelop(poruka, sender, receiver, failCount),  
>-agent pri prijemu envelopa moze procitati sadrzaj poruke kao i posiljaoca i remote adresu preko koje mu se pristupa.   
>-Trebra razraditi formiranje poruka, pri remote slanju se trenutno zahteva da korisnik sve isprati i definise kako treba   

