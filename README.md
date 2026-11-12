# Learning Go with Advent of Code

## English

This repository is my learning journal as I study Go, with the goal of becoming comfortable enough with the language to solve Advent of Code problems. It is a personal record of what I learn, what I enjoy, what I find challenging, and how my approach changes over time.

### Repository structure

I plan to keep language notes and Advent of Code work in separate areas:

```text
learning/
  <topic>/
    README.md       # Notes on the topic and my experience with it

advent-of-code/
  2026/
    day-01/
      main.go       # Solution
      notes.md      # Reasoning, problem notes, and Go observations
```

Each topic under `learning/` can have its own Markdown file for notes about the parts of Go I like, dislike, or want to understand better. Each Advent of Code day can have its own directory with the solution and notes about my thought process, the puzzle, and the language.

This is a suggested layout; I may adjust it as the project grows.

### Running a solution

The exact convention is still to be decided. One straightforward option is to make each day a Go package with its own `main.go`, then run a day from the repository root with:

```sh
go run ./advent-of-code/2026/day-01
```

Once the directory structure and entry points are in place, I’ll update this section with the commands that work for the project.

### Why Go?

I’ve spent a long time looking for a modern alternative to C++. I considered Rust, but I prefer a language with garbage collection. Go seems like a good fit, so I’m using this project to learn it through practice.

---

## Español

Este repositorio es mi bitácora de aprendizaje mientras estudio Go, con la meta de sentirme suficientemente cómodo con el lenguaje para resolver problemas de Advent of Code. Aquí iré registrando lo que aprendo, lo que me gusta, lo que me cuesta y cómo cambia mi forma de abordar los problemas.

### Estructura del repositorio

La idea es mantener las notas sobre el lenguaje y el trabajo de Advent of Code en áreas separadas:

```text
learning/
  <tema>/
    README.md       # Notas sobre el tema y mi experiencia

advent-of-code/
  2026/
    day-01/
      main.go       # Solución
      notes.md      # Razonamiento, notas del problema y observaciones sobre Go
```

Cada tema dentro de `learning/` puede tener su propio archivo Markdown con notas sobre las partes de Go que me gustan, no me gustan o quiero entender mejor. Cada día de Advent of Code puede tener su propio directorio con la solución y notas sobre mi razonamiento, el problema y el lenguaje.

Esta estructura es una propuesta; puedo ajustarla a medida que crezca el proyecto.

### Cómo ejecutar una solución

Todavía no he definido una convención. Una opción sencilla es que cada día sea un paquete de Go con su propio `main.go`; desde la raíz del repositorio se podría ejecutar así:

```sh
go run ./advent-of-code/2026/day-01
```

Cuando estén definidos los directorios y puntos de entrada, actualizaré esta sección con los comandos que use el proyecto.

### ¿Por qué Go?

Llevo tiempo buscando una alternativa moderna a C++. Consideré Rust, pero prefiero un lenguaje con recolector de basura. Go me parece una buena opción, así que estoy aprendiendo mediante la práctica en este proyecto.
