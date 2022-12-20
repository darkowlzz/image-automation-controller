package source

import (
	"context"

	"github.com/fluxcd/pkg/git"
	"k8s.io/apimachinery/pkg/types"

	sourcev1 "github.com/fluxcd/source-controller/api/v1beta2"

	imagev1 "github.com/fluxcd/image-automation-controller/api/v1beta1"
)

type SourceManager struct {
	gitClient git.RepositoryClient
	workDir   string
	gitRepo   *sourcev1.GitRepository
}

type sourceOptions struct {
	gitSpec             *imagev1.GitSpec
	noCrossNamespaceRef bool
}

type Option func(*sourceOptions)

func WithGitSpec(spec *imagev1.GitSpec) Option {
	return func(so *sourceOptions) {
		so.gitSpec = spec
	}
}

func WithNoCrossNamespaceRef(noCrossNSRef bool) Option {
	return func(so *sourceOptions) {
		so.noCrossNamespaceRef = noCrossNSRef
	}
}

func NewSourceManager(ctx context.Context, srcKind string, srcKey types.NamespacedName, options ...Option) (*SourceManager, error) {
	sm := &SourceManager{}

	return sm, nil
}

func (sm SourceManager) Push() error {
	return nil
}
