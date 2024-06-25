/* Code generated by github.com/srdtrk/go-codegen, DO NOT EDIT. */
package codegentests_test

import (
	"encoding/json"
	"errors"
)

type InstantiateMsg struct {
	// The admin address for instantiating new cw721 contracts. In case of None, contract is immutable.
	Cw721Admin *string `json:"cw721_admin,omitempty"`
	/*
	   Code ID of cw721-ics contract. A new cw721-ics will be instantiated for each new IBCd NFT classID.

	   NOTE: this _must_ correspond to the cw721-base contract. Using a regular cw721 may cause the ICS 721 interface implemented by this contract to stop working, and IBCd away NFTs to be unreturnable as cw721 does not have a mint method in the spec.
	*/
	Cw721BaseCodeId int `json:"cw721_base_code_id"`
	// An optional proxy contract. If an incoming proxy is set, the contract will call it and pass IbcPacket. The proxy is expected to implement the Ics721ReceiveIbcPacketMsg for message execution.
	IncomingProxy *ContractInstantiateInfo `json:"incoming_proxy,omitempty"`
	// An optional proxy contract. If an outging proxy is set, the contract will only accept NFTs from that proxy. The proxy is expected to implement the cw721 proxy interface defined in the cw721-outgoing-proxy crate.
	OutgoingProxy *ContractInstantiateInfo `json:"outgoing_proxy,omitempty"`
	// Address that may pause the contract. PAUSER may pause the contract a single time; in pausing the contract they burn the right to do so again. A new pauser may be later nominated by the CosmWasm level admin via a migration.
	Pauser *string `json:"pauser,omitempty"`
}

type ExecuteMsg struct {
	// Receives a NFT to be IBC transfered away. The `msg` field must be a binary encoded `IbcOutgoingMsg`.
	ReceiveNft *ExecuteMsg_ReceiveNft `json:"receive_nft,omitempty"`
	// Pauses the ICS721 contract. Only the pauser may call this. In pausing the contract, the pauser burns the right to do so again.
	Pause *ExecuteMsg_Pause `json:"pause,omitempty"`
	// Mesages used internally by the contract. These may only be called by the contract itself.
	Callback *ExecuteMsg_Callback `json:"callback,omitempty"`
	// Admin msg in case something goes wrong. As a minimum it clean up states (incoming channel and token metadata), and burn NFT if exists.
	AdminCleanAndBurnNft *ExecuteMsg_AdminCleanAndBurnNft `json:"admin_clean_and_burn_nft,omitempty"`
	// Admin msg in case something goes wrong. As a minimum it clean up state (outgoing channel), and transfer NFT if exists. - transfer NFT if exists
	AdminCleanAndUnescrowNft *ExecuteMsg_AdminCleanAndUnescrowNft `json:"admin_clean_and_unescrow_nft,omitempty"`
}

