package elog

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

/**
 * @desc Tipos de errores manejados
 */
const (
	// Genera un PANIC, el error es mostrado en consola y guardado en el .log
	PANIC uint8 = 0
	// El error es mostrado en consola y guardado en el .log
	ERROR uint8 = 1
)

/**
 * @desc Escribe en el fichero de logs
 */
func writeLog(mtx *sync.Mutex, err_str string) {
	mtx.Lock()
	defer mtx.Unlock()

	// Creamos el directorio del fichero si no existe
	log_dir := "./logs/"
	err := os.MkdirAll(log_dir, 0755)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(-1)
	}

	// Nombre del fichero .log
	strTime := time.Now().Format("02-January-2006")
	f_name := log_dir + strTime + ".log"

	// Archivo nuevo cada día, si existe uno escribe en él
	l_file, err := os.OpenFile(f_name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		os.Stderr.WriteString(err.Error() + f_name + "\n")
		os.Exit(-1)
	}

	// Dentro del .log solo nos interesa el tiempo
	strTime = time.Now().Format("03:04:05.00000")

	// Escribimos en el fichero
	fmt.Fprintf(l_file, strTime+" %s\n", err_str)

	// Lo liberamos
	defer l_file.Close()
}

/**
 * @desc Crea un nuevo log si existe un error
 *
 * @param level puede ser 0 o 1, pero recomiendo usar elog.PANIC / elog.ERROR
 * @param str un string cualquiera para identificar nuestro error con más facilidad
 * @param err el error a identifiar, si es nil no realizará ninguna acción
 */
func New(level uint8, str string, err error) {
	if err != nil {
		// El formato del log
		levels := [2]string{"FATAL: ", "ERROR: "}
		err_str := levels[level] + str + " -> " + err.Error()

		// Escribir
		mtx := &sync.Mutex{}
		writeLog(mtx, err_str)

		// Causar pánico
		if level == PANIC {
			panic(err_str)
		}

		// Mostrar
		log.Println(err_str)
	}
}
