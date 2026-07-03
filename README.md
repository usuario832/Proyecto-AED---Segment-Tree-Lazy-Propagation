# Proyecto AED: Segment Tree con Lazy Propagation

## Configuración de la Base de Datos
El proyecto requiere una base de datos SQLite para funcionar. Sigue estos pasos para configurarla:

1. **Ubicación del archivo:** Asegúrate de que el archivo de base de datos llamado `VentasTambo.db` se encuentre en la carpeta raíz del proyecto (donde está el archivo `main.go`).

2. **Dependencias:** El proyecto utiliza `modernc.org/sqlite` como driver puro de Go. Si es la primera vez que clonas el proyecto, ejecuta el siguiente comando en la terminal dentro de la carpeta del proyecto para instalar las dependencias necesarias:

   ```bash
   go get modernc.org/sqlite
    ```

3. **Ejecuta en la Terminal:**

    ```bash
   go run .
    ```