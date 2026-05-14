# Contenedores

Un contenedor es básicamente un programa en ejecución que ya tiene todo lo necesario para ejecutar tu programa. 

Es posible tener varios contenedores ejecutándose en una sola máquina y no que interferirán entre sí ya que las aplicaciones dentro de ese contenedor nunca sabrán qué más se está ejecutando fuera de ellas.  Solo sabe lo que le has indicado sobre el mundo exterior.

El codigo de este repositorio es el de un "conenedor" básico escrito en el lenguaje "Go"

🏗️Instalar compilador de "Go":

```
sudo apt update
sudo apt install golang-go

//Verificar
$ go version
go version go1.13.8 linux/amd64
```

⛏️Compilar <br>

```
go build -o contenedor main.go namespace.go
```
    
🟢Forma de uso: <br>

Abra dos ventanas de terminal conectadas a su máquina virtual.

:runner: En la primera ventana ejecutar el programa: 

```
sudo ./contenedor
```

🔎 Revisión: <br>

> En la primera ventana donde esta ejecutando el contenedor:
```
ps --forest aux
```
> En la segunda ventana ejecute el comando:
```
ps --forest aux | grep contenedor
```
❓¿Que cosas puede apreciar con respecto de la información desplegada.
