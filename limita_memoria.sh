#!/bin/bash
#
# URL - Sistemas Operativos 
#
# Script para demostrar limites de asignación de memoria. 
#

# Directorio para realizar la asignación de memoria
ALLOC_DIR="/dev/shm/mem_alloc_$$"

#Desplegar directorio creado
echo "Directorio: $ALLOC_DIR"

#Crear directorio
mkdir -p "$ALLOC_DIR"

# Limpiar al salir
trap 'echo "Liberando memoria..."; rm -rf "$ALLOC_DIR"' EXIT

# Asignar 10MB cada segundo hasta que se alcanzan 200Mb...
for i in {1..20}; do
  dd if=/dev/zero of="$ALLOC_DIR/block_$i" bs=1M count=10 &>/dev/null
  allocated_mb=$((i * 10))
  echo "Asignados ${allocated_mb}Mb"
  sleep 1
done

echo "¿200 Mb asignados en total o cuantos?  Se mantienen durante 30 segundos antes de liberar."
sleep 30
