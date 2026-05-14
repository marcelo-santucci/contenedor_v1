package main

/**
 * 
 * Sistemas Operativos Universidad Rafael Landivar 
 * Lenguage Go
 *
 * Este archivo implementa el aislamiento de espacios de nombres para un "contenedor" que ejecuta bash en un entorno aislado. 
 
 * El aislamiento de espacios de nombres "Namespace" solo funciona para archivos en memoria, lo que significa que el "contenedor" 
 * aún puede acceder directamente al sistema de archivos del host y a los archivos reales almacenados en el disco.
 *
 * Para implementar el aislamiento de recursos como CPU, disco, etc., se utiliza cgroups.
 */

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func correNameSpace() {
	args := os.Args

	// El proceso actual es para el hijo.
	if len(args) > 1 && strings.EqualFold(args[1], "hijo") {
		ArrancarNameSpaceHijo()
		return
	}

	// Si no se recibe argumento es el padre
	ArrancarNameSpacePadre()

}

/* En Linux, /proc/self/exe apunta al binario que se está ejecutando actualmente, 
 * lo que impide la suplantación de identidad a través de la línea de comandos.
 * 
 * Aquí se referencia al proceso hijo como parámetro para poder reconocer a dicho proceso. 
 * Es puramente sintáctico y puedes reemplazarlo con cualquier cosa para detectar si el proceso es hijo.
*/
func ArrancarNameSpacePadre() {

	cmd := exec.Command("/proc/self/exe", "hijo", "/bin/sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Pdeathsig:  syscall.SIGTERM,
		Setsid:     true,
	}

	// Mapeo de las tuberías del host para poder ver la salida.
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Run()

	if err != nil {
		fmt.Printf("No se pudo clonar el proceso hijo.: %v\n", err)
		os.Exit(1)
	}
}

func ArrancarNameSpaceHijo() {
	args := os.Args

	err := syscall.Sethostname([]byte("host-personalizado"))
	if err != nil {
		fmt.Println("No se pudo cambiar el nombre de host del espacio de nombres hijo (secundario).")
		os.Exit(1)
	}
	hname, err := os.Hostname()
	if err != nil {
		fmt.Printf("No se pudo obtener el nombre de host del espacio de nombres hijo (secundario).: %v", err)

	} else {
		fmt.Printf("El nombre de host ha cambiado. El nuevo nombre de host es: %s\n", hname)
	}

	// Evitar cualquier propagación de eventos al host.
	err = syscall.Mount("", "/", "", syscall.MS_PRIVATE|syscall.MS_REC, "")
	if err != nil {
		fmt.Printf("No se pudo hacer que el montaje raíz fuera privado. : %v", err)
		os.Exit(1)
	}

	// Añadir medidas de seguridad similares a las que aplica Docker ;-)
	err = syscall.Mount("proc", "/proc", "proc", syscall.MS_NOSUID|syscall.MS_NODEV|syscall.MS_NOEXEC, "")
	if err != nil {
		fmt.Printf("Fallo al montar el /proc: %v", err)
		os.Exit(1)
	}

	if len(args) > 2 && len(args[2]) > 0 {
		err := syscall.Exec(args[2], args[2:], os.Environ())
		if err != nil {
			fmt.Printf("No se pudo ejecutar el binario en el espacio de nombres hijo.: %v\n", err)
			os.Exit(1)
		}
		return
	}

}