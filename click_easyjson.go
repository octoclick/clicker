// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package clicker

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonE737ea52DecodeGithubComOctoclickClicker(in *jlexer.Lexer, out *ClickID) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "zone_id":
			out.ZoneID = int(in.Int())
		case "site_id":
			out.SiteID = uint64(in.Uint64())
		case "fingerprint_hash":
			out.FingerPrintHash = string(in.String())
		case "uuid":
			out.Uuid = string(in.String())
		case "webmaster_id":
			out.WebmasterID = int(in.Int())
		case "creative_id":
			out.CreativeID = int(in.Int())
		case "campaign_id":
			out.CampaignID = int(in.Int())
		case "advertiser_id":
			out.AdvertiserID = int(in.Int())
		case "page":
			out.Page = string(in.String())
		case "ref":
			out.Ref = string(in.String())
		case "ad_type":
			out.AdType = int(in.Int())
		case "link":
			out.Link = string(in.String())
		case "bid":
			out.Bid = int(in.Int())
		case "dsp_provider_id":
			out.DSPProviderID = int(in.Int())
		case "ssp_provider_id":
			out.SSPProviderID = int(in.Int())
		case "nurl":
			out.Nurl = string(in.String())
		case "burl":
			out.Burl = string(in.String())
		case "ip":
			out.Ip = string(in.String())
		case "os":
			out.Os = string(in.String())
		case "os_version":
			out.OsVersion = string(in.String())
		case "device_type":
			out.DeviceType = string(in.String())
		case "device_family":
			out.DeviceFamily = string(in.String())
		case "device_vendor":
			out.DeviceVendor = string(in.String())
		case "mobile_operator":
			out.MobileOperator = int(in.Int())
		case "country":
			out.Country = int(in.Int())
		case "region":
			out.Region = int(in.Int())
		case "city":
			out.City = int(in.Int())
		case "language":
			out.Language = string(in.String())
		case "browser":
			out.Browser = string(in.String())
		case "browser_version":
			out.BrowserVersion = string(in.String())
		case "seabus":
			out.Seabus = bool(in.Bool())
		case "ssp_domain_hash":
			out.SSPDomainHash = string(in.String())
		case "auction_number":
			out.AuctionNumber = int(in.Int())
		case "visitor_ip_ua":
			out.VisitorIPUA = string(in.String())
		case "visitor_ip_hardware":
			out.VisitorIPHardware = string(in.String())
		case "ssp_request_id":
			out.SSPRequestID = string(in.String())
		case "client_request_id":
			out.ClientRequestID = string(in.String())
		case "profile_id":
			out.ProfileID = string(in.String())
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "link_id":
			out.LinkID = int(in.Int())
		case "rtb_site_id":
			out.RTBSiteID = string(in.String())
		case "proxy":
			out.Proxy = bool(in.Bool())
		case "hosting":
			out.Hosting = bool(in.Bool())
		case "vpn":
			out.VPN = bool(in.Bool())
		case "mnc_id":
			out.MncID = string(in.String())
		case "mnc_name":
			out.MncName = string(in.String())
		case "mnc_country":
			out.MncCountry = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE737ea52EncodeGithubComOctoclickClicker(out *jwriter.Writer, in ClickID) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ZoneID != 0 {
		const prefix string = ",\"zone_id\":"
		first = false
		out.RawString(prefix[1:])
		out.Int(int(in.ZoneID))
	}
	if in.SiteID != 0 {
		const prefix string = ",\"site_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.SiteID))
	}
	{
		const prefix string = ",\"fingerprint_hash\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.FingerPrintHash))
	}
	{
		const prefix string = ",\"uuid\":"
		out.RawString(prefix)
		out.String(string(in.Uuid))
	}
	if in.WebmasterID != 0 {
		const prefix string = ",\"webmaster_id\":"
		out.RawString(prefix)
		out.Int(int(in.WebmasterID))
	}
	if in.CreativeID != 0 {
		const prefix string = ",\"creative_id\":"
		out.RawString(prefix)
		out.Int(int(in.CreativeID))
	}
	if in.CampaignID != 0 {
		const prefix string = ",\"campaign_id\":"
		out.RawString(prefix)
		out.Int(int(in.CampaignID))
	}
	if in.AdvertiserID != 0 {
		const prefix string = ",\"advertiser_id\":"
		out.RawString(prefix)
		out.Int(int(in.AdvertiserID))
	}
	{
		const prefix string = ",\"page\":"
		out.RawString(prefix)
		out.String(string(in.Page))
	}
	{
		const prefix string = ",\"ref\":"
		out.RawString(prefix)
		out.String(string(in.Ref))
	}
	{
		const prefix string = ",\"ad_type\":"
		out.RawString(prefix)
		out.Int(int(in.AdType))
	}
	{
		const prefix string = ",\"link\":"
		out.RawString(prefix)
		out.String(string(in.Link))
	}
	{
		const prefix string = ",\"bid\":"
		out.RawString(prefix)
		out.Int(int(in.Bid))
	}
	{
		const prefix string = ",\"dsp_provider_id\":"
		out.RawString(prefix)
		out.Int(int(in.DSPProviderID))
	}
	{
		const prefix string = ",\"ssp_provider_id\":"
		out.RawString(prefix)
		out.Int(int(in.SSPProviderID))
	}
	{
		const prefix string = ",\"nurl\":"
		out.RawString(prefix)
		out.String(string(in.Nurl))
	}
	{
		const prefix string = ",\"burl\":"
		out.RawString(prefix)
		out.String(string(in.Burl))
	}
	{
		const prefix string = ",\"ip\":"
		out.RawString(prefix)
		out.String(string(in.Ip))
	}
	{
		const prefix string = ",\"os\":"
		out.RawString(prefix)
		out.String(string(in.Os))
	}
	{
		const prefix string = ",\"os_version\":"
		out.RawString(prefix)
		out.String(string(in.OsVersion))
	}
	{
		const prefix string = ",\"device_type\":"
		out.RawString(prefix)
		out.String(string(in.DeviceType))
	}
	{
		const prefix string = ",\"device_family\":"
		out.RawString(prefix)
		out.String(string(in.DeviceFamily))
	}
	{
		const prefix string = ",\"device_vendor\":"
		out.RawString(prefix)
		out.String(string(in.DeviceVendor))
	}
	{
		const prefix string = ",\"mobile_operator\":"
		out.RawString(prefix)
		out.Int(int(in.MobileOperator))
	}
	{
		const prefix string = ",\"country\":"
		out.RawString(prefix)
		out.Int(int(in.Country))
	}
	{
		const prefix string = ",\"region\":"
		out.RawString(prefix)
		out.Int(int(in.Region))
	}
	{
		const prefix string = ",\"city\":"
		out.RawString(prefix)
		out.Int(int(in.City))
	}
	{
		const prefix string = ",\"language\":"
		out.RawString(prefix)
		out.String(string(in.Language))
	}
	{
		const prefix string = ",\"browser\":"
		out.RawString(prefix)
		out.String(string(in.Browser))
	}
	{
		const prefix string = ",\"browser_version\":"
		out.RawString(prefix)
		out.String(string(in.BrowserVersion))
	}
	{
		const prefix string = ",\"seabus\":"
		out.RawString(prefix)
		out.Bool(bool(in.Seabus))
	}
	{
		const prefix string = ",\"ssp_domain_hash\":"
		out.RawString(prefix)
		out.String(string(in.SSPDomainHash))
	}
	{
		const prefix string = ",\"auction_number\":"
		out.RawString(prefix)
		out.Int(int(in.AuctionNumber))
	}
	{
		const prefix string = ",\"visitor_ip_ua\":"
		out.RawString(prefix)
		out.String(string(in.VisitorIPUA))
	}
	{
		const prefix string = ",\"visitor_ip_hardware\":"
		out.RawString(prefix)
		out.String(string(in.VisitorIPHardware))
	}
	{
		const prefix string = ",\"ssp_request_id\":"
		out.RawString(prefix)
		out.String(string(in.SSPRequestID))
	}
	{
		const prefix string = ",\"client_request_id\":"
		out.RawString(prefix)
		out.String(string(in.ClientRequestID))
	}
	{
		const prefix string = ",\"profile_id\":"
		out.RawString(prefix)
		out.String(string(in.ProfileID))
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"link_id\":"
		out.RawString(prefix)
		out.Int(int(in.LinkID))
	}
	{
		const prefix string = ",\"rtb_site_id\":"
		out.RawString(prefix)
		out.String(string(in.RTBSiteID))
	}
	{
		const prefix string = ",\"proxy\":"
		out.RawString(prefix)
		out.Bool(bool(in.Proxy))
	}
	{
		const prefix string = ",\"hosting\":"
		out.RawString(prefix)
		out.Bool(bool(in.Hosting))
	}
	{
		const prefix string = ",\"vpn\":"
		out.RawString(prefix)
		out.Bool(bool(in.VPN))
	}
	{
		const prefix string = ",\"mnc_id\":"
		out.RawString(prefix)
		out.String(string(in.MncID))
	}
	{
		const prefix string = ",\"mnc_name\":"
		out.RawString(prefix)
		out.String(string(in.MncName))
	}
	{
		const prefix string = ",\"mnc_country\":"
		out.RawString(prefix)
		out.String(string(in.MncCountry))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ClickID) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE737ea52EncodeGithubComOctoclickClicker(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ClickID) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE737ea52EncodeGithubComOctoclickClicker(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ClickID) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE737ea52DecodeGithubComOctoclickClicker(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ClickID) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE737ea52DecodeGithubComOctoclickClicker(l, v)
}
