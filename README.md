#  Investment - Get Financial information


## Índice
- [Características](#características)
- [Instalación](#instalación)
- [Uso](#uso)
- [Contribución](#contribución)
- [Licencia](#licencia)

## Características
Obtiene informacion desde la bolsa de comercio de los siguientes topicos
-   Dividendos publicados
-   Resumen de mercado
-   other feat.
## Instalación
Proporciona instrucciones claras y concisas sobre cómo instalar tu proyecto. Puedes incluir comandos específicos, requisitos previos y pasos adicionales si es necesario.

Si quieres ver un ejemplo completo del archivo .env puedes ir a la siguiente ruta ./example/.env-example-full

```bash
# Ejemplo de comandos de instalación
go build .

```


GoogleSheets   
 
Si deseas guardar tus datos en un archivo de GoogleSheets (Aka Excel de google) puedes configurar en la consola una cuenta de servicio para que acceda a tu archivo y almacene tus datos de inversiones.
```
Habilitar la API de Google Sheets:  

    -   Ve a la Consola de Desarrolladores de Google (https://console.cloud.google.com/).
    -   Crea un nuevo proyecto o selecciona uno existente.
    -   Habilita la API de Google Sheets para tu proyecto.
    -   Crea credenciales para la API, selecciona el tipo de cuenta de servicio y descarga el archivo JSON con tus credenciales.
```
MongoDB   

Se puede generar un cluster free-tier de mongo DB para almacenar tus datos de inversiones.
```  
https://cloud.mongodb.com
```
## Uso
Para usar este programa puedes ejecutar el compilado  que se genera  de la siguiente forma :

```bash
#Debes tener en tu carpeta ./recursos el driver de chromiun 
# Ejecutar
webscrapping-investment.exe
```
## Contribución
¡Gracias por considerar contribuir al proyecto! A continuación, se presentan algunas pautas para contribuir:

Forkea el proyecto y clónalo localmente.
Crea una nueva rama para tu contribución.
Realiza tus cambios y asegúrate de seguir las pautas de estilo.
Haz pruebas para asegurarte de que todo funciona como se espera.
Envía una solicitud de extracción a la rama principal del proyecto.
## Licencia
Este proyecto está bajo la Licencia MIT. Consulta el archivo LICENSE para obtener más detalles.