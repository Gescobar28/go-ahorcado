🎮 **Programa de Consola en Go: Ahorcado** 🎮

¡Hola a todos! En esta ocasión me gustaría compartir un proyecto en el que he estado trabajando: un juego de consola desarrollado en Go, conocido como **Ahorcado**. Este juego clásico consiste en adivinar una palabra antes de que se complete la figura del ahorcado.

🔹 **Características del Juego:**

1. **Menú Principal:**
   - **Jugar:** Inicia una partida, mostrando una pantalla de bienvenida con una breve explicación, la longitud de la palabra a adivinar, el número de intentos restantes y la palabra oculta.
   - **Ayuda:** Proporciona una descripción detallada del juego y sus reglas.
   - **Salir:** Cierra el programa.

2. **Mecánica del Juego:**
   - El jugador ingresa letras para adivinar la palabra secreta.
   - Si la letra está en la palabra, se revela en su posición correcta.
   - Si no, se resta un intento y se comienza a dibujar la figura del ahorcado.
   - Las palabras son obtenidas aleatoriamente mediante una solicitud a la API de [Greenborn](https://greenborn.com.ar).
   - Se verifica que las palabras no contengan caracteres especiales; de lo contrario, se realiza una nueva solicitud.

🚀 **Próximos Pasos:**
   - Implementación de un modo multijugador.
   - Añadir opciones para ajustar la dificultad.
   - Introducir un sistema de puntuación para hacer el juego más competitivo.
