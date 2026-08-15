# template

Package placeholder da copiare per ogni nuovo esercizio.

## Come creare un nuovo esercizio

1. Copia la cartella `template` con un nuovo nome, es. `two_sum`:
   `cp -r template two_sum`
2. Rinomina i file dentro la nuova cartella:
   `mv two_sum/template.go two_sum/two_sum.go`
   `mv two_sum/template_test.go two_sum/two_sum_test.go`
3. Cambia il nome del package nei due file: `package template` → `package two_sum`.
4. Sostituisci la funzione `Add` di esempio con la tua soluzione e scrivi i test.
5. Esegui i test di tutto il progetto: `go test ./...`

## Regole Go che questo layout sfrutta

- Un package per cartella: il nome del package deve coincidere con la cartella
  (convenzione Go).
- I file che finiscono in `_test.go` contengono i test dello stesso package:
  possono accedere anche alle funzioni non esportate (minuscole).
- `go test ./...` esegue i test di tutti i package ricorsivamente.