type QueryMsg struct {
	// Gets the classID this contract has stored for a given NFT contract. If there is no class ID for the provided contract, returns None.
	ClassId *QueryMsg_ClassId `json:"class_id,omitempty"`
	// Gets the NFT contract associated wtih the provided class ID. If no such contract exists, returns None. Returns Option<Addr>.
	NftContract *QueryMsg_NftContract `json:"nft_contract,omitempty"`
	// Gets the class level metadata URI for the provided class_id. If there is no metadata, returns None. Returns `Option<Class>`.
	ClassMetadata *QueryMsg_ClassMetadata `json:"class_metadata,omitempty"`
	TokenMetadata *QueryMsg_TokenMetadata `json:"token_metadata,omitempty"`
	// Gets the owner of the NFT identified by CLASS_ID and TOKEN_ID. Errors if no such NFT exists. Returns `cw721::OwnerOfResonse`.
	Owner *QueryMsg_Owner `json:"owner,omitempty"`
	// Gets the address that may pause this contract if one is set.
	Pauser *QueryMsg_Pauser `json:"pauser,omitempty"`
	// Gets the current pause status.
	Paused *QueryMsg_Paused `json:"paused,omitempty"`
	// Gets this contract's outgoing cw721-outgoing-proxy if one is set.
	OutgoingProxy *QueryMsg_OutgoingProxy `json:"outgoing_proxy,omitempty"`
	// Gets this contract's incoming cw721-outgoing-proxy if one is set.
	IncomingProxy *QueryMsg_IncomingProxy `json:"incoming_proxy,omitempty"`
	// Gets the code used for instantiating new cw721s.
	Cw721CodeId *QueryMsg_Cw721CodeId `json:"cw721_code_id,omitempty"`
	// Gets the admin address for instantiating new cw721 contracts. In case of None, contract is immutable.
	Cw721Admin *QueryMsg_Cw721Admin `json:"cw721_admin,omitempty"`
	// Gets a list of classID as key (from NonFungibleTokenPacketData) and cw721 contract as value (instantiated for that classID).
	NftContracts *QueryMsg_NftContracts `json:"nft_contracts,omitempty"`
	// Gets a list of classID, tokenID, and local channelID. Used to determine the local channel that NFTs have been sent out on.
	OutgoingChannels *QueryMsg_OutgoingChannels `json:"outgoing_channels,omitempty"`
	// Gets a list of classID, tokenID, and local channel ID. Used to determine the local channel that NFTs have arrived at this contract.
	IncomingChannels *QueryMsg_IncomingChannels `json:"incoming_channels,omitempty"`
}

/*
A thin wrapper around u64 that is using strings for JSON encoding/decoding, such that the full u64 range can be used for clients that convert JSON numbers to floats, like JavaScript and jq.

# Examples

Use `from` to create instances of this and `u64` to get the value out:

``` # use cosmwasm_std::Uint64; let a = Uint64::from(42u64); assert_eq!(a.u64(), 42);

let b = Uint64::from(70u32); assert_eq!(b.u64(), 70); ```
*/
type Uint64 string

type Admin struct {
	Address      *Admin_Address      `json:"address,omitempty"`
	Instantiator *Admin_Instantiator `json:"instantiator,omitempty"`
}

/*
A thin wrapper around u128 that is using strings for JSON encoding/decoding, such that the full u128 range can be used for clients that convert JSON numbers to floats, like JavaScript and jq.

# Examples

Use `from` to create instances of this and `u128` to get the value out:

``` # use cosmwasm_std::Uint128; let a = Uint128::from(123u128); assert_eq!(a.u128(), 123);

let b = Uint128::from(42u64); assert_eq!(b.u128(), 42);

let c = Uint128::from(70u32); assert_eq!(c.u128(), 70); ```
*/
type Uint128 string

// A token ID according to the ICS-721 spec. The newtype pattern is used here to provide some distinction between token and class IDs in the type system.
type TokenId string

type QueryMsg_ClassMetadata struct {
	ClassId string `json:"class_id"`
}

type QueryMsg_Paused struct{}

/*
A human readable address.

In Cosmos, this is typically bech32 encoded. But for multi-chain smart contracts no assumptions should be made other than being UTF-8 encoded and of reasonable length.

This type represents a validated address. It can be created in the following ways 1. Use `Addr::unchecked(input)` 2. Use `let checked: Addr = deps.api.addr_validate(input)?` 3. Use `let checked: Addr = deps.api.addr_humanize(canonical_addr)?` 4. Deserialize from JSON. This must only be done from JSON that was validated before such as a contract's state. `Addr` must not be used in messages sent by the user because this would result in unvalidated instances.

This type is immutable. If you really need to mutate it (Really? Are you sure?), create a mutable copy using `let mut mutable = Addr::to_string()` and operate on that `String` instance.
*/
type Addr string

type VoucherCreation struct {
	// The class that these vouchers are being created for.
	Class Class `json:"class"`
	// The tokens to create debt-vouchers for.
	Tokens []Token `json:"tokens"`
}

type Class struct {
	// Optional base64 encoded metadata about the class.
	Data *Binary `json:"data,omitempty"`
	// A unique (from the source chain's perspective) identifier for the class.
	Id ClassId `json:"id"`
	// Optional URI pointing to off-chain metadata about the class.
	Uri *string `json:"uri,omitempty"`
}

