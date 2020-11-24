// Copyright 2017-2020 The Usacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by 'github.com/sacloud/usacloud/tools/gen-commands'; DO NOT EDIT

package services

import (
	service "github.com/sacloud/libsacloud/v2/helper/service/database"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/conv"
)

func init() {
	setDefaultServiceFunc("database", "boot",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())

			req := &service.BootRequest{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}

			err := svc.BootWithContext(ctx, req)
			if err != nil {
				return nil, err
			}

			return nil, nil

		},
	)
	setDefaultListAllFunc("database", "boot",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())
			res, err := svc.FindWithContext(ctx, &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()})
			if err != nil {
				return nil, err
			}

			var results []interface{}
			for _, v := range res {
				results = append(results, v)
			}
			return results, nil
		},
	)
	setDefaultServiceFunc("database", "shutdown",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())

			req := &service.ShutdownRequest{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}

			err := svc.ShutdownWithContext(ctx, req)
			if err != nil {
				return nil, err
			}

			return nil, nil

		},
	)
	setDefaultListAllFunc("database", "shutdown",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())
			res, err := svc.FindWithContext(ctx, &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()})
			if err != nil {
				return nil, err
			}

			var results []interface{}
			for _, v := range res {
				results = append(results, v)
			}
			return results, nil
		},
	)
	setDefaultServiceFunc("database", "reset",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())

			req := &service.ResetRequest{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}

			err := svc.ResetWithContext(ctx, req)
			if err != nil {
				return nil, err
			}

			return nil, nil

		},
	)
	setDefaultListAllFunc("database", "reset",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())
			res, err := svc.FindWithContext(ctx, &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()})
			if err != nil {
				return nil, err
			}

			var results []interface{}
			for _, v := range res {
				results = append(results, v)
			}
			return results, nil
		},
	)
	setDefaultServiceFunc("database", "list",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())

			req := &service.FindRequest{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}

			res, err := svc.FindWithContext(ctx, req)
			if err != nil {
				return nil, err
			}

			var results []interface{}
			for _, v := range res {
				results = append(results, v)
			}
			return results, nil

		},
	)
	setDefaultListAllFunc("database", "list",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())
			res, err := svc.FindWithContext(ctx, &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()})
			if err != nil {
				return nil, err
			}

			var results []interface{}
			for _, v := range res {
				results = append(results, v)
			}
			return results, nil
		},
	)
	setDefaultServiceFunc("database", "create",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())

			req := &service.CreateRequest{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}

			res, err := svc.CreateWithContext(ctx, req)
			if err != nil {
				return nil, err
			}

			return []interface{}{res}, nil

		},
	)
	setDefaultListAllFunc("database", "create",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())
			res, err := svc.FindWithContext(ctx, &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()})
			if err != nil {
				return nil, err
			}

			var results []interface{}
			for _, v := range res {
				results = append(results, v)
			}
			return results, nil
		},
	)
	setDefaultServiceFunc("database", "read",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())

			req := &service.ReadRequest{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}

			res, err := svc.ReadWithContext(ctx, req)
			if err != nil {
				return nil, err
			}

			return []interface{}{res}, nil

		},
	)
	setDefaultListAllFunc("database", "read",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())
			res, err := svc.FindWithContext(ctx, &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()})
			if err != nil {
				return nil, err
			}

			var results []interface{}
			for _, v := range res {
				results = append(results, v)
			}
			return results, nil
		},
	)
	setDefaultServiceFunc("database", "update",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())

			req := &service.UpdateRequest{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}

			res, err := svc.UpdateWithContext(ctx, req)
			if err != nil {
				return nil, err
			}

			return []interface{}{res}, nil

		},
	)
	setDefaultListAllFunc("database", "update",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())
			res, err := svc.FindWithContext(ctx, &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()})
			if err != nil {
				return nil, err
			}

			var results []interface{}
			for _, v := range res {
				results = append(results, v)
			}
			return results, nil
		},
	)
	setDefaultServiceFunc("database", "delete",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())

			req := &service.DeleteRequest{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}

			err := svc.DeleteWithContext(ctx, req)
			if err != nil {
				return nil, err
			}

			return nil, nil

		},
	)
	setDefaultListAllFunc("database", "delete",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())
			res, err := svc.FindWithContext(ctx, &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()})
			if err != nil {
				return nil, err
			}

			var results []interface{}
			for _, v := range res {
				results = append(results, v)
			}
			return results, nil
		},
	)
	setDefaultServiceFunc("database", "monitor-cpu",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())

			req := &service.MonitorCPURequest{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}

			res, err := svc.MonitorCPUWithContext(ctx, req)
			if err != nil {
				return nil, err
			}

			var results []interface{}
			for _, v := range res {
				results = append(results, v)
			}
			return results, nil

		},
	)
	setDefaultListAllFunc("database", "monitor-cpu",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())
			res, err := svc.FindWithContext(ctx, &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()})
			if err != nil {
				return nil, err
			}

			var results []interface{}
			for _, v := range res {
				results = append(results, v)
			}
			return results, nil
		},
	)
	setDefaultServiceFunc("database", "monitor-database",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())

			req := &service.MonitorDatabaseRequest{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}

			res, err := svc.MonitorDatabaseWithContext(ctx, req)
			if err != nil {
				return nil, err
			}

			var results []interface{}
			for _, v := range res {
				results = append(results, v)
			}
			return results, nil

		},
	)
	setDefaultListAllFunc("database", "monitor-database",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())
			res, err := svc.FindWithContext(ctx, &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()})
			if err != nil {
				return nil, err
			}

			var results []interface{}
			for _, v := range res {
				results = append(results, v)
			}
			return results, nil
		},
	)
	setDefaultServiceFunc("database", "monitor-disk",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())

			req := &service.MonitorDiskRequest{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}

			res, err := svc.MonitorDiskWithContext(ctx, req)
			if err != nil {
				return nil, err
			}

			var results []interface{}
			for _, v := range res {
				results = append(results, v)
			}
			return results, nil

		},
	)
	setDefaultListAllFunc("database", "monitor-disk",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())
			res, err := svc.FindWithContext(ctx, &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()})
			if err != nil {
				return nil, err
			}

			var results []interface{}
			for _, v := range res {
				results = append(results, v)
			}
			return results, nil
		},
	)
	setDefaultServiceFunc("database", "monitor-interface",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())

			req := &service.MonitorInterfaceRequest{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}

			res, err := svc.MonitorInterfaceWithContext(ctx, req)
			if err != nil {
				return nil, err
			}

			var results []interface{}
			for _, v := range res {
				results = append(results, v)
			}
			return results, nil

		},
	)
	setDefaultListAllFunc("database", "monitor-interface",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())
			res, err := svc.FindWithContext(ctx, &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()})
			if err != nil {
				return nil, err
			}

			var results []interface{}
			for _, v := range res {
				results = append(results, v)
			}
			return results, nil
		},
	)
	setDefaultServiceFunc("database", "wait-until-ready",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())

			req := &service.WaitBootRequest{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}

			err := svc.WaitBootWithContext(ctx, req)
			if err != nil {
				return nil, err
			}

			return nil, nil

		},
	)
	setDefaultListAllFunc("database", "wait-until-ready",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())
			res, err := svc.FindWithContext(ctx, &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()})
			if err != nil {
				return nil, err
			}

			var results []interface{}
			for _, v := range res {
				results = append(results, v)
			}
			return results, nil
		},
	)
	setDefaultServiceFunc("database", "wait-until-shutdown",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())

			req := &service.WaitBootRequest{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}

			err := svc.WaitBootWithContext(ctx, req)
			if err != nil {
				return nil, err
			}

			return nil, nil

		},
	)
	setDefaultListAllFunc("database", "wait-until-shutdown",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())
			res, err := svc.FindWithContext(ctx, &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()})
			if err != nil {
				return nil, err
			}

			var results []interface{}
			for _, v := range res {
				results = append(results, v)
			}
			return results, nil
		},
	)
}
