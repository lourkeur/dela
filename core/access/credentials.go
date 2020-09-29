package access

import "fmt"

// ContractCredential defines the credential for a contract. It contains the
// name of the contract and an associated command.
type ContractCredential struct {
	id       []byte
	contract string
	command  string
}

// NewContractCreds creates new credential from the associated identifier, the
// name of the contract and its command.
func NewContractCreds(id []byte, contract, command string) ContractCredential {
	return ContractCredential{
		id:       id,
		contract: contract,
		command:  command,
	}
}

// GetID returns the identifier for the credential.
func (cc ContractCredential) GetID() []byte {
	return append([]byte{}, cc.id...)
}

// GetRule returns the scope of the credential.
func (cc ContractCredential) GetRule() string {
	return fmt.Sprintf("%s:%s", cc.contract, cc.command)
}
