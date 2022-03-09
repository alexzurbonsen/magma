// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GatewayEpcConfigs EPC configuration for an LTE gateway
//
// swagger:model gateway_epc_configs
type GatewayEpcConfigs struct {

	// Flag to enable or disable congestion control on MME
	CongestionControlEnabled *bool `json:"congestion_control_enabled,omitempty"`

	// dns primary
	// Max Length: 45
	// Min Length: 5
	DNSPrimary string `json:"dns_primary,omitempty"`

	// dns secondary
	// Max Length: 45
	// Min Length: 5
	DNSSecondary string `json:"dns_secondary,omitempty"`

	// ip block
	// Required: true
	// Max Length: 49
	// Min Length: 5
	IPBlock string `json:"ip_block"`

	// IP address for IPv4 P-CSCF on the AGW
	// Format: ipv4
	IPV4pCscfAddr strfmt.IPv4 `json:"ipv4_p_cscf_addr,omitempty"`

	// IP address for IPv4 S1U endpoint on the AGW
	// Max Length: 49
	// Min Length: 5
	IPV4SgwS1uAddr string `json:"ipv4_sgw_s1u_addr,omitempty"`

	// ipv6 block
	IPV6Block string `json:"ipv6_block,omitempty"`

	// IPv6 DNS Server address on the AGW
	// Format: ipv6
	IPV6DNSAddr strfmt.IPv6 `json:"ipv6_dns_addr,omitempty"`

	// IP address for IPv6 P-CSCF on the AGW
	// Format: ipv6
	IPV6pCscfAddr strfmt.IPv6 `json:"ipv6_p_cscf_addr,omitempty"`

	// ipv6 prefix allocation mode
	// Enum: [RANDOM HASH]
	IPV6PrefixAllocationMode string `json:"ipv6_prefix_allocation_mode,omitempty"`

	// nat enabled
	// Required: true
	NatEnabled *bool `json:"nat_enabled"`

	// IP address of gateway for management interface on the AGW
	// Max Length: 49
	// Min Length: 5
	SgiManagementIfaceGw string `json:"sgi_management_iface_gw,omitempty"`

	// IPv6 address for management interface on the AGW in CIDR format
	SgiManagementIfaceIPV6Addr string `json:"sgi_management_iface_ipv6_addr,omitempty"`

	// IPv6 address of gateway for management interface on the AGW
	// Format: ipv6
	SgiManagementIfaceIPV6Gw strfmt.IPv6 `json:"sgi_management_iface_ipv6_gw,omitempty"`

	// IP address for management interface on the AGW, If not specified AGW uses DHCP to configure it.
	// Max Length: 49
	// Min Length: 5
	SgiManagementIfaceStaticIP string `json:"sgi_management_iface_static_ip,omitempty"`

	// VLAN ID for management interface traffic on the AGW
	// Max Length: 4
	// Min Length: 1
	SgiManagementIfaceVlan string `json:"sgi_management_iface_vlan,omitempty"`

	// subscriberdb sync interval
	SubscriberdbSyncInterval SubscriberdbSyncInterval `json:"subscriberdb_sync_interval,omitempty"`
}

