package proxy

import (
	sdkerrors "cosmossdk.io/errors"
)

var (
	codespace                                     = "relayer_proxy"
	ErrRelayerProxyUnsupportedRPCType             = sdkerrors.Register(codespace, 1, "unsupported rpc type")
	ErrRelayerProxyInvalidSession                 = sdkerrors.Register(codespace, 2, "invalid session in relayer request")
	ErrRelayerProxyInvalidSupplier                = sdkerrors.Register(codespace, 3, "supplier does not belong to session")
	ErrRelayerProxyUndefinedSigningKeyNames       = sdkerrors.Register(codespace, 4, "supplier signing key names are undefined")
	ErrRelayerServicesConfigsUndefined            = sdkerrors.Register(codespace, 5, "services configurations are undefined")
	ErrRelayerProxyInvalidRelayRequest            = sdkerrors.Register(codespace, 6, "invalid relay request")
	ErrRelayerProxyInvalidRelayResponse           = sdkerrors.Register(codespace, 7, "invalid relay response")
	ErrRelayerProxyServiceEndpointNotHandled      = sdkerrors.Register(codespace, 8, "service endpoint not handled by relayer proxy")
	ErrRelayerProxyUnsupportedTransportType       = sdkerrors.Register(codespace, 9, "unsupported proxy transport type")
	ErrRelayerProxyInternalError                  = sdkerrors.Register(codespace, 10, "internal error")
	ErrRelayerProxyMissingSupplierOperatorAddress = sdkerrors.Register(codespace, 11, "supplier operator address is missing")
	ErrRelayerProxyUnknownSession                 = sdkerrors.Register(codespace, 12, "relayer proxy encountered unknown session")
	ErrRelayerProxyRateLimited                    = sdkerrors.Register(codespace, 13, "offchain rate limit hit by relayer proxy")
	ErrRelayerProxyUnclaimRelayPrice              = sdkerrors.Register(codespace, 14, "failed to unclaim relay price")
	ErrRelayerProxySupplierNotReachable           = sdkerrors.Register(codespace, 15, "supplier(s) not reachable")
)
