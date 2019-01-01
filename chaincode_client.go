package fabclient

import (
	"fmt"
)

// ChaincodeClient contains data to interact with chaincode on behalf of user
type ChaincodeClient struct {
	chaincodeID string
	userClient  *UserClient
}

// CreateChaincodeClient is the same as  (c *FabricClient) CreateChaincodeClient(channelID string, name string, organization string) but it does not reuse Fabric Client
func CreateChaincodeClient(configPath string, ordererHost string, channelID string, chaincodeID string, name string, organization string) (*ChaincodeClient, error) {
	fabricClient, err := CreateFabricClient(configPath, ordererHost)
	if err != nil {
		return nil, err
	}
	return fabricClient.CreateChaincodeClient(channelID, chaincodeID, name, organization)
}

// Invoke triggers invokation of transaction
func (c *ChaincodeClient) Invoke(functionName string, args [][]byte) ([]byte, error) {
	resp, err := c.userClient.Invoke(c.chaincodeID, functionName, args)
	if err != nil {
		return nil, fmt.Errorf("Failed to invoke chaincode %s with funactions %s and arguments %v.\n Error: %v", c.chaincodeID, functionName, args, err)
	}
	logger.Debugf("Response on invoke chaincode: %s\n", resp)
	return resp, nil
}

// Query is the same as Invoke but without sending transaction to orderer so tx does not added to blockchain history. It is used for querying data
func (c *ChaincodeClient) Query(functionName string, args [][]byte) ([]byte, error) {
	resp, err := c.userClient.Query(c.chaincodeID, functionName, args)
	if err != nil {
		return nil, fmt.Errorf("Failed to query chaincode %s with funactions %s and arguments %v.\n Error: %v", c.chaincodeID, functionName, args, err)
	}
	logger.Debugf("Response on query chaincode: %s\n", resp)
	return resp, nil
}

// GetUserClient returns UserClient of ChaincodeClient
func (c *ChaincodeClient) GetUserClient() *UserClient {
	return c.userClient
}
