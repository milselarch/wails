//go:build windows
// +build windows

package desktop

import (
	"context"
	"github.com/milselarch/wails/v2/internal/binding"
	"github.com/milselarch/wails/v2/internal/frontend"
	"github.com/milselarch/wails/v2/internal/frontend/desktop/windows"
	"github.com/milselarch/wails/v2/internal/logger"
	"github.com/milselarch/wails/v2/pkg/options"
)

func NewFrontend(ctx context.Context, appoptions *options.App, logger *logger.Logger, appBindings *binding.Bindings, dispatcher frontend.Dispatcher) frontend.Frontend {
	return windows.NewFrontend(ctx, appoptions, logger, appBindings, dispatcher)
}
