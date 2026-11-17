# Learning Go

## English

This repository is my learning journal as I study Go. It contains my personal notes about the language, practice problems from online judges, and Advent of Code solutions.

### Repository structure

```text
learning/
  <topic>.md             # Notes about Go and my learning experience
  solving/
    <platform>/
      <problem>/
        main.go          # Problem solution

advent-of-code/
  <year>/                 # 2015 through 2026
    day-01/
      statement.md        # Puzzle statement
      input.txt           # Puzzle input
      part1/
        main.go           # Part 1 executable
      part2/
        main.go           # Part 2 executable
```

Files directly under `learning/` hold my notes about Go: what I like, what I dislike, concepts I want to understand, and observations from learning. The `learning/solving/` directory is for problems from platforms such as Codeforces and OmegaUp. Advent of Code work lives separately under `advent-of-code/`.

Each AoC day keeps its statement and input alongside independent executables for both parts. Separate directories let each part use `package main` and its own `func main` without conflicts.

### Running an Advent of Code solution

Run a part from the repository root and pass its input through standard input:

```sh
go run ./advent-of-code/2026/day-01/part1 < advent-of-code/2026/day-01/input.txt
go run ./advent-of-code/2026/day-01/part2 < advent-of-code/2026/day-01/input.txt
```

### Why Go?

I have spent a long time looking for a modern alternative to C++. I considered Rust, but I prefer a language with garbage collection. Go seems like a good fit, so I am learning through practice.

---

## Español

Este repositorio es mi bitácora mientras aprendo Go. Contiene mis notas personales sobre el lenguaje, ejercicios de plataformas en línea y soluciones de Advent of Code.

### Estructura del repositorio

```text
learning/
  <tema>.md               # Notas sobre Go y mi experiencia aprendiendo
  solving/
    <plataforma>/
      <problema>/
        main.go           # Solución del problema

advent-of-code/
  <año>/                   # 2015 hasta 2026
    day-01/
      statement.md        # Enunciado del problema
      input.txt           # Entrada del problema
      part1/
        main.go           # Ejecutable de la parte 1
      part2/
        main.go           # Ejecutable de la parte 2
```

Los archivos directamente dentro de `learning/` son para mis notas de Go: lo que me gusta, lo que no, los conceptos que quiero entender y mis observaciones mientras aprendo. `learning/solving/` es para problemas de plataformas como Codeforces y OmegaUp. Los ejercicios de Advent of Code están separados en `advent-of-code/`.

Cada día de AoC reúne el enunciado y la entrada junto con ejecutables independientes para ambas partes. Al estar en directorios separados, cada parte puede usar `package main` y su propia `func main` sin conflictos.

### Ejecutar una solución de Advent of Code

Desde la raíz del repositorio, ejecuta la parte y pásale la entrada por la entrada estándar:

```sh
go run ./advent-of-code/2026/day-01/part1 < advent-of-code/2026/day-01/input.txt
go run ./advent-of-code/2026/day-01/part2 < advent-of-code/2026/day-01/input.txt
```

### ¿Por qué Go?

Llevo tiempo buscando una alternativa moderna a C++. Consideré Rust, pero prefiero un lenguaje con recolector de basura. Go me parece una buena opción, así que estoy aprendiendo mediante la práctica.
