// Package engine implements a chess engine.
package engine

import (
	"context"
	"fmt"
	"io"

	"github.com/clfs/they/internal/encoding/uci"
)

const Banner = "they!"

type Engine struct {
	r io.Reader
	w io.Writer
}

func New(r io.Reader, w io.Writer) *Engine {
	return &Engine{
		r: r,
		w: w,
	}
}

func (e *Engine) Run(ctx context.Context) error {
	fmt.Fprintln(e.w, Banner)

	dec := uci.NewDecoder(e.r)

	for {
		m, err := dec.ReadMessage()
		if err != nil {
			return err
		}

		switch m.(type) {
		case *uci.UCI:
			// do something
		}
	}
}
