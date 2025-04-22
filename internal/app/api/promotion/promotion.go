package promotion

// Promotion is the model
type Promotion struct {
	ID          int      `json:"id" db:"id"`
	Code        string   `json:"code" db:"code"`
	Name        string   `json:"name" db:"name"`
	Priority    int      `json:"priority" db:"priority"`
	Exclusive   bool     `json:"exclusive" db:"exclusive"`
	Used        int      `json:"used" db:"used"`
	CouponBased bool     `json:"couponBased" db:"coupon_based"`
	Rules       []string `json:"rules" db:"rules"`
	Actions     []string `json:"actions" db:"actions"`
	CreatedAt   string   `json:"createdAt" db:"created_at"`
	UpdatedAt   string   `json:"updatedAt" db:"updated_at"`
	Channels    []string `json:"channels" db:"channels"`
	Links       struct {
		Self struct {
			Href string `json:"href" db:"href"`
		} `json:"self" db:"self"`
	} `json:"_links" db:"_links"`
}
