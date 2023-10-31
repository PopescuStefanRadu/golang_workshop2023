## Thread hardcodat la nivel de endpoint

Vreau un endpoint nou care sa intoarca o vizualizare a 
unui thread de whistleblower.

Un thread contine Titlu si o lista de mesaje text.

Implementarea sa fie cu date hardcodate dar sa foloseasca templating 
engineul html/template si sa fie preferabil cel gata integrat in gin.

Endpointul o sa arate ceva de genul asta: \(...)/threads/{thread-uuid}
Thread uuid vreau hardcodat.(sau poate fi orice, idc)

## Thread hardcodat la nivel de databasa

O sa avem doua tabele, una cu thread si una care sa tina mesajele aferente
fiecarui thread.

Exemplu de implementare: sqlc.

Install sqlc.

Use `go generate ./...` after configuring //go:generate
## Vineri

TODO:

 - arhitectura, structura proiect
 - teste cu mockuri
 - intrebari