# 📈 StocksRefferer

**StocksRefferer** es una aplicación web fullstack desarrollada con Go, Vue 3, TypeScript y TailwindCSS, cuyo propósito es consultar recomendaciones de acciones desde una API externa, almacenarlas en una base de datos remota y visualizarlas a través de una interfaz gráfica filtrable. Todo el proyecto puede desplegarse fácilmente en AWS mediante Terraform.

---

## 🚀 Características principales

- Backend en Go con manejo de rutas y conexión a base de datos usando GORM.
- Frontend en Vue 3 + TypeScript + TailwindCSS, completamente responsivo.
- Sistema de filtros dinámicos por columnas visibles y nivel de riesgo.
- Uso de Pinia como store para gestión de estado.
- Despliegue automatizado en AWS EC2 usando Terraform y NGINX como proxy inverso.
- Variables de entorno administradas con `godotenv`.

---

## 🗂️ Estructura del proyecto

```
stocksRefferer/
├── api/                 # Backend en Go
│   ├── dataRetrieve/    # Scripts para consumir la API externa y subir datos a la BD
│   ├── db/              # Conexión a base de datos con GORM
│   ├── handlers/        # Controladores de los endpoints
│   ├── models/          # Definición de structs para los modelos de datos
│   ├── main.go          # Punto de entrada del backend
│   └── .env             # Variables de entorno (no incluido en git)
│
├── frontend/            # Frontend en Vue 3 con TypeScript y TailwindCSS
│   ├── src/
│   │   ├── assets/      # Estilos personalizados con Tailwind (@apply)
│   │   ├── components/  # Header, FilterDropDown, Slider, StocksTable
│   │   ├── stores/      # Gestión de estado con Pinia
│   │   ├── type/        # Definición de tipos de datos
│   │   ├── views/       # HomeView y Recommendations
│   │   └── App.vue
│   └── index.html
│
├── deploy/              # Infraestructura como código con Terraform
│   ├── main.tf          #Definicion de proveedor y recurso que se crea
│   ├── key.tf           #Asociacion de clave publica al recurso
│   ├── security.tf      #Relgas de entrada y salida de trafico de red
│   ├── .env             #Archivo .env que se subira al EC2
│   └── startup.sh       # Script de instalación y configuración en EC2
│
└── README.md
```

---

## ⚙️ Tecnologías utilizadas

### Backend
- [Go](https://golang.org/)
- [GORM](https://gorm.io/)
- [godotenv](https://github.com/joho/godotenv)

### Frontend
- [Vue 3](https://vuejs.org/)
- [TypeScript](https://www.typescriptlang.org/)
- [TailwindCSS](https://tailwindcss.com/)
- [Pinia](https://pinia.vuejs.org/)

### DevOps / Deploy
- [Terraform](https://www.terraform.io/)
- [AWS EC2](https://aws.amazon.com/ec2/)
- [NGINX](https://www.nginx.com/)

---

## 📦 Instalación local

Debera tener previamente instalado Go, node.js y git para ejecutar estas instrucciones

### 1. Clonar el repositorio
```bash
git clone https://github.com/MoraDiego/stocksRefferer.git
cd stocksRefferer
```

### 2. Configurar el backend
```bash
cd api
nano .env   # Configuracion de variables de entorno
go mod tidy
cd dataRetrieve # Esto solo debes hacerlo una vez, para almacenar los datos en tu BD
go run .
cd ..
go run .
```

### 3. Configurar el frontend
```bash
cd frontend
npm install
npm run dev
```

### 4. Ver la aplicación
Abre tu navegador en `http://localhost:5173` (o el puerto configurado).

---

## ☁️ Despliegue en AWS

1. Asegúrate de tener configurado AWS CLI y permisos adecuados.
2. En la carpeta `deploy`, ejecuta:

```bash
terraform init
terraform apply
```

3. El script `startup.sh` se encargará de:
   - Instalar dependencias.
   - Ejecutar backend y frontend.
   - Configurar NGINX para enrutar las peticiones.

4. Actualmente el proyecto esta configurado para funcionar sin el servidor NGINX, deberas modificar las funciones de la carpteta `frontend/src/request/request.ts` para que las peticiones se hagan a la ruta relativa `/api/`, agregando en nombre del endpoint correspondiente y sus parametros necesarios, de lo contrario se ejecutara la peticion en la maquina del cliente y no se mostraran los datos del backend en go

---

## 📋 Variables de entorno

El archivo `.env` para el backend debe incluir las variables:

```
API_TOKEN=
API_URL=
DB_DSN=
```

(No incluyas este archivo en el control de versiones por seguridad.)

---

## 👨‍💻 Autor

**Diego Mora**  
[GitHub](https://github.com/MoraDiego)  
[LinkedIn](https://www.linkedin.com/in/diego-mora-1563a42bb/)