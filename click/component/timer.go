package component

import (
	"fmt"
	"image/color"
	"time"
)

// // 每帧时间，单位为秒
// var dt float64

// func init() {
// 	dt = 1.0 / ebiten.ActualTPS()
// }

type TimerMode int

const (
	TimerModeCountdown TimerMode = iota
	TimerModeCountup
)

type TimerStatus int

const (
	TimerStatusRunning TimerStatus = iota
	TimerStatusPaused
	TimerStatusStopped
)

// 定义时间格式类型
type TimerFormat int

const (
	TimerFormatHour        TimerFormat = 1 << iota // 小时
	TimerFormatMinute                              // 分钟
	TimerFormatSecond                              // 秒
	TimerFormatMillisecond                         // 毫秒
)

type Timer struct {
	TextArea
	mode      TimerMode
	status    TimerStatus
	startTime time.Time
	duration  time.Duration
	limit     time.Duration // 计时上限
	savePoint time.Duration // 暂停时的时间点
	format    TimerFormat   // 时间格式

	onTimerEnd func() // 计时结束回调
}

func NewTimer(x, y, width, height int) *Timer {
	textArea := NewTextArea(x, y, width, height, "")
	textArea.SetBgColor(color.NRGBA{}) // 透明
	timer := Timer{
		TextArea: *textArea,
		mode:     TimerModeCountdown,
		status:   TimerStatusStopped,
		duration: 0,
		limit:    60 * time.Second,
		format:   TimerFormatSecond | TimerFormatMillisecond,
	}
	timer.UpdateText(timer.FormatDuration())
	return &timer
}

func (t *Timer) SetMode(mode TimerMode) {
	t.mode = mode
}

func (t *Timer) SetDuration(duration time.Duration) {
	t.duration = duration
}

func (t *Timer) SetOnTimerEnd(onTimerEnd func()) {
	t.onTimerEnd = onTimerEnd
}

func (t *Timer) Start() {
	if t.status == TimerStatusStopped { // 重新开始
		t.Reset()
	}
	t.startTime = time.Now()
	t.status = TimerStatusRunning
}

func (t *Timer) Pause() {
	t.status = TimerStatusPaused
	t.savePoint = t.duration
}

func (t *Timer) Stop() {
	t.status = TimerStatusStopped
	if t.onTimerEnd != nil {
		t.onTimerEnd()
	}
}

func (t *Timer) Reset() {
	t.status = TimerStatusStopped

	if t.mode == TimerModeCountup {
		// 正计时模式，设置初始值
		t.duration = time.Duration(0)
		t.savePoint = time.Duration(0)
	} else {
		// 倒计时模式，设置初始值为上限
		t.duration = t.limit
		t.savePoint = t.limit
	}
}

func (t *Timer) Update() error {
	switch t.status {
	case TimerStatusRunning:
		if t.mode == TimerModeCountdown {
			// t.duration -= time.Duration(dt * float64(time.Second))
			elapsed := time.Since(t.startTime)
			t.duration = t.savePoint - elapsed
			if t.duration <= 0 {
				t.Stop()
			}
		} else {
			// t.duration += time.Duration(dt * float64(time.Second))
			elapsed := time.Since(t.startTime)
			t.duration = t.savePoint + elapsed
			if t.duration >= t.limit {
				t.Stop()
			}
		}
		t.UpdateText(t.FormatDuration())
	case TimerStatusPaused:
		// 暂停
		return nil
	case TimerStatusStopped:
		// 停止
		return nil
	}
	return nil
}

func (t *Timer) SetFormat(format TimerFormat) {
	t.format = format
}

func (t *Timer) FormatDuration() string {
	hours := int(t.duration.Hours())
	minutes := int(t.duration.Minutes()) % 60
	seconds := int(t.duration.Seconds()) % 60
	milliseconds := int(t.duration.Milliseconds()) % 1000

	var result string

	if t.format&TimerFormatHour != 0 {
		result += fmt.Sprintf("%02d:", hours)
	}
	if t.format&TimerFormatMinute != 0 {
		result += fmt.Sprintf("%02d:", minutes)
	}
	if t.format&TimerFormatSecond != 0 {
		result += fmt.Sprintf("%02d", seconds)
	}
	if t.format&TimerFormatMillisecond != 0 {
		result += fmt.Sprintf(".%03d", milliseconds)
	}

	// 如果结果以冒号结尾，去掉最后一个冒号
	if len(result) > 0 && result[len(result)-1] == ':' {
		result = result[:len(result)-1]
	}

	return result
}
