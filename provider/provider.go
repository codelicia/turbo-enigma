package provider

import (
	"turboenigma/model"
)

type Provider interface {
	NotifyMergeRequestCreated(mergeRequest model.MergeRequestInfo) error
	ReactToMessage(mergeRequest model.MergeRequestInfo, reactionRule model.ReactionRule) error
	GetReactionRules() []model.ReactionRule
}
