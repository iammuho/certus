package hub

import (
	"fmt"
	"sync"

	"github.com/iammuho/certus/cmd/app/context"
	"go.uber.org/zap"
)

var (
	driversMu sync.RWMutex              // Mutex to protect access to the drivers map
	drivers   = make(map[string]Driver) // Registered drivers
)

// Driver is the interface that all drivers must implement.
type Driver interface {
	Execute(ctx context.AppContext)
}

type hub struct {
	ctx     context.AppContext
	drivers map[string]Driver
}

// NewHub creates a new hub.
func NewHub(ctx context.AppContext) *hub {
	h := &hub{
		ctx:     ctx,
		drivers: make(map[string]Driver),
	}

	// Initialize drivers
	for _, name := range ListDrivers() {
		driver, err := GetDriver(name)
		if err != nil {
			ctx.GetLogger().Error("Failed to get driver %q: %v", zap.String("driver", name), zap.Error(err))
			continue
		}

		h.drivers[name] = driver
	}

	return h
}

// ExecuteDrivers executes all registered drivers.
func (h *hub) ExecuteDrivers() {
	for _, driver := range h.drivers {
		driver.Execute(h.ctx)
	}
}

// Register registers a driver by name.
// It panics if the driver is nil or if a driver with the same name is already registered.
func Register(name string, driver Driver) {
	driversMu.Lock()
	defer driversMu.Unlock()
	if driver == nil {
		panic("hub: Register driver is nil")
	}
	if _, dup := drivers[name]; dup {
		panic("hub: Register called twice for driver " + name)
	}
	drivers[name] = driver
}

// ListDrivers returns a sorted list of the names of the registered drivers.
func ListDrivers() []string {
	driversMu.RLock()
	defer driversMu.RUnlock()
	list := make([]string, 0, len(drivers))
	for name := range drivers {
		list = append(list, name)
	}
	return list
}

// GetDriver retrieves a registered driver by name.
// It returns an error if the driver is not found.
func GetDriver(name string) (Driver, error) {
	driversMu.RLock()
	defer driversMu.RUnlock()
	driver, exists := drivers[name]
	if !exists {
		return nil, fmt.Errorf("hub: driver %q not found", name)
	}
	return driver, nil
}
