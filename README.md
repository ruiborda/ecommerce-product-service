# Servicio de Productos de E-commerce

Este repositorio contiene el servicio API para la gestión de productos de e-commerce, desarrollado con Go y Gin framework.

## Requisitos

- Go 1.24 o superior
- Configuración de variables de entorno (ver `example.dev.sh`)

## Ejecución en local

1. Copia el archivo `example.dev.sh` a `dev.sh` y configura tus variables de entorno:

```bash
cp example.dev.sh dev.sh
# Edita dev.sh con tus valores privados
chmod +x dev.sh
```

2. Ejecuta el script:

```bash
./dev.sh
```

## Ejecución con Docker Compose

1. Copia el archivo `example.docker.env` a `.env` y configura tus variables:

```bash
cp example.docker.env .env
# Edita .env con tus valores privados
```

2. Levanta el contenedor:

```bash
docker-compose up -d
```

## CI/CD

Este proyecto utiliza CI/CD para automatizar el despliegue. La configuración se encuentra en `.github/workflows/ci.yml`. 

Asegúrate de configurar los secretos necesarios en tu repositorio GitHub para que el pipeline funcione correctamente.

## Documentación de la API

📄 Documentación de la API: http://localhost:8080/

Una vez que el servidor esté en ejecución, puedes acceder a la interfaz Swagger UI para explorar y probar todos los endpoints disponibles.

## Licencia

MIT