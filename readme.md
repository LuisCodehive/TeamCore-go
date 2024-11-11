# Despliegue de Servicio Go en Google Cloud Platform con Cloud Run

Este proyecto contiene un servicio en Go que se despliega en Google Cloud Platform (GCP) utilizando Cloud Run.

## Prerrequisitos

- **Cuenta de GCP**: Asegúrate de tener una cuenta en Google Cloud Platform.
- **Proyecto de GCP**: Crea un proyecto en GCP y anota el **Project ID**.
- **Google Cloud SDK**: Instala el [Google Cloud SDK](https://cloud.google.com/sdk/docs/install) y autentícate:
  ```bash
  gcloud auth login



## Habilita Cloud Run
  
  ```bash
    gcloud services enable run.googleapis.com

```


## Configura el proyecto con tu Project ID
  ```bash
gcloud config set project [PROJECT_ID]

```
## Construye y sube la imagen de Docker a GCR
  ```bash
gcloud builds submit --tag gcr.io/[PROJECT_ID]/my-go-service

```
### Despliega el Servicio en Cloud Run

  ```bash
gcloud run deploy my-go-service \
  --image gcr.io/[PROJECT_ID]/my-go-service \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated

```