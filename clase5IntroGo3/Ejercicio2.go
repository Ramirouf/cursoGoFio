/*
Crear un sistema que permita gestionar cursos.

La entidad Curso debe contar con: ID, nombre, aula asignada y cantidad máxima de estudiantes.
Debe contar con una lista de estudiantes y permitir su listado, alta, baja y modificación.
Debe contar con un profesor asignado: ID, nombre, antigüedad.
Cada alumno debe contar con: ID, nombre, lista de notas y cantidad de faltas.
Crear un método que permita listar y devolver qué alumnos se encuentran por debajo de 6 en su promedio de notas.
Crear un método que permita listar y devolver qué alumnos están libres por cantidad de faltas, siendo hasta 3 faltas las permitidas.
Crear un método que permita listar y devolver los 3 alumnos con mejor promedio.
Crear un método que devuelva cuán llena en porcentaje está el aula y cuántos alumnos tiene asignados.

*/

package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main(){

	course1 := course {
		id: 1,
		name: "Course 1, GO bootcamp",
		classroom: "D1",
		maxStudents: 40,
	}
	student1 := Student {
		id: 43758549,
		name: "Ramiro",
		grades: [4]float64{4,6,8,10}, //Creates slice, and then sets values
		absences: 0,
	}

	fmt.Println(course1.name)
	fmt.Println(student1.grades)
}
func createCourse() (createdCourse course){
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	createdCourse = course{
		id: r1.Intn(10),
		name: "Course number " + string(createdCourse.id),
		classroom: "A " + string(createdCourse.id),
		maxStudents: 40,
	}
	

	return createdCourse
}

type course struct{
	id int
	name string
	classroom string
	maxStudents int
	students []Student
	teacher Teacher
}
type Teacher struct{
	id int
	name string
	seniority int
}
type Student struct{
	id int
	name string
	grades [4]float64
	absences int
}