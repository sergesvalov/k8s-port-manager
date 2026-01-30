# Kubernetes NodePort Manager

**NodePort Manager** is a lightweight tool designed to help DevOps engineers and developers manage the chaos of Kubernetes NodePorts (range `30000-32767`).

It scans the cluster for occupied ports, allows manual reservation of ports for future use (or non-K8s services), and provides a clean Web UI to visualize port usage.

## 🚀 Features

* **Real-time Kubernetes Sync:** Automatically detects NodePorts occupied by `NodePort` or `LoadBalancer` services in the cluster.
* **Manual Reservation:** Reserve a specific port via UI or API to prevent others from using it.
* **Conflict Prevention:** Prevents reserving a port that is already taken by Kubernetes.
* **Web UI:** Simple interface built with Vue.js to search, filter, and view port status.
* **Persistence:** Stores manual reservations in a JSON file (persisted via PVC).
* **REST API:** Simple API to integrate with CI/CD pipelines.

---

## 📂 Project Structure

```text
port-manager/
├── charts/                 # Helm charts
│   └── port-manager/
├── templates/              # HTML frontend
│   └── index.html
├── Dockerfile              # Multi-stage build
├── main.go                 # Go backend logic
└── go.mod                  # Go dependencies

🛠 Prerequisites

    Kubernetes Cluster (v1.19+)

    Helm (v3+)

    Docker (to build the image)

    Go 1.21+ (optional, for local development)

📦 Installation & Deployment
1. Build and Push the Docker Image

Since this is a custom tool, you need to build the image and push it to your container registry.

# 1. Build the image
docker build -t [my-registry.com/ops/port-manager:latest](https://my-registry.com/ops/port-manager:latest) .

# 2. Push to your registry
docker push [my-registry.com/ops/port-manager:latest](https://my-registry.com/ops/port-manager:latest)

2. Configure Helm Chart

Edit charts/port-manager/values.yaml to point to your image:

image:
  repository: [my-registry.com/ops/port-manager](https://my-registry.com/ops/port-manager)
  tag: latest
  pullPolicy: Always

storage:
  size: 1Gi
  # storageClassName: "standard" # Uncomment if you need a specific class

  3. Deploy to Kubernetes

Install the chart using Helm. This will create the Deployment, Service, PVC, and RBAC roles required to read K8s Services.

helm install port-manager ./charts/port-manager \
  --create-namespace \
  --namespace tools

  4. Access the UI

To access the interface locally without an Ingress:

kubectl port-forward svc/port-manager 8080:80 -n tools

Open your browser at: http://localhost:8080

Вот полный файл README.md с корректной разметкой Markdown. Вы можете скопировать этот блок и сохранить его как README.md в корне вашего проекта.
Markdown

# Kubernetes NodePort Manager

**NodePort Manager** is a lightweight tool designed to help DevOps engineers and developers manage the chaos of Kubernetes NodePorts (range `30000-32767`).

It scans the cluster for occupied ports, allows manual reservation of ports for future use (or non-K8s services), and provides a clean Web UI to visualize port usage.

## 🚀 Features

* **Real-time Kubernetes Sync:** Automatically detects NodePorts occupied by `NodePort` or `LoadBalancer` services in the cluster.
* **Manual Reservation:** Reserve a specific port via UI or API to prevent others from using it.
* **Conflict Prevention:** Prevents reserving a port that is already taken by Kubernetes.
* **Web UI:** Simple interface built with Vue.js to search, filter, and view port status.
* **Persistence:** Stores manual reservations in a JSON file (persisted via PVC).
* **REST API:** Simple API to integrate with CI/CD pipelines.

---

## 📂 Project Structure

```text
port-manager/
├── charts/                 # Helm charts
│   └── port-manager/
├── templates/              # HTML frontend
│   └── index.html
├── Dockerfile              # Multi-stage build
├── main.go                 # Go backend logic
└── go.mod                  # Go dependencies

🛠 Prerequisites

    Kubernetes Cluster (v1.19+)

    Helm (v3+)

    Docker (to build the image)

    Go 1.21+ (optional, for local development)

📦 Installation & Deployment
1. Build and Push the Docker Image

Since this is a custom tool, you need to build the image and push it to your container registry.
Bash

