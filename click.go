package clicker

import (
	"math"
	"time"

	"github.com/google/uuid"
)

// ClickID ....
type ClickID struct {
	ZoneID            int       `json:"zone_id,omitempty" msgpack:"zone_id"`
	SiteID            uint64    `json:"site_id,omitempty" msgpack:"site_id"`
	FingerPrintHash   string    `json:"fingerprint_hash" msgpack:"fingerprint_hash"` //TODO:: remove
	Uuid              string    `json:"uuid" msgpack:"uuid"`
	WebmasterID       int       `json:"webmaster_id,omitempty" msgpack:"webmaster_id"`
	CreativeID        int       `json:"creative_id,omitempty" msgpack:"creative_id"`
	CampaignID        int       `json:"campaign_id,omitempty" msgpack:"campaign_id"`
	AdvertiserID      int       `json:"advertiser_id,omitempty" msgpack:"advertiser_id"`
	Page              string    `json:"page" msgpack:"page"`
	Ref               string    `json:"ref" msgpack:"ref"`
	AdType            int       `json:"ad_type" msgpack:"ad_type"`
	Link              string    `json:"link" msgpack:"link"`
	Bid               int       `json:"bid" msgpack:"bid"`
	DSPProviderID     int       `json:"dsp_provider_id" msgpack:"dsp_provider_id"`
	SSPProviderID     int       `json:"ssp_provider_id" msgpack:"ssp_provider_id"`
	Nurl              string    `json:"nurl" msgpack:"nurl"`
	Burl              string    `json:"burl" msgpack:"burl"`
	Ip                string    `json:"ip" msgpack:"ip"`
	Os                string    `json:"os" msgpack:"os"`
	OsVersion         string    `json:"os_version" msgpack:"os_version"`
	DeviceType        string    `json:"device_type" msgpack:"device_type"`
	DeviceFamily      string    `json:"device_family" msgpack:"hardware_family"`
	DeviceVendor      string    `json:"device_vendor" msgpack:"device_vendor"`
	MobileOperator    int       `json:"mobile_operator" msgpack:"mobile_operator"`
	Country           int       `json:"country" msgpack:"country"`
	Region            int       `json:"region" msgpack:"region"`
	City              int       `json:"city" msgpack:"city"`
	Language          string    `json:"language" msgpack:"language"`
	Browser           string    `json:"browser" msgpack:"browser"`
	BrowserVersion    string    `json:"browser_version" msgpack:"browser_version"`
	AsnID             int       `json:"asn_id" msgpack:"-"`
	Seabus            bool      `json:"seabus" msgpack:"seabus"`
	SSPDomainHash     string    `json:"ssp_domain_hash" msgpack:"ssp_domain_hash"`
	AuctionNumber     int       `json:"auction_number" msgpack:"auction_number"`
	VisitorIPUA       string    `json:"visitor_ip_ua" msgpack:"visitor_ip_ua"`             //hash from IP + UA,
	VisitorIPHardware string    `json:"visitor_ip_hardware" msgpack:"visitor_ip_hardware"` //hash from get params + IP
	SSPRequestID      string    `json:"ssp_request_id" msgpack:"ssp_request_id"`
	ClientRequestID   string    `json:"client_request_id" msgpack:"client_request_id"`
	ProfileID         string    `json:"profile_id" msgpack:"profile_id"`
	CreatedAt         time.Time `json:"created_at" msgpack:"-"`
	LinkID            int       `json:"link_id" msgpack:"-"`
	RTBSiteID         string    `json:"rtb_site_id" msgpack:"-"`
	Hosting           bool      `json:"hosting" msgpack:"-"`
	VPN               bool      `json:"vpn" msgpack:"-"`
	MncID             string    `json:"mnc_id" msgpack:"-"`
	MncName           string    `json:"mnc_name" msgpack:"-"`
	MncCountry        string    `json:"mnc_country" msgpack:"-"`
	BidType           int       `json:"bid_type" msgpack:"bid_type"`
	BidCPM            int       `json:"bid_cpm" msgpack:"bid_cpm"`
	BidCPC            int       `json:"bid_cpc" msgpack:"bid_cpc"`
	FraudScore        int       `json:"fraud_scope" msgpack:"-"`
	ProfitMargin      int       `json:"profit_margin" msgpack:"profit_margin"`
	ModelNames        []string  `json:"model_names" msgpack:"-"`
	Proxy             bool      `json:"proxy" msgpack:"-"`
}

// NewClickID ...
func NewClickID(
	_type int,
) *ClickID {
	return &ClickID{
		Uuid:      uuid.NewString(),
		AdType:    _type,
		CreatedAt: time.Now(),
	}
}

// CryptoSSPDomainHash ...
func (c *ClickID) CryptoSSPDomainHash(s string) *ClickID {
	c.SSPDomainHash = base64Encode([]byte(s))
	return c
}

// PrepareDSPProviderID ...
func (c *ClickID) PrepareDSPProviderID(v any) *ClickID {
	res, ok := v.(string)
	if !ok {
		return c
	}
	c.DSPProviderID = str2Int(res)
	return c
}

// PrepareZoneID ...
func (c *ClickID) PrepareZoneID(v any) *ClickID {
	res, ok := v.(string)
	if !ok {
		return c
	}
	c.ZoneID = str2Int(res)
	return c
}

// CalcBid ..
func (c *ClickID) CalcBid(f float64) *ClickID {
	c.Bid = int(math.Round(f * 100000000)) //iussik сказал что так по TRUE
	return c
}