type QueryMsg_Cw721CodeId struct{}

// Nullable_Addr is a nullable type of Addr
type Nullable_Addr = *Addr

// Nullable_Token is a nullable type of Token
type Nullable_Token = *Token

// Nullable_Class is a nullable type of Class
type Nullable_Class = *Class

/*
Binary is a wrapper around Vec<u8> to add base64 de/serialization with serde. It also adds some helper methods to help encode inline.

This is only needed as serde-json-{core,wasm} has a horrible encoding for Vec<u8>. See also <https://github.com/CosmWasm/cosmwasm/blob/main/docs/MESSAGE_TYPES.md>.
*/
type Binary string

// A token according to the ICS-721 spec.
type Token struct {
	// A unique identifier for the token.
	Id TokenId `json:"id"`
	// Optional URI pointing to off-chain metadata about the token.
	Uri *string `json:"uri,omitempty"`
	// Optional base64 encoded metadata about the token.
	Data *Binary `json:"data,omitempty"`
}

type Coin struct {
	Amount Uint128 `json:"amount"`
	Denom  string  `json:"denom"`
}

/*
The message types of the wasm module.

See https://github.com/CosmWasm/wasmd/blob/v0.14.0/x/wasm/internal/types/tx.proto
*/
type WasmMsg struct {
	/*
	   Dispatches a call to another contract at a known address (with known ABI).

	   This is translated to a [MsgExecuteContract](https://github.com/CosmWasm/wasmd/blob/v0.14.0/x/wasm/internal/types/tx.proto#L68-L78). `sender` is automatically filled with the current contract's address.
	*/
	Execute *WasmMsg_Execute `json:"execute,omitempty"`
	/*
	   Instantiates a new contracts from previously uploaded Wasm code.

	   The contract address is non-predictable. But it is guaranteed that when emitting the same Instantiate message multiple times, multiple instances on different addresses will be generated. See also Instantiate2.

	   This is translated to a [MsgInstantiateContract](https://github.com/CosmWasm/wasmd/blob/v0.29.2/proto/cosmwasm/wasm/v1/tx.proto#L53-L71). `sender` is automatically filled with the current contract's address.
	*/
	Instantiate *WasmMsg_Instantiate `json:"instantiate,omitempty"`
	/*
	   Instantiates a new contracts from previously uploaded Wasm code using a predictable address derivation algorithm implemented in [`cosmwasm_std::instantiate2_address`].

	   This is translated to a [MsgInstantiateContract2](https://github.com/CosmWasm/wasmd/blob/v0.29.2/proto/cosmwasm/wasm/v1/tx.proto#L73-L96). `sender` is automatically filled with the current contract's address. `fix_msg` is automatically set to false.
	*/
	Instantiate2 *WasmMsg_Instantiate2 `json:"instantiate2,omitempty"`
	/*
	   Migrates a given contracts to use new wasm code. Passes a MigrateMsg to allow us to customize behavior.

	   Only the contract admin (as defined in wasmd), if any, is able to make this call.

	   This is translated to a [MsgMigrateContract](https://github.com/CosmWasm/wasmd/blob/v0.14.0/x/wasm/internal/types/tx.proto#L86-L96). `sender` is automatically filled with the current contract's address.
	*/
	Migrate *WasmMsg_Migrate `json:"migrate,omitempty"`
	// Sets a new admin (for migrate) on the given contract. Fails if this contract is not currently admin of the target contract.
	UpdateAdmin *WasmMsg_UpdateAdmin `json:"update_admin,omitempty"`
	// Clears the admin on the given contract, so no more migration possible. Fails if this contract is not currently admin of the target contract.
	ClearAdmin *WasmMsg_ClearAdmin `json:"clear_admin,omitempty"`
}

type QueryMsg_Owner struct {
	TokenId string `json:"token_id"`
	ClassId string `json:"class_id"`
}

