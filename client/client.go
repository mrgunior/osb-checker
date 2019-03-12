package client

import (
	"encoding/json"
	"strconv"
	"strings"

	v2 "github.com/openservicebrokerapi/osb-checker/autogenerated/models"
	. "github.com/openservicebrokerapi/osb-checker/config"
)

var Default = NewClient(NewReceiver())

func NewClient(r Receiver) *Client {
	return &Client{
		Receiver: r,
	}
}

type Client struct {
	Receiver
}

func (c *Client) GetCatalog() (int, *v2.Catalog, error) {
	var catalog v2.Catalog
	params := &BrokerRequestParams{
		URL:    strings.Join([]string{CONF.URL, "v2/catalog"}, "/"),
		Method: "GET",
		HeaderOption: map[string]string{
			"X-Broker-API-Version": CONF.APIVersion,
		},
		Username: CONF.Authentication.Username,
		Password: CONF.Authentication.Password,
	}

	res, err := c.Recv(params)
	if err != nil {
		return -1, nil, err
	}
	json.Unmarshal(res.Message, &catalog)

	return res.Code, &catalog, nil
}

func (c *Client) Provision(
	instanceID string,
	req *v2.ServiceInstanceProvisionRequest,
	async bool,
) (int, *v2.ServiceInstanceAsyncOperation, error) {

	var asyncOperation v2.ServiceInstanceAsyncOperation
	params := &BrokerRequestParams{
		URL:    strings.Join([]string{CONF.URL, "v2/service_instances", instanceID}, "/"),
		Method: "PUT",
		HeaderOption: map[string]string{
			"X-Broker-API-Version": CONF.APIVersion,
		},
		QueryParam: map[string]string{
			"accepts_incomplete": strconv.FormatBool(async),
		},
		Username:  CONF.Authentication.Username,
		Password:  CONF.Authentication.Password,
		InputData: req,
	}

	res, err := c.Recv(params)
	if err != nil {
		return -1, nil, err
	}
	json.Unmarshal(res.Message, &asyncOperation)

	return res.Code, &asyncOperation, nil
}

func (c *Client) PollInstanceLastOperation(
	instanceID string,
) (int, *v2.LastOperationResource, error) {

	var lastOperation v2.LastOperationResource
	params := &BrokerRequestParams{
		URL:    strings.Join([]string{CONF.URL, "v2/service_instances", instanceID, "last_operation"}, "/"),
		Method: "GET",
		HeaderOption: map[string]string{
			"X-Broker-API-Version": CONF.APIVersion,
		},
		Username: CONF.Authentication.Username,
		Password: CONF.Authentication.Password,
	}

	res, err := c.Recv(params)
	if err != nil {
		return -1, nil, err
	}
	json.Unmarshal(res.Message, &lastOperation)

	return res.Code, &lastOperation, nil
}

func (c *Client) GetInstance(
	instanceID string,
) (int, *v2.ServiceInstanceResource, error) {

	var instance v2.ServiceInstanceResource
	params := &BrokerRequestParams{
		URL:    strings.Join([]string{CONF.URL, "v2/service_instances", instanceID}, "/"),
		Method: "GET",
		HeaderOption: map[string]string{
			"X-Broker-API-Version": CONF.APIVersion,
		},
		Username: CONF.Authentication.Username,
		Password: CONF.Authentication.Password,
	}

	res, err := c.Recv(params)
	if err != nil {
		return -1, nil, err
	}
	json.Unmarshal(res.Message, &instance)

	return res.Code, &instance, nil
}

func (c *Client) UpdateInstance(
	instanceID string,
	req *v2.ServiceInstanceUpdateRequest,
	async bool,
) (int, *v2.ServiceInstanceAsyncOperation, error) {

	var asyncOperation v2.ServiceInstanceAsyncOperation
	params := &BrokerRequestParams{
		URL:    strings.Join([]string{CONF.URL, "v2/service_instances", instanceID}, "/"),
		Method: "PATCH",
		HeaderOption: map[string]string{
			"X-Broker-API-Version": CONF.APIVersion,
		},
		QueryParam: map[string]string{
			"accepts_incomplete": strconv.FormatBool(async),
		},
		Username:  CONF.Authentication.Username,
		Password:  CONF.Authentication.Password,
		InputData: req,
	}

	res, err := c.Recv(params)
	if err != nil {
		return -1, nil, err
	}
	json.Unmarshal(res.Message, &asyncOperation)

	return res.Code, &asyncOperation, nil
}

