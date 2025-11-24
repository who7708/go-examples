package springboot

import (
	"os"
	"os/exec"
	"sync"
)

type ServiceManager struct {
	jarPath string
	cmd     *exec.Cmd
	mu      sync.Mutex
	running bool
}

func NewServiceManager(jarPath string) *ServiceManager {
	return &ServiceManager{jarPath: jarPath}
}

func (m *ServiceManager) Start() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.running {
		return nil
	}
	m.cmd = exec.Command("java", "-jar", m.jarPath)
	m.cmd.Stdout = os.Stdout
	m.cmd.Stderr = os.Stderr
	if err := m.cmd.Start(); err != nil {
		return err
	}
	m.running = true
	go func() {
		m.cmd.Wait()
		m.mu.Lock()
		m.running = false
		m.mu.Unlock()
	}()
	return nil
}

func (m *ServiceManager) Stop() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if !m.running || m.cmd == nil {
		return nil
	}
	if err := m.cmd.Process.Kill(); err != nil {
		return err
	}
	m.running = false
	return nil
}

func (m *ServiceManager) Upgrade(newJarPath string) error {
	if err := m.Stop(); err != nil {
		return err
	}
	if err := os.Rename(newJarPath, m.jarPath); err != nil {
		return err
	}
	return m.Start()
}

func (m *ServiceManager) Status() bool {
	m.mu.Lock()
	running := m.running
	m.mu.Unlock()
	return running
}
