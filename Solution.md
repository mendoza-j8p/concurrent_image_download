# steps

**1. Importa los paquetes necesarios.**

**2. Define una función para descargar una imagen desde una URL.**

1. La función downloadImage recibe dos parámetros: url de tipo string, que representa la URL de la imagen a descargar, y wg
de tipo *sync.WaitGroup, que es un puntero a un objeto WaitGroup utilizado para la sincronización de goroutines.

2. La declaración defer wg.Done() se coloca al comienzo de la función utilizando la declaración defer. Esta línea marca la
finalización de la goroutine cuando la función downloadImage retorna. Es decir, cuando la función finaliza, se notifica
al WaitGroup que la goroutine ha terminado.

3. Se verifica si la URL no tiene los prefijos "http://" ni "https://".
Esto se hace utilizando la función strings.HasPrefix(url, "http://") && strings.HasPrefix(url, "https://").
Si la condición es falsa, significa que la URL no tiene uno de esos prefijos, lo que indica que la URL es inválida.

4. La función http.Get(url) realiza una solicitud HTTP GET a la URL especificada y devuelve una respuesta HTTP resp y un
posible error err. Esto descarga la imagen desde la URL.

5. Se utiliza la declaración defer resp.Body.Close() para asegurarse de que el cuerpo de la respuesta HTTP se cierre al
final de la función. Esto libera los recursos asociados con la respuesta.

6. Se genera el nombre del archivo de destino concatenando "downloaded_" con el nombre del archivo extraído de la URL.
Para ello, se utiliza strings.LastIndex(url, "/")+1 para obtener el índice de la última aparición de "/" en la URL y
se añade 1 para obtener el índice del primer carácter después de "/". Luego, se utiliza este índice para extraer el
nombre del archivo de la URL.

7. Se crea un archivo utilizando os.Create(image_file) para guardar la imagen descargada. Se verifica si se produce algún
error y, en caso afirmativo, se muestra un mensaje de error y se retorna desde la función.
En esa línea de código, el guion bajo ("_") se utiliza para descartar el valor de retorno de io.Copy(), que es el número
total de bytes copiados. Esto significa que no estamos interesados en ese valor y simplemente lo estamos descartando.

8. Utilizando defer, se cierra el archivo al final de la función para asegurarse de que se liberen los recursos asociados.

9. Se utiliza io.Copy(file, resp.Body) para copiar el contenido de la respuesta HTTP (la imagen descargada) al archivo. Se verifica si se produce
algún error y, en caso de que se produzca algún error durante esta operación se muestra un mensaje de error y se retorna desde la función.

En resumen, la función downloadImage es responsable de descargar una imagen desde una URL, guardarla en un archivo y marcar su finalización en el
WaitGroup. Esto permite que el programa principal coordine y espere a que todas las descargas se completen antes de continuar.

**3. Define una función auxiliar readURLsFromFile para leer las URLs desde el archivo.**

1. Se define la función readURLsFromFile que recibe un parámetro filename de tipo string que representa el nombre del archivo a leer. La función devuelve
un slice de strings que contiene las URLs leídas y un error en caso de que ocurra algún problema.

2. Se intenta abrir el archivo utilizando os.Open(filename). Si ocurre algún error en la apertura del archivo, se retorna un error utilizando fmt.Errorf("no se pudo abrir el archivo: %v", err).

3. Se utiliza la declaración defer file.Close() para asegurarse de que el archivo se cierre al finalizar la función. Esto garantiza que los recursos asociados al archivo se liberen adecuadamente.

4. Se crea un escáner (scanner) para leer el archivo línea por línea utilizando bufio.NewScanner(file). El escáner se inicializa con el archivo abierto anteriormente.

5. Se crea un slice llamado urls de tipo []string que se utilizará para almacenar las URLs leídas del archivo. Al inicio, el slice está vacío.

6. Se comienza a leer el archivo línea por línea utilizando un bucle for y scanner.Scan(). En cada iteración del bucle, se lee una línea del archivo.

7. Dentro del bucle, se obtiene la URL de la línea actual utilizando scanner.Text(). Esta función devuelve el contenido de la línea como un string.

8. Se agrega la URL al slice de urls utilizando la función append(urls, url). Esto permite almacenar la URL en el slice.

9. Después de leer todas las líneas del archivo, se verifica si hubo algún error durante la lectura utilizando scanner.Err(). Si se produce algún error, se retorna un error 
utilizando fmt.Errorf("error al leer el archivo: %v", err).

10. Finalmente, si no hubo ningún error, se devuelve el slice de URLs (urls) y se retorna nil como error para indicar que la lectura del archivo se realizó correctamente.

En resumen, esta función abre un archivo, lee las URLs línea por línea y las almacena en un slice. 
Luego verifica si hubo algún error durante la lectura y devuelve el slice de URLs junto con un error en caso de ser necesario.

**4. Crea una función principal main para leer las URLs desde el archivo y realizar la descarga concurrente.**

1. Se llama a la función readURLsFromFile para leer las URLs desde el archivo "img_url". El resultado se guarda en las variables urls y err.
Si ocurre un error durante la lectura del archivo, se muestra un mensaje de error y se finaliza la ejecución del programa.

2. Se crea un WaitGroup llamado wg para esperar a que todas las goroutines finalicen. Esto se utilizará para coordinar y sincronizar el final de las 
descargas de imágenes.

3. Se crea un canal done para comunicar el estado de finalización. Este canal se utilizará para indicar cuando todas las descargas hayan finalizado.

4. Se itera sobre las URLs obtenidas y con el ciclo for se crea una goroutine para descargar cada imagen en paralelo.
Para cada URL, se llama a la función downloadImage en una goroutine separada, pasando la URL y el puntero al WaitGroup wg.

5. Se crea una goroutine anónima para cerrar el canal done cuando todas las descargas hayan finalizado. Esto se logra esperando a que el WaitGroup wg se 
complete utilizando el método Wait(), y luego cerrando el canal.

6. Se espera a que todas las descargas finalicen o se cierre el canal done utilizando la operación <-done. Esto bloqueará la ejecución hasta que se cierre 
el canal, lo que indica que todas las descargas han finalizado.

7. Se imprime un mensaje indicando que todas las descargas han finalizado.

En resumen, el código en la función main lee las URLs desde un archivo, crea goroutines para descargar las imágenes en paralelo, y luego espera a que 
todas las descargas finalicen antes de imprimir un mensaje de finalización.

**5. Ejecuta el programa y verifica la descarga concurrente de las imágenes.**

Ejecuta el programa con el comando: go run <nombre-archivo.go>

Se realizará la descarga concurrente de las imágenes.
Verifica los mensajes de registro para confirmar que las descargas se han realizado correctamente.
Se mostrarán mensajes de error en caso de que ocurra algún problema durante la descarga o si se encuentra una URL inválida.

En resumen, ejecutar el programa y observar la descarga concurrente de las imágenes te permitirá comprobar el funcionamiento correcto de la lógica de descarga,la sincronización de goroutines y la captura de errores, asegurando así que las imágenes sean descargadas de manera eficiente y precisa.
