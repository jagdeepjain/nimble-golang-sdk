// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

// NsFcSessionInitiator - Information of the Fibre Channel Session Initiator.
// Export NsFcSessionInitiatorFields for advance operations like search filter etc.
var NsFcSessionInitiatorFields *NsFcSessionInitiator

func init() {
	InitiatorAliasfield := "initiator_alias"
	InitiatorWwpnfield := "initiator_wwpn"
	InitiatorWwnnfield := "initiator_wwnn"
	InitiatorSwitchNamefield := "initiator_switch_name"
	InitiatorSwitchPortfield := "initiator_switch_port"
	InitiatorSymbolicPortnamefield := "initiator_symbolic_portname"
	InitiatorSymbolicNodenamefield := "initiator_symbolic_nodename"
	InitiatorFcidfield := "initiator_fcid"

	NsFcSessionInitiatorFields = &NsFcSessionInitiator{
		InitiatorAlias:            &InitiatorAliasfield,
		InitiatorWwpn:             &InitiatorWwpnfield,
		InitiatorWwnn:             &InitiatorWwnnfield,
		InitiatorSwitchName:       &InitiatorSwitchNamefield,
		InitiatorSwitchPort:       &InitiatorSwitchPortfield,
		InitiatorSymbolicPortname: &InitiatorSymbolicPortnamefield,
		InitiatorSymbolicNodename: &InitiatorSymbolicNodenamefield,
		InitiatorFcid:             &InitiatorFcidfield,
	}
}

type NsFcSessionInitiator struct {
	// InitiatorAlias - Alias of the Fibre Channel initiator.
	InitiatorAlias *string `json:"initiator_alias,omitempty"`
	// InitiatorWwpn - WWPN (World Wide Port Name) of the Fibre Channel initiator.
	InitiatorWwpn *string `json:"initiator_wwpn,omitempty"`
	// InitiatorWwnn - WWNN (World Wide Node Name) of the Fibre Channel initiator.
	InitiatorWwnn *string `json:"initiator_wwnn,omitempty"`
	// InitiatorSwitchName - Name of the switch used by the Fibre Channel initiator.
	InitiatorSwitchName *string `json:"initiator_switch_name,omitempty"`
	// InitiatorSwitchPort - Port on the switch used by the Fibre Channel initiator.
	InitiatorSwitchPort *string `json:"initiator_switch_port,omitempty"`
	// InitiatorSymbolicPortname - Symbolic port name associated with the Fibre Channel initiator.
	InitiatorSymbolicPortname *string `json:"initiator_symbolic_portname,omitempty"`
	// InitiatorSymbolicNodename - Symbolic node name associated with the Fibre Channel initiator.
	InitiatorSymbolicNodename *string `json:"initiator_symbolic_nodename,omitempty"`
	// InitiatorFcid - FCID assigned to the Fibre Channel initiator.
	InitiatorFcid *string `json:"initiator_fcid,omitempty"`
}
