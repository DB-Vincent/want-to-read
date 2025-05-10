# Want to Read

![Example of the want to read tool](./docs/example.png)

A web application to manage your reading list and track books you want to read.

## Backend

### OpenAPI docs

Generate docs: `~/go/bin/swag init -g cmd/api/main.go`

### Running locally

Start the backend app: `cd backend && go run cmd/api/main.go`

## Frontend

### Running locally

Start the frontend app: `cd frontend && npm run start`

## Deploying

This directory contains Kubernetes manifests for deploying want-to-read using [Kustomize](https://kustomize.io/) (supported natively by `kubectl`).

### Folder Structure

```
deploy/
├── base/ # Generic Kubernetes manifests
│   ├── backend-deployment.yaml
│   ├── backend-pvc.yaml
│   ├── backend-service.yaml
│   ├── frontend-deployment.yaml
│   ├── frontend-service.yaml
│   ├── ingress.yaml
│   └── kustomization.yaml
└── overlays/
    └── production/ # Production-specific kustomizations
        └── kustomization.yaml
```

### Prerequisites

- Access to a Kubernetes cluster.
- `kubectl` installed (with Kustomize support, v1.14+).

### Configuration

Before deploying, you **must** update the hostname in `overlays/production/kustomization.yaml` to match your desired domain.

You may also set the `ingressClassName` in the overlay’s `kustomization.yaml` if your cluster uses a specific ingress controller.

## Deployment Steps

1. **Clone the repository** (if you haven’t already):

    ```
    git clone git@github.com:DB-Vincent/want-to-read.git
    cd want-to-read/deploy
    ```

2. **Edit the hostname** in `overlays/production/kustomization.yaml`:

    ```
    # Example snippet
    - op: replace
      path: /spec/rules/0/host
      value: some-hostname.local
    ...
    ```

3. **(Optional) Set the ingress class** in `overlays/production/kustomization.yaml` if needed:

    ```
    # Example snippet
    - op: add
      path: /spec/ingressClassName
      value: nginx
    ...
    ```

4. **Deploy to your cluster** (replace `<namespace>` if you want to use a specific namespace):

    ```
    kubectl apply -k overlays/production
    # or, to deploy into a specific namespace:
    kubectl apply -k overlays/production -n <namespace>
    ```

### Accessing the Application

Once deployed, your application will be available at the hostname you configured in the ingress resource. Make sure your DNS points to your cluster’s ingress controller.