type QueryMsg_NftContracts struct {
	Limit      *int     `json:"limit,omitempty"`
	StartAfter *ClassId `json:"start_after,omitempty"`
}

// Nullable_Nullable_Addr is a nullable type of Nullable_Addr
type Nullable_Nullable_Addr = *Nullable_Addr

type QueryMsg_OutgoingProxy struct{}

type QueryMsg_Cw721Admin struct{}

type Approval struct {
	// When the Approval expires (maybe Expiration::never)
	Expires Expiration `json:"expires"`
	// Account that can transfer/send the token
	Spender string `json:"spender"`
}

type CallbackMsg struct {
	CreateVouchers *CallbackMsg_CreateVouchers `json:"create_vouchers,omitempty"`
	RedeemVouchers *CallbackMsg_RedeemVouchers `json:"redeem_vouchers,omitempty"`
	// Redeem all entries in outgoing channel.
	RedeemOutgoingChannelEntries *CallbackMsg_RedeemOutgoingChannelEntries `json:"redeem_outgoing_channel_entries,omitempty"`
	// Save all entries in incoming channel.
	AddIncomingChannelEntries *CallbackMsg_AddIncomingChannelEntries `json:"add_incoming_channel_entries,omitempty"`
	// Mints a NFT of collection class_id for receiver with the provided id and metadata. Only callable by this contract.
	Mint *CallbackMsg_Mint `json:"mint,omitempty"`
	/*
	   In submessage terms, say a message that results in an error "returns false" and one that succedes "returns true". Returns the logical conjunction (&&) of all the messages in operands.

	   Under the hood this just executes them in order. We use this to respond with a single ACK when a message calls for the execution of both `CreateVouchers` and `RedeemVouchers`.
	*/
	Conjunction *CallbackMsg_Conjunction `json:"conjunction,omitempty"`
}

// A class ID according to the ICS-721 spec. The newtype pattern is used here to provide some distinction between token and class IDs in the type system.
type ClassId string

// Cw721ReceiveMsg should be de/serialized under `Receive()` variant in a ExecuteMsg
type Cw721ReceiveMsg struct {
	Msg     Binary `json:"msg"`
	Sender  string `json:"sender"`
	TokenId string `json:"token_id"`
}

type QueryMsg_IncomingProxy struct{}

type QueryMsg_OutgoingChannels struct {
	Limit      *int        `json:"limit,omitempty"`
	StartAfter *ClassToken `json:"start_after,omitempty"`
}
type ExecuteMsg_ReceiveNft Cw721ReceiveMsg

type (
	ExecuteMsg_Pause    struct{}
	ExecuteMsg_Callback CallbackMsg
)

type ExecuteMsg_AdminCleanAndBurnNft struct {
	ClassId    string `json:"class_id"`
	Collection string `json:"collection"`
	Owner      string `json:"owner"`
	TokenId    string `json:"token_id"`
}

type QueryMsg_NftContract struct {
	ClassId string `json:"class_id"`
}

type QueryMsg_TokenMetadata struct {
	ClassId string `json:"class_id"`
	TokenId string `json:"token_id"`
}

/*
A point in time in nanosecond precision.

This type can represent times from 1970-01-01T00:00:00Z to 2554-07-21T23:34:33Z.

## Examples

``` # use cosmwasm_std::Timestamp; let ts = Timestamp::from_nanos(1_000_000_202); assert_eq!(ts.nanos(), 1_000_000_202); assert_eq!(ts.seconds(), 1); assert_eq!(ts.subsec_nanos(), 202);

let ts = ts.plus_seconds(2); assert_eq!(ts.nanos(), 3_000_000_202); assert_eq!(ts.seconds(), 3); assert_eq!(ts.subsec_nanos(), 202); ```
*/
type Timestamp Uint64

type ExecuteMsg_AdminCleanAndUnescrowNft struct {
	TokenId    string `json:"token_id"`
	ClassId    string `json:"class_id"`
	Collection string `json:"collection"`
	Recipient  string `json:"recipient"`
}

