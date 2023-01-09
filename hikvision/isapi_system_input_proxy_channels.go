package hikvision

import (
	"encoding/xml"
	"net/url"
)

type InputProxyChannelList struct {
	InputProxyChannel *InputProxyChannel `xml:"InputProxyChannel"`
}

type InputProxyChannel struct {
	XMLName                      xml.Name                   `xml:"InputProxyChannel,omitempty"`
	Id                           int                        `xml:"id"`
	Name                         string                     `xml:"name"`
	SourceInputPortDescriptor    *SourceInputPortDescriptor `xml:"SourceInputPortDescriptor"`
	CertificateValidationEnabled bool                       `xml:"certificateValidationEnabled"`
	DefaultAdminPortEnabled      bool                       `xml:"defaultAdminPortEnabled"`
	EnableAnr                    bool                       `xml:"enableAnr"`
	EnableTiming                 bool                       `xml:"enableTiming"`
	DevIndex                     string                     `xml:"devIndex"`
}

type SourceInputPortDescriptor struct {
	ProxyProtocol        string `xml:"proxyProtocol"`
	AddressingFormatType string `xml:"addressingFormatType"`
	IpAddress            string `xml:"ipAddress"`
	ManagePortNo         int    `xml:"managePortNo"`
	SrcInputPort         int    `xml:"srcInputPort"`
	UserName             string `xml:"userName"`
	StreamType           string `xml:"streamType"`
	DeviceID             int    `xml:"deviceID"`
}

func (c *Client) GetInputProxyChannels() (resp *InputProxyChannelList, err error) {
	path := "/ISAPI/ContentMgmt/InputProxy/channels"
	u, err := url.Parse(c.BaseURL + path)
	if err != nil {
		return nil, err
	}
	body, err := c.Get(u)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
