## module personal

The JSON-RPC API's `personal` namespace has historically been used to manage accounts and sign transactions and data over RPC.

### personal_listAccounts

The `listAccounts` method displays the addresses of all accounts in the keystore. It is identical to eth.accounts.

Parameters:

None

Returns:

result: array of address

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "personal_listAccounts"
}' | jq
```

### personal_deriveAccount

The `deriveAccount` method requests a hardware wallet to derive a new account, optionally pinning it for later use.

Parameters:

- url: string, required
- path: string, required
- pin: bool, optional

Returns:

result: object of Account

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "personal_deriveAccount",
  "params": [
    "http://127.0.0.1/",
    "m/44/60/0/0/0",
    true
  ]
}' | jq
```

### personal_ecRecover

The `ecRecover` method returns the address for the account that was used to create a signature.

Parameters:

- data: array of byte, required
- sig: array of byte, required

Returns:

result: address

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "personal_ecRecover",
  "params": [
    "0xaabbccdd",
    "0x5b6693f153b48ec1c706ba4169960386dbaa6903e249cc79a8e6ddc434451d417e1e57327872c7f538beeb323c300afa9999a3d4a5de6caf3be0d5ef832b67ef1c"
  ]
}' | jq
```

### personal_importRawKey

The `importRawKey` method was used to create a new account in the keystore from a raw private key.

Parameters:

- privkey: string, required, private key
- password: string, required, password

Returns:

result: address

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "personal_importRawKey",
  "params": [
    "cd87934ee007b7a458fa00dc0314fff8b2bd43b3071f46c820c379e483b4fd8e",
    "password"
  ]
}' | jq
```

### personal_listWallets

The `listWallets` method lists full details, including usb path or keystore-file paths.

Parameters:

- privkey: string, required, private key
- password: string, required, password

Returns:

result: address

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "personal_listWallets",
  "params": [
    "cd87934ee007b7a458fa00dc0314fff8b2bd43b3071f46c820c379e483b4fd8e",
    "password"
  ]
}' | jq
```

### personal_lockAccount

The `lockAccount` method removes the private key with a given address from memory. The account can no longer be used to send transactions.

Parameters:

- addr: address, required

Returns:

result: bool

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "personal_lockAccount",
  "params": [
    "0xD4CE02705041F04135f1949Bc835c1Fe0885513c"
  ]
}' | jq
```

### personal_newAccount

The `newAccount` method was used to create a new account and save it in the keystore.

Parameters:

- password: string, required

Returns:

result: address

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "personal_newAccount",
  "params": [
    "password"
  ]
}' | jq
```

### personal_openWallet

The `openWallet` method initiates a hardware wallet opening procedure by establishing a USB connection and then attempting to authenticate via the provided passphrase. Note, the method may return an extra challenge requiring a second open (e.g. the Trezor PIN matrix challenge).

Parameters:

- url: string, required
- passphrase: string, optional

Returns:

result: error

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "personal_openWallet",
  "params": [
    "keystore://",
    "password"
  ]
}' | jq
```

### personal_sendTransaction

The `sendTransaction` method is used to sign and submit a transaction. This can be done using eth_sendTransaction

Parameters:

- args: object TransactionArgs, required
- passwd: string, required

Returns:

result: hash

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "personal_sendTransaction",
  "params": [
    {
      "from": "0x1d4e05bb72677cb8fa576149c945b57d13f855e4",
      "to": "0xafa3f8684e54059998bc3a7b0d2b0da075154d66",
      "value": "0x1230000000"
    },
    "password"
  ]
}' | jq
```

### personal_sign

The `sign` method calculates an Ethereum specific signature with sign(keccak256("\x19Ethereum Signed Message:\n" + len(message) + message)). Adding a prefix to the message makes the calculated signature recognisable as an Ethereum specific signature. This prevents misuse where a malicious DApp can sign arbitrary data (e.g. transaction) and use the signature to impersonate the victim.

Parameters:

- data: array of bytes, required
- addr: address, required
- passwd: string, required

Returns:

result: array of byte

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "personal_sign",
  "params": [
    "0xdeadbeaf",
    "0x413ba0e5f6f00664598b5c80042b1308f4ff1408",
    "password"
  ]
}' | jq
```

### personal_signTransaction

The `signTransaction` was used to create and sign a transaction from the given arguments. The transaction was returned in RLP-form, not broadcast to other nodes.

Parameters:

- args: object TransactionArgs, required
- passwd: string, required

Returns:

result: ojbect Transaction

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "personal_signTransaction",
  "params": [
    {
      "from": "0x77982323172e5b6182539d3522d5a33a944206d4",
      "to": "0xcd6bfdb523a4d030890d28bf1eb6ef36307c9aaa",
      "value": "0x10000",
      "gas": "0x1000000",
      "nonce": "0x2",
      "gasPrice": "0x25000000000"
    }
  ]
}' | jq
```

### personal_unlockAccount

The `unlockAccount` method decrypts the key with the given address from the key store.

Parameters:

- addr: address, required
- password: string, required
- duration: uint64, optional

Returns:

result: bool

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "personal_unlockAccount",
  "params": [
    "0xD4CE02705041F04135f1949Bc835c1Fe0885513c",
    "password",
    30
  ]
}' | jq
```
