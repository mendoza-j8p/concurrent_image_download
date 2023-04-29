package main

import (
	"fmt"      // Para imprimir y formatear datos en la salida estándar.
	"io"       // Para trabajar con la entrada/salida, como copiar datos de un lugar a otro.
	"net/http" // Para realizar solicitudes HTTP y recibir respuestas.
	"os"       // Para interactuar con el sistema operativo, como crear y cerrar archivos.
	"sync"     // Para la sincronización de goroutines utilizando WaitGroup.
	"strings"  // Para proporcionar funciones para que manipulen cadenas de texto.
	"bufio"    // Para proporcionar un escáner para leer el archivo línea por línea.
)
func downloadImage(url string, wg *sync.WaitGroup) {
	defer wg.Done() // Marca la finalización de la goroutine cuando la función retorna.

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") { // Verifica si la URL no tiene los prefijos "http://" ni "https://"
		fmt.Printf("URL inválida: %s\n", url) // Muestra un mensaje de URL inválida
		return
	}

	resp, err := http.Get(url) // Realiza una solicitud HTTP GET a la URL especificada.
	if err != nil {
		fmt.Printf("Error al descargar la imagen de %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close() // Cierra la respuesta de la solicitud HTTP al final de la función.

	filename := "downloaded_" + url[strings.LastIndex(url, "/")+1:] // Obtiene el nombre de archivo a partir de la URL.

	file, err := os.Create(filename) // Crea un archivo para guardar la imagen descargada.
	if err != nil {
		fmt.Printf("Error al crear el archivo de imagen: %v\n", err)
		return
	}
	defer file.Close() // Cierra el archivo al final de la función.

	_, err = io.Copy(file, resp.Body) // Copia el contenido de la respuesta HTTP al archivo.
	if err != nil {
		fmt.Printf("Error al guardar la imagen: %v\n", err)
		return
	}

	fmt.Printf("Imagen descargada desde %s\n", url) // Imprime un mensaje indicando que la imagen ha sido descargada.
}

func readURLsFromFile(filename string) ([]string, error) {
	// Abrir el archivo
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("no se pudo abrir el archivo: %v", err)
	}
	defer file.Close() // Cerrar el archivo al finalizar la función

	// Crear un escáner para leer el archivo línea por línea
	scanner := bufio.NewScanner(file)
	urls := make([]string, 0) // Crear un slice para almacenar las URLs

	// Leer el archivo línea por línea
	for scanner.Scan() {
		url := scanner.Text() // Obtener la URL de la línea actual
		urls = append(urls, url) // Agregar la URL al slice
	}

	// Verificar si hubo algún error durante la lectura del archivo
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error al leer el archivo: %v", err)
	}

	return urls, nil // Devolver el slice de URLs y ningún error
}

func main() {
	// Leer las URLs desde el archivo
	urls, err := readURLsFromFile("img_url.txt")
	if err != nil {
		fmt.Printf("Error al leer las URLs: %v\n", err)
		return
	}

	// Crear un WaitGroup para esperar a que todas las goroutines finalicen
	var wg sync.WaitGroup

	// Crear un canal para comunicar el estado de finalización
	done := make(chan struct{})

	// Iterar sobre las URLs y crear una goroutine por cada una
	for _, url := range urls {
		wg.Add(1)
		go downloadImage(url, &wg)
	}

	// Goroutine anonima para cerrar el canal cuando todas las descargas finalicen
	go func() {
		wg.Wait()
		close(done)
	}()

	// Esperar a que todas las descargas finalicen o se cierre el canal
	<-done

	fmt.Println("Todas las descargas finalizadas")
}

