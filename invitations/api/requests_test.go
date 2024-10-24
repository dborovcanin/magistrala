// Copyright (c) Abstract Machines
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"fmt"
	"testing"

	"github.com/absmach/magistrala/invitations"
	"github.com/absmach/magistrala/pkg/apiutil"
	"github.com/absmach/magistrala/pkg/policies"
	"github.com/stretchr/testify/assert"
)

var valid = "valid"

func TestSendInvitationReqValidation(t *testing.T) {
	cases := []struct {
		desc string
		req  sendInvitationReq
		err  error
	}{
		{
			desc: "valid request",
			req: sendInvitationReq{
				token:    valid,
				UserID:   valid,
				DomainID: valid,
				Relation: policies.DomainRelation,
				Resend:   true,
			},
			err: nil,
		},
		{
			desc: "empty token",
			req: sendInvitationReq{
				token:    "",
				UserID:   valid,
				DomainID: valid,
				Relation: policies.DomainRelation,
				Resend:   true,
			},
			err: apiutil.ErrBearerToken,
		},
		{
			desc: "empty user ID",
			req: sendInvitationReq{
				token:    valid,
				UserID:   "",
				DomainID: valid,
				Relation: policies.DomainRelation,
				Resend:   true,
			},
			err: apiutil.ErrMissingID,
		},
		{
			desc: "empty domain_id",
			req: sendInvitationReq{
				token:    valid,
				UserID:   valid,
				DomainID: "",
				Relation: policies.DomainRelation,
				Resend:   true,
			},
			err: apiutil.ErrMissingDomainID,
		},
		{
			desc: "missing relation",
			req: sendInvitationReq{
				token:    valid,
				UserID:   valid,
				DomainID: valid,
				Relation: "",
				Resend:   true,
			},
			err: apiutil.ErrMissingRelation,
		},
		{
			desc: "invalid relation",
			req: sendInvitationReq{
				token:    valid,
				UserID:   valid,
				DomainID: valid,
				Relation: "invalid",
				Resend:   true,
			},
			err: apiutil.ErrInvalidRelation,
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.req.validate()
			assert.Equal(t, tc.err, err, fmt.Sprintf("%s: expected %s got %s\n", tc.desc, tc.err, err))
		})
	}
}

func TestListInvitationsReq(t *testing.T) {
	cases := []struct {
		desc string
		req  listInvitationsReq
		err  error
	}{
		{
			desc: "valid request",
			req: listInvitationsReq{
				token: valid,
				Page: invitations.Page{
					Limit:    1,
					DomainID: valid,
				},
			},
			err: nil,
		},
		{
			desc: "empty domainID",
			req: listInvitationsReq{
				token: valid,
				Page:  invitations.Page{Limit: 1},
			},
			err: apiutil.ErrMissingDomainID,
		},
		{
			desc: "empty token",
			req: listInvitationsReq{
				token: "",
				Page: invitations.Page{
					Limit:    1,
					DomainID: valid,
				},
			},
			err: apiutil.ErrBearerToken,
		},
		{
			desc: "invalid limit",
			req: listInvitationsReq{
				token: valid,
				Page: invitations.Page{
					Limit:    1000,
					DomainID: valid,
				},
			},
			err: apiutil.ErrLimitSize,
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.req.validate()
			assert.Equal(t, tc.err, err, fmt.Sprintf("%s: expected %s got %s\n", tc.desc, tc.err, err))
		})
	}
}

func TestAcceptInvitationReq(t *testing.T) {
	cases := []struct {
		desc string
		req  acceptInvitationReq
		err  error
	}{
		{
			desc: "valid request",
			req: acceptInvitationReq{
				token:    valid,
				domainID: valid,
			},
			err: nil,
		},
		{
			desc: "empty token",
			req: acceptInvitationReq{
				token: "",
			},
			err: apiutil.ErrBearerToken,
		},
		{
			desc: "empty domain_id",
			req: acceptInvitationReq{
				token:    valid,
				domainID: "",
			},
			err: apiutil.ErrMissingDomainID,
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.req.validate()
			assert.Equal(t, tc.err, err, fmt.Sprintf("%s: expected %s got %s\n", tc.desc, tc.err, err))
		})
	}
}

func TestInvitationReqValidation(t *testing.T) {
	cases := []struct {
		desc string
		req  invitationReq
		err  error
	}{
		{
			desc: "valid request",
			req: invitationReq{
				token:    valid,
				userID:   valid,
				domainID: valid,
			},
			err: nil,
		},
		{
			desc: "empty token",
			req: invitationReq{
				token:    "",
				userID:   valid,
				domainID: valid,
			},
			err: apiutil.ErrBearerToken,
		},
		{
			desc: "empty user ID",
			req: invitationReq{
				token:    valid,
				userID:   "",
				domainID: valid,
			},
			err: apiutil.ErrMissingID,
		},
		{
			desc: "empty domain",
			req: invitationReq{
				token:    valid,
				userID:   valid,
				domainID: "",
			},
			err: apiutil.ErrMissingDomainID,
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.req.validate()
			assert.Equal(t, tc.err, err, fmt.Sprintf("%s: expected %s got %s\n", tc.desc, tc.err, err))
		})
	}
}
