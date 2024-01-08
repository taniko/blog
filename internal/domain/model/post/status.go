package post

type (
	Visibility string // 公開範囲
	State      string // 投稿状態
)

const (
	VisibilityPublic  Visibility = "public"  // 公開
	VisibilityPrivate Visibility = "private" // 非公開
	VisibilityDraft   Visibility = "draft"   // 下書き
)

const (
	StatePublished State = "published" // 公開
	StateArchived  State = "archived"  // 削除済み
)
