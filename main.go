package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// Constants
const (
	DBFile      = "/data/reservations.json"
	TemplateDir = "./templates" // Path to the templates directory
	MinNodePort = 30000
	MaxNodePort = 32767
)

// PortInfo represents the data structure for a port
type PortInfo struct {
	Port        int32  `json:"port"`
	Status      string `json:"status"`       // "free", "occupied_k8s", "reserved_manual"
	ServiceName string `json:"service_name"` // Service name (k8s) or reserver name
	Namespace   string `json:"namespace"`
	UpdatedAt   string `json:"updated_at"`
}

// Store handles storage (in-memory + file)
type Store struct {
	sync.Mutex
	Reservations map[int32]PortInfo // Data from file
	K8sPorts     map[int32]PortInfo // Data from K8s (real-time)
}

var store = &Store{
	Reservations: make(map[int32]PortInfo),
	K8sPorts:     make(map[int32]PortInfo),
}

// --- File DB Operations ---

func (s *Store) Load() {
	s.Lock()
	defer s.Unlock()
	data, err := ioutil.ReadFile(DBFile)
	if err == nil {
		json.Unmarshal(data, &s.Reservations)
	}
}

func (s *Store) Save() {
	s.Lock()
	defer s.Unlock()
	data, _ := json.MarshalIndent(s.Reservations, "", "  ")
	ioutil.WriteFile(DBFile, data, 0644)
}

func (s *Store) Reserve(port int32, name string) error {
	s.Lock()
	defer s.Unlock()

	// Check if occupied by K8s
	if _, exists := s.K8sPorts[port]; exists {
		return fmt.Errorf("port %d occupied by K8s", port)
	}
	// Check if manually reserved
	if val, exists := s.Reservations[port]; exists {
		return fmt.Errorf("port %d reserved by %s", port, val.ServiceName)
	}

	s.Reservations[port] = PortInfo{
		Port:        port,
		Status:      "reserved_manual",
		ServiceName: name,
		UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}

	// Save immediately (for simplicity)
	data, _ := json.MarshalIndent(s.Reservations, "", "  ")
	ioutil.WriteFile(DBFile, data, 0644)
	return nil
}

// --- Synchronization with K8s ---

func syncK8s(clientset *kubernetes.Clientset) {
	for {
		services, err := clientset.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			log.Printf("Error listing services: %v", err)
			time.Sleep(10 * time.Second)
			continue
		}

		newK8sPorts := make(map[int32]PortInfo)

		for _, svc := range services.Items {
			if svc.Spec.Type == corev1.ServiceTypeNodePort || svc.Spec.Type == corev1.ServiceTypeLoadBalancer {
				for _, p := range svc.Spec.Ports {
					if p.NodePort != 0 {
						newK8sPorts[p.NodePort] = PortInfo{
							Port:        p.NodePort,
							Status:      "occupied_k8s",
							ServiceName: fmt.Sprintf("%s/%s", svc.Namespace, svc.Name),
							Namespace:   svc.Namespace,
							UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
						}
					}
				}
			}
		}

		store.Lock()
		store.K8sPorts = newK8sPorts
		store.Unlock()

		// log.Printf("Synced K8s ports. Found %d NodePorts.", len(newK8sPorts))
		time.Sleep(10 * time.Second)
	}
}

// --- Data Aggregation for API ---

func getAllPorts() []PortInfo {
	store.Lock()
	defer store.Unlock()

	var result []PortInfo
	// Iterate through the entire range
	for p := int32(MinNodePort); p <= MaxNodePort; p++ {
		info := PortInfo{Port: p, Status: "free", ServiceName: "-"}

		// Priority 1: Occupied by K8s
		if k8sInfo, ok := store.K8sPorts[p]; ok {
			info = k8sInfo
		} else if resInfo, ok := store.Reservations[p]; ok {
			// Priority 2: Reserved manually
			info = resInfo
		}
		result = append(result, info)
	}
	return result
}

// --- Main ---

func main() {
	// 1. Initialize K8s client (In-Cluster Config)
	config, err := rest.InClusterConfig()
	if err != nil {
		// Fallback for local testing (optional, if kubeconfig is present)
		log.Println("Failed to get in-cluster config, trying logical fallback or exit:", err)
		// return // Uncomment for production to enforce K8s environment
	}

	var clientset *kubernetes.Clientset
	if config != nil {
		clientset, _ = kubernetes.NewForConfig(config)
	}

	// 2. Load DB and start synchronization
	store.Load()
	if clientset != nil {
		go syncK8s(clientset)
	} else {
		log.Println("Warning: Running without K8s connection (Mock Mode)")
	}

	// 3. Web Server
	r := gin.Default()

	// Serve HTML file
	r.GET("/", func(c *gin.Context) {
		c.File(filepath.Join(TemplateDir, "index.html"))
	})

	r.GET("/api/ports", func(c *gin.Context) {
		c.JSON(200, getAllPorts())
	})

	r.POST("/api/reserve", func(c *gin.Context) {
		var req struct {
			Port int32  `json:"port"`
			Name string `json:"name"`
		}
		if err := c.BindJSON(&req); err != nil {
			c.String(400, "Bad Request")
			return
		}
		if err := store.Reserve(req.Port, req.Name); err != nil {
			c.String(409, err.Error())
			return
		}
		c.String(200, "Reserved")
	})

	// Health check
	r.GET("/health", func(c *gin.Context) { c.String(200, "OK") })

	r.Run(":8080")
}