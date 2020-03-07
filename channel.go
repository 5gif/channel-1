// Package channel provides mechanisms to model wireless medium used for simulating wireless links
// Supports M.2412 based environments and generation of fading generation for each links (SISO and MIMO)
package channel

import (
	"fmt"
	"log"
	"math/rand"
)

// Stores the wireless environment related paramters
// Each environment can have multiple wireless link-pairs (SRC-DEST)
type Env struct {
	EnvParams TestEnvironment //TestEnvironment Parameters based on in M.2412
	NLinks    int             // Number of wireless links
	Links     []WirelessLink
	base      BaseParam
}

// NewChannel creates a Default Envinronment
func NewEnv() *Env {
	env := new(Env)
	env.EnvParams = DefaultEnv()
	env.EnvParams.Initialize()
	return env
}

// NewSimpleEnv creates a Default Rural based Wireless Channel Env, with single Tx-Rx links, with IID generator
func NewSimpleEnv() *Env {
	env := new(Env)
	env.EnvParams = DefaultEnv()
	env.EnvParams.Initialize()
	return env
}

func (e *Env) Setup(fGHz float64, bwMHz float64) {
	e.base.FcGHz = fGHz
	e.base.BWMHz = bwMHz
}

// Create creates nlinks with each link having NxM mimo configuration where mimo=[N,M], mimo=[] then 1x1 system is assumed
func (e *Env) Create(nlinks int, ntx, nrx int) {
	e.NLinks = nlinks
	e.Links = make([]WirelessLink, nlinks)

	log.Printf("\n Creating %d x %d Links ", ntx, nrx)
	for i, link := range e.Links {
		link.ID = i
		link.BaseParam = e.base
		link.NTx, link.NRx = ntx, nrx
		fmt.Printf("\r %d ", i)
	}
}

// AttachGenerator attaches the fading generator fg,
// if clone=true all fading generator has same seed
func (e *Env) AttachGeneratorIID() {

	for i := 0; i < len(e.Links); i++ {
		state := rand.Uint64()
		iid := NewGeneratorIID(state)
		e.Links[i].generator = iid
	}
}
