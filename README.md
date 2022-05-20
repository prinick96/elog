# ERROR LOG

Una librería simple que escribí para utilizar en un proyecto, con la única función de registar errores en un fichero .log y evitar el spam de if's para revisar cada error

### EXAMPLE

```go
package main

import (
	//...... some libs
	"github.com/prinick96/elog"
)

func main() {
	//...... some code 
	querystring := `SELECT blab FROM blabla WHERE bla = ? LIMIT 1`
	err := db.QueryRowContext(ctx, querystring, bla).Scan(&tmp)

	// cuando usamos elog.ERROR ocupar go para hacerlo en segundo plano
	go elog.New(elog.ERROR, "In SELECT blabla", err)
	
	// cuando queremos que el error sea un panic, no usar go
	elog.New(elog.PANIC, "In SELECT blabla", err)
	//...... some code
}
```
- Genera un fichero .log por día
- Se llama una única función elog.New() para realizar el log
	- PRIMER PARÁMETRO:
		- Recibe elog.ERROR || 0, si queremos mostar en consola y dejar el registro
		- Recibe elog.PANIC || 1, si queremos generar un panic además de mostrar en consola y dejar registro
	- SEGUNDO PARÁMETRO:
		- Un string para indiciar más información sobre el error, si se pasa "" se omite este paso
	- TERCER PARÁMETRO:
		- El error, si no hay error no realiza ninguna acción

### ./logs/20-May-2022.log
```log
05:32:27.74832 ERROR: In SELECT blabla -> sql: no rows in result set
05:32:27.74838 FATAL: In SELECT blabla -> sql: no rows in result set
```