// Validate validates this gateway epc configs
func (m *GatewayEpcConfigs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDNSPrimary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNSSecondary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPBlock(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPV4pCscfAddr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPV4SgwS1uAddr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPV6DNSAddr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPV6pCscfAddr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPV6PrefixAllocationMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNatEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSgiManagementIfaceGw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSgiManagementIfaceIPV6Gw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSgiManagementIfaceStaticIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSgiManagementIfaceVlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriberdbSyncInterval(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GatewayEpcConfigs) validateDNSPrimary(formats strfmt.Registry) error {

	if swag.IsZero(m.DNSPrimary) { // not required
		return nil
	}

	if err := validate.MinLength("dns_primary", "body", string(m.DNSPrimary), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("dns_primary", "body", string(m.DNSPrimary), 45); err != nil {
		return err
	}

	return nil
}

func (m *GatewayEpcConfigs) validateDNSSecondary(formats strfmt.Registry) error {

	if swag.IsZero(m.DNSSecondary) { // not required
		return nil
	}

	if err := validate.MinLength("dns_secondary", "body", string(m.DNSSecondary), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("dns_secondary", "body", string(m.DNSSecondary), 45); err != nil {
		return err
	}

	return nil
}

func (m *GatewayEpcConfigs) validateIPBlock(formats strfmt.Registry) error {

	if err := validate.RequiredString("ip_block", "body", string(m.IPBlock)); err != nil {
		return err
	}

	if err := validate.MinLength("ip_block", "body", string(m.IPBlock), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("ip_block", "body", string(m.IPBlock), 49); err != nil {
		return err
	}

	return nil
}

func (m *GatewayEpcConfigs) validateIPV4pCscfAddr(formats strfmt.Registry) error {

	if swag.IsZero(m.IPV4pCscfAddr) { // not required
		return nil
	}

	if err := validate.FormatOf("ipv4_p_cscf_addr", "body", "ipv4", m.IPV4pCscfAddr.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GatewayEpcConfigs) validateIPV4SgwS1uAddr(formats strfmt.Registry) error {

	if swag.IsZero(m.IPV4SgwS1uAddr) { // not required
		return nil
	}

	if err := validate.MinLength("ipv4_sgw_s1u_addr", "body", string(m.IPV4SgwS1uAddr), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("ipv4_sgw_s1u_addr", "body", string(m.IPV4SgwS1uAddr), 49); err != nil {
		return err
	}

	return nil
}

func (m *GatewayEpcConfigs) validateIPV6DNSAddr(formats strfmt.Registry) error {

	if swag.IsZero(m.IPV6DNSAddr) { // not required
		return nil
	}

	if err := validate.FormatOf("ipv6_dns_addr", "body", "ipv6", m.IPV6DNSAddr.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GatewayEpcConfigs) validateIPV6pCscfAddr(formats strfmt.Registry) error {

	if swag.IsZero(m.IPV6pCscfAddr) { // not required
		return nil
	}

	if err := validate.FormatOf("ipv6_p_cscf_addr", "body", "ipv6", m.IPV6pCscfAddr.String(), formats); err != nil {
		return err
	}

	return nil
}

var gatewayEpcConfigsTypeIPV6PrefixAllocationModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RANDOM","HASH"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gatewayEpcConfigsTypeIPV6PrefixAllocationModePropEnum = append(gatewayEpcConfigsTypeIPV6PrefixAllocationModePropEnum, v)
	}
}

const (

	// GatewayEpcConfigsIPV6PrefixAllocationModeRANDOM captures enum value "RANDOM"
	GatewayEpcConfigsIPV6PrefixAllocationModeRANDOM string = "RANDOM"

	// GatewayEpcConfigsIPV6PrefixAllocationModeHASH captures enum value "HASH"
	GatewayEpcConfigsIPV6PrefixAllocationModeHASH string = "HASH"
)

// prop value enum
func (m *GatewayEpcConfigs) validateIPV6PrefixAllocationModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, gatewayEpcConfigsTypeIPV6PrefixAllocationModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GatewayEpcConfigs) validateIPV6PrefixAllocationMode(formats strfmt.Registry) error {

	if swag.IsZero(m.IPV6PrefixAllocationMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateIPV6PrefixAllocationModeEnum("ipv6_prefix_allocation_mode", "body", m.IPV6PrefixAllocationMode); err != nil {
		return err
	}

	return nil
}

func (m *GatewayEpcConfigs) validateNatEnabled(formats strfmt.Registry) error {

	if err := validate.Required("nat_enabled", "body", m.NatEnabled); err != nil {
		return err
	}

	return nil
}

func (m *GatewayEpcConfigs) validateSgiManagementIfaceGw(formats strfmt.Registry) error {

	if swag.IsZero(m.SgiManagementIfaceGw) { // not required
		return nil
	}

	if err := validate.MinLength("sgi_management_iface_gw", "body", string(m.SgiManagementIfaceGw), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("sgi_management_iface_gw", "body", string(m.SgiManagementIfaceGw), 49); err != nil {
		return err
	}

	return nil
}

func (m *GatewayEpcConfigs) validateSgiManagementIfaceIPV6Gw(formats strfmt.Registry) error {

	if swag.IsZero(m.SgiManagementIfaceIPV6Gw) { // not required
		return nil
	}

	if err := validate.FormatOf("sgi_management_iface_ipv6_gw", "body", "ipv6", m.SgiManagementIfaceIPV6Gw.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GatewayEpcConfigs) validateSgiManagementIfaceStaticIP(formats strfmt.Registry) error {

	if swag.IsZero(m.SgiManagementIfaceStaticIP) { // not required
		return nil
	}

	if err := validate.MinLength("sgi_management_iface_static_ip", "body", string(m.SgiManagementIfaceStaticIP), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("sgi_management_iface_static_ip", "body", string(m.SgiManagementIfaceStaticIP), 49); err != nil {
		return err
	}

	return nil
}

func (m *GatewayEpcConfigs) validateSgiManagementIfaceVlan(formats strfmt.Registry) error {

	if swag.IsZero(m.SgiManagementIfaceVlan) { // not required
		return nil
	}

	if err := validate.MinLength("sgi_management_iface_vlan", "body", string(m.SgiManagementIfaceVlan), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("sgi_management_iface_vlan", "body", string(m.SgiManagementIfaceVlan), 4); err != nil {
		return err
	}

	return nil
}

func (m *GatewayEpcConfigs) validateSubscriberdbSyncInterval(formats strfmt.Registry) error {

	if swag.IsZero(m.SubscriberdbSyncInterval) { // not required
		return nil
	}

	if err := m.SubscriberdbSyncInterval.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("subscriberdb_sync_interval")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GatewayEpcConfigs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GatewayEpcConfigs) UnmarshalBinary(b []byte) error {
	var res GatewayEpcConfigs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
