# Gestión de Tareas con API REST en Go y Gin

La empresa PragmaTask necesita una API REST para gestionar tareas en su plataforma de gestión de proyectos. La API debe permitir crear, listar y actualizar tareas. El sistema tiene los siguientes actores: el Administrador de Proyectos, que crea y actualiza las tareas, y el Visualizador de Tareas, que lista las tareas disponibles. La API debe asegurar que las tareas creadas tengan un título único y una descripción válida. En caso de error, debe retornar un mensaje claro indicando el problema.

## Informacion General

| Campo | Valor |
|-------|-------|
| **Tema** | API REST de tareas con Go y Gin |
| **Nivel** | senior-l2 |
| **Tipo** | practical |
| **Tiempo estimado** | 8 horas |

## Fases del Reto

### Fase 0: Configuración del Proyecto

**Objetivo:** Obtener el proyecto base funcional enviando el Código Base a un asistente de IA, que lo analizará, corregirá errores y entregará cada archivo con su ruta.

**Tiempo estimado:** 15-30 minutos

**Instrucciones:**

- Asegúrate de tener instalado para ejecutar el proyecto: Go 1.21+, VS Code o GoLand.
- Copia todo el contenido del campo **Código Base** de este reto — incluyendo el texto de instrucciones que aparece al inicio.
- Abre un asistente de IA (Claude en claude.ai, ChatGPT o Gemini — se recomienda Claude), pega el contenido copiado en el chat y envíalo.
- El asistente analizará los archivos, corregirá errores y entregará cada archivo con su ruta. Guárdalos en la carpeta donde quieras trabajar, respetando las rutas.
- Ejecuta `go build ./...`. Si no hay errores, estás listo.

**Entregable:** El proyecto compila/arranca sin errores.

<details>
<summary>Pistas de conocimiento</summary>

- Copia el Código Base completo incluyendo el texto de instrucciones al inicio — esas instrucciones le indican al asistente exactamente qué hacer con los archivos.
- Si el asistente no entrega todos los archivos al terminar el análisis, escríbele: "entrégame todos los archivos con su ruta".
- Si el proyecto tiene errores al arrancar, comparte el mensaje de error con el mismo asistente para que lo corrija.

</details>

### Fase 1: Creación de la API

**Objetivo:** Implementar los endpoints básicos para crear y listar tareas.

**Tiempo estimado:** 3 horas

**Instrucciones:**

- Diseña y crea los endpoints REST para crear y listar tareas.
- Asegura que las tareas creadas tengan un título único y una descripción válida.
- Maneja los errores de validación y retorna mensajes claros.

**Entregable:** Endpoints REST funcionales para crear y listar tareas.

<details>
<summary>Pistas de conocimiento</summary>

- Recuerda que el título de la tarea debe ser único.
- La descripción de la tarea debe ser válida y no estar vacía.

</details>

### Fase 2: Actualización de Tareas

**Objetivo:** Implementar el endpoint para actualizar tareas.

**Tiempo estimado:** 3 horas

**Instrucciones:**

- Diseña y crea el endpoint REST para actualizar tareas.
- Asegura que las actualizaciones respeten las reglas de validación establecidas.
- Maneja los errores de validación y retorna mensajes claros.

**Entregable:** Endpoint REST funcional para actualizar tareas.

<details>
<summary>Pistas de conocimiento</summary>

- Recuerda que las actualizaciones deben respetar las reglas de validación.
- Maneja los errores de validación y retorna mensajes claros.

</details>

### Fase 3: Optimización y Refactorización

**Objetivo:** Optimizar y refactorizar el código para mejorar la legibilidad y mantenibilidad.

**Tiempo estimado:** 2 horas

**Instrucciones:**

- Revisa el código implementado y realiza las optimizaciones necesarias.
- Refactoriza el código para mejorar la legibilidad y mantenibilidad.
- Asegura que el código cumpla con las mejores prácticas de desarrollo.

**Entregable:** Código refactorizado y optimizado.

<details>
<summary>Pistas de conocimiento</summary>

- Recuerda las mejores prácticas de desarrollo en Go.
- Utiliza patrones de diseño adecuados para mejorar la legibilidad y mantenibilidad del código.

</details>

## Dimensiones Evaluadas

- **queEs**: ¿Qué son los endpoints REST y cuál es su propósito en la API de gestión de tareas?
- **paraQueSirve**: ¿Para qué sirve cada uno de los endpoints implementados en la API?
- **comoSeUsa**: ¿Cómo se usan los endpoints para crear, listar y actualizar tareas?
- **erroresComunes**: ¿Cuáles son los errores comunes que pueden ocurrir al usar los endpoints y cómo se manejan?
- **queDecisionesImplica**: ¿Qué decisiones de diseño implica la implementación de los endpoints y cómo afecta la legibilidad y mantenibilidad del código?

## Criterios de Evaluacion

- Implementación correcta de los endpoints REST para crear, listar y actualizar tareas.
- Aseguramiento de la unicidad del título y validez de la descripción de las tareas.
- Manejo adecuado de los errores de validación y retorno de mensajes claros.
- Optimización y refactorización del código para mejorar la legibilidad y mantenibilidad.

## Como trabajar con un asistente de IA

Hay dos caminos, elegi uno:

- **AGENTS.md** (recomendado) — instrucciones nativas del repo. Abri esta carpeta con tu agente local (Claude Code, Cursor, Codex, Copilot, Gemini) y las carga solo. Sabe que archivos faltan y con que comando se verifica, y completa el scaffold escribiendo en disco.
- **PROMPT_MEJORA.md** — para copiar y pegar en un chat (claude.ai, ChatGPT) si no tenes un agente en el IDE. Declara lo mismo que AGENTS.md; los archivos que te devuelva el chat los guardas en esta carpeta respetando las rutas.

Ninguno de los dos resuelve las fases del reto: eso es tu trabajo.

## Verificacion

El proyecto esta listo para trabajar cuando este comando corre sin errores:

```bash
go mod tidy && go build ./...
```

---

*Reto generado automaticamente por Challenge Generator - Pragma*