func (c *Client) Deprovision(
	instanceID string,
	serviceID, planID string,
	async bool,
) (int, error) {

	params := &BrokerRequestParams{
		URL:    strings.Join([]string{CONF.URL, "v2/service_instances", instanceID}, "/"),
		Method: "DELETE",
		HeaderOption: map[string]string{
			"X-Broker-API-Version": CONF.APIVersion,
		},
		QueryParam: map[string]string{
			"service_id":         serviceID,
			"plan_id":            planID,
			"accepts_incomplete": strconv.FormatBool(async),
		},
		Username: CONF.Authentication.Username,
		Password: CONF.Authentication.Password,
	}

	res, err := c.Recv(params)
	if err != nil {
		return -1, err
	}

	return res.Code, nil
}

func (c *Client) Bind(
	instanceID, bindingID string,
	req *v2.ServiceBindingRequest,
	async bool,
) (int, *v2.ServiceBinding, error) {

	var binding v2.ServiceBinding
	params := &BrokerRequestParams{
		URL:    strings.Join([]string{CONF.URL, "v2/service_instances", instanceID, "service_bindings", bindingID}, "/"),
		Method: "PUT",
		HeaderOption: map[string]string{
			"X-Broker-API-Version": CONF.APIVersion,
		},
		QueryParam: map[string]string{
			"accepts_incomplete": strconv.FormatBool(async),
		},
		Username:  CONF.Authentication.Username,
		Password:  CONF.Authentication.Password,
		InputData: req,
	}

	res, err := c.Recv(params)
	if err != nil {
		return -1, nil, err
	}
	json.Unmarshal(res.Message, &binding)

	return res.Code, &binding, nil
}

func (c *Client) PollBindingLastOperation(
	instanceID, bindingID string,
) (int, *v2.LastOperationResource, error) {

	var lastOperation v2.LastOperationResource
	params := &BrokerRequestParams{
		URL:    strings.Join([]string{CONF.URL, "v2/service_instances", instanceID, "service_bindings", bindingID, "last_operation"}, "/"),
		Method: "GET",
		HeaderOption: map[string]string{
			"X-Broker-API-Version": CONF.APIVersion,
		},
		Username: CONF.Authentication.Username,
		Password: CONF.Authentication.Password,
	}

	res, err := c.Recv(params)
	if err != nil {
		return -1, nil, err
	}
	json.Unmarshal(res.Message, &lastOperation)

	return res.Code, &lastOperation, nil
}

func (c *Client) GetBinding(
	instanceID, bindingID string,
) (int, *v2.ServiceBindingResource, error) {

	var binding v2.ServiceBindingResource
	params := &BrokerRequestParams{
		URL:    strings.Join([]string{CONF.URL, "v2/service_instances", instanceID, "service_bindings", bindingID}, "/"),
		Method: "GET",
		HeaderOption: map[string]string{
			"X-Broker-API-Version": CONF.APIVersion,
		},
		Username: CONF.Authentication.Username,
		Password: CONF.Authentication.Password,
	}

	res, err := c.Recv(params)
	if err != nil {
		return -1, nil, err
	}
	json.Unmarshal(res.Message, &binding)

	return res.Code, &binding, nil
}

func (c *Client) Unbind(
	instanceID, bindingID,
	serviceID, planID string,
	async bool,
) (int, error) {

	params := &BrokerRequestParams{
		URL:    strings.Join([]string{CONF.URL, "v2/service_instances", instanceID, "service_bindings", bindingID}, "/"),
		Method: "DELETE",
		HeaderOption: map[string]string{
			"X-Broker-API-Version": CONF.APIVersion,
		},
		QueryParam: map[string]string{
			"service_id":         serviceID,
			"plan_id":            planID,
			"accepts_incomplete": strconv.FormatBool(async),
		},
		Username: CONF.Authentication.Username,
		Password: CONF.Authentication.Password,
	}

	res, err := c.Recv(params)
	if err != nil {
		return -1, err
	}

	return res.Code, nil
}