type QueryMsg_IncomingChannels struct {
	Limit      *int        `json:"limit,omitempty"`
	StartAfter *ClassToken `json:"start_after,omitempty"`
}

type ClassToken struct {
	ClassId ClassId `json:"class_id"`
	TokenId TokenId `json:"token_id"`
}

// Nullable_ClassId is a nullable type of ClassId
type Nullable_ClassId = *ClassId

type ContractInstantiateInfo struct {
	Msg    Binary `json:"msg"`
	Admin  *Admin `json:"admin,omitempty"`
	CodeId int    `json:"code_id"`
	Label  string `json:"label"`
}

type VoucherRedemption struct {
	// The class that these vouchers are being redeemed from.
	Class Class `json:"class"`
	// The tokens belonging to `class` that ought to be redeemed.
	TokenIds []TokenId `json:"token_ids"`
}

type QueryMsg_ClassId struct {
	Contract string `json:"contract"`
}

type QueryMsg_Pauser struct{}

// Expiration represents a point in time when some event happens. It can compare with a BlockInfo and will return is_expired() == true once the condition is hit (and for every block in the future)
type Expiration struct {
	// AtHeight will expire when `env.block.height` >= height
	AtHeight *Expiration_AtHeight `json:"at_height,omitempty"`
	// AtTime will expire when `env.block.time` >= time
	AtTime *Expiration_AtTime `json:"at_time,omitempty"`
	// Never will never expire. Used to express the empty variant
	Never *Expiration_Never `json:"never,omitempty"`
}

type OwnerOfResponse struct {
	// If set this address is approved to transfer/send the token as well
	Approvals []Approval `json:"approvals"`
	// Owner of the token
	Owner string `json:"owner"`
}

type Admin_Instantiator struct{}

type CallbackMsg_RedeemVouchers struct {
	// The address that should receive the tokens.
	Receiver string `json:"receiver"`
	// Information about the vouchers been redeemed.
	Redeem VoucherRedemption `json:"redeem"`
}

type Admin_Address struct {
	Addr string `json:"addr"`
}

type WasmMsg_Instantiate2 struct {
	Salt   Binary  `json:"salt"`
	Admin  *string `json:"admin,omitempty"`
	CodeId int     `json:"code_id"`
	Funds  []Coin  `json:"funds"`
	/*
	   A human-readable label for the contract.

	   Valid values should: - not be empty - not be bigger than 128 bytes (or some chain-specific limit) - not start / end with whitespace
	*/
	Label string `json:"label"`
	// msg is the JSON-encoded InstantiateMsg struct (as raw Binary)
	Msg Binary `json:"msg"`
}
type Expiration_AtTime Timestamp

type CallbackMsg_Mint struct {
	// The address that ought to receive the NFTs. This is a local address, not a bech32 public key.
	Receiver string `json:"receiver"`
	// The tokens to mint on the collection.
	Tokens []Token `json:"tokens"`
	// The class_id to mint for. This must have previously been created with `SaveClass`.
	ClassId ClassId `json:"class_id"`
}

type WasmMsg_Migrate struct {
	ContractAddr string `json:"contract_addr"`
	// msg is the json-encoded MigrateMsg struct that will be passed to the new code
	Msg Binary `json:"msg"`
	// the code_id of the new logic to place in the given contract
	NewCodeId int `json:"new_code_id"`
}

type CallbackMsg_CreateVouchers struct {
	// Information about the vouchers being created.
	Create VoucherCreation `json:"create"`
	// The address that ought to receive the NFT. This is a local address, not a bech32 public key.
	Receiver string `json:"receiver"`
}

type CallbackMsg_Conjunction struct {
	Operands []WasmMsg `json:"operands"`
}

type WasmMsg_UpdateAdmin struct {
	Admin        string `json:"admin"`
	ContractAddr string `json:"contract_addr"`
}

type CallbackMsg_AddIncomingChannelEntries []Tuple_of_Tuple_of_ClassId_and_TokenId_and_string

type Expiration_Never struct{}

type Expiration_AtHeight int

