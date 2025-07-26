
# 🕸️ Diferentes Estilos

![Badge](https://img.shields.io/badge/License-MIT-blue) ![Badge](https://img.shields.io/badge/Go-1.20+-blue) ![Badge](https://img.shields.io/badge/Status-Working-green)

Servidor web minimalista en Go que actúa como proxy para `https://senchpimy.xyz`, sirviendo archivos HTML y CSS locales o remotos con una lógica sencilla.

## ✨ Características

- Servidor HTTP local en el puerto `3002`.
- Redirige solicitudes `.html` y `/` al contenido remoto en `https://senchpimy.xyz`.
- Sirve archivos CSS locales dependiendo del path solicitado.
- Añade automáticamente una etiqueta `<meta viewport>` para mejorar la visualización en móviles.

## 🚀 Uso


### 1. Clona el repositorio

```bash
git clone https://github.com/tuusuario/senchpimy-proxy.git
cd senchpimy-proxy
````

### 2. Ejecuta el servidor

```bash
go run main.go
```

Visita [`http://localhost:3002`](http://localhost:3002) en tu navegador.

## 🧠 Funcionamiento

* Si la ruta solicitada termina en `.css`, se sirve el archivo local correspondiente (`style.css` o `style_ascci.css`).
* Si termina en `.html` o es la raíz (`/`), se obtiene el HTML desde el dominio remoto y se le añade la etiqueta de viewport.
* Otras rutas se redirigen tal cual al contenido remoto y se devuelven como texto plano.

## 📁 Estructura

```
.
├── main.go
├── style.css
└── style_ascci.css
```

## ⚠️ Notas

* El dominio de origen está fijado como constante: `https://senchpimy.xyz`.
* No se hace validación de contenido o cabeceras MIME, esto se puede mejorar para producción.

## 📄 Licencia

Este proyecto está bajo la licencia [MIT](LICENSE).