# 1. Build the image
docker build -t [my-registry.com/ops/port-manager:latest](https://my-registry.com/ops/port-manager:latest) .

# 2. Push to your registry
docker push [my-registry.com/ops/port-manager:latest](https://my-registry.com/ops/port-manager:latest)

2. Configure Helm Chart

Edit charts/port-manager/values.yaml to point to your image:
YAML

image:
  repository: [my-registry.com/ops/port-manager](https://my-registry.com/ops/port-manager)
  tag: latest
  pullPolicy: Always

storage:
  size: 1Gi
  # storageClassName: "standard" # Uncomment if you need a specific class

3. Deploy to Kubernetes

Install the chart using Helm. This will create the Deployment, Service, PVC, and RBAC roles required to read K8s Services.
Bash

helm install port-manager ./charts/port-manager \
  --create-namespace \
  --namespace tools

4. Access the UI

To access the interface locally without an Ingress:
Bash

kubectl port-forward svc/port-manager 8080:80 -n tools

Open your browser at: http://localhost:8080
💻 Local Development

You can run the service locally on your machine.

    With K8s connection: If you have ~/.kube/config set up, the app will try to connect to your current cluster context.

    Without K8s: If no cluster is found, it starts in "Mock Mode" (empty K8s data).

# Install dependencies
go mod tidy

# Run the app
go run main.go

Вот полный файл README.md с корректной разметкой Markdown. Вы можете скопировать этот блок и сохранить его как README.md в корне вашего проекта.
Markdown

# Kubernetes NodePort Manager

**NodePort Manager** is a lightweight tool designed to help DevOps engineers and developers manage the chaos of Kubernetes NodePorts (range `30000-32767`).

It scans the cluster for occupied ports, allows manual reservation of ports for future use (or non-K8s services), and provides a clean Web UI to visualize port usage.

## 🚀 Features

* **Real-time Kubernetes Sync:** Automatically detects NodePorts occupied by `NodePort` or `LoadBalancer` services in the cluster.
* **Manual Reservation:** Reserve a specific port via UI or API to prevent others from using it.
* **Conflict Prevention:** Prevents reserving a port that is already taken by Kubernetes.
* **Web UI:** Simple interface built with Vue.js to search, filter, and view port status.
* **Persistence:** Stores manual reservations in a JSON file (persisted via PVC).
* **REST API:** Simple API to integrate with CI/CD pipelines.

---

## 📂 Project Structure

```text
port-manager/
├── charts/                 # Helm charts
│   └── port-manager/
├── templates/              # HTML frontend
│   └── index.html
├── Dockerfile              # Multi-stage build
├── main.go                 # Go backend logic
└── go.mod                  # Go dependencies

🛠 Prerequisites

    Kubernetes Cluster (v1.19+)

    Helm (v3+)

    Docker (to build the image)

    Go 1.21+ (optional, for local development)

📦 Installation & Deployment
1. Build and Push the Docker Image

Since this is a custom tool, you need to build the image and push it to your container registry.
Bash

# 1. Build the image
docker build -t [my-registry.com/ops/port-manager:latest](https://my-registry.com/ops/port-manager:latest) .

# 2. Push to your registry
docker push [my-registry.com/ops/port-manager:latest](https://my-registry.com/ops/port-manager:latest)

2. Configure Helm Chart

Edit charts/port-manager/values.yaml to point to your image:
YAML

image:
  repository: [my-registry.com/ops/port-manager](https://my-registry.com/ops/port-manager)
  tag: latest
  pullPolicy: Always

storage:
  size: 1Gi
  # storageClassName: "standard" # Uncomment if you need a specific class

3. Deploy to Kubernetes

Install the chart using Helm. This will create the Deployment, Service, PVC, and RBAC roles required to read K8s Services.
Bash

helm install port-manager ./charts/port-manager \
  --create-namespace \
  --namespace tools

4. Access the UI

To access the interface locally without an Ingress:
Bash

kubectl port-forward svc/port-manager 8080:80 -n tools

Open your browser at: http://localhost:8080
💻 Local Development

You can run the service locally on your machine.

    With K8s connection: If you have ~/.kube/config set up, the app will try to connect to your current cluster context.

    Without K8s: If no cluster is found, it starts in "Mock Mode" (empty K8s data).

Bash

# Install dependencies
go mod tidy

# Run the app
go run main.go

📡 API Reference

The service exposes a simple JSON API.
Get All Ports

Returns the status of all ports in the range 30000-32767.

    URL: /api/ports

    Method: GET

    Response:

[
  {
    "port": 30001,
    "status": "occupied_k8s",
    "service_name": "default/my-app",
    "namespace": "default",
    "updated_at": "2023-10-25 12:00:00"
  },
  {
    "port": 30002,
    "status": "free",
    "service_name": "-"
  }
]

Вот полный файл README.md с корректной разметкой Markdown. Вы можете скопировать этот блок и сохранить его как README.md в корне вашего проекта.
Markdown

# Kubernetes NodePort Manager

**NodePort Manager** is a lightweight tool designed to help DevOps engineers and developers manage the chaos of Kubernetes NodePorts (range `30000-32767`).

It scans the cluster for occupied ports, allows manual reservation of ports for future use (or non-K8s services), and provides a clean Web UI to visualize port usage.

## 🚀 Features

* **Real-time Kubernetes Sync:** Automatically detects NodePorts occupied by `NodePort` or `LoadBalancer` services in the cluster.
* **Manual Reservation:** Reserve a specific port via UI or API to prevent others from using it.
* **Conflict Prevention:** Prevents reserving a port that is already taken by Kubernetes.
* **Web UI:** Simple interface built with Vue.js to search, filter, and view port status.
* **Persistence:** Stores manual reservations in a JSON file (persisted via PVC).
* **REST API:** Simple API to integrate with CI/CD pipelines.

---

## 📂 Project Structure

```text
port-manager/
├── charts/                 # Helm charts
│   └── port-manager/
├── templates/              # HTML frontend
│   └── index.html
├── Dockerfile              # Multi-stage build
├── main.go                 # Go backend logic
└── go.mod                  # Go dependencies

🛠 Prerequisites

    Kubernetes Cluster (v1.19+)

    Helm (v3+)

    Docker (to build the image)

    Go 1.21+ (optional, for local development)

📦 Installation & Deployment
1. Build and Push the Docker Image

Since this is a custom tool, you need to build the image and push it to your container registry.
Bash

# 1. Build the image
docker build -t [my-registry.com/ops/port-manager:latest](https://my-registry.com/ops/port-manager:latest) .

# 2. Push to your registry
docker push [my-registry.com/ops/port-manager:latest](https://my-registry.com/ops/port-manager:latest)

2. Configure Helm Chart

Edit charts/port-manager/values.yaml to point to your image:
YAML

image:
  repository: [my-registry.com/ops/port-manager](https://my-registry.com/ops/port-manager)
  tag: latest
  pullPolicy: Always

storage:
  size: 1Gi
  # storageClassName: "standard" # Uncomment if you need a specific class

3. Deploy to Kubernetes

Install the chart using Helm. This will create the Deployment, Service, PVC, and RBAC roles required to read K8s Services.
Bash

helm install port-manager ./charts/port-manager \
  --create-namespace \
  --namespace tools

4. Access the UI

To access the interface locally without an Ingress:
Bash

kubectl port-forward svc/port-manager 8080:80 -n tools

Open your browser at: http://localhost:8080
💻 Local Development

You can run the service locally on your machine.

    With K8s connection: If you have ~/.kube/config set up, the app will try to connect to your current cluster context.

    Without K8s: If no cluster is found, it starts in "Mock Mode" (empty K8s data).

Bash

# Install dependencies
go mod tidy

# Run the app
go run main.go

📡 API Reference

The service exposes a simple JSON API.
Get All Ports

Returns the status of all ports in the range 30000-32767.

    URL: /api/ports

    Method: GET

    Response:

JSON

[
  {
    "port": 30001,
    "status": "occupied_k8s",
    "service_name": "default/my-app",
    "namespace": "default",
    "updated_at": "2023-10-25 12:00:00"
  },
  {
    "port": 30002,
    "status": "free",
    "service_name": "-"
  }
]

Reserve a Port

Manually lock a port.

    URL: /api/reserve

    Method: POST

    Body:

{
  "port": 30005,
  "name": "Jira-Test-Instance"
}

Response: 200 OK or 409 Conflict (if already taken).

⚙️ Configuration

Environment Variable	Description	Default
PORT	Server Port	8080
DB_FILE	Path to JSON DB	/data/reservations.json

Note: Currently, these are defined as constants in main.go, but can be easily refactored to env vars if needed.

🔐 Permissions (RBAC)

The Helm chart creates a ClusterRole with the following permissions:

    apiGroups: [""]

    resources: ["services"]

    verbs: ["get", "list", "watch"]

This is necessary for the application to scan all namespaces for NodePort usage.