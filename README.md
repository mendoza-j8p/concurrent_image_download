# Práctica: Descarga concurrente de imágenes

**Descripción:**

Escribe un programa en Go que realice la descarga concurrente de imágenes desde una lista de URLs. El programa utilizará goroutines y sincronización basada en sync.WaitGroup para descargar las imágenes de forma concurrente. Se utilizará el método Done del sync.WaitGroup para marcar la finalización de cada goroutine de descarga.

## Pasos para la descarga concurrente de imágenes

1. Crea una lista de URLs de imágenes que deseas descargar. Puedes tener una lista de URLs predefinida o leerlas desde un archivo o entrada del usuario.

2. Define una función que tome una URL como argumento y descargue la imagen correspondiente utilizando la biblioteca net/http. Puedes utilizar http.Get() para realizar la solicitud HTTP y guardar el contenido de la respuesta en un archivo local.

3. Crea un objeto de tipo `sync.WaitGroup` para rastrear el estado de finalización de las descargas.

4. Por cada URL en la lista, incrementa el contador del `sync.WaitGroup` antes de iniciar una goroutine de descarga. Dentro de cada goroutine, después de completar la descarga, llama al método `Done()` del `sync.WaitGroup` para indicar que la descarga se ha completado.

5. Llama al método `Wait()` del `sync.WaitGroup` después de iniciar todas las goroutines para bloquear el programa principal hasta que todas las descargas hayan finalizado.

6. Al finalizar todas las descargas, muestra un mensaje indicando el número total de descargas realizadas.

Este enfoque utiliza el `sync.WaitGroup` para esperar a que todas las goroutines de descarga se completen sin la necesidad de utilizar canales para la comunicación entre goroutines.

..

..

## Puntos clave

El programa que desarrollamos es un sistema para descargar imágenes desde URLs proporcionadas en un archivo.

A continuación, se detallan los puntos más importantes de su funcionamiento:

- **Lectura de URLs:** El programa lee las URLs desde un archivo. Para ello, implementamos una función "readURLsFromFile" que utiliza la librería **"os"** para abrir y leer el archivo
línea por línea. Las URLs se almacenan en un slice para su posterior procesamiento.

- **Descarga de imágenes:** Definimos una función llamada "downloadImage" que utiliza la librería "net/http" para descargar una imagen desde una URL. La función toma la URL y un puntero
a un objeto **sync.WaitGroup**, que nos permite esperar a que todas las goroutines finalicen antes de continuar. La descarga se realiza utilizando las funciones http.Get para obtener la
respuesta HTTP y io.Copy para copiar el contenido de la respuesta en un archivo.

- **Concurrencia:** El programa aprovecha la concurrencia para descargar las imágenes de forma paralela. Utilizamos un **sync.WaitGroup** para esperar a que todas las descargas finalicen antes
de continuar. Iteramos sobre las URLs y **creamos una goroutine para cada una**, utilizando la función "downloadImage".
Cada goroutine descarga una imagen y marca su finalización utilizando **wg.Done()** en el objeto WaitGroup.

- **Canal de finalización:** Para controlar el momento en que todas las descargas han finalizado, **creamos un canal llamado done.** Después de iniciar las goroutines de descarga, **creamos una goroutine adicional** que espera a que todas las goroutines finalicen utilizando **wg.Wait()** y luego cierra el canal done mediante **close(done).**
En la función principal, utilizamos **<-done** para esperar hasta que se cierre el canal y se indique que todas las descargas han finalizado.

- **Manejo de errores:** El programa implementa manejo de errores para detectar problemas durante la lectura de URLs desde el archivo, la descarga de imágenes y otros errores potenciales.
Se utilizan instrucciones **if err != nil para verificar si se producen errores y se devuelve un error detallado si es necesario.

En resumen, el programa lee las URLs desde un archivo, descarga las imágenes de manera concurrente utilizando goroutines, y utiliza un canal y un WaitGroup para sincronizar y controlar el estado de finalización.
