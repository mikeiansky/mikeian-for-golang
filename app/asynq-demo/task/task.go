// task/task.go
package task

import "fmt"

// Task 任务类型常量
const (
	TypeEmailDelivery = "email:deliver"
	TypeImageResize   = "image:resize"
)

// EmailDeliveryPayload 邮件任务负载
type EmailDeliveryPayload struct {
	UserID     int
	Email      string
	TemplateID string
}

// ImageResizePayload 图片处理任务负载
type ImageResizePayload struct {
	ImageURL  string
	Width     int
	Height    int
	OutputURL string
}

// String 方便打印
func (p EmailDeliveryPayload) String() string {
	return fmt.Sprintf("UserID=%d, Email=%s, TemplateID=%s", p.UserID, p.Email, p.TemplateID)
}

func (p ImageResizePayload) String() string {
	return fmt.Sprintf("ImageURL=%s, Size=%dx%d, Output=%s", p.ImageURL, p.Width, p.Height, p.OutputURL)
}
