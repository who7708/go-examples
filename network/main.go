// main.go
package main

import (
	"bufio"
	"fmt"
	"network/network"
	"os/exec"
	"strconv"
	"strings"
)

func getIanaInterfaceTypeFromPowerShell() (uint32, error) {
	// PowerShell 命令：启用 WinRT + 获取 IanaInterfaceType
	cmd := exec.Command("powershell", "-NoProfile", "-Command", `
Add-Type -AssemblyName 'Windows, Version=255.255.255.255, Culture=neutral, PublicKeyToken=null, ContentType=WindowsRuntime' -ErrorAction SilentlyContinue
$profile = [Windows.Networking.Connectivity.NetworkInformation]::GetInternetConnectionProfile()
if ($profile -and $profile.NetworkAdapter) {
    $profile.NetworkAdapter.IanaInterfaceType
} else {
    Write-Host "0"
}
`)

	output, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("PowerShell execution failed: %w", err)
	}

	line := strings.TrimSpace(string(output))
	if line == "" || line == "0" {
		return 0, fmt.Errorf("no active network profile found")
	}

	// PowerShell 可能输出多行（如错误信息），取第一行数字
	scanner := bufio.NewScanner(strings.NewReader(line))
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}
		if num, err := strconv.ParseUint(text, 10, 32); err == nil {
			return uint32(num), nil
		}
	}

	return 0, fmt.Errorf("failed to parse IanaInterfaceType from output: %s", line)
}

func main() {
	typ, err := getIanaInterfaceTypeFromPowerShell()
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}

	name := network.GetIanaInterfaceTypeName(typ)
	wireless := network.IsWireless(typ)

	fmt.Printf("✅ IANA Interface Type ID: %d\n", typ)
	fmt.Printf("✅ Name: %s\n", name)
	fmt.Printf("✅ Is Wireless: %t\n", wireless)
}
