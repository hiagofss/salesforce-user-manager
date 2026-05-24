package infrastructure

import (
	"context"
	"time"
)

// ServiceStatus represents the status of a single service
type ServiceStatus struct {
	Status string `json:"status"` // "ok" or "down"
}

// HealthStatus represents the overall health status
type HealthStatus struct {
	Status    string                   `json:"status"` // "healthy" or "degraded"
	Timestamp time.Time                `json:"timestamp"`
	Services  map[string]ServiceStatus `json:"services"`
	IsHealthy bool                     `json:"-"`
}

// HealthChecker handles health check logic
type HealthChecker struct {
	container *AppContainer
}

// NewHealthChecker creates a new HealthChecker
func NewHealthChecker(container *AppContainer) *HealthChecker {
	return &HealthChecker{
		container: container,
	}
}

// Check performs health checks on all services
func (hc *HealthChecker) Check(ctx context.Context) HealthStatus {
	status := HealthStatus{
		Timestamp: time.Now(),
		Services:  make(map[string]ServiceStatus),
	}

	// Check user repository
	userRepoOk := hc.checkService(func() error {
		_, err := hc.container.UserRepository.GetAll()
		return err
	})

	// Check org repository
	orgRepoOk := hc.checkService(func() error {
		_, err := hc.container.OrgRepository.GetAll()
		return err
	})

	status.Services["user_repository"] = ServiceStatus{Status: statusString(userRepoOk)}
	status.Services["org_repository"] = ServiceStatus{Status: statusString(orgRepoOk)}

	// Overall status: healthy only if all services are ok
	allHealthy := userRepoOk && orgRepoOk
	status.IsHealthy = allHealthy

	if allHealthy {
		status.Status = "healthy"
	} else {
		status.Status = "degraded"
	}

	return status
}

// checkService safely executes a service check with timeout
func (hc *HealthChecker) checkService(check func() error) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	done := make(chan error, 1)
	go func() {
		done <- check()
	}()

	select {
	case err := <-done:
		return err == nil
	case <-ctx.Done():
		return false // timeout means service is down
	}
}

// statusString converts bool to status string
func statusString(ok bool) string {
	if ok {
		return "ok"
	}
	return "down"
}
