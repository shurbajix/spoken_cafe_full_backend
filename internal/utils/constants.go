package utils

import "time"

const (
	StatusSuccess = "success"
	StatusError   = "error"
	StatusFailed  = "failed"

	UserTypeStudent = "student"
	UserTypeTeacher = "teacher"
	UserTypeAdmin   = "admin"

	LessonStatusPending   = "pending"
	LessonStatusActive    = "active"
	LessonStatusCompleted = "completed"
	LessonStatusCancelled = "cancelled"

	BookingStatusPending    = "pending"
	BookingStatusConfirmed  = "confirmed"
	BookingStatusCancelled  = "cancelled"
	BookingStatuesCompleted = "completed"

	PaymentStatusPending   = "pending"
	PaymentStatusCompleted = "completed"
	PaymentStatusFailed    = "failed"
	PaymentStatusRefunded  = "refunded"

	PaymentMethodStripe = "stripe"
	PaymentMethodPaypal = "paypal"
	PaymentMethodWallet = "wallet"

	MessageTypeText  = "text"
	MessageTypeImage = "image"
	MessageTypeVideo = "vidoe"
	MessageTypeAudio = "audio"
	MessageTypeFile  = "file"

	NotificationTypeLesson    = "lesson"
	NotificationTypeBooking   = "booking"
	NotificationTypePayment   = "payment"
	NotificationTypeSystem    = "system"
	NotificationTypePromotion = "promotion"

	MaxFileSize  = 10 * 1024 * 1024  //10MB
	MaxImageSize = 5 * 1024 * 1024   //5MB
	MaxVideoSize = 100 * 1024 * 1024 //100MB
	MaxAudioSize = 20 * 1024 * 1024  // 20MB

	AccessTokenExpiry  = 24 * time.Hour
	RefreshTokenExpiry = 7 * 24 * time.Hour

	RateLimitReqest = 100
	RateLimitWindow = time.Hour

	CollectionUser         = "users"
	ColloectionLesson      = "lessons"
	CollerctionBooking     = "bookings"
	CollectionPayment      = "payments"
	CollectionChat         = "chats"
	CollectionMessage      = "messages"
	CollectionNotification = "notifications"
	CollectionPosts        = "posts"
	CollectionCommonts     = "commonts"
	CollectionAdmin        = "admin"

	// Google Cloud Storage Buckets

	BaucketUserAvatar = "spoken-cafe-user-avatars"
	BaucketLessonFile = "spoken-cafe-lesson-files"
	BaucketChatFile   = "spoken-cafe-chat-files"
	BaucketPostMedia  = "spoken-cafe-post-media"
	// WebSocket Events
	EventMessageSent     = "message-sent"
	EventMessageReceived = "message-received"
	EventUserJoined      = "user-joined"
	EventUserLeft        = "user-left"
	EventTyping          = "typing"
	EventNotification    = "notification"
	EventLessonStart     = "lesson-start"
	EventLessonEnd       = "lesson-end"

	// Validation Rules

	MinPasswordLength = 8
	MaxUsernameLength = 30
	MinUsernameLength = 3
	MaxBioLength      = 500
	MaxLessonTitle    = 100
	MaxLessonDesc     = 1000
	// Pagination
	DefaultPageSize = 20
	MaxPageSize     = 100

	// Languages (for language learning app)
	LangEnglish    = "english"
	LangSpanish    = "spanish"
	LangFrench     = "french"
	LangGerman     = "german"
	LangItalian    = "italian"
	LangPortuguese = "portuguese"
	LangChinese    = "chinese"
	LangJapanese   = "japanese"
	LangKorean     = "korean"
	LangArabic     = "arabic"

	// Proficiency Levels
	LevelBeginner     = "beginner"
	LevelElementary   = "elementary"
	LevelIntermediate = "intermediate"
	LevelUpperInter   = "upper_intermediate"
	LevelAdvanced     = "advanced"
	LevelNative       = "native"
)

// Error Messages
const (
	ErrInvalidCredentials = "invalid credentials"
	ErrUserNotFound       = "user not found"
	ErrUserExists         = "user already exists"
	ErrInvalidToken       = "invalid token"
	ErrTokenExpired       = "token expired"
	ErrUnauthorized       = "unauthorized access"
	ErrForbidden          = "forbidden access"
	ErrInvalidInput       = "invalid input data"
	ErrInternalServer     = "internal server error"
	ErrLessonNotFound     = "lesson not found"
	ErrBookingNotFound    = "booking not found"
	ErrPaymentFailed      = "payment failed"
	ErrFileTooLarge       = "file too large"
	ErrInvalidFileType    = "invalid file type"
	ErrChatNotFound       = "chat not found"
	ErrMessageNotFound    = "message not found"
)

// Success Messages

const (
	MsgUserCreated      = "user created successfully"
	MsgUserUpdated      = "user updated successfully"
	MsgLoginSuccess     = "login successful"
	MsgLogoutSuccess    = "logout successful"
	MsgLessonCreated    = "lesson created successfully"
	MsgLessonUpdated    = "lesson updated successfully"
	MsgBookingCreated   = "booking created successfully"
	MsgBookingCancelled = "booking cancelled successfully"
	MsgPaymentProcessed = "payment processed successfully"
	MsgFileUploaded     = "file uploaded successfully"
	MsgMessageSent      = "message sent successfully"
	MsgNotificationSent = "notification sent successfully"
)

//API Endpoints Prefixes

const (
	APIVersionV1 = "/api/v1"

	//Auth Routes
	AuthPrefix = "/auth"

	//User Routes
	UserPrefix = "/users"

	//Lesson Routes
	LessonPrefix = "/lessons"

	//Booking Routes
	BookingPrefix = "/bookings"

	//payment Routes
	PaymentPrefix = "/payments"

	//Chat Routes
	ChatPrefix = "/chats"

	//File Routes
	FilePrefix = "/files"

	// Admin Routes
	AdminPrefix = "/admin"

	// WebSocket Routes
	WebSocketPrefix = "/ws"
)

// Time Formats

const (
	TimeFormatISO     = "2006-01-02T15:04:05Z07:00"
	TimeFormatDate    = "2006-01-02"
	TimeFormatTime    = "15:04:05"
	TimeFormatDisplay = "January 2, 2006 at 3:04 PM"
)

// Supported File Types

var supportedFileTypes = []string{
	"images/jpeg",
	"images/png",
	"images/jpg",
	"images/gif",
	"images/webp",
}

var SupportedVideoTypes = []string{
	"video/mp4",
	"video/webm",
	"video/ogg",
	"video/avi",
	"video/mov",
}

var SupportedAudioTypes = []string{
	"audio/mp3",
	"audio/mav",
	"audio/ogg",
	"audio/m4a",
}

var SupportedDocTypes = []string{
	"application/pdf",
	"application/msword",
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
}

// Supported Languages Map
var SupportedLanguages = map[string]string{
	LangEnglish:    "English",
	LangSpanish:    "Spanish",
	LangFrench:     "French",
	LangGerman:     "German",
	LangItalian:    "Italian",
	LangPortuguese: "Portuguese",
	LangChinese:    "Chinese",
	LangJapanese:   "Japanese",
	LangKorean:     "Korean",
	LangArabic:     "Arabic",
}

// Proficiency Levels Map
var ProficiencyLevels = map[string]string{
	LevelBeginner:     "Beginner (A1)",
	LevelElementary:   "Elementary (A2)",
	LevelIntermediate: "Intermediate (B1)",
	LevelUpperInter:   "Upper Intermediate (B2)",
	LevelAdvanced:     "Advanced (C1)",
	LevelNative:       "Native (C2)",
}