type WasmMsg_Execute struct {
	ContractAddr string `json:"contract_addr"`
	Funds        []Coin `json:"funds"`
	// msg is the json-encoded ExecuteMsg struct (as raw Binary)
	Msg Binary `json:"msg"`
}

type CallbackMsg_RedeemOutgoingChannelEntries []Tuple_of_ClassId_and_TokenId

type WasmMsg_ClearAdmin struct {
	ContractAddr string `json:"contract_addr"`
}

type WasmMsg_Instantiate struct {
	CodeId int    `json:"code_id"`
	Funds  []Coin `json:"funds"`
	/*
	   A human-readable label for the contract.

	   Valid values should: - not be empty - not be bigger than 128 bytes (or some chain-specific limit) - not start / end with whitespace
	*/
	Label string `json:"label"`
	// msg is the JSON-encoded InstantiateMsg struct (as raw Binary)
	Msg   Binary  `json:"msg"`
	Admin *string `json:"admin,omitempty"`
}

var (
	_ json.Marshaler   = (*Tuple_of_Tuple_of_ClassId_and_TokenId_and_string)(nil)
	_ json.Unmarshaler = (*Tuple_of_Tuple_of_ClassId_and_TokenId_and_string)(nil)
)

// Tuple_of_Tuple_of_ClassId_and_TokenId_and_string is a tuple with custom marshal and unmarshal methods
type Tuple_of_Tuple_of_ClassId_and_TokenId_and_string struct {
	F0 Tuple_of_ClassId_and_TokenId
	F1 string
}

// MarshalJSON implements the json.Marshaler interface for Tuple_of_Tuple_of_ClassId_and_TokenId_and_string
func (t Tuple_of_Tuple_of_ClassId_and_TokenId_and_string) MarshalJSON() ([]byte, error) {
	f0, err := json.Marshal(t.F0)
	if err != nil {
		return nil, err
	}

	f1, err := json.Marshal(t.F1)
	if err != nil {
		return nil, err
	}

	return []byte("[" + string(f0) + "," + string(f1) + "]"), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for Tuple_of_Tuple_of_ClassId_and_TokenId_and_string
func (t *Tuple_of_Tuple_of_ClassId_and_TokenId_and_string) UnmarshalJSON(data []byte) error {
	var arr []json.RawMessage
	if err := json.Unmarshal(data, &arr); err != nil {
		return err
	}
	if len(arr) != 2 {
		return errors.New("expected 2 elements in the tuple")
	}

	if err := json.Unmarshal(arr[0], &t.F0); err != nil {
		return err
	}

	if err := json.Unmarshal(arr[1], &t.F1); err != nil {
		return err
	}

	return nil
}

var (
	_ json.Marshaler   = (*Tuple_of_ClassId_and_TokenId)(nil)
	_ json.Unmarshaler = (*Tuple_of_ClassId_and_TokenId)(nil)
)

// Tuple_of_ClassId_and_TokenId is a tuple with custom marshal and unmarshal methods
type Tuple_of_ClassId_and_TokenId struct {
	F0 ClassId
	F1 TokenId
}

// MarshalJSON implements the json.Marshaler interface for Tuple_of_ClassId_and_TokenId
func (t Tuple_of_ClassId_and_TokenId) MarshalJSON() ([]byte, error) {
	f0, err := json.Marshal(t.F0)
	if err != nil {
		return nil, err
	}

	f1, err := json.Marshal(t.F1)
	if err != nil {
		return nil, err
	}

	return []byte("[" + string(f0) + "," + string(f1) + "]"), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for Tuple_of_ClassId_and_TokenId
func (t *Tuple_of_ClassId_and_TokenId) UnmarshalJSON(data []byte) error {
	var arr []json.RawMessage
	if err := json.Unmarshal(data, &arr); err != nil {
		return err
	}
	if len(arr) != 2 {
		return errors.New("expected 2 elements in the tuple")
	}

	if err := json.Unmarshal(arr[0], &t.F0); err != nil {
		return err
	}

	if err := json.Unmarshal(arr[1], &t.F1); err != nil {
		return err
	}

	return nil
}
