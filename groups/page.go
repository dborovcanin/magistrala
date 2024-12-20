// Copyright (c) Abstract Machines
// SPDX-License-Identifier: Apache-2.0

package groups

import "github.com/absmach/supermq/pkg/clients"

// PageMeta contains page metadata that helps navigation.
type PageMeta struct {
	Total      uint64           `json:"total"`
	Offset     uint64           `json:"offset"`
	Limit      uint64           `json:"limit"`
	Name       string           `json:"name,omitempty"`
	ID         string           `json:"id,omitempty"`
	Level      int              `json:"level,omitempty"`
	Path       string           `json:"path,omitempty"`
	DomainID   string           `json:"domain_id,omitempty"`
	Tag        string           `json:"tag,omitempty"`
	Metadata   clients.Metadata `json:"metadata,omitempty"`
	Status     clients.Status   `json:"status,omitempty"`
	Permission string
	ListPerms  bool
}
