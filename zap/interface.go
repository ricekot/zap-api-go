// Zed Attack Proxy (ZAP) and its related class files.
//
// ZAP is an HTTP/HTTPS proxy for assessing web application security.
//
// Copyright 2017 the ZAP development team
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// TODO: auto generate this file
package zap

// Interface defines the interface a ZAP client should implement
type Interface interface {
	AccessControl() *AccessControl
	Acsrf() *Acsrf
	AjaxSpider() *AjaxSpider
	AlertFilter() *AlertFilter
	Ascan() *Ascan
	Authentication() *Authentication
	Authorization() *Authorization
	Autoupdate() *Autoupdate
	Break() *Break
	Context() *Context
	Core() *Core
	Exportreport() *Exportreport
	ForcedUser() *ForcedUser
	HttpSessions() *HttpSessions
	ImportLogFiles() *ImportLogFiles
	Importurls() *Importurls
	Openapi() *Openapi
	Params() *Params
	Pnh() *Pnh
	Pscan() *Pscan
	Replacer() *Replacer
	Reveal() *Reveal
	Revisit() *Revisit
	Script() *Script
	Search() *Search
	Selenium() *Selenium
	Soap() *Soap
	Spider() *Spider
	Stats() *Stats
	Users() *Users
	Wappalyzer() *Wappalyzer
	Websocket() *Websocket
}

// AccessControl() returns a AccessControl client
func (c *Client) AccessControl() *AccessControl {
	return &AccessControl{c}
}

// Acsrf() returns a Acsrf client
func (c *Client) Acsrf() *Acsrf {
	return &Acsrf{c}
}

// AjaxSpider() returns a AjaxSpider client
func (c *Client) AjaxSpider() *AjaxSpider {
	return &AjaxSpider{c}
}

// AlertFilter() returns a AlertFilter client
func (c *Client) AlertFilter() *AlertFilter {
	return &AlertFilter{c}
}

// Ascan() returns a Ascan client
func (c *Client) Ascan() *Ascan {
	return &Ascan{c}
}

// Authentication() returns a Authentication client
func (c *Client) Authentication() *Authentication {
	return &Authentication{c}
}

// Authorization() returns a Authorization client
func (c *Client) Authorization() *Authorization {
	return &Authorization{c}
}

// Autoupdate returns an Autoupdate client
func (c *Client) Autoupdate() *Autoupdate {
	return &Autoupdate{c}
}

// Break() returns a Break client
func (c *Client) Break() *Break {
	return &Break{c}
}

// Context() returns a Context client
func (c *Client) Context() *Context {
	return &Context{c}
}

// Core() returns a Core client
func (c *Client) Core() *Core {
	return &Core{c}
}

// Exportreport() returns a Exportreport client
func (c *Client) Exportreport() *Exportreport {
	return &Exportreport{c}
}

// ForcedUser() returns a ForcedUser client
func (c *Client) ForcedUser() *ForcedUser {
	return &ForcedUser{c}
}

// HttpSessions() returns a HttpSessions client
func (c *Client) HttpSessions() *HttpSessions {
	return &HttpSessions{c}
}

// ImportLogFiles() returns a ImportLogFiles client
func (c *Client) ImportLogFiles() *ImportLogFiles {
	return &ImportLogFiles{c}
}

// Importurls() returns a Importurls client
func (c *Client) Importurls() *Importurls {
	return &Importurls{c}
}

// Openapi() returns a Openapi clinet
func (c *Client) Openapi() *Openapi {
	return &Openapi{c}
}

// Params() returns a Params client
func (c *Client) Params() *Params {
	return &Params{c}
}

// Pnh() returns a Pnh client
func (c *Client) Pnh() *Pnh {
	return &Pnh{c}
}

// Pscan() returns a Pscan client
func (c *Client) Pscan() *Pscan {
	return &Pscan{c}
}

// Replacer() returns a Replacer client
func (c *Client) Replacer() *Replacer {
	return &Replacer{c}
}

// Reveal() returns a Reveal client
func (c *Client) Reveal() *Reveal {
	return &Reveal{c}
}

// Revisit() returns a Revisit client
func (c *Client) Revisit() *Revisit {
	return &Revisit{c}
}

// Script() returns a Script client
func (c *Client) Script() *Script {
	return &Script{c}
}

// Search() returns a Search client
func (c *Client) Search() *Search {
	return &Search{c}
}

// Selenium() returns a Selenium client
func (c *Client) Selenium() *Selenium {
	return &Selenium{c}
}

// Soap() returns a Soap client
func (c *Client) Soap() *Soap {
	return &Soap{c}
}

// Spider() returns a Spider client
func (c *Client) Spider() *Spider {
	return &Spider{c}
}

// Stats() returns a Stats client
func (c *Client) Stats() *Stats {
	return &Stats{c}
}

// Users() returns a Users client
func (c *Client) Users() *Users {
	return &Users{c}
}

// Wappalyzer() returns a Wappalyzer client
func (c *Client) Wappalyzer() *Wappalyzer {
	return &Wappalyzer{c}
}

// Websocket() returns a Websocket client
func (c *Client) Websocket() *Websocket {
	return &Websocket{c}
}
