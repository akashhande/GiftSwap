package schema

type GiftExchange struct {
	AssignerID  uint `json:"assigner_id"`
	RecipientID uint `json:"recipient_id"`
	Year        int  `json:"assigned_year"`
}
