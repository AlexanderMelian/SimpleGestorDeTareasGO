package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type listaaTareas struct {
	tareas []*tarea
}

func (lt *listaaTareas) agregarTarea(t *tarea) {
	lt.tareas = append(lt.tareas, t)
}

func (lt *listaaTareas) quitarTarea(index int) {
	lt.tareas = append(lt.tareas[:index], lt.tareas[index+1:]...)
}

func (t *tarea) marcarCompletado() {
	t.estado = true
}

type tarea struct {
	titulo      string
	descripcion string
	estado      bool
}

func menu() int {
	fmt.Println("Menu")
	fmt.Println("1- Agregar tarea   2- Quitar Tarea   3- Marcar tarea finalizada   4- Mostrar tareas   5- Salir")
	dato := leerDatos()
	opcion, _ := strconv.Atoi(dato)
	return opcion
}

func leerDatos() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	lista := &listaaTareas{
		tareas: []*tarea{},
	}
	for {
		op := menu()
		if op == 5 {
			break
		}
		switch op {
		case 1:
			fmt.Println("Ingrese el título:")
			tit := leerDatos()
			fmt.Println("Ingrese la descripcion:")
			desc := leerDatos()
			ta := &tarea{
				titulo: tit, descripcion: desc,
			}
			lista.agregarTarea(ta)
		case 2:
			fmt.Println("Ingrese el número de tarea que quiere quitar")
			dato := leerDatos()
			opcion, _ := strconv.Atoi(dato)
			lista.quitarTarea(opcion)
		case 3:
			fmt.Println("Ingrese el número de la tarea que quiere marcar")
			dato := leerDatos()
			opcion, _ := strconv.Atoi(dato)
			lista.tareas[opcion-1].marcarCompletado()
		case 4:
			fmt.Println("-------------------------------------------------")
			for index, tarea := range lista.tareas {
				fmt.Println("Tarea N°: ", (index + 1), "  Título: ", tarea.titulo)
				if tarea.estado == false {
					est := "Incompleto"
					fmt.Println("Descripción: ", tarea.descripcion, "  Estado: ", est)
				} else {
					est := "Completo"
					fmt.Println("Descripción: ", tarea.descripcion, "  Estado: ", est)
				}
				fmt.Println("-------------------------------------------------")
			}
		}

	}
}
