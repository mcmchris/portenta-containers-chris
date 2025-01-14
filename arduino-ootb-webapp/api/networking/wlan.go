package networking

import (
	"errors"
	"time"

	utils "x8-ootb/utils"

	"github.com/Wifx/gonetworkmanager/v2"
	"github.com/google/uuid"
)

var NetworkConnectionFailed = errors.New("cannot connect wifi network")

const WLAN_INTERFACE_NAME = "wlan0"

func WlanNetworks() ([]Network, error) {
	nm, err := gonetworkmanager.NewNetworkManager()
	if err != nil {
		return nil, err
	}
	device, err := nm.GetDeviceByIpIface(WLAN_INTERFACE_NAME)
	if err != nil {
		return nil, err
	}
	deviceWireless, err := gonetworkmanager.NewDeviceWireless(device.GetPath())
	if err != nil {
		return nil, err
	}
	err = deviceWireless.RequestScan()
	if err != nil {
		return nil, err
	}
	time.Sleep(1 * time.Second)
	accessPoints, err := deviceWireless.GetAllAccessPoints()
	if err != nil {
		return nil, err
	}
	networksMap := make(map[string]*Network)
	for _, ap := range accessPoints {
		ssid, err := ap.GetPropertySSID()
		if err != nil {
			return nil, err
		}
		signal, err := ap.GetPropertyStrength()
		if err != nil {
			return nil, err
		}
		securityNum, err := ap.GetPropertyRSNFlags()
		if err != nil {
			return nil, err
		}
		bssid, err := ap.GetPropertyHWAddress()
		if err != nil {
			return nil, err
		}
		security := ""
		if securityNum != 0 {
			security = "protected"
		}
		network := Network{
			SSID:     ssid,
			BSSID:    bssid,
			Signal:   signal,
			Security: security,
		}
		networkID := ssid
		if ssid == "" {
			networkID = bssid
		}
		if networksMap[networkID] == nil {
			networksMap[networkID] = &network
		}
	}

	networks := []Network{}
	for _, n := range networksMap {
		networks = append(networks, *n)

	}
	return networks, nil
}

func WlanConnect(ssid string, password string) error {
	settings, err := gonetworkmanager.NewSettings()
	if err != nil {
		return err
	}

	err = utils.DeleteConnectionByInterfaceName(WLAN_INTERFACE_NAME)
	if err != nil {
		return err
	}
	time.Sleep(1 * time.Second)

	connection := make(map[string]map[string]interface{})

	connection["802-11-wireless"] = make(map[string]interface{})
	connection["802-11-wireless"]["ssid"] = []byte(ssid)

	connection["802-11-wireless-security"] = make(map[string]interface{})
	connection["802-11-wireless-security"]["key-mgmt"] = "wpa-psk"
	connection["802-11-wireless-security"]["psk"] = password

	connection["connection"] = make(map[string]interface{})
	connection["connection"]["id"] = ssid
	connection["connection"]["type"] = "802-11-wireless"
	connectionUUID, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	connection["connection"]["uuid"] = connectionUUID.String()
	connection["connection"]["interface-name"] = WLAN_INTERFACE_NAME
	connection["connection"]["autoconnect"] = true
	_, err = settings.AddConnection(connection)
	if err != nil {
		return err
	}

	return nil
}

func GetWlanConnection() (*Connection, error) {
	return GetConnection(WLAN_INTERFACE_NAME)
}
