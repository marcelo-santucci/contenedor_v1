# Contenedores

Un contenedor es básicamente un programa en ejecución que ya tiene todo lo necesario para ejecutar tu programa. 

Es posible tener varios contenedores ejecutándose en una sola máquina y no que interferirán entre sí ya que las aplicaciones dentro de ese contenedor nunca sabrán qué más se está ejecutando fuera de ellas.  Solo sabe lo que le has indicado sobre el mundo exterior.

El codigo de este repositorio es el de un "conenedor" básico escrito en el lenguaje "Go"

🏗️Instalar compilador de "Go":

```c

sudo apt update
sudo apt install golang-go

//Verificar
go version

```

⛏️Compilar <br>
    👉🏼 go build -o contenedor main.go namespace.go <br>
    
🟢Forma de uso: <br>
  
  > Abra una ventana adicional donde revisara estado de procesos
  > Ejecute primero el programa con el comando: sudo ./contenedor
 
🔎 Revisión: <br>

  > En la ventana donde ejecuto el comando ejecute el comando: ps --forest aux <br>
  > En la segunda ventana ejecute el comando:  ps --forest aux | grep contenedor <br>

❓¿Que cosas puede apreciar con respecto de la información desplegada.
