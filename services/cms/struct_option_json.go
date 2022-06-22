package cms

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// OptionJson is a nested struct in cms response
type OptionJson struct {
	RequestFormat          string  `json:"request_format" xml:"request_format"`
	ResponseContent        string  `json:"response_content" xml:"response_content"`
	Port                   int     `json:"port" xml:"port"`
	ProxyProtocol          bool    `json:"proxy_protocol" xml:"proxy_protocol"`
	Authentication         int     `json:"authentication" xml:"authentication"`
	MatchRule              int     `json:"match_rule" xml:"match_rule"`
	DnsMatchRule           string  `json:"dns_match_rule" xml:"dns_match_rule"`
	Ipv6Task               string  `json:"ipv6_task" xml:"ipv6_task"`
	RequestContent         string  `json:"request_content" xml:"request_content"`
	AcceptableResponseCode string  `json:"acceptable_response_code" xml:"acceptable_response_code"`
	Username               string  `json:"username" xml:"username"`
	Traceroute             int64   `json:"traceroute" xml:"traceroute"`
	DnsType                string  `json:"dns_type" xml:"dns_type"`
	ResponseFormat         string  `json:"response_format" xml:"response_format"`
	Password               string  `json:"password" xml:"password"`
	ExpectValue            string  `json:"expect_value" xml:"expect_value"`
	TimeOut                int64   `json:"time_out" xml:"time_out"`
	FailureRate            float64 `json:"failure_rate" xml:"failure_rate"`
	Header                 string  `json:"header" xml:"header"`
	Cookie                 string  `json:"cookie" xml:"cookie"`
	PingNum                int     `json:"ping_num" xml:"ping_num"`
	HttpMethod             string  `json:"http_method" xml:"http_method"`
	UnfollowRedirect       bool    `json:"unfollow_redirect" xml:"unfollow_redirect"`
	CertVerify             bool    `json:"cert_verify" xml:"cert_verify"`
	DnsServer              string  `json:"dns_server" xml:"dns_server"`
	EnableOperatorDns      bool    `json:"enable_operator_dns" xml:"enable_operator_dns"`
	Attempts               int64   `json:"attempts" xml:"attempts"`
	Protocol               string  `json:"protocol" xml:"protocol"`
	IsBase64Encode         string  `json:"isBase64Encode" xml:"isBase64Encode"`
	DiagnosisMtr           bool    `json:"diagnosis_mtr" xml:"diagnosis_mtr"`
	DiagnosisPing          bool    `json:"diagnosis_ping" xml:"diagnosis_ping"`
	RetryDelay             int     `json:"retry_delay" xml:"retry_delay"`
}
